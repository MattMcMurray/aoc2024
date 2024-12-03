package day1

import (
	"aoc/lib"
	"fmt"
	"math"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Run the solution for day 1",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(_, cmd string) error {
	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "cmd", cmd, "input.txt")
	inputArray, err := lib.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	var lhs, rhs []int64
	rhMap := make(map[int64]int64)

	for _, line := range inputArray {
		parts := regexp.MustCompile(`\s+`).Split(line, -1)
		left, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return err
		}
		right, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return err
		}
		lhs = append(lhs, left)
		rhs = append(rhs, right)

		if rhMap[right] > 0 {
			rhMap[right] += 1
		} else {
			rhMap[right] = 1
		}
	}

	sort.Slice(lhs, func(i, j int) bool {
		return lhs[i] < lhs[j]
	})

	sort.Slice(rhs, func(i, j int) bool {
		return rhs[i] < rhs[j]
	})

	if len(rhs) != len(lhs) {
		panic("Input lengths do not match")
	}

	total := 0
	var similarityScore int64 = 0

	for i := 0; i < len(lhs); i++ {
		diff := math.Abs(float64(lhs[i]) - float64(rhs[i]))
		total += int(diff)

		similarityScore += lhs[i] * rhMap[lhs[i]]
	}

	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", similarityScore)

	return nil
}
