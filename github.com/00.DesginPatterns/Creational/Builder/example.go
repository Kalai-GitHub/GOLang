package main

import (
	"fmt"
)

//Create a NotificationBuilder and use it to set properties
var buldr = newNotificationBuilder()

func main() {

	//use the builder to set some properties
	buldr.SetTitle("New Notification")
	buldr.SetSubTitle("New subtitle")
	buldr.SetImage("image.jpg")
	buldr.SetIcon("icon.png")
	buldr.SetPriority(3)
	buldr.SetMessage("This is the basic notification with some text in it")
	buldr.SetNotType("alert")

	//use the build function to create a finished object
	notif, error := buldr.build()
	if error != nil {
		fmt.Println("Error creating the notification", error)
	} else {
		fmt.Printf("The notification object %+v", notif)
	}

}
