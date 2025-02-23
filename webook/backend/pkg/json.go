package pkg

import (
	"encoding/json"
	"fmt"
)

func JsonMarshal(v interface{}) string {
	res, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}
