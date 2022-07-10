package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	fmt.Println("Enter name:")
	fmt.Scan(&name)

	var address string
	fmt.Println("Enter address:")
	fmt.Scanln(&address)

	contacts := map[string]string{
		"name":    name,
		"address": address}

	fmt.Println(contacts)

	json.Marshal(contacts)
	bytes, _ := json.Marshal(contacts)

	fmt.Println(string(bytes))

}
