#include <stdio.h>
#include <stdlib.h>

typedef struct Board
{
	int m_values[25];
	int m_columnScores[5];
	int m_rowScores[5];
	int m_score;
} Board;

void boardInit(Board* const pBoard)
{
	int i;
	for (i = 0; i < 25; i++)
	{
		pBoard->m_values[i] = 0;
	}
	for (i = 0; i < 5; i++)
	{
		pBoard->m_columnScores[i] = 0;
		pBoard->m_rowScores[i] = 0;
	}
	pBoard->m_score = 0;
}

void boardPrint(const Board* const pBoard)
{
	int x;
	int y;
	int i = 0;
	for (y = 0; y < 5; y++)
	{
		for (x = 0; x < 5; x++)
		{
			printf("%3d ", pBoard->m_values[i]);
			i++;
		}
		printf("= %3d", pBoard->m_rowScores[y]);
		printf("\n");
	}
	printf(" ||  ||  ||  ||  ||\n");
	for (x = 0; x < 5; x++)
	{
		printf("%3d ", pBoard->m_columnScores[x]);
	}
	printf("\n");
	printf("Score = %d\n", pBoard->m_score);
}

int ComputeScore(const int values[5])
{
	int counts[6];
	int i;
	int score;

	for (i = 0; i < 6; i++)
	{
		counts[i] = 0;
	}

	for (i = 0; i < 5; i++)
	{
		int v = values[i];
		counts[v]++;
	}

	score = 0;

	for (i = 0; i < 6; i++)
	{
		int count = counts[i];
		if (count == 1)
		{
			score += i;
		}
		else if (count == 2)
		{
			score += 10*i;
		}
		else if (count >= 3)
		{
			score += 100;
		}
	}
	return score;
}

int boardValid(const Board* const pBoard)
{
	int counts[6];
	int i;
	for (i = 0; i < 6; i++)
	{
		counts[i] = 0;
	}
	for (i = 0; i < 25; i++)
	{
		const int v = pBoard->m_values[i];
		counts[v]++;
		if (counts[v] > 6)
		{
			return 0;
		}
	}
	return 1;
}

int boardGetValue(const Board* const pBoard, const int index)
{
	const int value = pBoard->m_values[index];
	return value;
}

void boardSetValue(Board* const pBoard, const int index, const int value) 
{
	pBoard->m_values[index] = value;
}

void boardGetRow(const Board* const pBoard, const int row, int values[5])
{
	int x;
	int index = row * 5;
	for (x = 0; x < 5; x++)
	{
		values[x] = pBoard->m_values[index];
		index++;
	}
}

void boardGetColumn(const Board* const pBoard, const int column, int values[5])
{
	int y;
	int index = column;
	for (y = 0; y < 5; y++)
	{
		values[y] = pBoard->m_values[index];
		index += 5;
	}
}

int boardBetterThan(Board* const pBoard, const int bestScore)
{
	int minScore = 1000;
	int i;
	int x;
	int y;
	int counts[6];
	for (i = 0; i < 6; i++)
	{
		counts[i] = 0;
	}

	for (y = 0; y < 5; y++)
	{
		int score = 0;
		int index = y * 5;

		/* ComputeScore */
		for (x = 0; x < 5; x++)
		{
			const int v = pBoard->m_values[index];
			index++;
			counts[v]++;
		}

		for (i = 0; i < 6; i++)
		{
			const int count = counts[i];
			if (count == 1)
			{
				score += i;
			}
			else if (count == 2)
			{
				score += 10*i;
			}
			else if (count >= 3)
			{
				score += 100;
			}
			counts[i] = 0;
		}
		pBoard->m_rowScores[y] = score;
		if (score < minScore)
		{
			minScore = score;
		}
		if (score < bestScore)
		{
			pBoard->m_score = minScore;
			return 0;
		}
	}
	for (x = 0; x < 5; x++)
	{
		int score = 0;
		int index = x;
		/* ComputeScore */
		for (y = 0; y < 5; y++)
		{
			int v = pBoard->m_values[index];
			index += 5;
			counts[v]++;
		}

		for (i = 0; i < 6; i++)
		{
			int count = counts[i];
			if (count == 1)
			{
				score += i;
			}
			else if (count == 2)
			{
				score += 10*i;
			}
			else if (count >= 3)
			{
				score += 100;
			}
			counts[i] = 0;
		}
		pBoard->m_columnScores[x] = score;
		if (score < minScore)
		{
			minScore = score;
		}
		if (score < bestScore)
		{
			pBoard->m_score = minScore;
			return 0;
		}
	}
	pBoard->m_score = minScore;
	return minScore;
}

void boardComputeScores(Board* const pBoard)
{
	int minScore = 1000;
	int x;
	int y;
	for (y = 0; y < 5; y++)
	{
		int score;
		int values[5];

		boardGetRow(pBoard, y, values);
		score = ComputeScore(values);
		pBoard->m_rowScores[y] = score;
		if (score < minScore)
		{
			minScore = score;
		}
	}
	for (x = 0; x < 5; x++)
	{
		int score;
		int values[5];

		boardGetColumn(pBoard, x, values);
		score = ComputeScore(values);
		pBoard->m_columnScores[x] = score;
		if (score < minScore)
		{
			minScore = score;
		}
	}
	pBoard->m_score = minScore;
}

void boardRandomBoard(Board* const pBoard)
{
	int i;
	int j;
	int num;
	int deck[6*6];
	int deckSize = 36;

	j = 0;
	for (num = 0; num < 6; num++)
	{
		for (i = 0; i < 6; i++)
		{
			deck[j] = num;
			j++;
		}
	}

	for (i = 0; i < 25; i++)
	{
		int deckPos = (rand() % deckSize);
		int card = deck[deckPos];
		pBoard->m_values[i] = card;
		deckSize--;
		for (j = deckPos; j < deckSize; j++)
		{
			deck[j] = deck[j+1];
		}
	}
}

void monteCarlo(void)
{
	int i = 0;
	Board bestBoard;
	boardInit(&bestBoard);
	do
	{
		Board board;
		boardInit(&board);

		boardRandomBoard(&board);
		boardComputeScores(&board);
		if (board.m_score > bestBoard.m_score)
		{
			bestBoard = board;
		}
		if (i > 100000) 
		{
			printf("------------\n");
			printf("Board\n");
			boardPrint(&board);
			printf("------------\n");
			printf("BestBoard\n");
			boardPrint(&bestBoard);
			printf("------------\n");
			i = 0;
		}
		i++;
	}
	while (1);

	printf("------------\n");
	printf("BestBoard\n");
	printf("------------\n");
	boardPrint(&bestBoard);
	printf("------------\n");
}

void findStartingBoard(Board* const pBoard)
{
	int value = 0;
	int count = 0;
	int i;
	for (i = 0; i < 25; i++)
	{
		boardSetValue(pBoard, 24-i, value);
		count++;
		if (count == 6)
		{
			count = 0;
			value++;
		}
	}
}

int nextBoardInSearch(Board* const pBoard)
{
	int i = 0;
	int doMore = 1;
	while (doMore == 1)
	{
		int index = 0;
		int carry = 1;
		while (carry == 1)
		{
			const int oldValue = pBoard->m_values[index];
			int newValue = oldValue + 1;
			carry = 0;
			if (newValue > 5)
			{
				carry = 1;
				newValue = 0;
			}
			pBoard->m_values[index] = newValue;
			index += carry;
			if (index > 24)
			{
				printf("------------\n");
				printf("Big Index\n");
				printf("------------\n");
				printf("Try Board\n");
				boardPrint(pBoard);
				printf("------------\n");
				return 0;
			}
		}
		doMore = (boardValid(pBoard) == 0) ? 1 : 0;
		i++;
		if (i > 100000000)
		{
			printf("------------\n");
			printf("Try Next Board\n");
			printf("------------\n");
			boardPrint(pBoard);
			printf("------------\n");
			i = 0;
		}
	}
	return 1;
}

void fullSearch(void)
{
	Board bestBoard;
	Board board;
	int i;
	int bestScore = 0;

	boardInit(&board);
	for (i = 0; i < 25; i++)
	{
		boardSetValue(&board, i, 0);
	}
	findStartingBoard(&board);
	boardPrint(&board);

	bestBoard = board;
	bestScore = bestBoard.m_score;

	for (i = 0; ; i++)
	{
		int valid = nextBoardInSearch(&board);
		int newScore = 0;
		if (valid == 0)
		{
			printf("------------\n");
			printf("Finished\n");
			printf("------------\n");
			printf("Board\n");
			boardPrint(&board);
			printf("------------\n");
			printf("BestBoard\n");
			boardPrint(&bestBoard);
			printf("------------\n");
			return;
		}
		newScore = boardBetterThan(&board, bestScore);
		if (newScore > bestScore)
		{
			bestBoard = board;
			bestScore = newScore;
			printf("------------\n");
			printf("New BestBoard\n");
			boardPrint(&bestBoard);
			printf("------------\n");
		}
		if (i > 100000)
		{
			printf("------------\n");
			printf("Board\n");
			boardPrint(&board);
			printf("------------\n");
			printf("BestBoard\n");
			boardPrint(&bestBoard);
			printf("------------\n");
			i = 0;
		}
	}
}

int main(int argc, char* argv[])
{
	int i = 0;

	for (i = 0; i < argc; i++)
	{
		printf("argv[%d] '%s'\n", i, argv[i]);
	}

/*	monteCarlo();*/
	fullSearch();

	{
		Board board;
		boardInit(&board);
		boardRandomBoard(&board);
		if (boardValid(&board) == 0)
		{
			printf("------------\n");
			printf("INVALID BOARD\n");
			boardPrint(&board);
			printf("------------\n");
		}
	}
	return -1;
}
