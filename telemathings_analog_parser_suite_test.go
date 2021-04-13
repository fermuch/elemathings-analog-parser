package telemathings_analog_parser_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTelemathingsAnalogParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TelemathingsAnalogParser Suite")
}
