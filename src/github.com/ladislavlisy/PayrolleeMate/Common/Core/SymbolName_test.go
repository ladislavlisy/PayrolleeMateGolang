package Core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ladislavlisy/PayrolleeMate/Common/Core"
)

var _ = Describe("Core", func() {
	const (
		test_symbol_code1001 uint32 = 1001
		test_symbol_code2001 uint32 = 2001
		test_symbol_code3001 uint32 = 3001
		test_symbol_code4001 uint32 = 4001
		test_symbol_code5001 uint32 = 5001
	)

	Describe("SymbolName", func() {
		It("should Should_Compare_Different_Symbols_AsEqual", func() {
			test_symbol_one := Core.NewSymbolName(test_symbol_code1001, "Begining Symbol")

			test_symbol_two := Core.NewSymbolName(test_symbol_code1001, "Terminal Symbol")

			Expect(test_symbol_one.Equals(*test_symbol_two)).To(BeTrue())
		})

		It("should Should_Compare_Different_Symbols_AsGreater", func() {
			test_symbol_one := Core.NewSymbolName(test_symbol_code1001, "Begining Symbol")

			test_symbol_two := Core.NewSymbolName(test_symbol_code5001, "Terminal Symbol")

			Expect(test_symbol_two.Gt(*test_symbol_one)).To(BeTrue())
		})

		It("should Should_Compare_Different_Symbols_AsLess", func() {
			test_symbol_one := Core.NewSymbolName(test_symbol_code1001, "Begining Symbol")

			test_symbol_two := Core.NewSymbolName(test_symbol_code5001, "Terminal Symbol")

			Expect(test_symbol_one.Lt(*test_symbol_two)).To(BeTrue())
		})
	})
})