package strFunc

func ReverseStr(str string) (reverseStr string) {

	for index := len(str) - 1; index >= 0; index-- {
		reverseStr += string(str[index])
	} 

	return reverseStr

}
