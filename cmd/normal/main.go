package main

import (
	"fmt"
	"flag"
	"strconv"
	"ras-robotics.com/golife"
)

func main() {
	var rand = flag.Bool("rand", false, "create a random seed")
	flag.Parse()
	argsWithoutProg := flag.Args()
	w, _ := strconv.Atoi(argsWithoutProg[0])
	h, _ := strconv.Atoi(argsWithoutProg[1])
	maxGen, _ := strconv.Atoi(argsWithoutProg[2])
	var l *golife.Life
	if *rand {
		l = golife.NewLifeRandSeed(w, h, golife.Next)

	} else {
		seedStr := argsWithoutProg[3]
		seedArray, err := golife.StringToSeed(seedStr, h, w)
		if err != nil {
			fmt.Println("Error ", err)
			return
		}
		l = golife.NewLife(w, h, seedArray, golife.Next)
	}

	fmt.Println(l)
	for i := 0; i < maxGen; i++ {
		l.Step()
		fmt.Println(l)
	}
}
