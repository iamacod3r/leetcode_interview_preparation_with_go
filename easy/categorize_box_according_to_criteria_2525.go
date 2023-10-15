package easy

import "interview_go/common"

type CategorizeBoxAccordingToCriteria_2525 struct {
}

// https://leetcode.com/problems/categorize-box-according-to-criteria/
func (c *CategorizeBoxAccordingToCriteria_2525) categorizeBox(length int, width int, height int, mass int) string {
	volume := length * width * height
	isBulky := length >= 10_000 || width >= 10_000 || height >= 10_000 || volume >= 1_000_000_000
	isHeavy := mass >= 100
	if isBulky != isHeavy {
		if isBulky {
			return "Bulky"
		}
		return "Heavy"
	}
	if isBulky {
		return "Both"
	}
	return "Neither"

}

func (c *CategorizeBoxAccordingToCriteria_2525) Test() {
	length, width, height, mass := 1000, 35, 700, 300
	expected := "Heavy"
	result := c.categorizeBox(length, width, height, mass)
	common.PrintStr(result, expected)

	length, width, height, mass = 200, 50, 800, 50
	expected = "Neither"
	result = c.categorizeBox(length, width, height, mass)
	common.PrintStr(result, expected)
}
