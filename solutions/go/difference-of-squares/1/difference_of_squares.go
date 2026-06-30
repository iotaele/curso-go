package differenceofsquares

func SquareOfSum(n int) int {
    soma := 0
	for i := 1; i <= n; i++ {
        soma += i
    }
    return soma * soma
}

func SumOfSquares(n int) int {
	soma := 0
    for i:= 1; i<= n; i++ {
        soma += (i * i)
    }
    return soma
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
