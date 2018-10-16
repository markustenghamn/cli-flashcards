package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

var langStrings = map[string]string{
	"ich [heißen]":       "ich heiße",
	"du [heißen]":        "du heißt",
	"er/sie/es [heißen]": "er/sie/es heißt",
	"wir [heißen]":       "wir heißen",
	"ihr [heißen]":       "ihr heißt",
	"sie/Sie [heißen]":   "sie/Sie heißen",
}

func main() {
	for {
		question := randString(langStrings)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(question)
		_, _ = reader.ReadString('\n')
		fmt.Println(langStrings[question])
		_, _ = reader.ReadString('\n')
		callClear()
	}
}

func randString(m map[string]string) string {
	i := rand.Intn(len(m))
	for k := range m {
		if i == 0 {
			return k
		}
		i--
	}
	panic("Out of questions...")
}

// Code for clearing the console
// https://stackoverflow.com/a/22896706

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func callClear() {
	println(runtime.GOOS)
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
