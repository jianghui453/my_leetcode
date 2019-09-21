package bfs

import "testing"

func TestSolve(t *testing.T) {
	var board, h [][]byte

	board = [][]byte{
		{byte('X'), byte('X'), byte('X'), byte('X')},
		{byte('X'), byte('O'), byte('O'), byte('X')},
		{byte('X'), byte('X'), byte('O'), byte('X')},
		{byte('X'), byte('O'), byte('X'), byte('X')},
	}
	h = [][]byte{
		{byte('X'), byte('X'), byte('X'), byte('X')},
		{byte('X'), byte('X'), byte('X'), byte('X')},
		{byte('X'), byte('X'), byte('X'), byte('X')},
		{byte('X'), byte('O'), byte('X'), byte('X')},
	}
	solve(board)
	t.Logf("\nboard=%v \n    h=%v\n", board, h)

	board = [][]byte{
		{byte('X'), byte('X'), byte('X')},
		{byte('X'), byte('O'), byte('X')},
		{byte('X'), byte('X'), byte('X')},
	}
	h = [][]byte{
		{byte('X'), byte('X'), byte('X')},
		{byte('X'), byte('X'), byte('X')},
		{byte('X'), byte('X'), byte('X')},
	}
	solve(board)
	t.Logf("\nboard=%v \n    h=%v\n", board, h)

	board = [][]byte{
		{byte('X'), byte('X'), byte('X')},
		{byte('X'), byte('O'), byte('X')},
		{byte('X'), byte('O'), byte('X')},
	}
	h = [][]byte{
		{byte('X'), byte('X'), byte('X')},
		{byte('X'), byte('O'), byte('X')},
		{byte('X'), byte('O'), byte('X')},
	}
	solve(board)
	t.Logf("\nboard=%v \n    h=%v\n", board, h)

	//[["O","X","X","O","X"],["X","O","O","X","O"],["X","O","X","O","X"],["O","X","O","O","O"],["X","X","O","X","O"]]
	//[["O","X","X","O","X"],["X","X","X","X","O"],["X","X","X","O","X"],["O","X","O","O","O"],["X","X","O","X","O"]]
	board = [][]byte{
		{byte('O'), byte('X'), byte('X'), byte('O'), byte('X')},
		{byte('X'), byte('O'), byte('O'), byte('X'), byte('O')},
		{byte('X'), byte('O'), byte('X'), byte('O'), byte('X')},
		{byte('O'), byte('X'), byte('O'), byte('O'), byte('O')},
		{byte('X'), byte('X'), byte('O'), byte('X'), byte('O')},
	}
	h = [][]byte{
		{byte('O'), byte('X'), byte('X'), byte('O'), byte('X')},
		{byte('X'), byte('X'), byte('X'), byte('X'), byte('O')},
		{byte('X'), byte('X'), byte('X'), byte('O'), byte('X')},
		{byte('O'), byte('X'), byte('O'), byte('O'), byte('O')},
		{byte('X'), byte('X'), byte('O'), byte('X'), byte('O')},
	}
	solve(board)
	t.Logf("\nboard=%v \n    h=%v\n", board, h)
}
