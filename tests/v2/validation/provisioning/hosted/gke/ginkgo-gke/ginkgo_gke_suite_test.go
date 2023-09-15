package ginkgo_gke_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoGke(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoGke Suite")
}
