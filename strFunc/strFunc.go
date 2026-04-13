package strFunc

//import "fmt"

func ReverseStr(str string) (reverseStr string) {

	for index := len(str) - 1; index >= 0; index-- {
		reverseStr += string(str[index])
	}

	return

}

func AddElement(str string, position int, newElement string) (newStr string) {

	for index := 0; index <= len(str); index++ {

		if index == position {
			newStr += newElement
		}

		if index <= len(str)-1 {
			newStr += string(str[index])
		}

	}

	return

}

func DeleteElement(str string, position int) (newStr string) {

	for index := 0; index <= len(str) - 1; index++ {
		if index != position {
			newStr += string(str[index])
		}
	}

	return

}

func ReplaceElement(str string, position int, element string) (newStr string) {

	for index := 0; index <= len(str) - 1; index++ {
		if index != position {
			newStr += string(str[index])
		} else {
			newStr += element
		}
	}

	return

}

func CheckSubStringInclude(str string, subString string) (isCheck bool) {

	for a := 0; a < len(str); a++ {
		checkString := string(str[a]);
		for b := a + 1; b < len(str); b++ {
			checkString += string(str[b])
			if checkString == subString {
				return true
			}
		}
	}

	return false

}

