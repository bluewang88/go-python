package gorouting_test

import "time"

func say(s string, c int) {
	for i := 0; i < c; i++ {
		time.Sleep(time.Millisecond * 100)
		println(s, i)
	}
}
func HelloGoRoutineWorld() {
	go say("world", 10)
	say("hello", 5)

}
