package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type Subject struct {
	Title string `json:"title"`
	Cards []Card `json:"cards"`
}

type Card struct {
	Phrase      string `json:"phrase"`
	Example     string `json:"example"`
	Answer      string `json:"answer"`
	Description string `json:"description"`
}

var pathname = "./cards"
var filename = ""

func main() {
	for !fileNavigator() {
	}

	// Open our jsonFile
	jsonFile, err := os.Open(path.Join(pathname, filename))
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened " + filename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var subject Subject

	err = json.Unmarshal(byteValue, &subject)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		callClear()
		println("Press ctrl+c to quit")
		card := randString(subject.Cards)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(card.Phrase)
		if card.Example != "" {
			fmt.Print("\n", card.Example)
		}
		_, _ = reader.ReadString('\n')
		fmt.Println(card.Answer)
		_, _ = reader.ReadString('\n')
	}
}

func fileNavigator() bool {
	log.Println("Please type the name or number of a file you would like to practice")

	files, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println(err)
		return false
	}

	for i, f := range files {
		fmt.Println(i, f.Name())
		filename = f.Name()
	}

	reader := bufio.NewReader(os.Stdin)
	typedString, _ := reader.ReadString('\n')
	typedString = strings.TrimSuffix(typedString, "\n")
	fmt.Println("Input " + typedString)

	for i, f := range files {
		if f.Name() == typedString || strconv.FormatInt(int64(i), 10) == typedString {
			fmt.Println("You have selected " + f.Name())
			filename = f.Name()
			return true
		}
	}
	return false

}

func randString(m []Card) Card {
	i := rand.Intn(len(m))
	for _, k := range m {
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
		err := cmd.Run()
		if err != nil {
			log.Println(err)
			return
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Println(err)
			return
		}
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Println(err)
			return
		}
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
