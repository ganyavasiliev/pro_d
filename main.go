package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Я загадаю число от 1 до 100")
	fmt.Println("Сможешь yгадать?")
	second := time.Now().Unix()
	rand.Seed(second)
	target := rand.Intn(100) + 1

	fmt.Println(target)
	succes := false
	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("Попыток осталось", 10-guesses)
		fmt.Println("Ведите ваше число")
		input, err := reader.ReadString(('\n'))
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace((input))

		guess, err := strconv.Atoi((input))
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Ваше число меньше загаданного")
		} else if guess > target {
			fmt.Println("Ваше число больше загаданного")
		} else {
			succes = true
			fmt.Println("Вы победили")
			break
		}
	}
	if !succes {
		fmt.Println("Вы проиграли. Загаданное число:", target)
	}
}
