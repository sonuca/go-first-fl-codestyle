package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

func attack(character, class string) string {
	if class == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", character, 5+randint(3, 5))
	}
	if class == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", character, 5+randint(5, 10))
	}
	if class == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", character, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func defence(character, class string) string {
	if class == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", character, 10+randint(5, 10))
	} else if class == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", character, 10+randint(-2, 2))
	} else if class == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", character, 10+randint(2, 5))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func special(character, class string) string {
	if class == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", character, 80+25)
	} else if class == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", character, 5+40)
	} else if class == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", character, 10+30)
	}
	return "неизвестный класс персонажа"
}

// здесь обратите внимание на имена параметров
func startTraining(character, class string) string {
	if class == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", character)
	}
	if class == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", character)
	}
	if class == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", character)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var command string
	for command != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &command)
		if command == "attack" {
			fmt.Println(attack(character, class))
		}
		if command == "defence" {
			fmt.Println(defence(character, class))
		}
		if command == "special" {
			fmt.Println(special(character, class))
		}
	}
	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func selectionСlass() string {
	var approveChoice string
	var class string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &class)
		if class == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if class == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if class == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return class
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var character string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &character)
	fmt.Printf("Здравствуй, %s\n", character)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	class := selectionСlass()
	fmt.Println(startTraining(character, class))
}
