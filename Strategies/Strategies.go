package Strategies

func (p *Player) updateYears(opponent Player) {
	if p.Cooperate && opponent.Cooperate {
		p.Years += 2
	} else if !p.Cooperate && !opponent.Cooperate {
		p.Years += 4
	} else if !p.Cooperate {
		p.Years += 0
		opponent.Years += 5
	} else {
		p.Years += 5
		opponent.Years += 0
	}
}

func TitForTat() (int, int) {
	player1 := Player{Cooperate: true, Years: 0}
	player2 := Player{Cooperate: false, Years: 0}

	for i := 0; i < numRounds; i++ {
		player1.updateYears(player2)
		player2.updateYears(player1)

		player1.Cooperate = player2.Cooperate
		player2.Cooperate = player1.Cooperate
	}

	return player1.Years, player2.Years
}

func AlwaysCooperate() (int, int) {
	player1 := Player{Cooperate: true, Years: 0}
	player2 := Player{Cooperate: true, Years: 0}

	for i := 0; i < numRounds; i++ {
		player1.updateYears(player2)
		player2.updateYears(player1)
	}

	return player1.Years, player2.Years
}

func AlwaysDefect() (int, int) {
	player1 := Player{Cooperate: false, Years: 0}
	player2 := Player{Cooperate: false, Years: 0}

	for i := 0; i < numRounds; i++ {
		player1.updateYears(player2)
		player2.updateYears(player1)
	}

	return player1.Years, player2.Years
}

func PeriodicallyCCD() (int, int) {
	player1 := Player{Cooperate: true, Years: 0}
	player2 := Player{Cooperate: true, Years: 0}

	swapPos := 1

	for i := 0; i < numRounds; i++ {
		isThirdRound := i%3 == 0

		if isThirdRound {
			swapPos++
		}

		if isThirdRound && swapPos%2 != 0 {
			player1.Cooperate = false
		} else {
			player1.Cooperate = true
		}

		if isThirdRound && swapPos%2 == 0 {
			player2.Cooperate = false
		} else {
			player2.Cooperate = true
		}

		player1.updateYears(player2)
		player2.updateYears(player1)
	}

	return player1.Years, player2.Years
}

func PeriodicallyDDC() (int, int) {
	player1 := Player{Cooperate: false, Years: 0}
	player2 := Player{Cooperate: false, Years: 0}

	swapPos := 1

	for i := 1; i <= numRounds; i++ {
		isThirdRound := i%3 == 0

		if isThirdRound {
			swapPos++
		}

		if isThirdRound && swapPos%2 != 0 {
			player1.Cooperate = true
		} else {
			player1.Cooperate = false
		}

		if isThirdRound && swapPos%2 == 0 {
			player2.Cooperate = true
		} else {
			player2.Cooperate = false
		}

		player1.updateYears(player2)
		player2.updateYears(player1)
	}

	return player1.Years, player2.Years
}
