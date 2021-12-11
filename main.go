package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type challenge func(string) (string, error)

var challenges = []challenge{
	challenge1,
	challenge2,
}

var (
	toRun = flag.Int("run", 1, "which day's challenge to run")
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(fmt.Sprintf("./inputs/%d", *toRun))
	if err != nil {
		log.Fatalf("failed to read input file: %v\n", err)
	}
	output, err := challenges[*toRun-1](string(inputBytes))
	if err != nil {
		log.Fatalf("challenge function failed: %v\n", err)
	}
	fmt.Println(output)
}

func sumSlice(s []int) int {
	t := 0
	for _, v := range s {
		t += v
	}
	return t
}
