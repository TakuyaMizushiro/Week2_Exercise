package main

import "testing"

func Testdigits01(t *testing.T) {
	str := digits(1)
	if str != "1" {
		t.Error("Test01 is failed")
	}
}

// func Testdigits02(t *testing.T) {
// 	str := digits(4)
// 	if str != "IV" {
// 		t.Error("Test02 is failed")
// 	}
//}
// func Testdigits03(t *testing.T) {
// 	str := digits(9)
// 	if str != "IX" {
// 		t.Error("Test03 is failed")
// 	}
//}

//add test code
