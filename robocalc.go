package main

import "os"
import "log"
import "fmt"
import "rand"

type Board struct {
	m_values[25] int
	m_columnScores[5] int
	m_rowScores[5] int
	m_score int
	m_worstScoreRow int
	m_worstScoreColumn int
	m_counts[6] int
	m_scoreCounts[6] int
}

func (board *Board) Init() {
	for i := 0; i < 25; i++ {
		board.m_values[i] = 0;
	}
	for i := 0; i < 5; i++ {
		board.m_columnScores[i] = 0;
		board.m_rowScores[i] = 0;
	}
	for i := 0; i < 6; i++ {
		board.m_counts[i] = 0;
	}
	board.m_score = 0;
	board.m_worstScoreRow = 0;
	board.m_worstScoreColumn = 0;
}

func (board Board) String() string {
	var ret string = ""
	index := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			ret += fmt.Sprintf("%3d ", board.m_values[index])
			index++
		}
		ret += fmt.Sprintf("= %3d", board.m_rowScores[y])
		ret += "\n"
	}
	ret += " ||  ||  ||  ||  ||\n"
	for x := 0; x < 5; x++ {
		ret += fmt.Sprintf("%3d ", board.m_columnScores[x])
	}
	ret += "\n"
	ret += fmt.Sprintf("0s = %1d ", board.m_counts[0])
	ret += fmt.Sprintf("1s = %1d ", board.m_counts[1])
	ret += fmt.Sprintf("2s = %1d ", board.m_counts[2])
	ret += fmt.Sprintf("3s = %1d ", board.m_counts[3])
	ret += fmt.Sprintf("4s = %1d ", board.m_counts[4])
	ret += fmt.Sprintf("5s = %1d ", board.m_counts[5])
	ret += "\n"
	ret += fmt.Sprintf("Score = %d", board.m_score)
	ret += fmt.Sprintf(" WorstRow = %d", board.m_worstScoreRow)
	ret += fmt.Sprintf(" WorstColumn = %d", board.m_worstScoreColumn)

	return ret
}

func ComputeScore(values[5] int) (score int) {
	var counts[6] int
	counts[0] = 0
	counts[1] = 0
	counts[2] = 0
	counts[3] = 0
	counts[4] = 0
	counts[5] = 0

	counts[values[0]]++
	counts[values[1]]++
	counts[values[2]]++
	counts[values[3]]++
	counts[values[4]]++

	score = 0
	for i := 0; i < 6; i++ {
		var count int = counts[i]
		switch count {
			case 1:
				score += i
			case 2:
				score += 10*i
			case 3, 4, 5:
				score += 100
		}
	}
	return
}

func (board *Board) Valid() (valid bool) {
	valid = false

/*
	var counts[6] int
	counts[0] = 0
	counts[1] = 0
	counts[2] = 0
	counts[3] = 0
	counts[4] = 0
	counts[5] = 0

	for i := 0; i < 25; i++ {
		var v = board.m_values[24-i]
		counts[v]++
		if counts[v] > 6 {
			return
		}
	}
*/
	board.m_counts[0] = 0
	board.m_counts[1] = 0
	board.m_counts[2] = 0
	board.m_counts[3] = 0
	board.m_counts[4] = 0
	board.m_counts[5] = 0

	for i := 0; i < 25; i++ {
		var v = board.m_values[24-i]
		board.m_counts[v]++
		if board.m_counts[v] > 6 {
			return
		}
	}
	valid = true
	return
}

func (board *Board) GetValue(index int) (value int) {
	value = board.m_values[index]
	return
}

func (board *Board) SetValue(index int, value int) {
	board.m_values[index] = value
	return
}

func (board *Board) GetRow(row int) (values[5] int) {
	index := row * 5
	values[0] = board.m_values[index+0]
	values[1] = board.m_values[index+1]
	values[2] = board.m_values[index+2]
	values[3] = board.m_values[index+3]
	values[4] = board.m_values[index+4]
	return
}

func (board *Board) GetColumn(column int) (values[5] int) {
	index := column
	values[0] = board.m_values[index+0]
	values[1] = board.m_values[index+5]
	values[2] = board.m_values[index+10]
	values[3] = board.m_values[index+15]
	values[4] = board.m_values[index+20]
	return
}

func (board *Board) BetterThan(bestScore int) (boardScore int) {
	worstScore := 1000
	var counts[6] int
	counts[0] = 0
	counts[1] = 0
	counts[2] = 0
	counts[3] = 0
	counts[4] = 0
	counts[5] = 0

	for x := 0; x < 5; x++ {
		score := 0

		index := x
		counts[board.m_values[index+0]]++
		counts[board.m_values[index+5]]++
		counts[board.m_values[index+10]]++
		counts[board.m_values[index+15]]++
		counts[board.m_values[index+20]]++

		for i := 0; i < 6; i++ {
			var count int = counts[i]
			counts[i] = 0
			switch count {
				case 1:
					score += i
				case 2:
					score += 10*i
				case 3, 4, 5:
					score += 100
			}
		}

		board.m_columnScores[x] = score
		if score <= worstScore {
			worstScore = score
			board.m_worstScoreRow = -1
			board.m_worstScoreColumn = x
		}
	}
	for y := 0; y < 5; y++ {
		score := 0

		index := y * 5
		counts[board.m_values[index+0]]++
		counts[board.m_values[index+1]]++
		counts[board.m_values[index+2]]++
		counts[board.m_values[index+3]]++
		counts[board.m_values[index+4]]++
		for i := 0; i < 6; i++ {
			var count int = counts[i]
			counts[i] = 0
			switch count {
				case 1:
					score += i
				case 2:
					score += 10*i
				case 3, 4, 5:
					score += 100
			}
		}

		board.m_rowScores[y] = score
		if score <= worstScore {
			worstScore = score
			board.m_worstScoreRow = y
			board.m_worstScoreColumn = -1
		}
	}
	board.m_score = worstScore
	boardScore = worstScore
	return
}

func (board *Board) ComputeScores() {
	worstScore := 1000

	for y := 0; y < 5; y++ {
		score := ComputeScore( board.GetRow(y) )
		board.m_rowScores[y] = score
		if score < worstScore {
			worstScore = score
			board.m_worstScoreRow = y
			board.m_worstScoreColumn = -1
		}
	}
	for x := 0; x < 5; x++ {
		score := ComputeScore( board.GetColumn(x) )
		board.m_columnScores[x] = score
		if score < worstScore {
			worstScore = score
			board.m_worstScoreRow = -1
			board.m_worstScoreColumn = x
		}
	}
	board.m_score = worstScore
}

func (board *Board) RandomBoard() {
	var deck[6*6] int
	j := 0
	for num := 0; num < 6; num++ {
		for i := 0; i < 6; i++ {
			deck[j] = num
			j++
		}
	}
	deckSize := len(deck)

	index := 0
	for i := 0; i < 25; i++ {
		deckPos := rand.Intn(deckSize)
		card := deck[deckPos]
		board.m_values[index] = card
		deckSize--
		for j:= deckPos; j < deckSize; j++ {
			deck[j] = deck[j+1]
		}
	}
}

func monteCarlo() {
	var bestBoard Board
	for i := 0; ; i++ {
		var board Board
		board.RandomBoard()
		board.ComputeScores()
		if board.m_score > bestBoard.m_score {
			bestBoard = board
		}
		if i > 100000 {
			i = 0
			fmt.Println("------------")
			fmt.Println("Board")
			fmt.Println(board)
			fmt.Println("------------")
			fmt.Println("BestBoard")
			fmt.Println(bestBoard)
			fmt.Println("------------")
		}
	}
	fmt.Println("------------")
	fmt.Println("BestBoard")
	fmt.Println("------------")
	fmt.Println(bestBoard)
	fmt.Println("------------")
}

func findStartingBoard(board *Board) {

	value := 0
	count := 0
	for i := 0; i < 25; i++ {
		board.SetValue(24-i, value)
		count++
		if count == 6 {
			count = 0
			value++
		}
	}
}

func nextBoardInSearch(board *Board) (valid bool) {
	i := 0
	valid = true
	worstRow := board.m_worstScoreRow
	worstColumn := board.m_worstScoreColumn
	changedWorstScore := false
	//changedWorstScore := true
	for doMore:= true; doMore == true; {
		index := 0
		row := 0
		column := 0
		for carry := 1; carry == 1; {
			if column > 4 {
				row++
				column = 0
			}
			carry = 0
			oldValue := board.m_values[index]
			newValue := oldValue + 1

			// stop skipping when we make a change to the worst scoring row/column change
			if worstRow == row {
				changedWorstScore = true
			}
			if worstColumn == column {
				changedWorstScore = true
			}
			if newValue > 5 {
				carry = 1
				newValue = 0
				column++
			}
			board.m_values[index] = newValue
			board.m_counts[oldValue]--
			board.m_counts[newValue]++
			index++
			if index > 24 {
				fmt.Println("---------")
				fmt.Println("Big Index")
				fmt.Println("---------")
				fmt.Println("Try Board")
				fmt.Println(board)
				valid = false
				return
			}
		}
		if changedWorstScore == true {
			doMore = false
			for j:= 0; j < 6; j++ {
				if board.m_counts[j] > 6 {
					doMore = true
					break
				}
			}
/*
			var saveCounts[6] int
			saveCounts = board.m_counts
			doMore2 := (board.Valid() == false)
			if doMore2 != doMore {
				fmt.Println(saveCounts)
				fmt.Println("doMore2 != doMore ", doMore2, doMore)
				fmt.Println(board)
				return false
			}
			board.m_counts = saveCounts
*/
/*
			doMore = false
			board.m_counts[0] = 0
			board.m_counts[1] = 0
			board.m_counts[2] = 0
			board.m_counts[3] = 0
			board.m_counts[4] = 0
			board.m_counts[5] = 0

			for i := 0; i < 25; i++ {
				var v = board.m_values[24-i]
				board.m_counts[v]++
				if board.m_counts[v] > 6 {
					doMore = true
					break
				}
			}
*/
		}
		i++
		showBoardDelay := 100000000
		//showBoardDelay := 1
		if i >= showBoardDelay {
			i = 0
			fmt.Println("------------")
			fmt.Println("worstRow:", worstRow, " worstColumn:", worstColumn, " changedWorstScore:", changedWorstScore)
			fmt.Println("Try Next Board")
			fmt.Println(board)
			fmt.Println("------------")
		}
	}
	return
}

func fullSearch() {
	var bestBoard Board
	var board Board
	for i := 0; i < 25; i++ {
		board.SetValue(i, 0)
	}
	findStartingBoard(&board)
	board.ComputeScores()
	board.Valid()
	fmt.Println(board)

	bestBoard = board
	bestScore := bestBoard.m_score

	for i := 0; ; i++ {
		valid := nextBoardInSearch(&board)
		if valid == false {
			fmt.Println("------------")
			fmt.Println("Finished")
			fmt.Println("------------")
			fmt.Println("Board")
			fmt.Println(board)
			fmt.Println("------------")
			fmt.Println("BestBoard")
			fmt.Println(bestBoard)
			fmt.Println("------------")
			return
		}
		newScore := board.BetterThan(bestScore)
		if (newScore > bestScore) {
			bestBoard = board
			bestScore = newScore
			fmt.Println("------------")
			fmt.Println("New BestBoard")
			fmt.Println(bestBoard)
			fmt.Println("------------")
		}
		showBoardDelay := 100000
		//showBoardDelay := 1
		if i >= showBoardDelay {
			i = 0
			fmt.Println("------------")
			fmt.Println("Board")
			fmt.Println(board)
			fmt.Println("------------")
			fmt.Println("BestBoard")
			fmt.Println(bestBoard)
			fmt.Println("------------")
		}
		if bestScore > 1000 {
			return
		}
	}
}

func main() {
	var error os.Error
	log.SetFlags(log.Ltime|log.Lmicroseconds)

	if error != nil {
		log.Fatalf("%s", error.String())
	}

//	monteCarlo()
	fullSearch()

	if false {
		var board Board
		board.RandomBoard()
		if board.Valid() == false {
			fmt.Println("-----------------")
			fmt.Println("INVALID BOARD")
			fmt.Println(board)
			fmt.Println("-----------------")
			log.Fatalf("INVALID BOARD")
		}
	}
}
