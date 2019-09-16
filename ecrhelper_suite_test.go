package ecrhelper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEcrhelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ecrhelper Suite")
}
