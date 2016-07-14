package lk

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lk Suite")
}
