package main

import "fmt"

//constant declaration
const LoginToken string ="krishna" //Public

func main(){
	var username string = "Hitesh"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n",username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n",isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type: %T \n",smallVal)

	var smallFloat float32 = 255.1234567876543
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type: %T \n",smallFloat)

	//default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variables is of type: %T \n",anothervariable)

	//implicit type
	var website="learncodeonline.in"
	fmt.Println(website)
	
	//no var style
	varname := 30000
	fmt.Println(varname)
	//valorus operator should be declared inside a main function

	//Printing Constant
	fmt.Println(LoginToken)
	fmt.Printf("Variables is of type:%T \n",LoginToken)
}