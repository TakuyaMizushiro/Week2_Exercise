package main

import (
	"testing"
)

func TestDigits01(t *testing.T) {
	expected := "I"
	result := digits(1)
	if expected != result {
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

//add test code
