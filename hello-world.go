package main

import "fmt" // formatted I/O

func smain() {

	// 1. Hello World !!
	fmt.Println(" 1. Hello World---------------------")
	fmt.Println("Hello World!")
	fmt.Println("------------------------------------")

	// 2. Variable
	fmt.Println(" 2. Variable -----------------------")
	// Variable declaration
	// frmt : var variableName dataType
	var framework string
	var programmngLanguage string = "GoLang"
	framework = "Nothing"
	fmt.Println("Framework : ", framework)
	fmt.Println("Language : ", programmngLanguage)
	// Multiple variable declaration
	// frmt : var variableName1, variableName2, variableName3
	var name string
	var firstName, lastName string = "Andika", "Nugraha"
	name = firstName + lastName
	fmt.Println("Name : ", name)
	// Varibale without defining dataType
	// frmt : var variableName = value
	var nationality = "Indonesia"
	fmt.Println("Nationality : ", nationality)

	// variable without var. but use := to assign value
	// frmt : variableName := value
	address := "Bogor"
	fmt.Println("Address : ", address)

	// Facts about variable
	// - := only can use in function
	// - var is used to define for global variable
	// - _ is spesific type of variable. Any value assigned to it like _:= 12 will be ignored
	// - if a variable declared in a program is not used will throw an error

	// constants variable is variable but it's value can'tbe changed at the runtime
	// frmt : const variableName = value
	const appName = "Let's GO"
	fmt.Println("------------------------------------")

	// 3 Data Types

	// 3.1 Boolean
	// default is false
	// frmt : var variableName bool
	var isILiveInJakarta bool = false
	isMyNameAndika := true
	fmt.Println("Live in Jakarta : ", isILiveInJakarta)
	fmt.Println("Name is Andika : ", isMyNameAndika)

	// 3.2 Numerical Types
	// int(integer) –2147483648 to 2147483647 & uint(un-signed integer) 0 to 4294967295 note : depends upon the bit of operating system
	var firtsNumber int = 20
	var secondNumber = 30
	thirdNumber := 40
	additionOfthreeNumberAbove := firtsNumber + secondNumber + thirdNumber
	fmt.Println("Additional : ", additionOfthreeNumberAbove)
	// 3.3 String & Error Types
	// String can be use "" or ``
	fullName := firstName + " " + lastName
	fmt.Println("My fullname is: ", fullName)

	// 3.4 Amazing declaration of Variable
	const (
		constantName = "Andika Nugraha"
		phi          = 3.1415
		lang         = "Go"
	)

	var (
		variableName1 int
		variableName2 float32
		variableName3 string
	)
	variableName1 = 1
	variableName2 = 1.0
	variableName3 = "Amazing variable"
	fmt.Println(variableName1, variableName2, variableName3)
	// 3.5 Array
	// frmt var arrayName = [size] datatype
	array1 := [3]int{1, 2, 3}
	array2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array1)
	fmt.Println(array2)
	// twoDimensional := [n][m]{{},{}}

	// 3.6 Map
	// Map are key value pair
	// each value have a unique key
	// frmt var mapName map[dataType1] dataType2
	// dataType1 : the data type for key
	// dataType2 : the data type for value
	var userData = make(map[int]string)
	userData[0] = "Andika"
	userData[1] = "Nugraha"

	// 4 Control statement
	// 4.1 if statement
	if additionOfthreeNumberAbove < 0 {
		fmt.Println(`the result is negative`)
	} else if additionOfthreeNumberAbove == 0 {
		fmt.Println(`the result is equals to Zero`)
	} else {
		fmt.Println(`the result is posotive`)
	}

	// 4.2 GOTO statement
	i := 10
LOOP:
	for i < 20 {
		if i == 15 {
			i++
			goto LOOP
		}
		fmt.Println(`value : `, i)
		i++
	}

	// 4.3 for loop
	sum := 0
	for index := 0; index < 5; index++ {
		sum += index
	}
	fmt.Println("Sum: ", sum)

	// 4.4 switch
	exmaple := 3

	switch exmaple {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Exclude")
	}

	// 5. Function
	// in new file.
}
