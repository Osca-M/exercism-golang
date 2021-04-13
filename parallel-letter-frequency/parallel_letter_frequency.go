package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given text using goroutines and returns this
// data as a FreqMap.
func ConcurrentFrequency(str []string) FreqMap {
	c := make(chan FreqMap)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(str))
		for _, s := range str {
			go func(s string) {
				defer wg.Done()
				c <- Frequency(s)
			}(s)
		}
		wg.Wait()
		close(c)
	}()
	m := FreqMap{}
	for mp := range c {
		for k, v := range mp {
			m[k] += v
		}
	}
	return m
}
