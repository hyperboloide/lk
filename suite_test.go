package lk_test

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/suite"
)

type (
	Suite struct {
		suite.Suite
	}
)

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) RandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	s.Require().NoError(err)
	return b
}

func (s *Suite) RandomB64String(n int) string {
	b := s.RandomBytes(n)
	return base64.RawStdEncoding.EncodeToString(b)
}

func (s *Suite) RandomB32String(n int) string {
	b := s.RandomBytes(n)
	return base32.StdEncoding.EncodeToString(b)
}

func (s *Suite) RandomHexString(n int) string {
	b := s.RandomBytes(n)
	return hex.EncodeToString(b)
}
