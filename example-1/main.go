package main

import (
	"fmt"
	"strconv"
)

// The program starts with 10 Goroutines. We created a ch
// string channel and writing data to that channel by running 10 goroutines concurrently.
// The direction of the arrow with respect to the channel specifies whether the data is
// sent or received. The arrow points towards ch specifies we are writing to channel ch.
// The arrow points outwards from ch specifies we are reading from channel ch.

// Программа начинается с 10 Горутинов. Мы создали канал ch string и записали данные в этот канал,
// запустив одновременно 10 подпрограмм. Направление стрелки относительно канала указывает,
// отправляются или принимаются данные. Стрелка, указывающая на ch, указывает,
// что мы записываем в канал ch. Стрелка, указывающая наружу от ch, указывает,
// что мы читаем из канала ch.

func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- "Gourutine: " + strconv.Itoa(i)
			}
		}(i)
	}

	for k := 1; k <= 100; k++ {
		fmt.Println(k, <-ch)
	}
}
