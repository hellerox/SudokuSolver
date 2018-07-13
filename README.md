# Go Sudoku Solver

Sudoku solver, so far tested with some boards found online and it has solved all of them.

My next goal is to add some concurrency using goroutines and channels.

## Add Board

Add board directly to main, assigning it to variable **bi** like:

```
bi := boardInt{
	{0, 0, 0, 6, 0, 0, 4, 0, 0},
	{7, 0, 0, 0, 0, 3, 6, 0, 0},
	{0, 0, 0, 0, 9, 1, 0, 8, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 5, 0, 1, 8, 0, 0, 0, 3},
	{0, 0, 0, 3, 0, 6, 0, 4, 5},
	{0, 4, 0, 2, 0, 0, 0, 6, 0},
	{9, 0, 3, 0, 0, 0, 0, 0, 0},
	{0, 2, 0, 0, 0, 0, 1, 0, 0},
}
```

## Execute

Like any go code.

```
go run main.go
```

## TEST

**still missing**

```
go test
```

## Used Puzzles

I got some of these boards from these links:

- [Example Puzzles and Solutions](http://elmo.sbs.arizona.edu/sandiway/sudoku/examples.html)
- [Websudoku](https://www.websudoku.com/)
