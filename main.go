package main

import (
	"fmt"
	"os"

	"github.com/Kwynto/mech-exp/pkg/incolor"
)

var nIn int

func fileRead(name string) (string, error) {
	bRead, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(bRead), nil
}

func info() {
	fmt.Println(incolor.StringBlue("Введите команду"))
	fmt.Println(incolor.StringMagenta("0"), "- выход")
	fmt.Println(incolor.StringMagenta("1"), "- информация")
	fmt.Println(incolor.StringMagenta("2"), "- показать базу тиражей")
	fmt.Println(incolor.StringMagenta("3"), "- добавить данные последнего тиража")
	fmt.Println(incolor.StringMagenta("4"), "- сделать прогноз на следующий тираж")
}

func showBase() {
	fmt.Println(incolor.StringBlue("Данные тиражей:"))

	data, err := fileRead("data.txt")
	if err != nil {
		return
	}

	fmt.Println(incolor.StringBlueH(data))

}

func enterData() {
	var sNum, sWins, sWrongs string
	fmt.Println(incolor.StringBlue("Ввод данных:"))

	fmt.Print(incolor.StringGreen("Тираж > "))
	fmt.Scanf("%v\n", &sNum)
	fmt.Print(incolor.StringGreen("Выпали > "))
	fmt.Scanf("%v\n", &sWins)
	fmt.Print(incolor.StringGreen("Не выпали > "))
	fmt.Scanf("%v\n", &sWrongs)
	fmt.Println(" ")
	fmt.Println(incolor.StringBlue("Вы ввели"))
	fmt.Println(incolor.StringBlueH("Тираж:"), sNum, incolor.StringGreenH("Выиграли:"), sWins, incolor.StringRedH("Остались:"), sWrongs)
	fmt.Println(" ")
}

func main() {
	info()

	for i := 1; i < 5; i++ {
		i = i - 1

		fmt.Print(incolor.StringGreen("> "))

		_, err := fmt.Scanf("%d\n", &nIn)

		if err != nil {
			info()
			continue
		}

		if nIn == 0 {
			break
		}

		switch nIn {
		case 1:
			info()
		case 2:
			showBase()
		case 3:
			enterData()
		}

	}
}
