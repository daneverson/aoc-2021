package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type windowScanner struct {
	s               *bufio.Scanner
	windowSize      int
	currWindow      []int
	prevWindowTotal int
	increased       bool
}

func (s *windowScanner) fill() bool {
	for currLen := len(s.currWindow); currLen < s.windowSize; currLen++ {
		if !s.s.Scan() {
			return false
		}
		v, _ := strconv.ParseInt(s.s.Text(), 10, 64)
		s.currWindow = append(s.currWindow, int(v))
	}
	s.prevWindowTotal = sumSlice(s.currWindow)
	return true
}

func (s *windowScanner) scan() bool {
	if !s.s.Scan() {
		return false
	}
	v, _ := strconv.ParseInt(s.s.Text(), 10, 64)
	s.currWindow = append(s.currWindow[1:], int(v))
	currTotal := sumSlice(s.currWindow)
	if currTotal > s.prevWindowTotal {
		s.increased = true
	} else {
		s.increased = false
	}
	s.prevWindowTotal = currTotal
	return true
}

func challenge1(input string) (string, error) {
	scanner := windowScanner{
		s:          bufio.NewScanner(strings.NewReader(input)),
		windowSize: 3,
	}

	if !scanner.fill() {
		return "0", nil
	}
	totalIncreases := 0
	for scanner.scan() {
		if scanner.increased {
			totalIncreases++
		}
	}

	return fmt.Sprint(totalIncreases), nil
}
