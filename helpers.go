package lk

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

func toBytes(obj interface{}) ([]byte, error) {
	var buffBin bytes.Buffer

	encoderBin := gob.NewEncoder(&buffBin)
	if err := encoderBin.Encode(obj); err != nil {
		return nil, err
	}

	return buffBin.Bytes(), nil
}

func toB64String(obj interface{}) (string, error) {
	b, err := toBytes(obj)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func fromBytes(obj interface{}, b []byte) error {
	buffBin := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buffBin)

	return decoder.Decode(obj)
}

func fromB64String(obj interface{}, s string) error {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}

	return fromBytes(obj, b)
}
