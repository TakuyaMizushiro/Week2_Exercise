package main

//import "strconv"
//import "fmt"

func digits(i int) string {

	var digitNum int
	var j int

	digitNum = digitNumCheck(i)

	var startNum int
	var tmpD int      //一時的に商を保存
	var result string //結果の出力の為の文字列

	if digitNum == 4 {
		startNum = 1000
	} else if digitNum == 3 {
		startNum = 100
	} else if digitNum == 2 {
		startNum = 10
	} else if digitNum == 1 {
		startNum = 1
	} else {
		return "Error"
	}

	for j = digitNum; j > 0; j-- {
		tmpD = i / startNum
		//tmpA = i & startNum

		switch {

		case j == 1 && tmpD == 1:
			result += "I"
		case j == 1 && tmpD == 2:
			result += "II"
		case j == 1 && tmpD == 3:
			result += "III"
		case j == 1 && tmpD == 4:
			result += "IV"
		case j == 1 && tmpD == 5:
			result += "V"
		case j == 1 && tmpD == 6:
			result += "VI"
		case j == 1 && tmpD == 7:
			result += "VII"
		case j == 1 && tmpD == 8:
			result += "VIII"
		case j == 1 && tmpD == 9:
			result += "IX"

		case j == 2 && tmpD == 1:
			result += "X"
		case j == 2 && tmpD == 2:
			result += "XX"
		case j == 2 && tmpD == 3:
			result += "XXX"
		case j == 2 && tmpD == 4:
			result += "XL"
		case j == 2 && tmpD == 5:
			result += "L"
		case j == 2 && tmpD == 6:
			result += "LX"
		case j == 2 && tmpD == 7:
			result += "LXX"
		case j == 2 && tmpD == 8:
			result += "LXXX"
		case j == 2 && tmpD == 9:
			result += "XC"

		case j == 3 && tmpD == 1:
			result += "C"
		case j == 3 && tmpD == 2:
			result += "CC"
		case j == 3 && tmpD == 3:
			result += "CCC"
		case j == 3 && tmpD == 4:
			result += "CD"
		case j == 3 && tmpD == 5:
			result += "D"
		case j == 3 && tmpD == 6:
			result += "DC"
		case j == 3 && tmpD == 7:
			result += "DCC"
		case j == 3 && tmpD == 8:
			result += "DCCC"
		case j == 3 && tmpD == 9:
			result += "CM"

		case j == 4 && tmpD == 1:
			result += "M"
		case j == 4 && tmpD == 2:
			result += "MM"
		case j == 4 && tmpD == 3:
			result += "MMM"
		case j == 4 && tmpD == 4:
			result += "MMMM"

		}

		i -= startNum * tmpD
		if i == 0 {
			break
		}
		startNum /= 10
	}

	return result
}

func digitNumCheck(i int) int { // 入力された数字の桁数を返す（”1”なら1を，85なら”2”を，1520なら”4”，0以下，5000以上の数字は-1を返す）

	var j int
	k := 10

	if i <= 0 || i > 4999 {
		return -1
	}

	for j = 1; j < 5; j++ {
		if i < k {
			return j
		}
		k *= 10
	}

	return -1

}

func main() {

	// var result string

	// result = digits(47)
	// fmt.Println(result)

}

//add test code
