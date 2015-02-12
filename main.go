package main

import (
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	timer := time.Duration(60000) //60,000 Milliseconds (One minute)

	println("You have left the Crystal Room, Croter is center")

	go enrage()

	for {
		croter(timer)

	}

}

func croter(timer time.Duration) {

	time.Sleep(timer * time.Millisecond)
	println("Croter is moving right")

	time.Sleep(timer * time.Millisecond)
	println("Croter is moving center")

	time.Sleep(timer * time.Millisecond)
	println("Croter is moving left")

	time.Sleep(timer * time.Millisecond)
	println("Croter is moving center")

}

func enrage() {

	pid := os.Getpid()
	str := strconv.Itoa(pid)

	start := time.Now().Unix()
	enrageTimer := int64(600) + start //10 Minutes (10x 60 Seconds)

	println("Enrage Timer Started: 10 Minutes Remain")

	for {
		time.Sleep(1000 * time.Millisecond)
		now := time.Now().Unix()

		if now > enrageTimer {
			println("Croter is ENRAGED! You suck. Wipe.")
			exec.Command("kill", "-9", str).Output()
		}
	}
}
