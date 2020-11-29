package main

import (
	Tinder "TinderClient"
	"fmt"
)

func main() {
	//Tinder.LoginSms("90xxxxxxxxx")
	TinderApp := Tinder.NewTinderClient("06ed7c8a-b173-40a2-801e-6421633687b7")
	user, _ := TinderApp.GetUser("5f9d2a87842123010053f7ef")
	fmt.Println(user)
	//data, err := TinderApp.MyLikes()
	//fmt.Println(data)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for i := 0; i < len(data); i++ {
	//	item := data[i]
	//	userId := item.User.ID
	//
	//	go func() {
	//		go TinderApp.LikeUser(userId)
	//		go fmt.Println("Like : "+userId, item.User.Name)
	//	}()
	//	time.Sleep(time.Second)
	//}
	//fmt.Scanln("")

}
