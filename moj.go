package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func flip(b int, l string) {
	fn := "/sys/class/leds/beaglebone:green:%s/brightness"
	f, err := os.OpenFile(fmt.Sprintf(fn, l), os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(fmt.Sprintf("%d", b))
	f.Close()
}

func main() {
	leds := [...]string{"usr2", "usr3", "mmc0", "heartbeat"}

	rand.Seed(time.Now().UTC().UnixNano())

	for {
		for l := 0; l < len(leds); l++ {
			b := 0
			for i := rand.Intn(9) + 2; i > 0; i-- {
				flip(b, leds[l])
				time.Sleep(1 * 1000000000 / 75)

				if b == 0 {
					b = 1
				} else {
					b = 0
				}
			}
			flip(0, leds[l])
		}
	}
}
