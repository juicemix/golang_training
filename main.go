package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input first name : ")
	firstName, _ := reader.ReadString('\n')
	fmt.Print("Input last name : ")
	lastName, _ := reader.ReadString('\n')
	fmt.Print("Input delivery distance : ")
	dist, _ := reader.ReadString('\n')

	firstName = strings.TrimSuffix(firstName, "\n")

	lastName = strings.TrimSuffix(lastName, "\n")

	fmt.Println("Hello, " + firstName + " " + lastName)

	distInt, _ := strconv.ParseInt(dist, 10, 8)

	if distInt < 2 {
		fmt.Println("Shipping : Rp 5000")
	} else if distInt < 5 {
		fmt.Println("Shipping : Rp 10000")
	} else {
		fmt.Println("Shipping : Rp 15000")
	}

	reader.ReadString('\n')

}
