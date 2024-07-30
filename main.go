package main

import "fmt"

var GameOver bool
var game [9]string = [9]string{"_", "_", "_", "_", "_", "_", "_", "_", "_"}
var counter int = 1

func main() {
	for GameOver != true {
		if counter > 9 {
			fmt.Println("Ничья")
			GameOver = true
			break
		}

		displayGame(game)

		id := identificateUserID(counter)
		counter++
		fieldNum := listenUser()
		executeUser(fieldNum, id)
		id = calculateResult()

		switch id {
		case 1:
			fmt.Println("Игрок 1 одержал победу")
			GameOver = true
		case 2:
			fmt.Println("Игрок 2 одержал победу")
			GameOver = true
		}
	}
}

func displayGame(game [9]string) {
	for i := 0; i < 9; i += 3 {
		fmt.Printf("%s | %s | %s\n", game[i], game[i+1], game[i+2])
	}
}

func listenUser() int {
	fmt.Println("Введите номер поля")
	number := 0
	fmt.Scan(&number)

	if number > 9 || number < 1 {
		fmt.Println("Поля не существует")
		number = listenUser()
	}

	return number
}

func executeUser(fieldNum, id int) {
	if game[fieldNum-1] != "_" {
		fmt.Printf("Поле занято: %d\n", fieldNum)
		fieldNum = listenUser()
		executeUser(fieldNum, id)
	}

	switch id {
	case 2:
		game[fieldNum-1] = "X"
	case 1:
		game[fieldNum-1] = "O"
	}
}

func identificateUserID(counter int) int {
	if counter%2 == 0 {
		return 1
	}
	return 2
}

func calculateResult() int {
	b := [9]int{}
	for i := 0; i < len(game); i++ {
		if game[i] == "X" {
			b[i] = 10
		} else if game[i] == "O" {
			b[i] = 1
		}
	}

	resultCount := [9]int{}
	resultCount[0] = b[2] + b[4] + b[6]
	resultCount[1] = b[0] + b[3] + b[6]
	resultCount[2] = b[1] + b[4] + b[7]
	resultCount[3] = b[2] + b[5] + b[8]
	resultCount[4] = b[0] + b[4] + b[8]
	resultCount[5] = b[6] + b[7] + b[8]
	resultCount[6] = b[3] + b[4] + b[5]
	resultCount[7] = b[0] + b[1] + b[2]

	for i := 0; i < len(game); i++ {
		if resultCount[i] == 30 {
			return 1
		} else if resultCount[i] == 3 {
			return 2
		}

	}
	return 0
}
