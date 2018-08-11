package main

import (
	"os"
	"fmt"
	"flag"
	"strconv"
	"strings"
	"ras-robotics.com/golife"
)

func NextV2(f *golife.Field, x, y int) bool {
	//if cell is alive and ha no living horizontal neighbors return false
	if f.Alive(x, y) && !f.Alive(x+1, y) && !f.Alive(x-1, y) && !f.Alive(x, y+1) && !f.Alive(x, y-1) {
		return false
	}
	//if cell is dead but has exactly one living neighbor return true
	if !f.Alive(x, y) && f.CountLiveNeighbors(x, y) == 1 {
		return true
	}
	//state remains the same
	return f.Alive(x, y)
}

func main() {
	flag.Parse()
	argsWithoutProg := os.Args[1:]
	w, _ := strconv.Atoi(argsWithoutProg[0])
	h, _ := strconv.Atoi(argsWithoutProg[1])
	infect, _ := strconv.Atoi(argsWithoutProg[2])
	maxGen, _ := strconv.Atoi(argsWithoutProg[3])
	seed := strings.Split(argsWithoutProg[4], " ")
	if len(seed) != w*h {
		fmt.Println("bad seed langth ")
		return
	}
	seedArray := make([]bool, 0, h*w)
	for _, e := range seed {
		num, err := strconv.Atoi(e)
		if err != nil || (num != 0 && num != 1) {
			fmt.Println("bad seed content")
			return
		}
		seedArray = append(seedArray, num == 1)
	}
	l := golife.NewLife(w, h, seedArray, golife.Next)
	fmt.Println(l)
	for i := 0; i < infect; i++ {
		l.Step()
		fmt.Println(l)
	}
	l.SetNextFunc(NextV2)
	for i := infect; i < maxGen; i++ {
		l.Step()
		fmt.Println(l)
	}
}