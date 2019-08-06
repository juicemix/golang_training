package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var products map[int]Product
var customers map[int]Customer
var carts map[int]Cart

type Address struct {
	street, no, city, region string
}

type Product struct {
	id, stock, price int
	name             string
}

type Customer struct {
	id                                int
	firstname, lastname, phone, email string
	shippingaddress                   Address
	defaultaddress                    Address
}

type Cart struct {
	id            int
	customer      Customer
	item          Product
	qty, subtotal int
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
		sel, err := strconv.Atoi(scanner.Text())
		if err == nil {
			if sel != 1 && sel != 2 && sel != 3 && sel != 4 {
				fmt.Println("Wrong number!")
			} else {
				return int(sel)
			}
		} else {
			fmt.Println("Wrong input!")
		}
	}
}

func viewProducts() {
	if len(products) == 0 {
		fmt.Println("There's no products!")
	} else {
		fmt.Printf("|ID\t|NAME\t|PRICE\t|STOCK\t|\n")
		for _, v := range products {
			fmt.Printf("|%d.\t|%s\t|%d\t|%d\t|\n", v.id, v.name, v.price, v.stock)
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
		sel, err := strconv.Atoi(scanner.Text())
		if err == nil {
			if sel != 1 && sel != 2 && sel != 3 && sel != 4 {
				fmt.Println("Wrong input!")
			} else if sel == 1 {
				viewProducts()
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
				scanner.Scan()
				stock, _ := strconv.Atoi(scanner.Text())
				products[id] = Product{id, stock, price, name}
				fmt.Println("Item inserted")
			} else if sel == 3 {
				fmt.Printf("Input ID : ")
				scanner.Scan()
				id, _ := strconv.Atoi(scanner.Text())
				_, ok := products[id]
				if ok {
					delete(products, id)
					fmt.Println("Item deleted")
				} else {
					fmt.Println("Product not found!")
				}
			} else if sel == 4 {
				return
			}
		} else {
			fmt.Println("Wrong input!")
		}
	}
}

func viewCustomers() {
	if len(customers) == 0 {
		fmt.Println("There's no products!")
	} else {
		fmt.Printf("|NO\t|FIRST NAME\t|LAST NAME\t|PHONE\t|EMAIL\t|DEFAULT ADDRESS\t|SHIPPING ADDRESS\t|\n")
		for _, v := range customers {
			fmt.Printf("|%d.\t|%s\t|%s\t|%s\t|%s\t|%s\t|%s\t|\n", v.id, v.firstname, v.lastname, v.phone, v.email, v.defaultaddress.street+" "+v.defaultaddress.no+" "+v.defaultaddress.city+" "+v.defaultaddress.region, v.shippingaddress.street+" "+v.shippingaddress.no+" "+v.shippingaddress.city+" "+v.shippingaddress.region)
		}
	}
}

func customerMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Select Menu : ")
		fmt.Println("1. View Customers")
		fmt.Println("2. Add Customers")
		fmt.Println("3. Remove Customers")
		fmt.Println("4. Back")
		fmt.Printf("Input : ")
		scanner.Scan()
		sel, err := strconv.Atoi(scanner.Text())
		if err == nil {
			if sel != 1 && sel != 2 && sel != 3 && sel != 4 {
				fmt.Println("Wrong input!")
			} else if sel == 1 {
				viewCustomers()
			} else if sel == 2 {
				fmt.Printf("Input ID : ")
				scanner.Scan()
				id, _ := strconv.Atoi(scanner.Text())
				fmt.Printf("Input First Name : ")
				scanner.Scan()
				firstname := scanner.Text()
				fmt.Printf("Input Last Name : ")
				scanner.Scan()
				lastname := scanner.Text()
				fmt.Printf("Input Phone : ")
				scanner.Scan()
				phone := scanner.Text()
				fmt.Printf("Input Email : ")
				scanner.Scan()
				email := scanner.Text()
				fmt.Printf("Input Default Address Street : ")
				scanner.Scan()
				defstreet := scanner.Text()
				fmt.Printf("Input Default Address No : ")
				scanner.Scan()
				defno := scanner.Text()
				fmt.Printf("Input Default Address City : ")
				scanner.Scan()
				defcity := scanner.Text()
				fmt.Printf("Input Default Address Region : ")
				scanner.Scan()
				defregion := scanner.Text()

				var shipstreet, shipno, shipcity, shipregion, same string

				for {
					fmt.Printf("Is your shipping address same as default address? (Y/N) : ")
					scanner.Scan()
					same = scanner.Text()
					if same != "Y" && same != "N" {
						fmt.Println("Wrong input!")
					} else {
						break
					}
				}

				if same == "Y" {
					shipstreet = defstreet
					shipno = defno
					shipregion = defregion
					shipcity = defcity
				} else {
					fmt.Printf("Input Shipping Address Street : ")
					scanner.Scan()
					shipstreet = scanner.Text()
					fmt.Printf("Input Shipping Address No : ")
					scanner.Scan()
					shipno = scanner.Text()
					fmt.Printf("Input Shipping Address City : ")
					scanner.Scan()
					shipcity = scanner.Text()
					fmt.Printf("Input Shipping Address Region : ")
					scanner.Scan()
					shipregion = scanner.Text()
				}
				customers[id] = Customer{id, firstname, lastname, phone, email, Address{defstreet, defno, defcity, defregion}, Address{shipstreet, shipno, shipcity, shipregion}}
				fmt.Println("Success")
			} else if sel == 3 {
				fmt.Printf("Input ID : ")
				scanner.Scan()
				id, _ := strconv.Atoi(scanner.Text())
				_, ok := customers[id]
				if ok {
					delete(customers, id)
					fmt.Println("Customer deleted")
				} else {
					fmt.Println("Customer not found!")
				}
			} else if sel == 4 {
				return
			}
		} else {
			fmt.Println("Wrong input!")
		}
	}
}

func cartMenu() {

start:
	var selected Customer
	userFound := false
	scanner := bufio.NewScanner(os.Stdin)
	viewCustomers()
	for !userFound {

		fmt.Printf("Select Customer (type \"-1\" to go back) : ")
		scanner.Scan()
		user, _ := strconv.Atoi(scanner.Text())

		if user == -1 {
			return
		}

		selected, userFound = customers[user]

		if !userFound {
			fmt.Println("Customer not found")
		}
	}

	for {
		fmt.Printf("Selected Customer : %s %s\n", selected.firstname, selected.lastname)
		fmt.Println("Select Menu : ")
		fmt.Println("1. View Cart")
		fmt.Println("2. Add Cart")
		fmt.Println("3. Remove Cart")
		fmt.Println("4. Change Customer")
		fmt.Println("5. Back")
		fmt.Printf("Input : ")
		scanner.Scan()
		sel, err := strconv.Atoi(scanner.Text())
		if err == nil {
			if sel != 1 && sel != 2 && sel != 3 && sel != 4 && sel != 5 {
				fmt.Println("Wrong input!")
			} else if sel == 1 {
				fmt.Println("|NO\t|PRODUCT\t|PRICE\t|QTY\t|SUBTOTAL\t|")
				for i, v := range carts {
					if v.customer.id == selected.id {
						fmt.Printf("|%d\t|%s\t|%d\t|%d\t|%d\t|\n", i, v.item.name, v.item.price, v.qty, v.subtotal)
					}
				}
			} else if sel == 2 {
				var selectedProduct Product
				productFound := false

				viewProducts()

				for !productFound {
					fmt.Printf("Input ID (type \"-1\" to cancel) : ")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())
					if id == -1 {
						break
					}
					fmt.Printf("Select Product : ")
					scanner.Scan()
					product, _ := strconv.Atoi(scanner.Text())
					fmt.Printf("Input quantity : ")
					scanner.Scan()
					qty, _ := strconv.Atoi(scanner.Text())

					selectedProduct, productFound = products[product]

					if !productFound {
						fmt.Println("Product not found")
					} else {
						if qty > selectedProduct.stock {
							fmt.Println("Insufficent stock")
						} else {
							carts[id] = Cart{id, selected, selectedProduct, qty, calcTotal(selectedProduct.price, qty)}
							selectedProduct.stock = selectedProduct.stock - qty
							products[selectedProduct.id] = selectedProduct
							fmt.Println("Success")
						}
					}
				}
			} else if sel == 3 {
				fmt.Printf("Input ID : ")
				scanner.Scan()
				id, _ := strconv.Atoi(scanner.Text())
				_, ok := carts[id]

				if ok {
					delete(carts, id)
					fmt.Printf("Success")
				} else {
					fmt.Printf("ID not found ")
				}
			} else if sel == 4 {
				goto start
			} else if sel == 5 {
				return
			}
		} else {
			fmt.Println("Wrong input!")
		}
	}
}

func main() {
	products = make(map[int]Product)
	customers = make(map[int]Customer)
	carts = make(map[int]Cart)
	for {
		fmt.Println("Welcome!")
		i := mainMenu()
		if i == 1 {
			productMenu()
		} else if i == 2 {
			customerMenu()
		} else if i == 3 {
			cartMenu()
		} else {
			fmt.Println("Thank you!")
			return
		}
	}
}
