package basics

type BasicProblems struct {
}

func (b *BasicProblems) FindHFC(x int, y int) int {
	hcf := 1
	for i := 1; i <= x && i <= y; i++ {
		if x%i == 0 && y%i == 0 {
			hcf = i
		}
	}
	return hcf
}

func (b *BasicProblems) FactorialNum(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else {
		return n * b.FactorialNum(n-1)
	}
}

func (b *BasicProblems) FibonaciSeries(n int) int {
	if n <= 1 {
		return 1
	} else {
		return b.FibonaciSeries(n-1) + b.FibonaciSeries(n-2)
	}
}
