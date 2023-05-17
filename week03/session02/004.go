package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	f := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	err := json.Unmarshal([]byte(`{"name":"John", "age": 20}`), &f)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(f.Name, f.Age)
}
