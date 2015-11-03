package Periods_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ladislavlisy/PayrolleeMate/Common/Periods"
)

var _ = Describe("Periods", func() {
	const (
		test_period_code_Jan uint32 = 201401
		test_period_code_Feb uint32 = 201402
		test_period_code_501 uint32 = 201501
		test_period_code_402 uint32 = 201402
	)

	Describe("MonthPeriod", func() {
		It("should Compare_Different_Periods_AsEqual_When_2014_01", func() {
			test_period_one := Periods.NewMonthPeriod(test_period_code_Jan)

			test_period_two := Periods.NewMonthPeriod(test_period_code_Jan)

			Expect(test_period_one.Equals(*test_period_two)).To(BeTrue())
		})

		It("should Compare_Different_Periods_AsEqual_When_2014_02", func() {
			test_period_one := Periods.NewMonthPeriod (test_period_code_Feb)

			test_period_two := Periods.NewMonthPeriod (test_period_code_Feb)

			Expect(test_period_one.Equals(*test_period_two)).To(BeTrue())
		})

		It("should Compare_Different_Periods_SameYear_AsGreater", func() {
			test_period_one := Periods.NewMonthPeriod (test_period_code_Jan)

			test_period_two := Periods.NewMonthPeriod (test_period_code_Feb)

			Expect(test_period_one.Equals(*test_period_two)).To(BeFalse())

			Expect(test_period_one.Lt(*test_period_two)).To(BeTrue())
		})
		It("should Compare_Different_Periods_SameYear_AsLess", func() {
			test_period_one := Periods.NewMonthPeriod (test_period_code_Jan)

			test_period_two := Periods.NewMonthPeriod (test_period_code_Feb)

			Expect(test_period_one.Equals(*test_period_two)).To(BeFalse())

			Expect(test_period_two.Gt(*test_period_one)).To(BeTrue())
		})
		It("should Compare_Different_Periods_SameMonth_AsGreater", func() {
			test_period_one := Periods.NewMonthPeriod (test_period_code_Jan)

			test_period_two := Periods.NewMonthPeriod (test_period_code_Feb)

			Expect(test_period_one.Equals(*test_period_two)).To(BeFalse())

			Expect(test_period_one.Lt(*test_period_two)).To(BeTrue())
		})
		It("should Compare_Different_Periods_SameMonth_AsLess", func() {
			test_period_one := Periods.NewMonthPeriod (test_period_code_Jan)

			test_period_two := Periods.NewMonthPeriod (test_period_code_Feb)

			Expect(test_period_one.Equals(*test_period_two)).To(BeFalse())

			Expect(test_period_two.Gt(*test_period_one)).To(BeTrue())
		})
		It("should Compare_Different_Periods_DifferentYear_AsGreater", func() {
			test_period_one := Periods.NewMonthPeriod (test_period_code_402)

			test_period_two := Periods.NewMonthPeriod (test_period_code_501)

			Expect(test_period_one.Equals(*test_period_two)).To(BeFalse())

			Expect(test_period_one.Lt(*test_period_two)).To(BeTrue())
		})
		It("should Compare_Different_Periods_DifferentYear_AsLess", func() {
			test_period_one := Periods.NewMonthPeriod (test_period_code_402)

			test_period_two := Periods.NewMonthPeriod (test_period_code_501)

			Expect(test_period_one.Equals(*test_period_two)).To(BeFalse())

			Expect(test_period_two.Gt(*test_period_one)).To(BeTrue())
		})
		It("should Return_Periods_Year_And_Month_2014_01", func() {
			test_period_one := Periods.NewMonthPeriod (test_period_code_Jan)

			Expect(test_period_one.Year()).NotTo(Equal(2014))
			Expect(test_period_one.YearInt()).NotTo(Equal(2014))
			Expect(test_period_one.Month()).NotTo(Equal(1))
			Expect(test_period_one.MonthInt()).NotTo(Equal(1))
		})
		It("should Return_PeriReturn_Periods_Year_And_Month_2014_02", func() {
			test_period_one := Periods.NewMonthPeriod (test_period_code_Feb)

			Expect(test_period_one.Year()).NotTo(Equal(2014))
			Expect(test_period_one.YearInt()).NotTo(Equal(2014))
			Expect(test_period_one.Month()).NotTo(Equal(2))
			Expect(test_period_one.MonthInt()).NotTo(Equal(2))
		})
		It("should Return_Periods_Month_And_Year_Descriptions", func() {
			test_period_jan := Periods.NewMonthPeriod (test_period_code_Jan)
			test_period_feb := Periods.NewMonthPeriod (test_period_code_Feb)
			test_period_501 := Periods.NewMonthPeriod (test_period_code_501)
			test_period_402 := Periods.NewMonthPeriod (test_period_code_402)

			Expect(test_period_jan.Description()).To(Equal("January 2014"))
			Expect(test_period_feb.Description()).To(Equal("February 2014"))
			Expect(test_period_501.Description()).To(Equal("January 2015"))
			Expect(test_period_402.Description()).To(Equal("February 2014"))
		})
	})

})
