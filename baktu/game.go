package baktu

//func gameHost() {
//	c := make(chan string)
//	go request("boring!", c)
//	for i := 0; i < 15; i++ {
//		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
//	}
//	//time.Sleep(3000 * time.Millisecond)
//	fmt.Println("You're boring; I'm leaving.")
//}
//
//func request(msg string, c chan string) {
//	for i := 0; ; i++ {
//		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
//		//fmt.Printf("%s %d\n", msg, i)
//		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
//	}
//}
