package main

import (
	"os"
)

func setEnv() {
	os.Setenv("FOO", "1")
}

func swap(string1 string, string2 string) (string, string) {
	return string1, string2
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
