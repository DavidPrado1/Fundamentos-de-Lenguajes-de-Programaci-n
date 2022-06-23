package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func insertionSort(arr []int) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i -= 1
		}
		arr[i+1] = key
	}
}

func printArray(arr []int) {
	fmt.Print("[")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("]")
}

func main() {
	nombres := []string{"test1.txt", "test2.txt", "test3.txt", "test4.txt", "test5.txt", "test6.txt", "test7.txt", "test8.txt", "test9.txt", "test10.txt", "test11.txt", "test12.txt", "test13.txt", "test14.txt", "test15.txt"}
	tamaños := []int{100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000}
	// open the file
	for i := 0; i < 15; i++ {
		var ltiempos []float64
		for j := 0; j < 5; j++ {
			var arr []int
			file, err := os.Open(nombres[i])

			//handle errors while opening
			if err != nil {
				log.Fatalf("Error when opening file: %s", err)
			}

			fileScanner := bufio.NewScanner(file)
			for fileScanner.Scan() {
				sv := fileScanner.Text()
				if sv, err := strconv.Atoi(sv); err == nil {
					arr = append(arr, sv)
				}
			}
			file.Close()
			insertionSort(arr)
			start := time.Now()
			if j == 4 {
				printArray(arr)
			}
			duration := time.Since(start).Seconds()
			ltiempos = append(ltiempos, float64(duration))

		}
		var promf float64 = 0
		for k := 0; k < 5; k++ {
			promf += ltiempos[k]
		}
		promf = promf / 5
		var dvs float64 = 0
		for p := 0; p < 5; p++ {
			dvs += math.Pow(ltiempos[p]-promf, 2)
		}
		dvs = math.Sqrt(dvs / 4)
		fmt.Println("[Array Tamaño:", tamaños[i], "][Promedio de Tiempo de ejecucion:", promf, "][Desviacion Estandar:", dvs, "]")
	}

}
