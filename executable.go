package hedera

import (
	"context"
	protobuf "google.golang.org/protobuf/proto"
	"math"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/pkg/errors"

	"github.com/hashgraph/hedera-protobufs-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var logCtx zerolog.Logger

// A required init function to setup logging at the correct level
func init() { // nolint
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if os.Getenv("HEDERA_SDK_GO_LOG_PRETTY") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	switch os.Getenv("HEDERA_SDK_GO_LOG_LEVEL") {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	logCtx = log.With().Str("module", "hedera-sdk-go").Logger()
}

const maxAttempts = 10

type _ExecutionState uint32

const (
	executionStateRetry    _ExecutionState = 0
	executionStateFinished _ExecutionState = 1
	executionStateError    _ExecutionState = 2
	executionStateExpired  _ExecutionState = 3
)

type _Method struct {
	query func(
		context.Context,
		*services.Query,
		...grpc.CallOption,
	) (*services.Response, error)
	transaction func(
		context.Context,
		*services.Transaction,
		...grpc.CallOption,
	) (*services.TransactionResponse, error)
}

type _Response struct {
	query       *services.Response
	transaction *services.TransactionResponse
}

type _IntermediateResponse struct {
	query       *services.Response
	transaction TransactionResponse
}

type _ProtoRequest struct {
	query       *services.Query
	transaction *services.Transaction
}

type QueryHeader struct {
	header *services.QueryHeader //nolint
}

type _Request struct {
	query       *Query
	transaction *Transaction
}

func _Execute(
	client *Client,
	request _Request,
	shouldRetry func(string, _Request, _Response) _ExecutionState,
	makeRequest func(request _Request) _ProtoRequest,
	advanceRequest func(_Request),
	getNodeAccountID func(_Request) AccountID,
	getMethod func(_Request, *_Channel) _Method,
	mapStatusError func(_Request, _Response) error,
	mapResponse func(_Request, _Response, AccountID, _ProtoRequest) (_IntermediateResponse, error),
	logID string,
	deadline *time.Duration,
) (_IntermediateResponse, error) {
	var maxAttempts int
	var minBackoff *time.Duration
	var maxBackoff *time.Duration

	if client.maxAttempts != nil {
		maxAttempts = *client.maxAttempts
	} else {
		if request.query != nil {
			maxAttempts = request.query.maxRetry
		} else {
			maxAttempts = request.transaction.maxRetry
		}
	}

	if request.query != nil {
		if request.query.maxBackoff == nil {
			maxBackoff = &client.maxBackoff
		} else {
			maxBackoff = request.query.maxBackoff
		}
		if request.query.minBackoff == nil {
			minBackoff = &client.minBackoff
		} else {
			minBackoff = request.query.minBackoff
		}
	} else {
		if request.transaction.maxBackoff == nil {
			maxBackoff = &client.maxBackoff
		} else {
			maxBackoff = request.transaction.maxBackoff
		}
		if request.transaction.minBackoff == nil {
			minBackoff = &client.minBackoff
		} else {
			minBackoff = request.transaction.minBackoff
		}
	}

	var attempt int64
	var errPersistent error

	for attempt = int64(0); attempt < int64(maxAttempts); attempt++ {
		var protoRequest _ProtoRequest
		var node *_Node
		var ok bool
		if request.transaction != nil {
			if request.transaction.nodeAccountIDs.locked && request.transaction.nodeAccountIDs._Length() > 0 {
				protoRequest = makeRequest(request)
				nodeAccountID := getNodeAccountID(request)
				node, ok = client.network._GetNodeForAccountID(nodeAccountID)
				if !ok {
					return _IntermediateResponse{}, ErrInvalidNodeAccountIDSet{nodeAccountID}
				}
				advanceRequest(request)
			} else {
				node = client.network._GetNode()
				request.transaction.nodeAccountIDs._Set(0, node.accountID)
				tx, _ := request.transaction._BuildTransaction(0)
				protoRequest = _ProtoRequest{
					transaction: tx,
				}
			}
		} else {
			if request.query.nodeAccountIDs.locked && request.query.nodeAccountIDs._Length() > 0 {
				protoRequest = makeRequest(request)
				nodeAccountID := getNodeAccountID(request)
				node, ok = client.network._GetNodeForAccountID(nodeAccountID)
				advanceRequest(request)
			} else {
				node = client.network._GetNode()
				if len(request.query.paymentTransactions) > 0 {
					var paymentTransaction services.TransactionBody
					protobuf.Unmarshal(request.query.paymentTransactions[0].BodyBytes, &paymentTransaction)
					paymentTransaction.NodeAccountID = node.accountID._ToProtobuf()
					request.query.paymentTransactions[0].BodyBytes, _ = protobuf.Marshal(&paymentTransaction)
				}
				//for _, s := range client.network.nodes {
				//	if node, ok := s.(*_Node); ok {
				//		println(node.accountID.String(), "node", node.readmitTime)
				//	}
				//}
				//for _, s := range client.network.goodNodes {
				//	if node, ok := s.(*_Node); ok {
				//		println(node.accountID.String(), "good node", node.readmitTime)
				//	}
				//}
				request.query.nodeAccountIDs._Set(0, node.accountID)
				protoRequest = makeRequest(request)
			}
		}

		node._InUse()

		logCtx.Trace().Str("requestId", logID).Str("nodeAccountID", node.accountID.String()).Str("nodeIPAddress", node.address._String())

		if !node._IsHealthy() {
			logCtx.Trace().Str("requestId", logID).Str("delay", node._Wait().String()).Msg("node is unhealthy, waiting before continuing")
			delay := node._Wait()
			time.Sleep(delay)
		}

		logCtx.Trace().Str("requestId", logID).Msg("updating node account ID index")
		advanceRequest(request)

		channel, err := node._GetChannel()
		if err != nil {
			client.network._IncreaseDelay(node)
			continue
		}

		method := getMethod(request, channel)

		resp := _Response{}

		ctx := context.TODO()
		var cancel context.CancelFunc
		if deadline != nil {
			grpcDeadline := time.Now().Add(*deadline)
			ctx, cancel = context.WithDeadline(ctx, grpcDeadline)
		}

		logCtx.Trace().Str("requestId", logID).Msg("executing gRPC call")
		if method.query != nil {
			resp.query, err = method.query(ctx, protoRequest.query)
		} else {
			resp.transaction, err = method.transaction(ctx, protoRequest.transaction)
		}

		if cancel != nil {
			cancel()
		}

		if err != nil {
			errPersistent = err
			if _ExecutableDefaultRetryHandler(logID, err) {
				client.network._IncreaseDelay(node)
				continue
			}
			if errPersistent == nil {
				errPersistent = errors.New("error")
			}
			return _IntermediateResponse{}, errors.Wrapf(errPersistent, "retry %d/%d", attempt, maxAttempts)
		}

		node._DecreaseDelay()

		retry := shouldRetry(logID, request, resp)

		switch retry {
		case executionStateRetry:
			errPersistent = mapStatusError(request, resp)
			_DelayForAttempt(logID, minBackoff, maxBackoff, attempt)
			client.network._IncreaseDelay(node)
			continue
		case executionStateExpired:
			if !client.GetOperatorAccountID()._IsZero() && request.transaction.regenerateTransactionID && !request.transaction.transactionIDs.locked {
				logCtx.Trace().Str("requestId", logID).Msg("received `TRANSACTION_EXPIRED` with transaction ID regeneration enabled; regenerating")
				_, err = request.transaction.transactionIDs._Set(request.transaction.nextTransactionIndex, TransactionIDGenerate(client.GetOperatorAccountID()))
				if err != nil {
					panic(err)
				}
				continue
			} else {
				return _IntermediateResponse{}, mapStatusError(request, resp)
			}
		case executionStateError:
			return _IntermediateResponse{}, mapStatusError(request, resp)
		case executionStateFinished:
			return mapResponse(request, resp, node.accountID, protoRequest)
		}
	}

	if errPersistent == nil {
		errPersistent = errors.New("error")
	}

	return _IntermediateResponse{}, errors.Wrapf(errPersistent, "retry %d/%d", attempt, maxAttempts)
}

func _DelayForAttempt(logID string, minBackoff *time.Duration, maxBackoff *time.Duration, attempt int64) {
	// 0.1s, 0.2s, 0.4s, 0.8s, ...
	ms := int64(math.Min(float64(minBackoff.Milliseconds())*math.Pow(2, float64(attempt)), float64(maxBackoff.Milliseconds())))
	logCtx.Trace().Str("requestId", logID).Dur("delay", time.Duration(ms)).Int64("attempt", attempt+1).Msg("retrying  request attempt")
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func _ExecutableDefaultRetryHandler(logID string, err error) bool {
	code := status.Code(err)
	logCtx.Trace().Str("requestId", logID).Str("status", code.String()).Msg("received gRPC error with status code")

	switch code {
	case codes.ResourceExhausted, codes.Unavailable:
		return true
	case codes.Internal:
		grpcErr, ok := status.FromError(err)

		if !ok {
			return false
		}

		return rstStream.Match([]byte(grpcErr.Message()))
	default:
		return false
	}
}
