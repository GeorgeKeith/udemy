package main

func main() {

	const n int = 100000000
	input := originate(n)
	output := fanOut(input)
	for i := 1; i <= n; i++ {
		<-output
	}
	close(output)
}

func originate(n int) (input chan int) {
	input = make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			input <- i % 12 + 1
		}
		close(input)
	}()
	return
}

func fanOut(input chan int) (output chan int) {
	output = make(chan int)
	go func(){
		for n := range input {
//			fan := fanIn(output)
			go func(n int) {
				f := 1
				for i := 2; i <= n; i++ {
					f *= i
				}
				output <- f
			} (n)
		}
	} ()
	return
}

//func fanIn(output chan int) (fan chan int) {
//	fan = make(chan int)
//	go func () {
//		output <- <-fan
//		close(fan)
//	} ()
//	return
//}
