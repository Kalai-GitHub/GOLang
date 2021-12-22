//notification.go file contains the definition of the notification of the struct that the builder will create
package main

//this is the finished product that is created by the builder
type Notification struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
	notType  string
}
