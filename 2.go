package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type sub struct {
	pos, depth, aim int
}

func (s *sub) move(command string) error {
	split := strings.SplitN(command, " ", 2)
	dir, amountStr := split[0], split[1]
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return fmt.Errorf("failed to parse movement amount: %v", err)
	}
	switch dir {
	case "forward":
		s.pos += amount
		s.depth += (s.aim * amount)
	case "down":
		s.aim += amount
	case "up":
		s.aim -= amount
	default:
		return fmt.Errorf("received unknown direction command '%s'", dir)
	}
	return nil
}

func challenge2(input string) (string, error) {
	s := sub{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		s.move(scanner.Text())
	}

	fmt.Printf("%d %d\n", s.depth, s.pos)

	return fmt.Sprint(s.depth * s.pos), nil
}
