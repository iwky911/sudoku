## How efficiently can we solve a sudoku ?

This repo contains to sudoku solver:

* a bruteforce solver, simple but not efficient
* an exact cover solver using [dancing links][dl]

[dl]: http://en.wikipedia.org/wiki/Dancing_Links

### Building

```
$ mkdir -p src/github.com/iwky911
$ export GOPATH=$(pwd)
$ cd src/github.com/iwky911
$ git clone https://github.com/iwky911/sudoku.git
$ go build dancinglinks/main.go
```

### Usage
You can pass the sudoku to be solved in a csv format.
Empty or non-integer cells will be considered blank.

For example, using the provided example:
```
$ ./dancinglinks -csv=testdata/largeinput.csv
Parsed a matrix of size 16
Sparse matrix created
Affecting, (0, 0) = 0
Affecting, (0, 1) = 1
Affecting, (0, 2) = 2
Affecting, (2, 10) = 2
Affecting, (2, 11) = 4
Affecting, (3, 5) = 5
Affecting, (3, 7) = 1
Affecting, (3, 10) = 0
sudoku is solvable!!

1,2,3,6,5,4,7,8,9,10,11,12,13,14,15,16,
5,4,7,8,3,1,13,14,2,6,15,16,9,10,11,12,
13,14,15,16,9,10,11,12,4,7,3,5,2,1,6,8,
9,10,11,12,15,6,16,2,8,13,1,14,3,5,4,7,
12,11,9,2,1,14,15,10,7,3,5,6,8,16,13,4,
15,8,1,4,11,13,3,7,16,14,12,9,10,6,5,2,
16,13,6,3,12,9,2,5,10,4,8,15,14,11,7,1,
10,7,14,5,4,8,6,16,13,1,2,11,12,15,9,3,
14,12,16,9,13,15,1,3,5,8,7,2,6,4,10,11,
3,5,4,1,2,16,12,6,11,15,13,10,7,8,14,9,
11,6,8,7,10,5,4,9,1,16,14,3,15,12,2,13,
2,15,10,13,8,7,14,11,6,12,9,4,1,3,16,5,
4,9,13,15,7,11,10,1,3,5,6,8,16,2,12,14,
8,1,2,14,6,12,5,4,15,9,16,7,11,13,3,10,
7,3,12,10,16,2,8,15,14,11,4,13,5,9,1,6,
6,16,5,11,14,3,9,13,12,2,10,1,4,7,8,15,

```

