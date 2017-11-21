package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	ímpar := make(chan int)
	converge := make(chan int)

	go envia(par, ímpar)
	go recebe(par, ímpar, converge)

	for v := range converge {
		fmt.Println("Valor recebido:", v)
	}

}

func envia(p, i chan int) {
	defer close(p)
	defer close(i)
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}

}

func recebe(p, i, c chan int) {
	defer close(c)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
}

// - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
// - Por fim um range retira todas as informações do canal converge.
