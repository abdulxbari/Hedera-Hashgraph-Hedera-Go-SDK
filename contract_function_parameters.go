package hedera

/*-
 *
 * Hedera Go SDK
 *
 * Copyright (C) 2020 - 2022 Hedera Hashgraph, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

// ContractFunctionParameters is a struct which builds a solidity function call
// Use the builder methods `Add<Type>()` to add a parameter. Not all solidity types
// are supported out of the box, but the most common types are. The larger variants
// of number types require the parameter to be `[]byte`. This is a little unintuitive,
// so here is an example of how to use those larger number variants using
// "github.com/ethereum/go-ethereum/common/math" and "math/big"
// ```
// AddUint88(math.PaddedBigBytes(n, 88 / 8))
// ```
// If you're using `Uint256` specifically you can opt into using
// ```
// AddUin256(math.PaddedBigBytes(math.U256(n), 32))
// ```
type ContractFunctionParameters struct {
	function  ContractFunctionSelector
	arguments []Argument
}

type Argument struct {
	value   []byte
	dynamic bool
}

// Builder for encoding parameters for a Solidity contract constructor/function call.
func NewContractFunctionParameters() *ContractFunctionParameters {
	return &ContractFunctionParameters{
		function:  NewContractFunctionSelector(""),
		arguments: []Argument{},
	}
}

// AddBool adds a bool parameter to the function call
func (contract *ContractFunctionParameters) AddBool(value bool) *ContractFunctionParameters {
	argument := _NewArgument()

	if value {
		argument.value[31] = 1
	} else {
		argument.value[31] = 0
	}

	contract.function.AddBool()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddFunction adds a Solidity function reference and a function selector.
func (contract *ContractFunctionParameters) AddFunction(address string, selector ContractFunctionSelector) (*ContractFunctionParameters, error) {
	if len(address) != 40 {
		return contract, errors.Unwrap(fmt.Errorf("address is required to be 40 characters"))
	}

	argument := _NewArgument()
	argument.dynamic = false

	addressBytes, err := hex.DecodeString(address)
	if err != nil {
		return contract, err
	}

	bytes := make([]byte, 12)
	bytes = append(bytes, addressBytes[0:20]...)

	function := selector._Build(nil)

	bytes = append(bytes, function[0:4]...)
	argument.value = bytes

	contract.function.AddFunction()
	contract.arguments = append(contract.arguments, argument)
	return contract, nil
}

// AddInt8 adds an int8 parameter to the function call
func (contract *ContractFunctionParameters) AddInt8(value int8) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value[31] = uint8(value)

	contract.function.AddInt8()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt16 adds an int16 parameter to the function call
func (contract *ContractFunctionParameters) AddInt16(value int16) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint16(argument.value[30:32], uint16(value))

	contract.function.AddInt16()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt24 adds an int24 parameter to the function call
func (contract *ContractFunctionParameters) AddInt24(value int32) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint32(argument.value[28:32], uint32(value))

	contract.function.AddInt24()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt32 adds an int32 parameter to the function call
func (contract *ContractFunctionParameters) AddInt32(value int32) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint32(argument.value[28:32], uint32(value))

	contract.function.AddInt32()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt40 adds an int40 parameter to the function call
func (contract *ContractFunctionParameters) AddInt40(value int64) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint64(argument.value[24:32], uint64(value))

	contract.function.AddInt40()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt48 adds an int48 parameter to the function call
func (contract *ContractFunctionParameters) AddInt48(value int64) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint64(argument.value[24:32], uint64(value))

	contract.function.AddInt48()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt56 adds an int56 parameter to the function call
func (contract *ContractFunctionParameters) AddInt56(value int64) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint64(argument.value[24:32], uint64(value))

	contract.function.AddInt56()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt64 adds an int64 parameter to the function call
func (contract *ContractFunctionParameters) AddInt64(value int64) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint64(argument.value[24:32], uint64(value))

	contract.function.AddInt64()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt72 adds an int72 parameter to the function call
func (contract *ContractFunctionParameters) AddInt72(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:23], value...)

	contract.function.AddInt72()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt80 adds an int80 parameter to the function call
func (contract *ContractFunctionParameters) AddInt80(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:22], value...)

	contract.function.AddInt80()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt88 adds an int88 parameter to the function call
func (contract *ContractFunctionParameters) AddInt88(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:21], value...)

	contract.function.AddInt88()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt96 adds an int96 parameter to the function call
func (contract *ContractFunctionParameters) AddInt96(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:20], value...)

	contract.function.AddInt96()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt104 adds an int104 parameter to the function call
func (contract *ContractFunctionParameters) AddInt104(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:19], value...)

	contract.function.AddInt104()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt112 adds an int112 parameter to the function call
func (contract *ContractFunctionParameters) AddInt112(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:18], value...)

	contract.function.AddInt112()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt120 adds an int120 parameter to the function call
func (contract *ContractFunctionParameters) AddInt120(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:17], value...)

	contract.function.AddInt120()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt128 adds an int128 parameter to the function call
func (contract *ContractFunctionParameters) AddInt128(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:16], value...)

	contract.function.AddInt128()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt136 adds an int136 parameter to the function call
func (contract *ContractFunctionParameters) AddInt136(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:15], value...)

	contract.function.AddInt136()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt144 adds an int144 parameter to the function call
func (contract *ContractFunctionParameters) AddInt144(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:14], value...)

	contract.function.AddInt144()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt152 adds an int152 parameter to the function call
func (contract *ContractFunctionParameters) AddInt152(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:13], value...)

	contract.function.AddInt152()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt160 adds an int160 parameter to the function call
func (contract *ContractFunctionParameters) AddInt160(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:12], value...)

	contract.function.AddInt160()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt168 adds an int168 parameter to the function call
func (contract *ContractFunctionParameters) AddInt168(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:11], value...)

	contract.function.AddInt168()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt176 adds an int176 parameter to the function call
func (contract *ContractFunctionParameters) AddInt176(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:10], value...)

	contract.function.AddInt176()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt184 adds an int184 parameter to the function call
func (contract *ContractFunctionParameters) AddInt184(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:9], value...)

	contract.function.AddInt184()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt192 adds an int192 parameter to the function call
func (contract *ContractFunctionParameters) AddInt192(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:8], value...)

	contract.function.AddInt192()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt200 adds an int200 parameter to the function call
func (contract *ContractFunctionParameters) AddInt200(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:7], value...)

	contract.function.AddInt200()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt208 adds an int208 parameter to the function call
func (contract *ContractFunctionParameters) AddInt208(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:6], value...)

	contract.function.AddInt208()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt216 adds an int216 parameter to the function call
func (contract *ContractFunctionParameters) AddInt216(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:5], value...)

	contract.function.AddInt216()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt224 adds an int224 parameter to the function call
func (contract *ContractFunctionParameters) AddInt224(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:4], value...)

	contract.function.AddInt224()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt232 adds an int232 parameter to the function call
func (contract *ContractFunctionParameters) AddInt232(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:3], value...)

	contract.function.AddInt232()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt240 adds an int240 parameter to the function call
func (contract *ContractFunctionParameters) AddInt240(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:2], value...)

	contract.function.AddInt240()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt248 adds an int248 parameter to the function call
func (contract *ContractFunctionParameters) AddInt248(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:1], value...)

	contract.function.AddInt248()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt256 adds an int256 parameter to the function call
func (contract *ContractFunctionParameters) AddInt256(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = value

	contract.function.AddInt256()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint8 adds a uint8 parameter to the function call
func (contract *ContractFunctionParameters) AddUint8(value uint8) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value[31] = value

	contract.function.AddUint8()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint16 adds a uint16 parameter to the function call
func (contract *ContractFunctionParameters) AddUint16(value uint16) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint16(argument.value[30:32], value)

	contract.function.AddInt16()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint24 adds a uint24 parameter to the function call
func (contract *ContractFunctionParameters) AddUint24(value uint32) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint32(argument.value[28:32], value)

	contract.function.AddInt24()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint32 adds a uint32 parameter to the function call
func (contract *ContractFunctionParameters) AddUint32(value uint32) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint32(argument.value[28:32], value)

	contract.function.AddUint32()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint40 adds a uint40 parameter to the function call
func (contract *ContractFunctionParameters) AddUint40(value uint64) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint64(argument.value[24:32], value)

	contract.function.AddInt40()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint48 adds a uint48 parameter to the function call
func (contract *ContractFunctionParameters) AddUint48(value uint64) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint64(argument.value[24:32], value)

	contract.function.AddInt48()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint56 adds a uint56 parameter to the function call
func (contract *ContractFunctionParameters) AddUint56(value uint64) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint64(argument.value[24:32], value)

	contract.function.AddInt56()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint64 adds a uint64 parameter to the function call
func (contract *ContractFunctionParameters) AddUint64(value uint64) *ContractFunctionParameters {
	argument := _NewArgument()

	binary.BigEndian.PutUint64(argument.value[24:32], value)

	contract.function.AddUint64()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint72 adds a uint72 parameter to the function call
func (contract *ContractFunctionParameters) AddUint72(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:23], value...)

	contract.function.AddUint72()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint80 adds a uint80 parameter to the function call
func (contract *ContractFunctionParameters) AddUint80(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:22], value...)

	contract.function.AddUint80()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint88 adds a uint88 parameter to the function call
func (contract *ContractFunctionParameters) AddUint88(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:21], value...)

	contract.function.AddUint88()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint96 adds a uint96 parameter to the function call
func (contract *ContractFunctionParameters) AddUint96(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:20], value...)

	contract.function.AddUint96()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint104 adds a uint104 parameter to the function call
func (contract *ContractFunctionParameters) AddUint104(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:19], value...)

	contract.function.AddUint104()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint112 adds a uint112 parameter to the function call
func (contract *ContractFunctionParameters) AddUint112(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:18], value...)

	contract.function.AddUint112()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint120 adds a uint120 parameter to the function call
func (contract *ContractFunctionParameters) AddUint120(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:17], value...)

	contract.function.AddUint120()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint128 adds a uint128 parameter to the function call
func (contract *ContractFunctionParameters) AddUint128(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:16], value...)

	contract.function.AddUint128()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint136 adds a uint136 parameter to the function call
func (contract *ContractFunctionParameters) AddUint136(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:15], value...)

	contract.function.AddUint136()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint144 adds a uint144 parameter to the function call
func (contract *ContractFunctionParameters) AddUint144(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:14], value...)

	contract.function.AddUint144()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint152 adds a uint152 parameter to the function call
func (contract *ContractFunctionParameters) AddUint152(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:13], value...)

	contract.function.AddUint152()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint160 adds a uint160 parameter to the function call
func (contract *ContractFunctionParameters) AddUint160(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:12], value...)

	contract.function.AddUint160()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint168 adds a uint168 parameter to the function call
func (contract *ContractFunctionParameters) AddUint168(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:11], value...)

	contract.function.AddUint168()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint176 adds a uint176 parameter to the function call
func (contract *ContractFunctionParameters) AddUint176(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:10], value...)

	contract.function.AddUint176()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint184 adds a uint184 parameter to the function call
func (contract *ContractFunctionParameters) AddUint184(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:9], value...)

	contract.function.AddUint184()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint192 adds a uint192 parameter to the function call
func (contract *ContractFunctionParameters) AddUint192(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:8], value...)

	contract.function.AddUint192()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint200 adds a uint200 parameter to the function call
func (contract *ContractFunctionParameters) AddUint200(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:7], value...)

	contract.function.AddUint200()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint208 adds a uint208 parameter to the function call
func (contract *ContractFunctionParameters) AddUint208(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:6], value...)

	contract.function.AddUint208()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint216 adds a uint216 parameter to the function call
func (contract *ContractFunctionParameters) AddUint216(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:5], value...)

	contract.function.AddUint216()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint224 adds a uint224 parameter to the function call
func (contract *ContractFunctionParameters) AddUint224(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:4], value...)

	contract.function.AddUint224()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint232 adds a uint232 parameter to the function call
func (contract *ContractFunctionParameters) AddUint232(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:3], value...)

	contract.function.AddUint232()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint240 adds a uint240 parameter to the function call
func (contract *ContractFunctionParameters) AddUint240(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:2], value...)

	contract.function.AddUint240()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint248 adds a uint248 parameter to the function call
func (contract *ContractFunctionParameters) AddUint248(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = append(argument.value[:1], value...)

	contract.function.AddUint248()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddUint256 adds a uint256 parameter to the function call
func (contract *ContractFunctionParameters) AddUint256(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = value

	contract.function.AddUint256()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddInt8Array adds an int8 array parameter to the function call
func (contract *ContractFunctionParameters) AddInt8Array(value []int8) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)*32+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		binary.BigEndian.PutUint32(result[i*32+32+28:i*32+32+32], uint32(v))
	}

	argument.value = result

	contract.function.AddInt32Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddInt16Array adds an int16 array parameter to the function call
func (contract *ContractFunctionParameters) AddInt16Array(value []int16) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)*32+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		binary.BigEndian.PutUint32(result[i*32+32+28:i*32+32+32], uint32(v))
	}

	argument.value = result

	contract.function.AddInt32Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddInt24Array adds an int24 array parameter to the function call
func (contract *ContractFunctionParameters) AddInt24Array(value []int32) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)*32+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		binary.BigEndian.PutUint32(result[i*32+32+28:i*32+32+32], uint32(v))
	}

	argument.value = result

	contract.function.AddInt32Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddInt32Array adds an int32 array parameter to the function call
func (contract *ContractFunctionParameters) AddInt32Array(value []int32) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		binary.BigEndian.PutUint32(result[i*32+32+28:i*32+32+32], uint32(v))
	}

	argument.value = result

	contract.function.AddInt32Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddInt64Array adds an int64 array parameter to the function call
func (contract *ContractFunctionParameters) AddInt64Array(value []int64) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)*32+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		binary.BigEndian.PutUint64(result[i*32+32+24:i*32+32+32], uint64(v))
	}

	argument.value = result

	contract.function.AddInt64Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddInt256Array adds an int256 array parameter to the function call
func (contract *ContractFunctionParameters) AddInt256Array(value [][32]byte) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)*32+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		copy(result[i*32+32:i*32+32+32], v[0:32])
	}

	argument.value = result

	contract.function.AddInt256Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddUint32Array adds a uint32 array parameter to the function call
func (contract *ContractFunctionParameters) AddUint32Array(value []uint32) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)*32+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		binary.BigEndian.PutUint32(result[i*32+32+28:i*32+32+32], v)
	}

	argument.value = result

	contract.function.AddUint32Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddUint64Array adds a uint64 array parameter to the function call
func (contract *ContractFunctionParameters) AddUint64Array(value []uint64) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)*32+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		binary.BigEndian.PutUint64(result[i*32+32+24:i*32+32+32], v)
	}

	argument.value = result

	contract.function.AddUint64Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddUint256Array adds a uint256 array parameter to the function call
func (contract *ContractFunctionParameters) AddUint256Array(value [][32]byte) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)*32+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		copy(result[i*32+32:i*32+32+32], v[0:32])
	}

	argument.value = result

	contract.function.AddUint256Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddAddressArray adds an address array parameter to the function call
func (contract *ContractFunctionParameters) AddAddressArray(value []string) (*ContractFunctionParameters, error) {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)*32+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		if len(v) != 40 {
			return contract, errors.Unwrap(fmt.Errorf("address is required to be 40 characters"))
		}

		addressBytes, err := hex.DecodeString(v)
		if err != nil {
			return contract, err
		}

		copy(result[i*32+32+12:i*32+32+32], addressBytes[0:20])
	}

	argument.value = result

	contract.function.AddAddressArray()
	contract.arguments = append(contract.arguments, argument)
	return contract, nil
}

// AddString ads a string parameter to the function call
func (contract *ContractFunctionParameters) AddString(value string) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	bytes := []byte(value)
	binary.BigEndian.PutUint64(argument.value[24:32], uint64(len(bytes)))
	argument.value = append(argument.value, bytes...)
	argument.value = append(argument.value, make([]byte, 32-len(bytes)%32)...)

	contract.function.AddString()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddBytes adds a bytes parameter to the function call
func (contract *ContractFunctionParameters) AddBytes(value []byte) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	binary.BigEndian.PutUint64(argument.value[24:32], uint64(len(value)))
	argument.value = append(argument.value, value...)
	argument.value = append(argument.value, make([]byte, uint64(32-len(value)%32))...)

	contract.function.AddBytes()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddBytes32 adds a bytes32 parameter to the function call
func (contract *ContractFunctionParameters) AddBytes32(value [32]byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.value = value[:]

	contract.function.AddBytes32()
	contract.arguments = append(contract.arguments, argument)

	return contract
}

// AddAddress adds an address parameter to the function call
func (contract *ContractFunctionParameters) AddAddress(value string) (*ContractFunctionParameters, error) {
	if len(value) != 40 {
		return contract, errors.Unwrap(fmt.Errorf("address is required to be 40 characters"))
	}

	addressBytes, err := hex.DecodeString(value)
	if err != nil {
		return contract, err
	}

	argument := _NewArgument()
	argument.dynamic = false

	bytes := make([]byte, 12)
	bytes = append(bytes, addressBytes...)

	argument.value = bytes

	contract.function.AddAddress()
	contract.arguments = append(contract.arguments, argument)
	return contract, nil
}

// AddBytesArray adds a bytes array parameter to the function call
func (contract *ContractFunctionParameters) AddBytesArray(value [][]byte) *ContractFunctionParameters {
	argument := _NewArgument()

	argument.dynamic = true
	argument.value = bytesArray(value)

	contract.function.AddBytesArray()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddBytes32Array adds a bytes32 array parameter to the function call
func (contract *ContractFunctionParameters) AddBytes32Array(value [][]byte) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	result := make([]byte, len(value)+32)

	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	for i, v := range value {
		copy(result[i*32+32:i*32+32+32], v[0:32])
	}

	argument.value = result

	contract.function.AddBytes32Array()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

// AddStringArray adds a string array parameter to the function call
func (contract *ContractFunctionParameters) AddStringArray(value []string) *ContractFunctionParameters {
	argument := _NewArgument()
	argument.dynamic = true

	var bytes [][]byte
	for _, s := range value {
		bytes = append(bytes, []byte(s))
	}

	argument.value = bytesArray(bytes)
	contract.function.AddStringArray()
	contract.arguments = append(contract.arguments, argument)
	return contract
}

func (contract *ContractFunctionParameters) _Build(functionName *string) []byte {
	length := uint64(0)

	functionOffset := uint64(0)
	if functionName != nil {
		functionOffset = uint64(4)
	}

	for _, argument := range contract.arguments {
		length += uint64(32)
		if argument.dynamic {
			length += uint64(len(argument.value))
		}
	}

	result := make([]byte, length+functionOffset)
	if functionName != nil {
		copy(result[0:4], contract.function._Build(functionName))
	}

	offset := uint64(len(contract.arguments) * 32)

	for i, argument := range contract.arguments {
		j := uint64(i)
		if argument.dynamic {
			binary.BigEndian.PutUint64(result[(j*32+functionOffset)+24:(j+1)*32+functionOffset], offset)
			copy(result[offset+functionOffset:], argument.value)
			offset += uint64(len(argument.value))
		} else {
			copy(result[j*32+functionOffset:((j+1)*32)+functionOffset], argument.value)
		}
	}

	return result
}

func _NewArgument() Argument {
	return Argument{
		value:   make([]byte, 32),
		dynamic: false,
	}
}

func bytesArray(value [][]byte) []byte {
	// Calculate Length of final result
	length := uint64(0)
	for _, s := range value {
		length += 32 + 32
		sbytes := s
		if len(sbytes)/32 == 0 {
			length += 32
		} else {
			length += uint64(((len(sbytes) / 32) + 1) * 32)
		}
	}

	// Zero initialize final resulting byte array
	result := make([]byte, length+32)

	// Write length of array into the first 32 bytes
	binary.BigEndian.PutUint64(result[24:32], uint64(len(value)))

	// Create array of byte arrays to hold each string value
	// Needed to concat later
	arguments := make([][]byte, len(value))

	// Convert each argument into bytes, and push each argument
	// into the argument list
	for i, s := range value {
		// Get the length of the current argument (again)
		var length uint64
		if len(s)/32 == 0 {
			length = 32
		} else {
			length = uint64(((len(s) / 32) + 1) * 32)
		}

		// Create byte array of correct size
		// Length of value to the nearest 32 byte boundary +
		// 32 bytes to store the length
		bytes := make([]byte, length+32)

		// Write length into first 32 bytes
		binary.BigEndian.PutUint64(bytes[24:32], uint64(len(s)))

		// Copy string as bytes to the rest of the buffer
		copy(bytes[32:], s)

		// Set the argument bytes to be used later
		arguments[i] = bytes
	}

	// Initialize offset to the number of strings
	offset := uint64(len(value) * 32)

	// For each argument, write the offset into result
	// and the argument value (which includes data and length already)
	for i, s := range arguments {
		binary.BigEndian.PutUint64(result[(i+1)*32+24:(i+2)*32], offset)
		copy(result[offset+32:offset+32+uint64(len(s))], s)
		offset += uint64(len(s))
	}

	return result
}
