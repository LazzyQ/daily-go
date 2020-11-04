package str

import "fmt"

func Traverse(str string)  {
	s := []rune(str)

	for i := 0; i < len(s); i++{
		fmt.Printf("%T", s[i])
	}
}

