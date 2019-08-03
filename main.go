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
		sel, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			if sel != 1 && sel != 2 && sel != 3 && sel != 4 {
				fmt.Println("Wrong input!")
			} else {
				return int(sel)
			}
		} else {
			fmt.Println("Wrong input!")
		}
	}
}

func productMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Select Menu : ")
		fmt.Println("1. View Products")
		fmt.Println("2. Add Products")
		fmt.Println("3. Remove Products")
		fmt.Println("4. Back")
		fmt.Printf("Input : ")
		scanner.Scan()
		sel, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			if sel != 1 && sel != 2 && sel != 3 && sel != 4 {
				fmt.Println("Wrong input!")
			} else if sel == 1 {
				if len(products) == 0 {
					fmt.Println("There's no products!")
				} else {
					for _, v := range products {
						fmt.Printf("%d.\t%s\t%d\t%d", v.id, v.name, v.price, v.stock)
					}
				}
			} else if sel == 2 {
				fmt.Printf("Input ID : ")
				scanner.Scan()
				id, _ := strconv.Atoi(scanner.Text())
				fmt.Printf("Input Name : ")
				scanner.Scan()
				name := scanner.Text()
				fmt.Printf("Input Price : ")
				scanner.Scan()
				price, _ := strconv.Atoi(scanner.Text())
				fmt.Printf("Input Stock : ")
				stock, _ := strconv.Atoi(scanner.Text())
				products[name] = Product{id, stock, price, name}
			} else if sel == 3 {
				fmt.Printf("Input Name : ")
				scanner.Scan()
				name := scanner.Text()
				_, ok := products[name]
				if ok {
					delete(products, name)
				} else {
					fmt.Printf("Product not found!")
				}
			} else if sel == 4 {
				return
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
