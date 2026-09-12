package intypes

type TStGame struct {
	Game  int
	Wins  []int
	Wrong []int
}

type TStStatNumber struct {
	PremiumWin  int
	OrdinariWin int
	Wrong       int
}

type TMapNembers map[int]TStStatNumber

type TPoint struct {
	Number int
	Score  int
}

type TMapPoints map[int]TPoint
