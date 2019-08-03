package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var products map[string]Product
var customers map[string]Customer
var carts map[string]Cart

type Address struct {
	street, city, region string
	no                   int
}

type Product struct {
	id, stock, price int
	name             string
}

type Customer struct {
	firstname, lastname, phone, email string
	shippingaddress                   Address
	defaultaddress                    Address
}

type Cart struct {
	customer Customer
	item     Product
	qty      int
}

func calcTotal(price int, qty int) int {
	return price * qty
}

func mainMenu() int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Select Menu : ")
		fmt.Println("1. Product")
		fmt.Println("2. Customers")
		fmt.Println("3. Order")
		fmt.Println("4. Exit")
		fmt.Printf("Input : ")
		scanner.Scan()
		sel, err := strconv.ParseInt(scanner.Text())
		if err != nil {
			if sel != 1 && sel != 2 && sel != 3 && sel != 4 {
				fmt.Println("Wrong input!")
			} else {
				return sel
			}
		} else {
			fmt.Println("Wrong input!")
		}
	}
}

func productMenu(products map[string]Product) map[string]Product {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Select Menu : ")
		fmt.Println("1. View Products")
		fmt.Println("2. Add Products")
		fmt.Println("3. Remove Products")
		fmt.Println("4. Back")
		fmt.Printf("Input : ")
		scanner.Scan()
		sel, err := strconv.ParseInt(scanner.Text())
		if err != nil {
			if sel != 1 && sel != 2 && sel != 3 && sel != 4 {
				fmt.Println("Wrong input!")
			} else if sel == 1 {

			} else if sel == 2 {

			} else if sel == 3 {

			} else if sel == 4 {

			}
		} else {
			fmt.Println("Wrong input!")
		}
	}
}

func main() {

	for {
		fmt.Println("Welcome!")
		i := mainMenu()
		if i == 1 {

		} else if i == 2 {

		} else if i == 3 {

		} else {
			fmt.Println("Thank you!")
			return
		}
	}
}
