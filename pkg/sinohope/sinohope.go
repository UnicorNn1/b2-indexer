package sinohope

import (
	"bytes"
	"encoding/hex"

	"github.com/b2network/b2-indexer/pkg/crypto"
	"github.com/b2network/b2-indexer/pkg/vsm"
)

func DecodeKey(key string, internalKeyIndex uint, localKey string, vsmIv string) (string, error) {
	localKeyByte, err := hex.DecodeString(localKey)
	if err != nil {
		return "", err
	}
	tassInputData, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	decKey, _, err := vsm.TassSymmKeyOperation(vsm.TaDec, vsm.AlgAes256, tassInputData, []byte(vsmIv), internalKeyIndex)
	if err != nil {
		return "", err
	}
	key = string(bytes.TrimRight(decKey, "\x00"))
	decodeLocalData, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	localEncData, err := crypto.AesDecrypt(decodeLocalData, localKeyByte)
	if err != nil {
		return "", err
	}
	key = string(localEncData)
	return key, nil
}
