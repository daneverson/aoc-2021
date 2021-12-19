package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// 3882564
// 3385170

type bitProcessor struct {
	data  []uint64
	freq  map[int]int
	width int
}

func challenge3(input string) (string, error) {
	// Create a bit processor, load it with data and track frequencies of each bit position.
	p := &bitProcessor{
		freq: make(map[int]int),
	}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		t := scanner.Text()
		if p.width == 0 {
			p.width = len(t)
		}
		v, _ := strconv.ParseUint(t, 2, 64)
		p.data = append(p.data, v)
		freqUpdate(v, p.freq, p.width)
	}

	// Retrieve values assembled from most and least common bits in each position
	g, e := gammaEpsilon(p.freq, p.width)
	fmt.Printf("gamma:\t%d\nepsil:\t%d\ng * e:\t%d\n", g, e, g*e)

	// Filter original data list based on bit frequencies.
	// o2: keep those with most common bit at each position. Keep 1 if frequencies are equal.
	// co2: keep those with least common bit at each position. Keep 0 if frequencies are equal.
	o2 := p.data
	co2 := p.data
	for i := p.width - 1; i >= 0; i-- {
		filter := func(list []uint64, cmp func(int, int, int) bool, name string) []uint64 {
			// If we've already filtered down to 1 elem, just return
			if len(list) == 1 {
				return list
			}

			// Get new frequencies
			freq := map[int]int{}
			for _, v := range list {
				freqUpdate(v, freq, p.width)
			}

			// Use new frequencies to filter list and create a new one with only valid entries.
			newList := []uint64{}
			for _, v := range list {
				if cmp(int(v), freq[i], i) {
					newList = append(newList, v)
				} else {
				}
			}
			return newList
		}

		mostCommon := func(v int, freq int, bitPos int) bool {
			mask := 1 << bitPos
			anded := v & mask
			if freq >= 0 && mask == anded {
				return true
			}
			if freq < 0 && mask != anded {
				return true
			}
			return false
		}

		leastCommon := func(v int, freq int, bitPos int) bool {
			mask := 1 << bitPos
			anded := v & mask
			if freq < 0 && mask == anded {
				return true
			}
			if freq >= 0 && mask != anded {
				return true
			}
			return false
		}

		o2 = filter(o2, mostCommon, "o2 ")
		co2 = filter(co2, leastCommon, "co2")
	}

	fmt.Printf("o2:\t%d\n", o2[0])
	fmt.Printf("co2:\t%d\n", co2[0])
	fmt.Printf("o * c:\t%d\n", o2[0]*co2[0])

	return "", nil
}

func gammaEpsilon(m map[int]int, w int) (int, int) {
	gammaStr, epsilonStr := "", ""
	for i := 0; i < w; i++ {
		bitPos := w - 1 - i
		if m[bitPos] > 0 {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	g, _ := strconv.ParseUint(gammaStr, 2, 64)
	e, _ := strconv.ParseUint(epsilonStr, 2, 64)

	return int(g), int(e)
}

func freqUpdate(v uint64, freq map[int]int, width int) {
	for i := width - 1; i >= 0; i-- {
		if (1<<i)&v == (1 << i) {
			freq[i]++
		} else {
			freq[i]--
		}
	}
}

///////////////////////////////////////

// type frequency struct {
// 	on, off int
// }
// type bitProcessor struct {
// 	bitWidth        int
// 	data            []uint64
// 	columnFrequency map[int]frequency
// }

// func challenge3(input string) (string, error) {
// 	p := bitProcessor{
// 		data:            []uint64{},
// 		columnFrequency: make(map[int]frequency),
// 	}
// 	p.load(input)

// 	g := p.gammaEpsilon(false)
// 	e := p.gammaEpsilon(true)
// 	o := p.filter(true, 1)
// 	c := p.filter(false, 0)

// 	if o == -1 || c == -1 {
// 		log.Fatalf("got a negative value for o2 (%d) or co2 (%d) reading", o, c)
// 	}

// 	fmt.Printf("gamma: %d\nepsilon: %d\no2: %d\nco2: %d\n", g, e, o, c)
// 	fmt.Printf("gamma * epsilon: %d\n", g*e)
// 	fmt.Printf("o2 * co2: %d\n", o*c)
// 	return "", nil
// }

// func (p *bitProcessor) filter(most bool, tie int) int {
// 	// Start with a copy of the entire data slice.
// 	dataCopy := make([]uint64, len(p.data))
// 	copy(dataCopy, p.data)

// 	checkZero := func(col int, value uint64) bool {
// 		return ((1<<col)^value)&(1<<col) == (1 << col)
// 	}
// 	checkOne := func(col int, value uint64) bool {
// 		return (1<<col)&value == 1<<col
// 	}

// 	// For each bit position, iterate over the existing data. If an entry is valid, keep it in the
// 	// data copy slice. Keep an index pointer to the end of actual "keepers" within the copy list.
// 	currFreq := p.columnFrequency
// 	for i := p.bitWidth - 1; i >= 0; i-- {
// 		// fmt.Printf("current freq counts: %d + %d (%d)\n", currFreq[i].on, currFreq[i].off, currFreq[i].on+currFreq[i].off)
// 		// keepIndex := 0
// 		newFreq := make(map[int]frequency, p.bitWidth)
// 		keepValue := p.valueToKeep(i, most, tie, currFreq)
// 		newCopy := []uint64{}
// 		for _, v := range dataCopy {
// 			if keepValue == 1 && checkOne(i, v) {
// 				// fmt.Printf("want %d in pos %d and have %0.12b (%0.12b)\n", 1, i, v, (1<<i)&v)
// 				// dataCopy[keepIndex] = v
// 				// keepIndex++
// 				newCopy = append(newCopy, v)
// 				freqUpdate(v, p.bitWidth, newFreq)
// 				// fmt.Printf("%0.12b\n", v)
// 			} else if keepValue == 0 && checkZero(i, v) {
// 				// fmt.Printf("want %d in pos %d and have %0.12b (%0.12b)\n", 0, i, v, (1<<i)&v)
// 				// dataCopy[keepIndex] = v
// 				// keepIndex++
// 				newCopy = append(newCopy, v)
// 				freqUpdate(v, p.bitWidth, newFreq)
// 				// fmt.Printf("%0.12b\n", v)
// 			}
// 			for ii := 0; ii < p.bitWidth; ii++ {
// 				// fmt.Printf("%d(%d,%d) ", ii, newFreq[ii].on, newFreq[ii].off)
// 			}
// 			// fmt.Println()
// 		}
// 		// fmt.Println("--------------------------------")
// 		currFreq = newFreq
// 		dataCopy = newCopy

// 		// After checking each bit, check for single entry in keep list. Single entry is base case.
// 		// fmt.Printf("Finishing computing bit %d (%d -> ", i, len(dataCopy))
// 		// dataCopy = dataCopy[:keepIndex]
// 		// fmt.Printf("%d)\n", len(dataCopy))
// 		if len(dataCopy) == 1 {
// 			// fmt.Printf("%+v\n", dataCopy)
// 			return int(dataCopy[0])
// 		}
// 	}

// 	// This shouldn't happen
// 	return -1
// }

// func (p *bitProcessor) valueToKeep(col int, most bool, tie int, freq map[int]frequency) int {
// 	ons, offs := freq[col].on, freq[col].off
// 	if ons == offs {
// 		return tie
// 	}
// 	if most {
// 		if ons > offs {
// 			return 1
// 		}
// 		return 0
// 	}
// 	if ons > offs {
// 		return 0
// 	}
// 	return 1
// }

// func (p *bitProcessor) gammaEpsilon(doEpsilon bool) int {
// 	s := ""
// 	for i := 0; i < p.bitWidth; i++ {
// 		v1, v2 := p.columnFrequency[i].on, p.columnFrequency[i].off
// 		if doEpsilon {
// 			v1, v2 = p.columnFrequency[i].off, p.columnFrequency[i].on
// 		}
// 		if v1 > v2 {
// 			s += "1"
// 		} else {
// 			s += "0"
// 		}
// 	}
// 	ret, err := strconv.ParseUint(s, 2, 64)
// 	if err != nil {
// 		log.Fatalf("failed to parse gamma string as binary int: %v", err)
// 	}
// 	return int(ret)
// }

// func (p *bitProcessor) load(input string) bool {
// 	s := bufio.NewScanner(strings.NewReader(input))
// 	for s.Scan() {
// 		// Parse out the integer.
// 		t := s.Text()
// 		value, err := strconv.ParseUint(t, 2, 64)
// 		if err != nil {
// 			log.Fatalf("bitProcessor failed to parse input line as integer: %v", err)
// 		}
// 		// If this processor doesn't know yet, establish the number of bits per entry based on input.
// 		if p.bitWidth == 0 {
// 			p.bitWidth = len(t)
// 		}
// 		// Keep all the raw data.
// 		p.data = append(p.data, value)
// 		// Update bit frequencies in each column
// 		p.updateFrequencies(value)
// 	}

// 	return true
// }

// func freqUpdate(val uint64, width int, m map[int]frequency) {
// 	for i := 0; i < width; i++ {
// 		f, found := m[i]
// 		if !found {
// 			f = frequency{}
// 		}
// 		c := width - 1 - i
// 		if 1<<c&val == 1<<c {
// 			f.on++
// 		} else {
// 			f.off++
// 		}
// 		m[i] = f
// 	}
// }

// func (p *bitProcessor) updateFrequencies(val uint64) {
// 	for i := 0; i < p.bitWidth; i++ {
// 		currentCount, found := p.columnFrequency[i]
// 		if !found {
// 			currentCount = frequency{}
// 		}
// 		columnToCheck := p.bitWidth - 1 - i
// 		if 1<<columnToCheck&val == 1<<columnToCheck {
// 			currentCount.on++
// 		} else {
// 			currentCount.off++
// 		}
// 		p.columnFrequency[i] = currentCount
// 	}
// }
