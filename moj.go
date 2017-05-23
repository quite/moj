package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func flip(b int, l int) {
	fn := "/sys/class/leds/beaglebone:green:%s/brightness"
	leds := [...]string{"usr2", "usr3", "mmc0", "heartbeat"}
	f, err := os.OpenFile(fmt.Sprintf(fn, leds[l]), os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(fmt.Sprintf("%d", b))
	f.Close()
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		l := rand.Intn(3 + 1)
		// for l := 0; l <= 3; l++ {
		b := 0
		for i := rand.Intn(9+1) + 2; i > 0; i-- {
			// for i := 0; i < 5; i++ {
			flip(b, l)
			time.Sleep(1 * 1000000000 / 75)

			if b == 0 {
				b = 1
			} else {
				b = 0
			}
		}
		flip(0, l)
		// }
	}
}
