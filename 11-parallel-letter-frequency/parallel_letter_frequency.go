package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given text and returns this data
func ConcurrentFrequency(sentences []string) FreqMap {
	ch := make(chan FreqMap)
	result := FreqMap{}
	for _, s := range sentences {
		go func(w string) {
			ch <- Frequency(w)
		}(s)
	}
	for range sentences {
		for k, v := range <-ch {
			result[k] += v
		}
	}
	return result
}
