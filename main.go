package main

import "fmt"

func main() {
	var action string
	water := 400
	milk := 540
	coffee := 120
	cups := 9
	money := 550
	run := true
	for run {
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)
		switch action {
		case "buy":
			buy(&water, &milk, &coffee, &cups, &money)
		case "fill":
			fill(&water, &milk, &coffee, &cups)
		case "take":
			fmt.Printf("I gave you $%d\n", money)
			money = 0
		case "remaining":
			info(water, milk, coffee, cups, money)
		case "exit":
			run = false
		}
	}
}

func buy(water *int, milk *int, coffee *int, cups *int, money *int) {
	var choice string
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		if *water < 250 {
			fmt.Println("Sorry, not enough water!")
			return
		}
		if *coffee < 16 {
			fmt.Println("Sorry, not enough coffee beans!")
			return
		}
		*water -= 250
		*coffee -= 16
		*money += 4
	case "2":
		if *water < 350 {
			fmt.Println("Sorry, not enough water!")
			return
		}
		if *milk < 75 {
			fmt.Println("Sorry, not enough coffee beans!")
			return
		}
		if *coffee < 20 {
			fmt.Println("Sorry, not enough coffee beans!")
			return
		}
		*water -= 350
		*milk -= 75
		*coffee -= 20
		*money += 7
	case "3":
		if *water < 200 {
			fmt.Println("Sorry, not enough water!")
			return
		}
		if *milk < 100 {
			fmt.Println("Sorry, not enough coffee beans!")
			return
		}
		if *coffee < 12 {
			fmt.Println("Sorry, not enough coffee beans!")
			return
		}
		*water -= 200
		*milk -= 100
		*coffee -= 12
		*money += 6
	case "back":
		return
	}
	fmt.Println("I have enough resources, making you a coffee!")
	*cups -= 1
}

func info(water int, milk int, coffee int, cups int, money int) {
	fmt.Println("\nThe coffee machine has:")
	fmt.Printf("%d water ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", coffee)
	fmt.Printf("%d disposable cups\n", cups)
	fmt.Printf("$%d of money\n", money)
}

func fill(water *int, milk *int, coffee *int, cups *int) {
	var temp int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&temp)
	*water += temp
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&temp)
	*milk += temp
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&temp)
	*coffee += temp
	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&temp)
	*cups += temp
}
