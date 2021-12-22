package main

import "fmt"

type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subTitle string) {
	nb.SubTitle = subTitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetNotType(notType string) {
	nb.NotType = notType
}

//build method returns the fully finished Notification object
func (nb *NotificationBuilder) build() (*Notification, error) {
	//Error checking can be done at the Build stage
	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, fmt.Errorf("Subtitle is required when using the icon")
	}
	if nb.Priority > 5 {
		return nil, fmt.Errorf("priority should not 0 to 5")
	}

	//Return a newly created Notification object using the current setting
	return &Notification{
		title:    nb.Title,
		subtitle: nb.SubTitle,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
