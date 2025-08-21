package main

// import (
// 	"fmt"
// )

func CamelToSnakeCase(s string) string {
	length := len(s)
	if length < 1 {
		return ""
	}

	for i := 0; i < length; i++ {
		if s[length - 1] >= 'A' && s[length - 1] <= 'Z' {
			return s //if it ends with capital
		}
		if s[i] >= 'A' && s[i] <= 'z' && s[i+1] >= 'A' && s[i+1] <= 'Z'{
			return s //it has two consecutive caps
		}
		if s[i] < 'A' && s[i] > 'Z' || s[i] < 'a' && s[i] > 'a' {
			return s
		}
	}
	result := ""
	for i := 0; i < length; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			if i != 0 {
				result += "_"
			}
			result += string(s[i] + 32)
		}else{
			result += string(s[i])
		}
	}
	return result
}

// func main() {
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))         // Hello_World
// 	fmt.Println(CamelToSnakeCase("helloWorld"))         // hello_World
// 	fmt.Println(CamelToSnakeCase("camelCase"))          // camel_Case
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))   // CAMELtoSnackCASE
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))   // camel_To_Snake_Case
// 	fmt.Println(CamelToSnakeCase("hey2"))               // hey2
// }

