package main

import "fmt"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	// initialize variables
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// string length
	var (
		name    string = "güliz"
		surname string = "ay"
	)

	fmt.Println(len(name))    // result: 6 (g: 1, ü: 2, l: 1, i: 1, z: 1 byte)
	fmt.Println(len(surname)) // result: 2

	// casting
	accountAge := 2.6

	// create a new "accountAgeInt" here
	// it should be the result of casting "accountAge" to an integer
	accountAgeInt := int8(accountAge)
	fmt.Println("Your account has existed for", accountAgeInt, "years")

	// formatting
	s := fmt.Sprintf("I am %.2f years old\n", 10.523)
	fmt.Printf(s)

	// conditions
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	if messageLen < maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	// iota
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	/* results:
	0 0.000000 false ""
	6
	2
	Your account has existed for 2 years
	I am 10.52 years old
	Trying to send a message of length: 10 and a max length of: 20
	Message sent
	0 1 2 3 4 5 6
	*/

}
