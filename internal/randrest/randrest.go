package randrest

// random restrictions

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"

	"github.com/Kwynto/mech-exp/pkg/incolor"
)

func startGen(slIBadNums []int, iTiket int) {
	iLen := len(slIBadNums)
	iTiketIn := iTiket + 1

	for i := 1; i < iTiketIn; i++ {
		var slIResNum []int
		for {
			iPos := rand.Intn(iLen)
			iNum := slIBadNums[iPos]
			slIResNum = append(slIResNum, iNum)
			slices.Sort(slIResNum)
			slIResNum = slices.Compact(slIResNum)
			if len(slIResNum) >= 5 {
				s := fmt.Sprint(i, " =")
				fmt.Println(incolor.StringGreen(s), slIResNum)
				break
			}
		}
	}
}

func Start() {
	var sBadNums, sTiket string
	var slIBadNums []int

	fmt.Println(incolor.StringBlue("Формирование случайных номеров из негативного набора для ограничений по отдельным билетам."))
	fmt.Println("")

	fmt.Print(incolor.StringMagenta("Кол-во билетов > "))
	fmt.Scanf("%v\n", &sTiket)
	iTiket, err := strconv.Atoi(sTiket)
	if err != nil {
		fmt.Println("Conversion failed:", err)
	}

	fmt.Print(incolor.StringMagenta("Все рисковые номера (через дефис) > "))
	fmt.Scanf("%v\n", &sBadNums)

	slSBadNums := strings.Split(sBadNums, "-")

	for _, sNum := range slSBadNums {
		iNum, err := strconv.Atoi(sNum)
		if err != nil {
			fmt.Println("Conversion failed:", err)
		}
		slIBadNums = append(slIBadNums, iNum)
	}

	slices.Sort(slIBadNums)
	slIBadNums = slices.Compact(slIBadNums)
	fmt.Println(incolor.StringGreen("Сортированные данные:"), slIBadNums)
	fmt.Println("")

	startGen(slIBadNums, iTiket)
}
