package puzzles

import (
	"strconv"
	"testing"
)

func isPalindrome(str string) bool {
	maxIdx := len(str) - 1
	for i := 0; i < maxIdx; i++ {
		if i > maxIdx/2 {
			break
		}
		if str[i] != str[maxIdx-i] {
			return false
		}
	}
	return true
}

func findMinPalindrome(max int) int {
	for i := 11; i <= max; i++ {
		two := strconv.FormatInt(int64(i), 2)
		eight := strconv.FormatInt(int64(i), 8)
		ten := strconv.FormatInt(int64(i), 10)
		if isPalindrome(ten) && isPalindrome(two) && isPalindrome(eight) {
			return i
		}
	}
	return 0
}

func TestFindPalindrome(t *testing.T) {
	t.Log(findMinPalindrome(1000))
}
