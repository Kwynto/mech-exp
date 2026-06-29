package pointdanger

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

var SlStGames []intypes.TStGame

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

func spaceSimbol(k int) string {
	if (k / 10) >= 1 {
		return ""
	}
	return " "
}

func startAnalize(slStInput []intypes.TStGame, iGame int) {
	var slWork, slWorkShort []intypes.TStGame

	slStPrepear := slices.Clone(slStInput)
	slices.Reverse(slStPrepear)

	if iGame < 0 {
		iGame = 0
	}

	iLenSlIn := len(slStPrepear)
	if iLenSlIn < iGame {
		iGame = iLenSlIn
	}

	if iGame == 0 {
		fmt.Println(incolor.StringBlueH("Анализ всех тиражей."))
		slWork = slStPrepear
	} else {
		sMsg1 := fmt.Sprintf("Анализ %d тиражей.", iGame)
		fmt.Println(incolor.StringBlueH(sMsg1))
		slWork = slStPrepear[0:iGame]
	}

	fmt.Println("")

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
	}

	for iNumber, stStatNumber := range mStatNumbers {
		tempScore := stStatNumber.PremiumWin * 3
		// tempScore := stStatNumber.PremiumWin
		tempScore = tempScore + stStatNumber.OrdinariWin
		tempScore = tempScore - (stStatNumber.Wrong * 14)
		mScorePoints[iNumber] = intypes.TPoint{
			Number: iNumber,
			Score:  tempScore,
		}
	}

	// Short - begin
	slWorkShort = slStPrepear[0:3]

	mStatNumbersShort := initMapNumbers()

	for _, stGame := range slWorkShort {
		for i01, iWinNum := range stGame.Wins {
			if i01 < PREMIUM_WIN {
				tempMStatNumber := mStatNumbersShort[iWinNum]
				tempMStatNumber.PremiumWin = tempMStatNumber.PremiumWin + 1
				mStatNumbersShort[iWinNum] = tempMStatNumber
			} else if (i01 >= PREMIUM_WIN) && (i01 < ORDINARI_WIN) {
				tempMStatNumber := mStatNumbersShort[iWinNum]
				tempMStatNumber.OrdinariWin = tempMStatNumber.OrdinariWin + 1
				mStatNumbersShort[iWinNum] = tempMStatNumber
			} else {
				tempMStatNumber := mStatNumbersShort[iWinNum]
				tempMStatNumber.Wrong = tempMStatNumber.Wrong + 1
				mStatNumbersShort[iWinNum] = tempMStatNumber
			}
		}

		for _, iWrongNum := range stGame.Wrong {
			tempMStatNumber := mStatNumbersShort[iWrongNum]
			tempMStatNumber.Wrong = tempMStatNumber.Wrong + 1
			mStatNumbersShort[iWrongNum] = tempMStatNumber
		}
	}

	for iNumber, stStatNumber := range mStatNumbersShort {
		tempScore := mScorePoints[iNumber].Score
		tempScore = tempScore - (stStatNumber.Wrong * 4)
		mScorePoints[iNumber] = intypes.TPoint{
			Number: iNumber,
			Score:  tempScore,
		}
	}
	// Short - end

	fmt.Println(incolor.StringBlue("Все номера:"))
	for i := range NUMBERS_IN_GAME {
		k1 := i + 1
		for k, stScorePoint := range mScorePoints {
			if k == k1 {
				sMsg := ""

				if stScorePoint.Score <= 0 {
					sMsg = fmt.Sprintf("Номер %s%s = %d", spaceSimbol(stScorePoint.Number), incolor.StringRedH("%d", stScorePoint.Number), stScorePoint.Score)
				} else if stScorePoint.Score >= 28 {
					sMsg = fmt.Sprintf("Номер %s%s = %d", spaceSimbol(stScorePoint.Number), incolor.StringGreenH("%d", stScorePoint.Number), stScorePoint.Score)
				} else {
					sMsg = fmt.Sprintf("Номер %s%s = %d", spaceSimbol(stScorePoint.Number), incolor.StringBlackH("%d", stScorePoint.Number), stScorePoint.Score)
				}

				fmt.Println(sMsg)
			}
		}
	}

	fmt.Println("")
	fmt.Println(incolor.StringGreen("Позитивный анализ:"))
	for i := range NUMBERS_IN_GAME {
		k1 := i + 1
		for k, stScorePoint := range mScorePoints {
			if k == k1 {
				if stScorePoint.Score >= 28 {
					sMsg := fmt.Sprintf("Номер %s%s = %d    %s  %s", spaceSimbol(stScorePoint.Number), incolor.StringGreenH("%d", stScorePoint.Number), stScorePoint.Score, incolor.StringBlackH("%s", "проигрыше было"), incolor.StringRedH("%d", mStatNumbers[k].Wrong))
					fmt.Println(sMsg)
				}
			}
		}
	}

	fmt.Println("")
	fmt.Println(incolor.StringRed("Номера зоны риска:"))
	for i := range NUMBERS_IN_GAME {
		k1 := i + 1
		for k, stScorePoint := range mScorePoints {
			if k == k1 {
				if stScorePoint.Score <= 0 {
					sMsg := fmt.Sprintf("Номер %s%s = %d", spaceSimbol(stScorePoint.Number), incolor.StringRedH("%d", stScorePoint.Number), stScorePoint.Score)
					fmt.Println(sMsg)
				}
			}
		}
	}

}

func Start(slStInput []intypes.TStGame) {
	var sGame string

	SlStGames = slStInput

	fmt.Println(incolor.StringBlue("Баловый анализ тенденции результатов:"))

	fmt.Print(incolor.StringMagenta("Кол-во последних тиражей для анализа (0 для всех тиражей) > "))
	fmt.Scanf("%v\n", &sGame)
	iGame, err1 := strconv.Atoi(sGame)
	if err1 != nil {
		fmt.Println("Conversion failed:", err1)
	}

	fmt.Println("")

	startAnalize(slStInput, iGame)
}
