package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(math.MaxInt) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	largest := data[0]
	for _, value := range data[1:] {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func maxChunks(data []int) int {
	if len(data) < CHUNKS {
		return maximum(data)
	}
	chunkSize := len(data) / CHUNKS
	maxima := make([]int, CHUNKS)
	var wg sync.WaitGroup
	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			maxima[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Последовательный поиск")
	start := time.Now()
	largest := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n\n", largest, elapsed)

	fmt.Printf("Параллельный поиск (%d горутин)\n", CHUNKS)
	start = time.Now()
	largest = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", largest, elapsed)
}
