package telemathings_analog_parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/telemathings/telemathings-analog-parser/parser"
)

var _ = Describe("Parser", func() {
	It("simple", func() {
		groups, err := parser.ParseString("2-5")
		Expect(err).To(BeNil())
		Expect(groups[0]).To(BeEquivalentTo(parser.Group{
			Name: "0",
			Left: 2,
			Divider: "-",
			Right: 5,
		}))
	})

	It("simple named", func() {
		groups, err := parser.ParseString("an1=2-5")
		Expect(err).To(BeNil())
		Expect(groups[0]).To(BeEquivalentTo(parser.Group{
			Name: "an1",
			Left: 2,
			Divider: "-",
			Right: 5,
		}))
	})

	It("simple group", func() {
		groups, err := parser.ParseString("2-5;5-9")
		Expect(err).To(BeNil())
		Expect(groups[0]).To(BeEquivalentTo(parser.Group{
			Name: "0",
			Left: 2,
			Divider: "-",
			Right: 5,
		}))
		Expect(groups[1]).To(BeEquivalentTo(parser.Group{
			Name: "1",
			Left: 5,
			Divider: "-",
			Right: 9,
		}))
	})

	It("simple named group", func() {
		groups, err := parser.ParseString("an1=2-5;an2=5-9")
		Expect(err).To(BeNil())
		Expect(groups[0]).To(BeEquivalentTo(parser.Group{
			Name: "an1",
			Left: 2,
			Divider: "-",
			Right: 5,
		}))
		Expect(groups[1]).To(BeEquivalentTo(parser.Group{
			Name: "an2",
			Left: 5,
			Divider: "-",
			Right: 9,
		}))
	})

	It("simple named/unnamed group", func() {
		groups, err := parser.ParseString("an1=2-5;5-9;an2=9-24")
		Expect(err).To(BeNil())
		Expect(groups[0]).To(BeEquivalentTo(parser.Group{
			Name: "an1",
			Left: 2,
			Divider: "-",
			Right: 5,
		}))
		Expect(groups[1]).To(BeEquivalentTo(parser.Group{
			Name: "1",
			Left: 5,
			Divider: "-",
			Right: 9,
		}))
		Expect(groups[2]).To(BeEquivalentTo(parser.Group{
			Name: "an2",
			Left: 9,
			Divider: "-",
			Right: 24,
		}))
	})
})
