# Codeforces in Go

Solutions to problem sets on [Codeforces](https://codeforces.com/) in Go.

## Structure

```text
[Range]/[ContestId]/[ProblemId]/main.go
```
Example: `1700-1799/1791/A/main.go`

## Creating a New Solution

```bash
./init_problem.sh https://codeforces.com/problemset/problem/69/A
# or shorthand:
./init_problem.sh 69 A
```

## Running a Solution

```bash
go run ./1-99/34/B/main.go < input.txt
```

## Template

Use `snippet.go` as a starting point for new solutions.
