package main

import (
	Tinder "TinderClient"
	"fmt"
)

func main() {
	//Tinder.LoginSms("90xxxxxxxxx")
	TinderApp := Tinder.NewTinderClient("token")
	//user, _ := TinderApp.GetUser("user id")
	recUsers, _ := TinderApp.RecsCore()
	var total int = 0
	for i := 0; i < len(recUsers); i++ {
		var photos = recUsers[i].User.Photos
		for _, photo := range photos {
			go Tinder.DownloadImage(photo, "/tinder")
			total++
		}
	}
	fmt.Println(total)
	fmt.Scanln()
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

}
