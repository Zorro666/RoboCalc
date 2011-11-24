#include <stdio.h>
#include <stdlib.h>

typedef struct Board
{
	int m_values[5][5];
	int m_columnScores[5];
	int m_rowScores[5];
	int m_score;
} Board;

void printBoard(Board* pBoard)
{
	int x;
	int y;
	for (y = 0; y < 5; y++)
	{
		for (x = 0; x < 5; x++)
		{
			printf("%3d ", pBoard->m_values[x][y]);
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

int ComputeScore(int values[5])
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

void GetRow(Board* pBoard, int row, int values[5])
{
	int x;
	for (x = 0; x < 5; x++)
	{
		values[x] = pBoard->m_values[x][row];
	}
}

void GetColumn(Board* pBoard, int column, int values[5])
{
	int y;
	for (y = 0; y < 5; y++)
	{
		values[y] = pBoard->m_values[column][y];
	}
}

void ComputeScores(Board* pBoard)
{
	int minScore = 1000;
	int x;
	int y;
	for (y = 0; y < 5; y++)
	{
		int score;
		int values[5];

		GetRow(pBoard, y, values);
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

		GetColumn(pBoard, x, values);
		score = ComputeScore(values);
		pBoard->m_columnScores[x] = score;
		if (score < minScore)
		{
			minScore = score;
		}
	}
	pBoard->m_score = minScore;
}

void RandomBoard(Board* pBoard)
{
	int x;
	int y;
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

	for (y = 0; y < 5; y++)
	{
		for (x = 0; x < 5; x++)
		{
			int deckPos = (rand() % deckSize);
			int card = deck[deckPos];
			pBoard->m_values[x][y] = card;
			deckSize--;
			for (i = deckPos; i < deckSize; i++)
			{
				deck[i] = deck[i+1];
			}
		}
	}
}

int main(int argc, char* argv[])
{
	Board bestBoard;
	int i = 0;

	for (i = 0; i < argc; i++)
	{
		printf("argv[%d] '%s'\n", i, argv[i]);
	}

	bestBoard.m_score = 0;
	do
	{
		Board board;

		RandomBoard(&board);
		ComputeScores(&board);
		if (board.m_score > bestBoard.m_score)
		{
			bestBoard = board;
		}
		if (i % 1000 == 0) 
		{
			printf("------------\n");
			printf("Board\n");
			printBoard(&board);
			printf("------------\n");
			printf("BestBoard\n");
			printBoard(&bestBoard);
			printf("------------\n");
		}
		i++;
	}
	while (1);

	printf("------------\n");
	printf("BestBoard\n");
	printf("------------\n");
	/*printf(bestBoard);*/
	printf("------------\n");
}
