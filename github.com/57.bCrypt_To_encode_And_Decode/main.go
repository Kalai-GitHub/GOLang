package main

import (
	"fmt"
)

func main() {
	pwd := `Password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Passwrod", pwd)
	fmt.Println(bs)

	fmt.Println("--------------------------------------")

	loginPwd := `Password123`

	//Incorrect passwrod
	//loginPwd := `1212#$$#$&@` // will show the error message

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd))

	if err != nil {
		fmt.Println("You Can't Login")
		return
	}

	fmt.Println("You are logged in")
}
