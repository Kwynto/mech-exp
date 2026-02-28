package defective

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/Kwynto/mech-exp/internal/intypes"
	"github.com/Kwynto/mech-exp/pkg/incolor"
)

const (
	MAX_NUMBER     = 40
	PREMIUM_WIN    = 16
	ORDINARI_WIN   = 17
	DEFAULT_BORDER = 2
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

func startAnalize(slStInput []intypes.TStGame, iGame, iBorder int) {
	var slWork []intypes.TStGame

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

	// for _, v := range slWork {
	// 	fmt.Println(v.Game)
	// }

	mStatNumbers := initMapNumbers()

	for _, stGame := range slWork {
		for _, iWrongNum := range stGame.Wrong {
			tempMStatNumber := mStatNumbers[iWrongNum]
			tempMStatNumber.Wrong = tempMStatNumber.Wrong + 1
			mStatNumbers[iWrongNum] = tempMStatNumber
		}
	}

	fmt.Println(incolor.StringRedH("Номера зоны риска:"))
	for k, v := range mStatNumbers {
		if v.Wrong >= iBorder {
			sMsg := fmt.Sprintf("Номер %s wrong = %d", incolor.StringRed("%d", k), v.Wrong)
			fmt.Println(sMsg)
		}
	}

}

func Start(slStInput []intypes.TStGame) {
	var sGame, sBorder string

	SlStGames = slStInput

	fmt.Println(incolor.StringBlue("Дефектный анализ:"))

	fmt.Print(incolor.StringMagenta("Кол-во последних тиражей для анализа (0 для всех тиражей) > "))
	fmt.Scanf("%v\n", &sGame)
	iGame, err1 := strconv.Atoi(sGame)
	if err1 != nil {
		fmt.Println("Conversion failed:", err1)
	}

	fmt.Print(incolor.StringMagenta("Граница повторений (%d или более) > ", DEFAULT_BORDER))
	fmt.Scanf("%v\n", &sBorder)
	iBorder, err2 := strconv.Atoi(sBorder)
	if err2 != nil {
		fmt.Println("Conversion failed:", err2)
	}

	if iBorder < DEFAULT_BORDER {
		iBorder = DEFAULT_BORDER
	}

	fmt.Println("")

	startAnalize(slStInput, iGame, iBorder)
}
