package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		row := make([]int, n)
		for j, p := range parts {
			num, _ := strconv.Atoi(p)
			row[j] = num
		}
		containers[i] = row
	}
	containerSums := make([]int, n)
	for i, row := range containers {
		sum := 0
		for _, num := range row {
			sum += num
		}
		containerSums[i] = sum
	}
	colorSums := make([]int, n)
	for j := 0; j < n; j++ {
		sum := 0
		for i := 0; i < n; i++ {
			sum += containers[i][j]
		}
		colorSums[j] = sum
	}
	sort.Ints(containerSums)
	sort.Ints(colorSums)
	result := "yes"
	for i := 0; i < n; i++ {
		if containerSums[i] != colorSums[i] {
			result = "no"
			break
		}
	}

	fmt.Println(result)
}
