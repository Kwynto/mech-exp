package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Kwynto/mech-exp/internal/intypes"
	"github.com/Kwynto/mech-exp/pkg/incolor"
)

const FILE_NAME = "data.txt"

var nIn int
var SlStGames []intypes.TStGame

func fileRead(sName string) (string, error) {
	bRead, err := os.ReadFile(sName)
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
	fmt.Println(incolor.StringMagenta("4"), "- сделать дефектный анализ")
}

func convInsideFormat(sData string) []intypes.TStGame {
	var SlStGamesTemp []intypes.TStGame

	slList := strings.Split(sData, "\n")

	for _, sLine := range slList {
		slGame := strings.Split(sLine, "|")

		if len(slGame) == 1 {
			break
		}

		slSWins := strings.Split(slGame[1], "-")
		slSWrongs := strings.Split(slGame[2], "-")

		var slWinNums []int
		var slWrongNums []int

		for _, sWin := range slSWins {
			iNum, err := strconv.Atoi(sWin)
			if err != nil {
				fmt.Println("Conversion failed:", err)
			}
			slWinNums = append(slWinNums, iNum)
		}

		for _, sWrong := range slSWrongs {
			iNum, err := strconv.Atoi(sWrong)
			if err != nil {
				fmt.Println("Conversion failed:", err)
			}
			slWrongNums = append(slWrongNums, iNum)
		}

		iGame, err := strconv.Atoi(slGame[0])
		if err != nil {
			fmt.Println("Conversion failed:", err)
		}

		stGameTemp := intypes.TStGame{
			Game:  iGame,
			Wins:  slWinNums,
			Wrong: slWrongNums,
		}

		SlStGamesTemp = append(SlStGamesTemp, stGameTemp)
	}

	return SlStGamesTemp
}

func readBase(sName string) {
	sData, err := fileRead(sName)
	if err != nil {
		return
	}

	clear(SlStGames)
	SlStGames = convInsideFormat(sData)
}

func showInsideFormat(SlStGamesInput []intypes.TStGame) {
	for _, stGame := range SlStGamesInput {
		fmt.Println(incolor.StringMagenta("Тираж:   "), incolor.StringMagentaH(strconv.Itoa(stGame.Game)))

		stOutWins := fmt.Sprint(stGame.Wins)
		fmt.Println(incolor.StringGreen("Выиграли:"), incolor.StringGreenH(stOutWins))

		stOutWwongs := fmt.Sprint(stGame.Wrong)
		fmt.Println(incolor.StringRed("Остались:"), incolor.StringRedH(stOutWwongs))
		fmt.Println(incolor.StringYellow("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -"))
	}
}

func showBase() {
	fmt.Println(incolor.StringBlue("Данные тиражей:"))

	readBase(FILE_NAME)

	showInsideFormat(SlStGames)
	fmt.Println("")
}

func enterData() {
	defer readBase(FILE_NAME)

	var sNum, sWins, sWrongs string
	fmt.Println(incolor.StringBlue("Ввод данных:"))

	fmt.Print(incolor.StringMagenta("Тираж > "))
	fmt.Scanf("%v\n", &sNum)
	fmt.Print(incolor.StringGreen("Выпали > "))
	fmt.Scanf("%v\n", &sWins)
	fmt.Print(incolor.StringRed("Не выпали > "))
	fmt.Scanf("%v\n", &sWrongs)
	fmt.Println(" ")
	fmt.Println(incolor.StringBlueH("Вы ввели"))
	fmt.Println(incolor.StringMagentaH("Тираж:"), sNum, incolor.StringGreenH("Выиграли:"), sWins, incolor.StringRedH("Остались:"), sWrongs)

	sRecord := fmt.Sprintf("%s|%s|%s\n", sNum, sWins, sWrongs)
	fFileName, err := os.OpenFile(FILE_NAME, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer fFileName.Close()

	if _, err := fFileName.WriteString(sRecord); err != nil {
		return
	}

	fmt.Println(incolor.StringGreen("Данные записаны."))
	fmt.Println(" ")
}

func main() {
	readBase(FILE_NAME)

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
