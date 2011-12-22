package main

import "os"
import "log"
import "fmt"
import "rand"

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

func (board *Board) Valid() (valid bool) {
	valid = false
	var counts[6] int
	for i := 0; i < 6; i++ {
		counts[i] = 0
	}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			var v = board.m_values[x][y]
			counts[v]++
			if counts[v] > 6 {
				return
			}
		}
	}
	valid = true
	return
}

func (board *Board) GetValue(index int) (value int) {
	x := index % 5
	y := index / 5
	value = board.m_values[x][y]
	return
}

func (board *Board) SetValue(index int, value int) {
	x := index % 5
	y := index / 5
	board.m_values[x][y] = value
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

func monteCarlo() {
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
	for doMore:= true; doMore == true; {
		index := 0
		for carry := 1; carry == 1; {
			carry = 0
			oldValue := board.GetValue(index)
			newValue := oldValue + 1
			if newValue > 5 {
				carry = 1
				newValue = 0
			}
			board.SetValue(index, newValue)
			index += carry
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
		doMore = (board.Valid() == false)
		i++
		if i % 100000000 == 0 {
			fmt.Println("------------")
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
	fmt.Println(board)

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
		board.ComputeScores()
		if board.m_score > bestBoard.m_score {
			bestBoard = board
			fmt.Println("------------")
			fmt.Println("New BestBoard")
			fmt.Println(bestBoard)
			fmt.Println("------------")
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
}

func main() {
	var error os.Error
	log.SetFlags(log.Ltime|log.Lmicroseconds)

	if error != nil {
		log.Fatalf("%s", error.String())
	}

//	monteCarlo()
	fullSearch()

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
