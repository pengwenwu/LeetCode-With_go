package recursion

import "testing"

func TestFibsMap_Fibonacci(t *testing.T) {
	fib := NewFibs(10)
	for i:=1; i<15; i++{
		fib.Fibonacci(i)
		fib.Print(i)
	}
}