package numerical

import (
	"fmt"
	"testing"
)

// TestBinPow tests binpow function
func TestBinPow(t *testing.T) {

	fmt.Println(BinPow0520(3, 11)/BinPow0520(3,9))

	// if BinPow(2, 10, 121323) != 1024 {
	// 	t.Error("[Error] BinPow(2, 10) is wrong")
	// }

	// if BinPow(1, 10, 121323) != 1 {
	// 	t.Error("[Error] BinPow(1, 10) is wrong")
	// }

	// if BinPow(0, 123123, 2) != 0 {
	// 	t.Error("[Error] BinPow(0, 123123) is wrong")
	// }
}
