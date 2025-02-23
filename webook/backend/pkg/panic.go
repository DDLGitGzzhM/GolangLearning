package pkg

import "fmt"

func PanicIf(err error, msg ...string) {
	if err != nil {
		fmt.Println(err, msg)
		panic(err)
	}
}
