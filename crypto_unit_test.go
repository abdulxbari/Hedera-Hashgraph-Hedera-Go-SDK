//go:build all || unit
// +build all unit

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
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const _Ed25519PubKeyPrefix = "302a300506032b6570032100"
const _ECDSAPubKeyPrefix = "3036301006072a8648ce3d020106052b8104000a0322000"

const testPrivateKeyStr = "302e020100300506032b657004220420db484b828e64b2d8f12ce3c0a0e93a0b8cce7af1bb8f39c97732394482538e10"

const testPublicKeyStr = "302a300506032b6570032100e0c8ec2758a5879ffac226a13c0c516b799e72e35141a0dd828f94d37988a4b7"

const testMnemonic3 = "obvious favorite remain caution remove laptop base vacant increase video erase pass sniff sausage knock grid argue salt romance way alone fever slush dune"

// generated by hedera-keygen-java, not used anywhere
const testMnemonic = "inmate flip alley wear offer often piece magnet surge toddler submit right radio absent pear floor belt raven price stove replace reduce plate home"
const testMnemonicKey = "302e020100300506032b657004220420853f15aecd22706b105da1d709b4ac05b4906170c2b9c7495dff9af49e1391da"

// backup phrase generated by the iOS wallet, not used anywhere
const iosMnemonicString = "tiny denial casual grass skull spare awkward indoor ethics dash enough flavor good daughter early hard rug staff capable swallow raise flavor empty angle"

// private key for "default account", should be index 0
const iosDefaultPrivateKey = "5f66a51931e8c99089472e0d70516b6272b94dd772b967f8221e1077f966dbda2b60cf7ee8cf10ecd5a076bffad9a7c7b97df370ad758c0f1dd4ef738e04ceb6"

// backup phrase generated by the Android wallet, also not used anywhere
const androidMnemonicString = "ramp april job flavor surround pyramid fish sea good know blame gate village viable include mixed term draft among monitor swear swing novel track"

// private key for "default account", should be index 0
const androidDefaultPrivateKey = "c284c25b3a1458b59423bc289e83703b125c8eefec4d5aa1b393c2beb9f2bae66188a344ba75c43918ab12fa2ea4a92960eca029a2320d8c6a1c3b94e06c9985"

// test pem key contests for the above testPrivateKeyStr
const pemString = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEINtIS4KOZLLY8SzjwKDpOguMznrxu485yXcyOUSCU44Q
-----END PRIVATE KEY-----
`

// const encryptedPem = `-----BEGIN ENCRYPTED PRIVATE KEY-----
// MIGbMFcGCSqGSIb3DQEFDTBKMCkGCSqGSIb3DQEFDDAcBAi8WY7Gy2tThQICCAAw
// DAYIKoZIhvcNAgkFADAdBglghkgBZQMEAQIEEOq46NPss58chbjUn20NoK0EQG1x
// R88hIXcWDOECttPTNlMXWJt7Wufm1YwBibrxmCq1QykIyTYhy1TZMyxyPxlYW6aV
// 9hlo4YEh3uEaCmfJzWM=
// -----END ENCRYPTED PRIVATE KEY-----`

const encryptedPem = `-----BEGIN ENCRYPTED PRIVATE KEY-----
MIGbMFcGCSqGSIb3DQEFDTBKMCkGCSqGSIb3DQEFDDAcBAi8WY7Gy2tThQICCAAw
DAYIKoZIhvcNAgkFADAdBglghkgBZQMEAQIEEOq46NPss58chbjUn20NoK0EQG1x
R88hIXcWDOECttPTNlMXWJt7Wufm1YwBibrxmCq1QykIyTYhy1TZMyxyPxlYW6aV
9hlo4YEh3uEaCmfJzWM=
-----END ENCRYPTED PRIVATE KEY-----
`
const pemPassphrase = "this is a passphrase"

func TestUnitPrivateKeyGenerate(t *testing.T) {
	key, err := GeneratePrivateKey()

	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(key.String(), _Ed25519PrivateKeyPrefix))
}

func TestUnitPrivateEd25519KeyGenerate(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()

	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(key.String(), _Ed25519PrivateKeyPrefix))
}

func TestUnitPrivateECDSAKeyGenerate(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()

	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(key.String(), _ECDSAPrivateKeyPrefix))
}

func TestUnitPrivateKeyExternalSerialization(t *testing.T) {
	key, err := PrivateKeyFromString(testPrivateKeyStr)

	require.NoError(t, err)
	assert.Equal(t, testPrivateKeyStr, key.String())
}

func TestUnitPrivateKeyExternalSerializationForConcatenatedHex(t *testing.T) {
	keyStr := "db484b828e64b2d8f12ce3c0a0e93a0b8cce7af1bb8f39c97732394482538e10e0c8ec2758a5879ffac226a13c0c516b799e72e35141a0dd828f94d37988a4b7"
	key, err := PrivateKeyFromStringEd25519(keyStr)

	require.NoError(t, err)
	assert.Equal(t, testPrivateKeyStr, key.String())
}

func TestUnitShouldMatchHbarWalletV1(t *testing.T) {
	mnemonic, err := MnemonicFromString("jolly kidnap tom lawn drunk chick optic lust mutter mole bride galley dense member sage neural widow decide curb aboard margin manure")
	require.NoError(t, err)

	key, err := mnemonic.ToLegacyPrivateKey()
	require.NoError(t, err)

	deriveKey, err := key.LegacyDerive(1099511627775)
	require.NoError(t, err)

	assert.Equal(t, "302a300506032b657003210045f3a673984a0b4ee404a1f4404ed058475ecd177729daa042e437702f7791e9", deriveKey.PublicKey().String())
}

func TestUnitLegacyPrivateKeyFromMnemonicDerive(t *testing.T) {
	mnemonic, err := MnemonicFromString("jolly kidnap tom lawn drunk chick optic lust mutter mole bride galley dense member sage neural widow decide curb aboard margin manure")
	require.NoError(t, err)

	key, err := mnemonic.ToLegacyPrivateKey()
	require.NoError(t, err)

	deriveKey, err := key.LegacyDerive(0)
	require.NoError(t, err)
	deriveKey2, err := key.LegacyDerive(-1)
	require.NoError(t, err)

	assert.Equal(t, "302e020100300506032b657004220420882a565ad8cb45643892b5366c1ee1c1ef4a730c5ce821a219ff49b6bf173ddf", deriveKey2.String())
	assert.Equal(t, "302e020100300506032b657004220420fae0002d2716ea3a60c9cd05ee3c4bb88723b196341b68a02d20975f9d049dc6", deriveKey.String())
}

func TestUnitPrivateKeyExternalSerializationForRawHex(t *testing.T) {
	keyStr := "db484b828e64b2d8f12ce3c0a0e93a0b8cce7af1bb8f39c97732394482538e10"
	key, err := PrivateKeyFromStringEd25519(keyStr)

	require.NoError(t, err)
	assert.Equal(t, testPrivateKeyStr, key.String())
}

func TestUnitPublicKeyExternalSerializationForDerEncodedHex(t *testing.T) {
	key, err := PublicKeyFromString(testPublicKeyStr)

	require.NoError(t, err)
	assert.Equal(t, testPublicKeyStr, key.String())
}

func TestUnitPublicKeyExternalSerializationForRawHex(t *testing.T) {
	keyStr := "e0c8ec2758a5879ffac226a13c0c516b799e72e35141a0dd828f94d37988a4b7"
	key, err := PublicKeyFromStringEd25519(keyStr)

	require.NoError(t, err)
	assert.Equal(t, testPublicKeyStr, key.String())
}

func TestUnitMnemonicToPrivateKey(t *testing.T) {
	mnemonic, err := MnemonicFromString(testMnemonic)
	require.NoError(t, err)

	key, err := mnemonic.ToPrivateKey("")
	require.NoError(t, err)

	assert.Equal(t, testMnemonicKey, key.String())
}

func TestUnitIOSPrivateKeyFromMnemonic(t *testing.T) {
	mnemonic, err := MnemonicFromString(iosMnemonicString)
	require.NoError(t, err)

	key, err := PrivateKeyFromMnemonic(mnemonic, "")
	require.NoError(t, err)

	derivedKey, err := key.Derive(0)
	require.NoError(t, err)

	expectedKey, err := PrivateKeyFromString(iosDefaultPrivateKey)
	require.NoError(t, err)

	assert.Equal(t, expectedKey.ed25519PrivateKey.keyData, derivedKey.ed25519PrivateKey.keyData)
}

func TestUnitAndroidPrivateKeyFromMnemonic(t *testing.T) {
	mnemonic, err := MnemonicFromString(androidMnemonicString)
	require.NoError(t, err)

	key, err := PrivateKeyFromMnemonic(mnemonic, "")
	require.NoError(t, err)

	derivedKey, err := key.Derive(0)
	require.NoError(t, err)

	expectedKey, err := PrivateKeyFromString(androidDefaultPrivateKey)
	require.NoError(t, err)

	assert.Equal(t, expectedKey.ed25519PrivateKey.keyData, derivedKey.ed25519PrivateKey.keyData)
}

func TestUnitMnemonic3(t *testing.T) {
	mnemonic, err := MnemonicFromString(testMnemonic3)
	require.NoError(t, err)

	key, err := mnemonic.ToLegacyPrivateKey()
	require.NoError(t, err)

	derivedKey, err := key.LegacyDerive(0)
	require.NoError(t, err)
	derivedKey2, err := key.LegacyDerive(-1)
	require.NoError(t, err)

	assert.Equal(t, "302e020100300506032b6570042204202b7345f302a10c2a6d55bf8b7af40f125ec41d780957826006d30776f0c441fb", derivedKey.String())
	assert.Equal(t, "302e020100300506032b657004220420caffc03fdb9853e6a91a5b3c57a5c0031d164ce1c464dea88f3114786b5199e5", derivedKey2.String())
}

func TestUnitSigning(t *testing.T) {
	priKey, err := PrivateKeyFromString(testPrivateKeyStr)
	require.NoError(t, err)

	pubKey, err := PublicKeyFromString(testPublicKeyStr)
	require.NoError(t, err)

	testSignData := []byte("this is the test data to sign")
	signature := priKey.Sign(testSignData)

	assert.True(t, ed25519.Verify(pubKey.Bytes(), []byte("this is the test data to sign"), signature))
}

func TestUnitGenerated24MnemonicToWorkingPrivateKey(t *testing.T) {
	mnemonic, err := GenerateMnemonic24()

	require.NoError(t, err)

	privateKey, err := mnemonic.ToPrivateKey("")

	require.NoError(t, err)

	message := []byte("this is a test message")

	signature := privateKey.Sign(message)

	assert.True(t, ed25519.Verify(privateKey.PublicKey().Bytes(), message, signature))
}

func TestUnitGenerated12MnemonicToWorkingPrivateKey(t *testing.T) {
	mnemonic, err := GenerateMnemonic12()

	require.NoError(t, err)

	privateKey, err := mnemonic.ToPrivateKey("")

	require.NoError(t, err)

	message := []byte("this is a test message")

	signature := privateKey.Sign(message)

	assert.True(t, ed25519.Verify(privateKey.PublicKey().Bytes(), message, signature))
}

func TestUnitPrivateKeyFromKeystore(t *testing.T) {
	privatekey, err := PrivateKeyFromKeystore([]byte(testKeystore), passphrase)
	require.NoError(t, err)

	actualPrivateKey, err := PrivateKeyFromStringEd25519(testKeystoreKeyString)
	require.NoError(t, err)

	assert.Equal(t, actualPrivateKey.ed25519PrivateKey.keyData, privatekey.ed25519PrivateKey.keyData)
}

func TestUnitPrivateKeyKeystore(t *testing.T) {
	privateKey, err := PrivateKeyFromString(testPrivateKeyStr)
	require.NoError(t, err)

	keystore, err := privateKey.Keystore(passphrase)
	require.NoError(t, err)

	ksPrivateKey, err := _ParseKeystore(keystore, passphrase)
	require.NoError(t, err)

	assert.Equal(t, privateKey.ed25519PrivateKey.keyData, ksPrivateKey.ed25519PrivateKey.keyData)
}

func TestUnitPrivateKeyReadKeystore(t *testing.T) {
	actualPrivateKey, err := PrivateKeyFromStringEd25519(testKeystoreKeyString)
	require.NoError(t, err)

	keystoreReader := bytes.NewReader([]byte(testKeystore))

	privateKey, err := PrivateKeyReadKeystore(keystoreReader, passphrase)
	require.NoError(t, err)

	assert.Equal(t, actualPrivateKey.ed25519PrivateKey.keyData, privateKey.ed25519PrivateKey.keyData)
}

func TestUnitPrivateKeyFromPem(t *testing.T) {
	actualPrivateKey, err := PrivateKeyFromString(testPrivateKeyStr)
	require.NoError(t, err)

	privateKey, err := PrivateKeyFromPem([]byte(pemString), "")
	require.NoError(t, err)

	assert.Equal(t, actualPrivateKey, privateKey)
}

func TestUnitPrivateKeyFromPemInvalid(t *testing.T) {
	_, err := PrivateKeyFromPem([]byte("invalid"), "")
	assert.Error(t, err)
}

func TestUnitPrivateKeyFromPemWithPassphrase(t *testing.T) {
	actualPrivateKey, err := PrivateKeyFromString(testPrivateKeyStr)
	require.NoError(t, err)

	privateKey, err := PrivateKeyFromPem([]byte(encryptedPem), pemPassphrase)
	require.NoError(t, err)

	assert.Equal(t, actualPrivateKey, privateKey)
}

func TestUnitPrivateKeyECDSASign(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)

	hash := crypto.Keccak256Hash([]byte("aaa"))
	sig := key.Sign([]byte("aaa"))
	s2 := crypto.VerifySignature(key.ecdsaPrivateKey._PublicKey()._BytesRaw(), hash.Bytes(), sig)
	require.True(t, s2)
}

func DisabledTestUnitPrivateKeyECDSASign(t *testing.T) {
	message := []byte("hello world")
	key, err := PrivateKeyFromStringECDSA("8776c6b831a1b61ac10dac0304a2843de4716f54b1919bb91a2685d0fe3f3048")
	require.NoError(t, err)

	sig := key.Sign(message)

	require.Equal(t, hex.EncodeToString(sig), "f3a13a555f1f8cd6532716b8f388bd4e9d8ed0b252743e923114c0c6cbfe414cf791c8e859afd3c12009ecf2cb20dacf01636d80823bcdbd9ec1ce59afe008f0")
	require.True(t, key.PublicKey().Verify(message, sig))
}

func TestUnitPrivateKeyEd25519FromString(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	key2, err := PrivateKeyFromString(key.String())
	require.NoError(t, err)

	require.Equal(t, key2.String(), key.String())
}

func TestUnitPrivateKeyEd25519FromStringRaw(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	key2, err := PrivateKeyFromStringEd25519(key.StringRaw())
	require.NoError(t, err)

	require.Equal(t, key2.String(), key.String())
}

func TestUnitPrivateKeyEd25519FromStringDer(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	key2, err := PrivateKeyFromStringEd25519(key.StringDer())
	require.NoError(t, err)

	require.Equal(t, key2.StringDer(), key.StringDer())
}

func TestUnitPublicKeyEd25519FromString(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	publicKey := key.PublicKey()
	publicKey2, err := PublicKeyFromStringEd25519(publicKey.String())
	require.NoError(t, err)
	require.Equal(t, publicKey2.String(), publicKey.String())
}

func TestUnitPublicKeyEd25519FromStringRaw(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	publicKey := key.PublicKey()
	publicKey2, err := PublicKeyFromStringEd25519(publicKey.StringRaw())
	require.NoError(t, err)

	require.Equal(t, publicKey2.String(), publicKey.String())
}

func TestUnitPublicKeyEd25519FromStringDer(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	publicKey := key.PublicKey()
	publicKey2, err := PublicKeyFromStringEd25519(publicKey.StringDer())
	require.NoError(t, err)

	require.Equal(t, publicKey2.StringDer(), publicKey.StringDer())
}

func TestUnitPrivateKeyECDSAFromString(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	key2, err := PrivateKeyFromString(key.String())
	require.NoError(t, err)

	require.Equal(t, key2.String(), key.String())
}

func TestUnitPrivateKeyECDSAFromStringRaw(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	key2, err := PrivateKeyFromStringECDSA(key.StringRaw())
	require.NoError(t, err)

	require.Equal(t, key2.String(), key.String())
}

func TestUnitPrivateKeyECDSAFromStringDer(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	key2, err := PrivateKeyFromStringECDSA(key.StringDer())
	require.NoError(t, err)

	require.Equal(t, key2.StringDer(), key.StringDer())
}

func TestUnitPublicKeyECDSAFromString(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	publicKey := key.PublicKey()
	publicKey2, err := PublicKeyFromStringECDSA(publicKey.String())
	require.NoError(t, err)
	require.Equal(t, publicKey2.String(), publicKey.String())
}

func TestUnitPublicKeyECDSAFromStringRaw(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	publicKey := key.PublicKey()
	publicKey2, err := PublicKeyFromStringECDSA(publicKey.StringRaw())
	require.NoError(t, err)

	require.Equal(t, publicKey2.String(), publicKey.String())
}

func TestUnitPublicKeyECDSAFromStringDer(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	publicKey := key.PublicKey()
	publicKey2, err := PublicKeyFromStringECDSA(publicKey.StringDer())
	require.NoError(t, err)

	require.Equal(t, publicKey2.StringDer(), publicKey.StringDer())
}

func TestUnitPrivateKeyFromBytesDerECDSA(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	bytes := key.BytesDer()
	key2, err := PrivateKeyFromBytesDer(bytes)
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(key2.String(), _ECDSAPrivateKeyPrefix))
}

func TestUnitPrivateKeyFromBytesDerEd25519(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	bytes := key.BytesDer()
	key2, err := PrivateKeyFromBytesDer(bytes)
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(key2.String(), _Ed25519PrivateKeyPrefix))
}

func TestUnitPublicKeyFromBytesDerECDSA(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	pkey := key.PublicKey()
	bytes := pkey.BytesDer()
	pkey2, err := PublicKeyFromBytesDer(bytes)
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(pkey2.String(), _ECDSAPubKeyPrefix))
}

func TestUnitPublicKeyFromBytesDerEd25519(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	pkey := key.PublicKey()
	bytes := pkey.BytesDer()
	pkey2, err := PublicKeyFromBytesDer(bytes)
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(pkey2.String(), _Ed25519PubKeyPrefix))
}

func TestUnitPrivateKeyFromStringDerEd25519(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	key2, err := PrivateKeyFromStringDer(key.StringDer())
	require.NoError(t, err)
	require.Equal(t, key2.String(), key.String())
}

func TestUnitPrivateKeyFromStringDerECDSA(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	key2, err := PrivateKeyFromStringDer(key.StringDer())
	require.NoError(t, err)
	require.Equal(t, key2.String(), key.String())
}

func TestUnitPrivateKeyECDSAFromBytes(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	key2, err := PrivateKeyFromBytes(key.Bytes())
	require.NoError(t, err)
	require.Equal(t, key2.String(), key.String())
}

func TestUnitPrivateKeyEd25519FromBytes(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	key2, err := PrivateKeyFromBytes(key.Bytes())
	require.NoError(t, err)
	require.Equal(t, key2.String(), key.String())
}

func TestUnitPrivateKeyFromBytesECDSA(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	key2, err := PrivateKeyFromBytesECDSA(key.Bytes())
	require.NoError(t, err)
	require.Equal(t, key2.String(), key.String())
}

func TestUnitPrivateKeyFromBytesEd25519(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	key2, err := PrivateKeyFromBytesEd25519(key.Bytes())
	require.NoError(t, err)
	require.Equal(t, key2.String(), key.String())
}

func TestUnitPublicKeyFromBytesECDSA(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	key2, err := PublicKeyFromBytesECDSA(key.PublicKey().Bytes())
	require.NoError(t, err)
	require.Equal(t, key.PublicKey().String(), key2.String())
}

func TestUnitPublicKeyFromBytesEd25519(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	key2, err := PublicKeyFromBytesEd25519(key.PublicKey().Bytes())
	require.NoError(t, err)
	require.Equal(t, key.PublicKey().String(), key2.String())
}

func TestUnitPublicKeyECDSAFromBytes(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	key2, err := PublicKeyFromBytes(key.PublicKey().Bytes())
	require.NoError(t, err)
	require.Equal(t, key.PublicKey().String(), key2.String())
}

func TestUnitPublicKeyEd25519FromBytes(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	key2, err := PublicKeyFromBytes(key.PublicKey().Bytes())
	require.NoError(t, err)
	require.Equal(t, key.PublicKey().String(), key2.String())
}

func TestUnitPrivateKeyBytesRawEd25519(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	require.Equal(t, key.ed25519PrivateKey.keyData[0:32], key.BytesRaw())
}

func TestUnitPrivateKeyBytesRawECDSA(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	require.Equal(t, key.ecdsaPrivateKey.keyData.D.Bytes(), key.BytesRaw())
}

func TestUnitPublicKeyBytesRawEd25519(t *testing.T) {
	key, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)
	require.Equal(t, key.PublicKey().ed25519PublicKey.keyData, key.PublicKey().BytesRaw())
}

func TestUnitPublicKeyBytesRawECDSA(t *testing.T) {
	key, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)
	require.Equal(t, crypto.CompressPubkey(&key.ecdsaPrivateKey.keyData.PublicKey), key.PublicKey().BytesRaw())
}

func TestUnitPrivateKeyECDSASignTransaction(t *testing.T) {
	newKey, err := PrivateKeyGenerateEcdsa()
	require.NoError(t, err)

	newBalance := NewHbar(2)

	txID := TransactionIDGenerate(AccountID{Account: 123})

	tx, err := NewAccountCreateTransaction().
		SetKey(newKey).
		SetNodeAccountIDs([]AccountID{{Account: 3}}).
		SetTransactionID(txID).
		SetInitialBalance(newBalance).
		SetMaxAutomaticTokenAssociations(100).
		Freeze()
	require.NoError(t, err)

	_, err = newKey.SignTransaction(&tx.Transaction)
	require.NoError(t, err)
}

func TestUnitPublicKeyFromPrivateKeyString(t *testing.T) {
	key, err := PrivateKeyFromStringECDSA("30540201010420ac318ea8ff8d991ab2f16172b4738e74dc35a56681199cfb1c0cb2e7cb560ffda00706052b8104000aa124032200036843f5cb338bbb4cdb21b0da4ea739d910951d6e8a5f703d313efe31afe788f4")
	require.NoError(t, err)
	require.Equal(t, "3036301006072a8648ce3d020106052b8104000a032200036843f5cb338bbb4cdb21b0da4ea739d910951d6e8a5f703d313efe31afe788f4", key.PublicKey().String())
}

func TestUnitPublicKeyToEthereumAddress(t *testing.T) {
	byt, err := hex.DecodeString("03af80b90d25145da28c583359beb47b21796b2fe1a23c1511e443e7a64dfdb27d")
	require.NoError(t, err)
	key, err := PublicKeyFromBytesECDSA(byt)
	ethereumAddress := key.ToEthereumAddress()
	require.Equal(t, ethereumAddress, "627306090abab3a6e1400e9345bc60c78a8bef57")
}

func TestSlip10Ed25519Vector1(t *testing.T) {
	// https://github.com/satoshilabs/slips/blob/master/slip-0010.md#test-vector-1-for-ed25519
	test1PrivateKey := "2b4be7f19ee27bbf30c667b642d5f4aa69fd169872f8fc3059c08ebae2eb19e7"
	test1PublicKey := "a4b2856bfec510abab89753fac1ac0e1112364e7d250545963f135f2a33188ed"
	test1ChainCode := "90046a93de5380a72b5e45010748567d5ea02bbf6522f979e05c0d8d8ca9fffb"
	test2PrivateKey := "68e0fe46dfb67e368c75379acec591dad19df3cde26e63b93a8e704f1dade7a3"
	test2PublicKey := "8c8a13df77a28f3445213a0f432fde644acaa215fc72dcdf300d5efaa85d350c"
	test2ChainCode := "8b59aa11380b624e81507a27fedda59fea6d0b779a778918a2fd3590e16e9c69"
	test3PrivateKey := "b1d0bad404bf35da785a64ca1ac54b2617211d2777696fbffaf208f746ae84f2"
	test3PublicKey := "1932a5270f335bed617d5b935c80aedb1a35bd9fc1e31acafd5372c30f5c1187"
	test3ChainCode := "a320425f77d1b5c2505a6b1b27382b37368ee640e3557c315416801243552f14"
	test4PrivateKey := "92a5b23c0b8a99e37d07df3fb9966917f5d06e02ddbd909c7e184371463e9fc9"
	test4PublicKey := "ae98736566d30ed0e9d2f4486a64bc95740d89c7db33f52121f8ea8f76ff0fc1"
	test4ChainCode := "2e69929e00b5ab250f49c3fb1c12f252de4fed2c1db88387094a0f8c4c9ccd6c"
	test5PrivateKey := "30d1dc7e5fc04c31219ab25a27ae00b50f6fd66622f6e9c913253d6511d1e662"
	test5PublicKey := "8abae2d66361c879b900d204ad2cc4984fa2aa344dd7ddc46007329ac76c429c"
	test5ChainCode := "8f6d87f93d750e0efccda017d662a1b31a266e4a6f5993b15f5c1f07f74dd5cc"
	test6PrivateKey := "8f94d394a8e8fd6b1bc2f3f49f5c47e385281d5c17e65324b0f62483e37e8793"
	test6PublicKey := "3c24da049451555d51a7014a37337aa4e12d41e485abccfa46b47dfb2af54b7a"
	test6ChainCode := "68789923a0cac2cd5a29172a475fe9e0fb14cd6adb5ad98a3fa70333e7afa230"

	seed, err := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	assert.NoError(t, err)

	//Chain m
	key1, err := PrivateKeyFromSeedEd25519(seed)
	assert.NoError(t, err)

	assert.Equal(t, key1.StringRaw(), test1PrivateKey)
	assert.Equal(t, key1.PublicKey().StringRaw(), test1PublicKey)
	assert.Equal(t, hex.EncodeToString(key1.ed25519PrivateKey.chainCode), test1ChainCode)

	// Chain m/0'
	key2, err := key1.Derive(0)
	assert.NoError(t, err)

	assert.Equal(t, key2.StringRaw(), test2PrivateKey)
	assert.Equal(t, key2.PublicKey().StringRaw(), test2PublicKey)
	assert.Equal(t, hex.EncodeToString(key2.ed25519PrivateKey.chainCode), test2ChainCode)

	// Chain m/0'/1'
	key3, err := key2.Derive(1)
	assert.NoError(t, err)
	assert.Equal(t, key3.StringRaw(), test3PrivateKey)
	assert.Equal(t, key3.PublicKey().StringRaw(), test3PublicKey)
	assert.Equal(t, hex.EncodeToString(key3.ed25519PrivateKey.chainCode), test3ChainCode)

	// Chain m/0'/1'/2'
	key4, err := key3.Derive(2)
	assert.NoError(t, err)
	assert.Equal(t, key4.StringRaw(), test4PrivateKey)
	assert.Equal(t, key4.PublicKey().StringRaw(), test4PublicKey)
	assert.Equal(t, hex.EncodeToString(key4.ed25519PrivateKey.chainCode), test4ChainCode)

	//Chain m/0'/1'/2'/2'
	key5, err := key4.Derive(2)
	assert.NoError(t, err)
	assert.Equal(t, key5.StringRaw(), test5PrivateKey)
	assert.Equal(t, key5.PublicKey().StringRaw(), test5PublicKey)
	assert.Equal(t, hex.EncodeToString(key5.ed25519PrivateKey.chainCode), test5ChainCode)

	// Chain m/0'/1'/2'/2'/1000000000'
	key6, err := key5.Derive(1000000000)
	assert.NoError(t, err)
	assert.Equal(t, key6.StringRaw(), test6PrivateKey)
	assert.Equal(t, key6.PublicKey().StringRaw(), test6PublicKey)
	assert.Equal(t, hex.EncodeToString(key6.ed25519PrivateKey.chainCode), test6ChainCode)
}

func TestSlip10Ed25519Vector2(t *testing.T) {
	// https://github.com/satoshilabs/slips/blob/master/slip-0010.md#test-vector-2-for-ed25519
	test1PrivateKey := "171cb88b1b3c1db25add599712e36245d75bc65a1a5c9e18d76f9f2b1eab4012"
	test1PublicKey := "8fe9693f8fa62a4305a140b9764c5ee01e455963744fe18204b4fb948249308a"
	test1ChainCode := "ef70a74db9c3a5af931b5fe73ed8e1a53464133654fd55e7a66f8570b8e33c3b"
	test2PrivateKey := "1559eb2bbec5790b0c65d8693e4d0875b1747f4970ae8b650486ed7470845635"
	test2PublicKey := "86fab68dcb57aa196c77c5f264f215a112c22a912c10d123b0d03c3c28ef1037"
	test2ChainCode := "0b78a3226f915c082bf118f83618a618ab6dec793752624cbeb622acb562862d"
	test3PrivateKey := "ea4f5bfe8694d8bb74b7b59404632fd5968b774ed545e810de9c32a4fb4192f4"
	test3PublicKey := "5ba3b9ac6e90e83effcd25ac4e58a1365a9e35a3d3ae5eb07b9e4d90bcf7506d"
	test3ChainCode := "138f0b2551bcafeca6ff2aa88ba8ed0ed8de070841f0c4ef0165df8181eaad7f"
	test4PrivateKey := "3757c7577170179c7868353ada796c839135b3d30554bbb74a4b1e4a5a58505c"
	test4PublicKey := "2e66aa57069c86cc18249aecf5cb5a9cebbfd6fadeab056254763874a9352b45"
	test4ChainCode := "73bd9fff1cfbde33a1b846c27085f711c0fe2d66fd32e139d3ebc28e5a4a6b90"
	test5PrivateKey := "5837736c89570de861ebc173b1086da4f505d4adb387c6a1b1342d5e4ac9ec72"
	test5PublicKey := "e33c0f7d81d843c572275f287498e8d408654fdf0d1e065b84e2e6f157aab09b"
	test5ChainCode := "0902fe8a29f9140480a00ef244bd183e8a13288e4412d8389d140aac1794825a"
	test6PrivateKey := "551d333177df541ad876a60ea71f00447931c0a9da16f227c11ea080d7391b8d"
	test6PublicKey := "47150c75db263559a70d5778bf36abbab30fb061ad69f69ece61a72b0cfa4fc0"
	test6ChainCode := "5d70af781f3a37b829f0d060924d5e960bdc02e85423494afc0b1a41bbe196d4"

	seed, err := hex.DecodeString("fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542")
	assert.NoError(t, err)

	// Chain m
	key1, err := PrivateKeyFromSeedEd25519(seed)
	assert.NoError(t, err)

	assert.Equal(t, key1.StringRaw(), test1PrivateKey)
	assert.Equal(t, key1.PublicKey().StringRaw(), test1PublicKey)
	assert.Equal(t, hex.EncodeToString(key1.ed25519PrivateKey.chainCode), test1ChainCode)

	// Chain m/0'
	key2, err := key1.Derive(0)
	assert.NoError(t, err)

	assert.Equal(t, key2.StringRaw(), test2PrivateKey)
	assert.Equal(t, key2.PublicKey().StringRaw(), test2PublicKey)
	assert.Equal(t, hex.EncodeToString(key2.ed25519PrivateKey.chainCode), test2ChainCode)

	// Chain m/0'/2147483647'
	key3, err := key2.Derive(2147483647)
	assert.NoError(t, err)
	assert.Equal(t, key3.StringRaw(), test3PrivateKey)
	assert.Equal(t, key3.PublicKey().StringRaw(), test3PublicKey)
	assert.Equal(t, hex.EncodeToString(key3.ed25519PrivateKey.chainCode), test3ChainCode)

	// Chain m/0'/2147483647'/1'
	key4, err := key3.Derive(1)
	assert.NoError(t, err)
	assert.Equal(t, key4.StringRaw(), test4PrivateKey)
	assert.Equal(t, key4.PublicKey().StringRaw(), test4PublicKey)
	assert.Equal(t, hex.EncodeToString(key4.ed25519PrivateKey.chainCode), test4ChainCode)

	// Chain m/0'/2147483647'/1'/2147483646'
	key5, err := key4.Derive(2147483646)
	assert.NoError(t, err)
	assert.Equal(t, key5.StringRaw(), test5PrivateKey)
	assert.Equal(t, key5.PublicKey().StringRaw(), test5PublicKey)
	assert.Equal(t, hex.EncodeToString(key5.ed25519PrivateKey.chainCode), test5ChainCode)

	// Chain m/0'/2147483647'/1'/2147483646'/2'
	key6, err := key5.Derive(2)
	assert.NoError(t, err)
	assert.Equal(t, key6.StringRaw(), test6PrivateKey)
	assert.Equal(t, key6.PublicKey().StringRaw(), test6PublicKey)
	assert.Equal(t, hex.EncodeToString(key6.ed25519PrivateKey.chainCode), test6ChainCode)
}

func TestSlip10ECDSAVector1(t *testing.T) {
	// https://github.com/satoshilabs/slips/blob/master/slip-0010.md#test-vector-1-for-secp256k1
	test1PrivateKey := "e8f32e723decf4051aefac8e2c93c9c5b214313817cdb01a1494b917c8436b35"
	test1PublicKey := "0339a36013301597daef41fbe593a02cc513d0b55527ec2df1050e2e8ff49c85c2"
	test1ChainCode := "873dff81c02f525623fd1fe5167eac3a55a049de3d314bb42ee227ffed37d508"
	test2PrivateKey := "edb2e14f9ee77d26dd93b4ecede8d16ed408ce149b6cd80b0715a2d911a0afea"
	test2PublicKey := "035a784662a4a20a65bf6aab9ae98a6c068a81c52e4b032c0fb5400c706cfccc56"
	test2ChainCode := "47fdacbd0f1097043b78c63c20c34ef4ed9a111d980047ad16282c7ae6236141"
	test3PrivateKey := "3c6cb8d0f6a264c91ea8b5030fadaa8e538b020f0a387421a12de9319dc93368"
	test3PublicKey := "03501e454bf00751f24b1b489aa925215d66af2234e3891c3b21a52bedb3cd711c"
	test3ChainCode := "2a7857631386ba23dacac34180dd1983734e444fdbf774041578e9b6adb37c19"
	test4PrivateKey := "cbce0d719ecf7431d88e6a89fa1483e02e35092af60c042b1df2ff59fa424dca"
	test4PublicKey := "0357bfe1e341d01c69fe5654309956cbea516822fba8a601743a012a7896ee8dc2"
	test4ChainCode := "04466b9cc8e161e966409ca52986c584f07e9dc81f735db683c3ff6ec7b1503f"
	test5PrivateKey := "0f479245fb19a38a1954c5c7c0ebab2f9bdfd96a17563ef28a6a4b1a2a764ef4"
	test5PublicKey := "02e8445082a72f29b75ca48748a914df60622a609cacfce8ed0e35804560741d29"
	test5ChainCode := "cfb71883f01676f587d023cc53a35bc7f88f724b1f8c2892ac1275ac822a3edd"
	test6PrivateKey := "471b76e389e528d6de6d816857e012c5455051cad6660850e58372a6c3e6e7c8"
	test6PublicKey := "022a471424da5e657499d1ff51cb43c47481a03b1e77f951fe64cec9f5a48f7011"
	test6ChainCode := "c783e67b921d2beb8f6b389cc646d7263b4145701dadd2161548a8b078e65e9e"

	seed, err := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	assert.NoError(t, err)

	// Chain m
	key1, err := PrivateKeyFromSeedECDSAsecp256k1(seed)
	assert.NoError(t, err)

	assert.Equal(t, key1.StringRaw(), test1PrivateKey)
	assert.Equal(t, key1.PublicKey().StringRaw(), test1PublicKey)
	assert.Equal(t, hex.EncodeToString(key1.ecdsaPrivateKey.chainCode), test1ChainCode)

	// Chain m/0'
	key2, err := key1.Derive(ToHardenedIndex(0))
	assert.NoError(t, err)

	assert.Equal(t, key2.StringRaw(), test2PrivateKey)
	assert.Equal(t, key2.PublicKey().StringRaw(), test2PublicKey)
	assert.Equal(t, hex.EncodeToString(key2.ecdsaPrivateKey.chainCode), test2ChainCode)

	// Chain m/0'/1
	key3, err := key2.Derive(1)
	assert.NoError(t, err)
	assert.Equal(t, key3.StringRaw(), test3PrivateKey)
	assert.Equal(t, key3.PublicKey().StringRaw(), test3PublicKey)
	assert.Equal(t, hex.EncodeToString(key3.ecdsaPrivateKey.chainCode), test3ChainCode)

	// Chain m/0'/1/2'
	key4, err := key3.Derive(ToHardenedIndex(2))
	assert.NoError(t, err)
	assert.Equal(t, key4.StringRaw(), test4PrivateKey)
	assert.Equal(t, key4.PublicKey().StringRaw(), test4PublicKey)
	assert.Equal(t, hex.EncodeToString(key4.ecdsaPrivateKey.chainCode), test4ChainCode)

	// Chain m/0'/1/2'/2
	key5, err := key4.Derive(2)
	assert.NoError(t, err)
	assert.Equal(t, key5.StringRaw(), test5PrivateKey)
	assert.Equal(t, key5.PublicKey().StringRaw(), test5PublicKey)
	assert.Equal(t, hex.EncodeToString(key5.ecdsaPrivateKey.chainCode), test5ChainCode)

	// Chain m/0'/1/2'/2/1000000000
	key6, err := key5.Derive(1000000000)
	assert.NoError(t, err)
	assert.Equal(t, key6.StringRaw(), test6PrivateKey)
	assert.Equal(t, key6.PublicKey().StringRaw(), test6PublicKey)
	assert.Equal(t, hex.EncodeToString(key6.ecdsaPrivateKey.chainCode), test6ChainCode)
}

func TestSlip10ECDSAVector2(t *testing.T) {
	// https://github.com/satoshilabs/slips/blob/master/slip-0010.md#test-vector-2-for-secp256k1
	test1PrivateKey := "4b03d6fc340455b363f51020ad3ecca4f0850280cf436c70c727923f6db46c3e"
	test1PublicKey := "03cbcaa9c98c877a26977d00825c956a238e8dddfbd322cce4f74b0b5bd6ace4a7"
	test1ChainCode := "60499f801b896d83179a4374aeb7822aaeaceaa0db1f85ee3e904c4defbd9689"
	test2PrivateKey := "abe74a98f6c7eabee0428f53798f0ab8aa1bd37873999041703c742f15ac7e1e"
	test2PublicKey := "02fc9e5af0ac8d9b3cecfe2a888e2117ba3d089d8585886c9c826b6b22a98d12ea"
	test2ChainCode := "f0909affaa7ee7abe5dd4e100598d4dc53cd709d5a5c2cac40e7412f232f7c9c"
	test3PrivateKey := "877c779ad9687164e9c2f4f0f4ff0340814392330693ce95a58fe18fd52e6e93"
	test3PublicKey := "03c01e7425647bdefa82b12d9bad5e3e6865bee0502694b94ca58b666abc0a5c3b"
	test3ChainCode := "be17a268474a6bb9c61e1d720cf6215e2a88c5406c4aee7b38547f585c9a37d9"
	test4PrivateKey := "704addf544a06e5ee4bea37098463c23613da32020d604506da8c0518e1da4b7"
	test4PublicKey := "03a7d1d856deb74c508e05031f9895dab54626251b3806e16b4bd12e781a7df5b9"
	test4ChainCode := "f366f48f1ea9f2d1d3fe958c95ca84ea18e4c4ddb9366c336c927eb246fb38cb"
	test5PrivateKey := "f1c7c871a54a804afe328b4c83a1c33b8e5ff48f5087273f04efa83b247d6a2d"
	test5PublicKey := "02d2b36900396c9282fa14628566582f206a5dd0bcc8d5e892611806cafb0301f0"
	test5ChainCode := "637807030d55d01f9a0cb3a7839515d796bd07706386a6eddf06cc29a65a0e29"
	test6PrivateKey := "bb7d39bdb83ecf58f2fd82b6d918341cbef428661ef01ab97c28a4842125ac23"
	test6PublicKey := "024d902e1a2fc7a8755ab5b694c575fce742c48d9ff192e63df5193e4c7afe1f9c"
	test6ChainCode := "9452b549be8cea3ecb7a84bec10dcfd94afe4d129ebfd3b3cb58eedf394ed271"

	seed, err := hex.DecodeString("fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542")
	assert.NoError(t, err)

	// Chain m
	key1, err := PrivateKeyFromSeedECDSAsecp256k1(seed)
	assert.NoError(t, err)

	assert.Equal(t, key1.StringRaw(), test1PrivateKey)
	assert.Equal(t, key1.PublicKey().StringRaw(), test1PublicKey)
	assert.Equal(t, hex.EncodeToString(key1.ecdsaPrivateKey.chainCode), test1ChainCode)

	// Chain m/0
	key2, err := key1.Derive(0)
	assert.NoError(t, err)

	assert.Equal(t, key2.StringRaw(), test2PrivateKey)
	assert.Equal(t, key2.PublicKey().StringRaw(), test2PublicKey)
	assert.Equal(t, hex.EncodeToString(key2.ecdsaPrivateKey.chainCode), test2ChainCode)

	// Chain m/0/2147483647'
	key3, err := key2.Derive(ToHardenedIndex(2147483647))
	assert.NoError(t, err)
	assert.Equal(t, key3.StringRaw(), test3PrivateKey)
	assert.Equal(t, key3.PublicKey().StringRaw(), test3PublicKey)
	assert.Equal(t, hex.EncodeToString(key3.ecdsaPrivateKey.chainCode), test3ChainCode)

	// Chain m/0/2147483647'/1
	key4, err := key3.Derive(1)
	assert.NoError(t, err)
	assert.Equal(t, key4.StringRaw(), test4PrivateKey)
	assert.Equal(t, key4.PublicKey().StringRaw(), test4PublicKey)
	assert.Equal(t, hex.EncodeToString(key4.ecdsaPrivateKey.chainCode), test4ChainCode)

	// Chain m/0/2147483647'/1/2147483646'
	key5, err := key4.Derive(ToHardenedIndex(2147483646))
	assert.NoError(t, err)
	assert.Equal(t, key5.StringRaw(), test5PrivateKey)
	assert.Equal(t, key5.PublicKey().StringRaw(), test5PublicKey)
	assert.Equal(t, hex.EncodeToString(key5.ecdsaPrivateKey.chainCode), test5ChainCode)

	// Chain m/0/2147483647'/1/2147483646'/2
	key6, err := key5.Derive(2)
	assert.NoError(t, err)
	assert.Equal(t, key6.StringRaw(), test6PrivateKey)
	assert.Equal(t, key6.PublicKey().StringRaw(), test6PublicKey)
	assert.Equal(t, hex.EncodeToString(key6.ecdsaPrivateKey.chainCode), test6ChainCode)
}
