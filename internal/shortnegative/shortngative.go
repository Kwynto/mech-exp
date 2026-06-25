package shortnegative

import (
	"fmt"
	"slices"

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

func spaceSimbol(k int) string {
	if (k / 10) >= 1 {
		return ""
	}
	return " "
}

func startAnalize(slStInput []intypes.TStGame, iGame int) {
	var slWork []intypes.TStGame

	slStPrepear := slices.Clone(slStInput)
	slices.Reverse(slStPrepear)

	iLenSlIn := len(slStPrepear)
	if iLenSlIn < iGame {
		iGame = iLenSlIn
	}

	slWork = slStPrepear[0:iGame]

	mStatNumbers := initMapNumbers()

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

	fmt.Println(incolor.StringRedH("Номера зоны риска:"))
	for i := range NUMBERS_IN_GAME {
		k1 := i + 1
		for k, v := range mStatNumbers {
			if (k == k1) && (v.Wrong > 0) {
				sMsg := fmt.Sprintf("Номер %s%s wrong = %d", spaceSimbol(k), incolor.StringRed("%d", k), v.Wrong)
				fmt.Println(sMsg)
			}
		}
	}

}

func Start(slStInput []intypes.TStGame) {
	SlStGames = slStInput

	fmt.Println(incolor.StringBlue("Анализ короткого негативного повторения:"))
	fmt.Println("")

	iGame := 3

	startAnalize(slStInput, iGame)
}
