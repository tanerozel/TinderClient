package main

import (
	Tinder "TinderClient"
	"fmt"
)

func main() {
	TinderApp := Tinder.NewTinderClient("9a8a541d-cb08-4a14-8151-25e36d3253ca")
	TinderApp.SetHeader("aa", "aa")
	data, err := TinderApp.RecsCore()
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
