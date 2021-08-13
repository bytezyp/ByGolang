package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	var data = []byte(`{"status":200}`)
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result);err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Printf("%t",result)
}
