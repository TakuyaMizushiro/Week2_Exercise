package main

import (
	"testing"
)

func TestDigits01(t *testing.T) {
	str := digits(1)
	if str != "I" {
		t.Error("Test01 is failed")
	}
}
func TestDigits02(t *testing.T) {
	str := digits(2)
	if str != "II" {
		t.Error("Test02 is failed")
	}
}
func TestDigits03(t *testing.T) {
	str := digits(3)
	if str != "III" {
		t.Error("Test03 is failed")
	}
}
func TestDigits04(t *testing.T) {
	str := digits(4)
	if str != "IV" {
		t.Error("Test04 is failed")
	}
}
func TestDigits05(t *testing.T) {
	str := digits(5)
	if str != "V" {
		t.Error("Test05 is failed")
	}
}
func TestDigits06(t *testing.T) {
	str := digits(6)
	if str != "VI" {
		t.Error("Test06 is failed")
	}
}
func TestDigits07(t *testing.T) {
	str := digits(7)
	if str != "VII" {
		t.Error("Test07 is failed")
	}
}
func TestDigits08(t *testing.T) {
	str := digits(8)
	if str != "VIII" {
		t.Error("Test08 is failed")
	}
}
func TestDigits09(t *testing.T) {
	str := digits(9)
	if str != "IX" {
		t.Error("Test09 is failed")
	}
}

// func TestDigits10(t *testing.T) {
// 	str := digits(11)
// 	if str != "11" {
// 		t.Error("Test10 is failed")
// 	}
// }

func TestDigitNumCheck11(t *testing.T) {
	tmp := digitNumCheck(5)
	if tmp != 1 {
		t.Error("Test11 is failed")
	}
}

func TestDigitNumCheck12(t *testing.T) {
	tmp := digitNumCheck(95)
	if tmp != 2 {
		t.Error("Test12 is failed")
	}
}

func TestDigitNumCheck13(t *testing.T) {
	tmp := digitNumCheck(800)
	if tmp != 3 {
		t.Error("Test13 is failed")
	}
}

func TestDigitNumCheck14(t *testing.T) {
	tmp := digitNumCheck(4998)
	if tmp != 4 {
		t.Error("Test14 is failed")
	}
}

func TestDigitNumCheck15(t *testing.T) {
	tmp := digitNumCheck(5000)
	if tmp != -1 {
		t.Error("Test15 is failed")
	}
}

func TestDigitNumCheck16(t *testing.T) {
	tmp := digitNumCheck(0)
	if tmp != -1 {
		t.Error("Test16 is failed")
	}
}

func TestDigitNumCheck17(t *testing.T) {
	tmp := digitNumCheck(-15)
	if tmp != -1 {
		t.Error("Test17 is failed")
	}
}

// func TestDigitNumCheck18(t *testing.T) {
// 	tmp := digits(321)
// 	if tmp != "CCCXXI" {
// 		t.Error("Test18 is failed")
// 	}
// }

func TestDigitNumCheck19(t *testing.T) {
	tmp := digits(5000)
	if tmp != "Error" {
		t.Error("Test19 is failed")
	}
}

func TestDigitNumCheck20(t *testing.T) {
	tmp := digits(0)
	if tmp != "Error" {
		t.Error("Test20 is failed")
	}
}

func TestDigitNumCheck21(t *testing.T) {
	tmp := digits(17)
	if tmp != "XVII" {
		t.Error("Test21 is failed")
	}
}

func TestDigitNumCheck22(t *testing.T) {
	tmp := digits(10)
	if tmp != "X" {
		t.Error("Test22 is failed")
	}
}

func TestDigitNumCheck23(t *testing.T) {
	tmp := digits(4999)
	if tmp != "MMMMCMXCIX" {
		t.Error("Test23 is failed")
	}
}

func TestDigitNumCheck24(t *testing.T) {
	tmp := digits(1000)
	if tmp != "M" {
		t.Error("Test24 is failed")
	}
}

//add test code
