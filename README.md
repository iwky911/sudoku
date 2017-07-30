# How efficiently can we solve a sudoku ?

It depends on the solver!

This repository contains two method to solve sudokus

* a bruteforce solver, simple but not efficient
* an exact cover solver using [dancing links][dl]

[dl]: http://en.wikipedia.org/wiki/Dancing_Links

As the dancing link implementation is pretty fast, we can
use some randomness and quick validity testing to generate
sudokus :)

## Building

```
$ mkdir -p src/github.com/iwky911
$ export GOPATH=$(pwd)
$ cd src/github.com/iwky911
$ git clone https://github.com/iwky911/sudoku.git
$ cd sudoku
$ go build fastsolver.go
$ go build generator.go
```

## Usage
You can pass the sudoku to be solved in a csv format.
Empty or non-integer cells will be considered blank.

For example, using the provided example:
```
$ ./fastsolver -csv=testdata/largeinput.csv
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
1,2,3,6,5,4,7,8,9,10,11,12,13,14,15,16
5,4,7,8,3,1,13,14,2,6,15,16,9,10,11,12
13,14,15,16,9,10,11,12,4,7,3,5,2,1,6,8
9,10,11,12,15,6,16,2,8,13,1,14,3,5,4,7
12,11,9,2,1,14,15,10,7,3,5,6,8,16,13,4
15,8,1,4,11,13,3,7,16,14,12,9,10,6,5,2
16,13,6,3,12,9,2,5,10,4,8,15,14,11,7,1
10,7,14,5,4,8,6,16,13,1,2,11,12,15,9,3
14,12,16,9,13,15,1,3,5,8,7,2,6,4,10,11
3,5,4,1,2,16,12,6,11,15,13,10,7,8,14,9
11,6,8,7,10,5,4,9,1,16,14,3,15,12,2,13
2,15,10,13,8,7,14,11,6,12,9,4,1,3,16,5
4,9,13,15,7,11,10,1,3,5,6,8,16,2,12,14
8,1,2,14,6,12,5,4,15,9,16,7,11,13,3,10
7,3,12,10,16,2,8,15,14,11,4,13,5,9,1,6
6,16,5,11,14,3,9,13,12,2,10,1,4,7,8,15
```

You can also generate new sudokus with generator with generator.go.
Both the size and the percentage of cell erased is ajustable

Example:
```
$ ./generator -n=25
Sparse matrix created
2,9,17,5,,,10,,,1,,,,,14,18,8,,,,,,3,13,
19,,6,,,8,13,,14,4,,2,,9,,,,22,,10,11,17,18,21,
,,,14,12,3,17,24,11,19,23,,,13,21,,,,,,,1,10,,
,,21,,,25,,,,20,,10,,5,,11,12,,13,,8,,,,
11,13,,10,,,,,6,,7,,,8,,,17,,3,,15,,12,25,2
22,,1,,4,,,,,12,15,,,16,23,5,10,20,,,14,,,7,21
,,16,,15,17,19,4,23,25,5,11,12,,,21,,,9,14,3,,8,,
,21,19,7,20,18,5,,10,2,22,14,25,,4,12,,8,,24,13,,9,15,1
23,14,13,,,,,,,,8,7,,,,,,,22,19,,,,,11
5,,8,,9,22,,7,,14,,,1,,,2,23,25,18,,16,,4,,
,24,,,,,21,5,8,,,15,,11,18,,,10,,,12,,,3,
12,10,20,3,2,11,14,13,19,16,,,5,,24,,21,,,15,,22,23,6,8
15,,18,,,,,6,,,,8,17,22,16,23,,24,12,5,21,20,13,,4
21,8,,,,,,,,23,19,25,13,7,,22,,6,,18,,,,,
7,25,23,,,,,15,18,,,,10,,,,11,16,8,4,,,,,
20,5,15,9,,13,22,12,,8,,,16,,,10,,21,14,,7,,,18,25
,,12,,19,14,25,1,,9,6,5,21,10,,15,,,23,,,13,,,
,7,,8,3,,,,,,25,1,2,14,19,6,13,,,9,,11,,12,
,,,,14,,,18,,3,12,,,15,,16,,,,,,,,23,
,23,2,25,21,,4,,,,13,,,,9,,,,,12,6,,,8,
,17,3,,8,,,25,,11,,,14,,2,19,22,15,21,,10,,7,,13
6,,,,1,,23,9,,,20,,8,,25,,18,13,,2,,3,,,16
13,,11,20,,,8,,22,18,1,,,,,3,4,23,,17,25,9,15,,
14,2,,4,22,,,,,,10,13,11,19,17,,,,16,,23,24,,,
,19,,21,18,1,24,10,,,16,,22,,,20,,12,,7,5,8,,,
```

