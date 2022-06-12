package letter

import (
	"runtime"
	"sync"
)

// FreqMap is rune frequency counter
type FreqMap map[rune]int

// Frequency calculates rune frequency in each word
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calls frequency on each word
func ConcurrentFrequency(ss []string) FreqMap {
	wordStream := make(chan string)

	go func() {
		for _, s := range ss {
			wordStream <- s
		}
		close(wordStream)
	}()

	var wg sync.WaitGroup
	freqMapStream := make(chan FreqMap)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			for s := range wordStream {
				freqMapStream <- Frequency(s)
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(freqMapStream)
	}()

	result := make(FreqMap)

	for fm := range freqMapStream {
		for r, c := range fm {
			result[r] += c
		}
	}

	return result
}
