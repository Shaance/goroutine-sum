package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ClassicSum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getNbChunks[T interface{}](arr []T, chunkSize int) int {
	chunks := len(arr) / chunkSize
	if len(arr)%chunkSize != 0 {
		chunks += 1
	}
	return chunks
}

func goroutineChunkSum(arr []int, c chan int) {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	c <- sum
}

func GoroutineSum(arr []int) int {
	// TODO choose a better bufferSize?
	const bufferSize = 100
	chunks := getNbChunks(arr, bufferSize)
	c := make(chan int, chunks)
	startIndex := 0
	for i := 0; i < chunks; i++ {
		endIndex := min(startIndex+bufferSize, len(arr))
		go goroutineChunkSum(arr[startIndex:endIndex], c)
		startIndex += bufferSize
	}

	sum := 0
	for i := 0; i < chunks; i++ {
		sum += <-c
	}
	return sum
}

func generateRandomIntArr(size int) []int {
	if size < 0 {
		size = 15000
	}
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}

	return arr
}

func main() {
	arr := generateRandomIntArr(50000)
	classicSum := ClassicSum(arr)
	goroutineSum := GoroutineSum(arr)

	fmt.Printf("Sum from classic algorithm %d\n", classicSum)
	fmt.Printf("Sum from go coroutine %d\n", goroutineSum)
}
