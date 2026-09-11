package waveanalysis

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/Kwynto/mech-exp/internal/intypes"
	"github.com/Kwynto/mech-exp/pkg/incolor"
)

const (
	MAX_NUMBER   = 40
	PREMIUM_WIN  = 16
	ORDINARI_WIN = 37

	DEFAULT_BORDER  = 2
	NUMBERS_IN_GAME = 40
	ADD_TO_BORER    = 1
)

// var SlStGames []intypes.TStGame

func initMapNumbers() intypes.TMapNembers {
	initMap := make(intypes.TMapNembers, 40)
	iCount := MAX_NUMBER + 1
	for i := 1; i < iCount; i++ {
		initMap[i] = intypes.TStStatNumber{
			PremiumWin:  0,
			OrdinariWin: 0,
			Wrong:       0,
		}
	}

	return initMap
}

func initMapPoint() intypes.TMapPoints {
	initMap := make(intypes.TMapPoints, 40)
	iCount := MAX_NUMBER + 1

	for i := 1; i < iCount; i++ {
		initMap[i] = intypes.TPoint{
			Number: i,
			Score:  0,
		}
	}

	return initMap
}

// func spaceSimbol(k int) string {
// 	if (k / 10) >= 1 {
// 		return ""
// 	}
// 	return " "
// }

func startAnalize(slStInput []intypes.TStGame, iGame int) {
	var slWork []intypes.TStGame

	slWork = slices.Clone(slStInput)

	mStatNumbers := initMapNumbers()
	mScorePoints := initMapPoint()

	for _, stGame := range slWork {
		for i01, iWinNum := range stGame.Wins {
			if i01 < PREMIUM_WIN {
				tempMStatNumber := mStatNumbers[iWinNum]
				tempMStatNumber.PremiumWin = tempMStatNumber.PremiumWin + 1
				mStatNumbers[iWinNum] = tempMStatNumber
			} else if (i01 >= PREMIUM_WIN) && (i01 < ORDINARI_WIN) {
				tempMStatNumber := mStatNumbers[iWinNum]
				tempMStatNumber.OrdinariWin = tempMStatNumber.OrdinariWin + 1
				mStatNumbers[iWinNum] = tempMStatNumber
			} else {
				tempMStatNumber := mStatNumbers[iWinNum]
				tempMStatNumber.Wrong = tempMStatNumber.Wrong + 1
				mStatNumbers[iWinNum] = tempMStatNumber
			}
		}

		for _, iWrongNum := range stGame.Wrong {
			tempMStatNumber := mStatNumbers[iWrongNum]
			tempMStatNumber.Wrong = tempMStatNumber.Wrong + 1
			mStatNumbers[iWrongNum] = tempMStatNumber
		}

		// статистика за один раунд зафиксирована в mStatNumbers
		// нужно посчитать баллы к этому раунду
		for iNumber, stStatNumber := range mStatNumbers {
			tempScore := stStatNumber.PremiumWin * 3
			tempScore = tempScore + stStatNumber.OrdinariWin
			tempScore = tempScore - (stStatNumber.Wrong * 14)
			mScorePoints[iNumber] = intypes.TPoint{
				Number: iNumber,
				Score:  tempScore,
			}
		}

		// формируем срез отчета по баллам за выбранный раунд
		// if iGame == stGame.Game {
		// 	// fmt.Println(mScorePoints)
		// 	// fmt.Println(" ")
		// 	for i := 1; i < MAX_NUMBER+1; i++ {
		// 		tempNum := mScorePoints[i]
		// 		// fmt.Println(tempNum.Number, ": ", tempNum.Score)
		// 		fmt.Println(tempNum.Score)
		// 	}
		// }

		// формируем срез отчета по балам за текущий раунд
		s1 := fmt.Sprint(stGame.Game)
		for i := 1; i < MAX_NUMBER+1; i++ {
			tempNum := mScorePoints[i]
			s1 = fmt.Sprint(s1, " ", tempNum.Score)
		}
		fmt.Println(s1)

	}

	_ = iGame
}

func Start(slStInput []intypes.TStGame) {
	var sGame string

	fmt.Println(incolor.StringBlue("Координаты для графика волнового анализа:"))

	fmt.Print(incolor.StringMagenta("Номер раунда > "))
	fmt.Scanf("%v\n", &sGame)
	iGame, err1 := strconv.Atoi(sGame)
	if err1 != nil {
		fmt.Println("Conversion failed:", err1)
	}

	fmt.Println("")

	startAnalize(slStInput, iGame)
}
