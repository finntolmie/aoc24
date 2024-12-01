package day1

import (
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func SolveFirstPart(input string) (ans int, err error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	leftList := make([]int, 0, len(lines))
	rightList := make([]int, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, "   ")
		left, right := split[0], split[1]
		leftNum, err := strconv.Atoi(left)
		if err != nil {
			return ans, err
		}
		rightNum, err := strconv.Atoi(right)
		if err != nil {
			return ans, err
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}
	slices.Sort(leftList)
	slices.Sort(rightList)
	for i := range leftList {
		ans += abs(leftList[i] - rightList[i])
	}
	return ans, nil
}

func SolveSecondPart(input string) (ans int, err error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	leftList := make([]int, 0, len(lines))
	rightFrequencies := make(map[int]int, 0)
	for _, line := range lines {
		split := strings.Split(line, "   ")
		left, right := split[0], split[1]
		leftNum, err := strconv.Atoi(left)
		if err != nil {
			return ans, err
		}
		leftList = append(leftList, leftNum)
		rightNum, err := strconv.Atoi(right)
		if err != nil {
			return ans, err
		}
		rightFrequencies[rightNum] += 1
	}
	for _, num := range leftList {
		ans += num * rightFrequencies[num]
	}
	return ans, nil
}

func abs[T constraints.Integer](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
