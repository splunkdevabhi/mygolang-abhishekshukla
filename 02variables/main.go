package main

import "fmt"

const LoginToken string = "abhishek" //Public value

func main(){
	var username string = "Abhishek"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)

	//default values and some aliases
	var testvar int
	fmt.Println(testvar)
	fmt.Printf("Variable is of type: %T \n",testvar)

	//implicit
    var website = "https://github.com"
	fmt.Println(website)

	//no var style
	numberofUser := 300000
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
}