package main

import "os"
import "log"
import "fmt"
import "rand"

var g_board Board

type Board struct {
	m_values[5][5] int
	m_columnScores[5] int
	m_rowScores[5] int
	m_score int
}

func (board Board) String() string {
	var ret string = ""
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			ret += fmt.Sprintf("%3d ", board.m_values[x][y])
		}
		ret += fmt.Sprintf("= %3d", board.m_rowScores[y])
		ret += "\n"
	}
	ret += " ||  ||  ||  ||  ||\n"
	for x := 0; x < 5; x++ {
		ret += fmt.Sprintf("%3d ", board.m_columnScores[x])
	}
	ret += "\n"
	ret += fmt.Sprintf("Score = %d\n", board.m_score)

	return ret
}

func ComputeScore(values[5] int) (score int) {
	var counts[6] int
	for i := 0; i < 6; i++ {
		counts[i] = 0
	}
	for i := 0; i < 5; i++ {
		var v = values[i]
		counts[v]++
	}
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

func (board *Board) GetRow(row int) (values[5] int) {
	for x := 0; x < 5; x++ {
		values[x] = board.m_values[x][row]
	}
	return
}

func (board *Board) GetColumn(column int) (values[5] int) {
	for y := 0; y < 5; y++ {
		values[y] = board.m_values[column][y]
	}
	return
}

func (board *Board) ComputeScores() {
	minScore := 1000
	for y := 0; y < 5; y++ {
		score := ComputeScore( board.GetRow(y) )
		board.m_rowScores[y] = score
		if score < minScore {
			minScore = score
		}
	}
	for x := 0; x < 5; x++ {
		score := ComputeScore( board.GetColumn(x) )
		board.m_columnScores[x] = score
		if score < minScore {
			minScore = score
		}
	}
	board.m_score = minScore
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

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			deckPos := rand.Intn(deckSize)
			card := deck[deckPos]
			board.m_values[x][y] = card
			deckSize--
			for i:= deckPos; i < deckSize; i++ {
				deck[i] = deck[i+1]
			}
		}
	}
}

func main() {
	var error os.Error
	log.SetFlags(log.Ltime|log.Lmicroseconds)

	if error != nil {
		log.Fatalf("%s", error.String())
	}

	var bestBoard Board
	for i := 0; ; i++ {
		var board Board
		board.RandomBoard()
		board.ComputeScores()
		if board.m_score > bestBoard.m_score {
			bestBoard = board
		}
		if i % 1000 == 0 {
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
