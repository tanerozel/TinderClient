package main

import (
	Tinder "TinderClient"
	"fmt"
	"time"
)

func main() {
	//Tinder.LoginSms("905380310832")
	TinderApp := Tinder.NewTinderClient("token")

	data, err := TinderApp.RecsCore()

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(data); i++ {
		item := data[i]
		userId := item.User.ID

		go func() {
			go TinderApp.LikeUser(userId)
			go fmt.Println("Like : "+userId, item.User.Name)
		}()
		time.Sleep(time.Second)
	}
	fmt.Scanln("")

}
