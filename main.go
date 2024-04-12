package main

import (
	"Axelrod_Tournaments/Strategies"
	"fmt"
)

func main() {
	strategies := []struct {
		Name         string
		PlayFunction func() (int, int)
		TotalYearsP1 int
		TotalYearsP2 int
		OverallAvg   float64
	}{
		{"TitForTat", Strategies.TitForTat, 0, 0, 0},
		{"Always Cooperate", Strategies.AlwaysCooperate, 0, 0, 0},
		{"Always Defect", Strategies.AlwaysDefect, 0, 0, 0},
		{"Periodically DDC", Strategies.PeriodicallyDDC, 0, 0, 0},
		{"Periodically CCD", Strategies.PeriodicallyCCD, 0, 0, 0},
	}
	y1, y2 := Strategies.PeriodicallyCCD()
	fmt.Println(y1, y2)
	for i, strategy := range strategies {
		strategies[i].TotalYearsP1, strategies[i].TotalYearsP2 = strategy.PlayFunction()
		strategies[i].OverallAvg = float64(strategies[i].TotalYearsP1+strategies[i].TotalYearsP2) / 2
	}

	fmt.Println("Výsledky Axelrodova turnaje:")
	for _, strategy := range strategies {
		fmt.Printf("Strategie: %s, Roky hráče 1: %d, Roky hráče 2: %d, Celkový průměr: %.2f\n", strategy.Name, strategy.TotalYearsP1, strategy.TotalYearsP2, strategy.OverallAvg)
	}
}
