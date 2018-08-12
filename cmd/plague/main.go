package main

import (
	"fmt"
	"flag"
	"strconv"
	"ras-robotics.com/golife"
)

func main() {
	random := flag.Bool("rand", false, "use random seed")
	flag.Parse()
	argsWithoutProg := flag.Args()
	w, _ := strconv.Atoi(argsWithoutProg[0])
	h, _ := strconv.Atoi(argsWithoutProg[1])
	infect, _ := strconv.Atoi(argsWithoutProg[2])
	maxGen, _ := strconv.Atoi(argsWithoutProg[3])
	var l *golife.Life
	if *random {
		l = golife.NewLifeRandSeed(w, h, golife.Next)
	} else {
		seedStr := argsWithoutProg[4]
		seedArray, err := golife.StringToSeed(seedStr, h, w)
		if err != nil {
			fmt.Println("Error ", err)
			return
		}
		l = golife.NewLife(w, h, seedArray, golife.Next)
	}
	fmt.Println(l)
	for i := 0; i < infect; i++ {
		l.Step()
		fmt.Println(l)
	}
	l.SetNextFunc(golife.NextV2)
	for i := infect; i < maxGen; i++ {
		l.Step()
		fmt.Println(l)
	}
}
