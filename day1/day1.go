package day1

import (
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func SolveFirstPart(input string) (ans int, err error) {
	leftList, rightList := formatInput(input)
	slices.Sort(leftList)
	slices.Sort(rightList)
	for i := range leftList {
		ans += abs(leftList[i] - rightList[i])
	}
	return ans, nil
}

func SolveSecondPart(input string) (ans int, err error) {
	leftList, rightList := formatInput(input)
	rightFrequencies := make(map[int]int, 0)
	for _, num := range rightList {
		rightFrequencies[num] += 1
	}
	for _, num := range leftList {
		ans += num * rightFrequencies[num]
	}
	return ans, nil
}

func formatInput(input string) (left []int, right []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	leftList := make([]int, 0, len(lines))
	rightList := make([]int, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		left, right := fields[0], fields[1]
		leftNum, _ := strconv.Atoi(left)
		rightNum, _ := strconv.Atoi(right)
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}
	return leftList, rightList
}

func abs[T constraints.Integer](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
