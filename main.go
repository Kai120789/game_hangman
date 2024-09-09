package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	fmt.Print("Начать новую игру [y - да/n - exit]? ")
	var enter string
	fmt.Scan(&enter)
	if enter == "y" {
		number := rand.Intn(32793)
		Words, err := os.Open("./russian-mnemonic-words.txt")
		if err != nil {
			log.Fatal(err)
		}

		var lines []string

		scanner := bufio.NewScanner(Words)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		word := lines[number]

		defer Words.Close()
		var word_len int = len([]rune(word))
		var word_quest []string

		for i := 0; i < word_len; i++ {
			word_quest = append(word_quest, "_ ")
		}
		fmt.Println(word_quest)

		hangman(word, word_quest)

		main()
	} else if enter == "n" {
		return
	} else {
		fmt.Println("Введите y или n")
		main()
	}

}
func hangman(word string, word_quest []string) {
	var atempts string
	var draw_elements [7]string = [7]string{"виселица", "голова", "туловище", "левая рука", "правая рука", "левая нога", "правая нога\nТы проиграл"}

	var count int
	for j := 0; j < len([]rune(word))+7; j++ {
		var letter string
		var check bool
		fmt.Scan(&letter)
		for i := 0; i < len([]rune(word)); i++ {
			if []rune(word)[i] == []rune(letter)[0] && valid_letter(letter) && repeat_letter(letter, atempts) {
				word_quest[i] = letter
				check = true
				flag := check_win(word_quest)
				if flag {
					j = len(word) + 7
				}

			}

		}
		if !valid_letter(letter) || !repeat_letter(letter, atempts) {
			fmt.Printf("Введенная буква не из кирилицы или уже была введена\n%s\n", word_quest)
		} else if check {
			fmt.Printf("Отгадана буква %s\n", letter)
			fmt.Println(word_quest)
		} else if !check {
			fmt.Println(word_quest)
			fmt.Printf("Ошибка номер %d, рисуется %s\n%s\n", count+1, draw_elements[count], variants[count])
			count++
			if count == 7 {
				j = len([]rune(word)) + 7
				fmt.Printf("Ответ: %s\n", word)
			}

		}

		atempts += string(letter)
	}

}
func check_win(word_quest []string) bool {
	var flag bool = true
	for i := 0; i < len(word_quest); i++ {
		if string(word_quest[i]) == "_ " {
			flag = false
		}
	}
	if flag {
		fmt.Println("Ты победил, слово:", word_quest)
	}
	return flag
}

func valid_letter(letter string) bool {
	vallid_values := "абвгдежзиклмнопрстуфхцчшщъыьэюя"

	for i := 0; i < len([]rune(vallid_values)); i++ {
		if len([]rune(letter)) == 1 && []rune(letter)[0] == []rune(vallid_values)[i] {
			return true
		}
	}
	return false

}

func repeat_letter(letter string, atempts string) bool {
	for i := 0; i < len([]rune(atempts)); i++ {
		if []rune(letter)[0] == []rune(atempts)[i] {
			return false
		}
	}
	return true
}
