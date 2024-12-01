package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  string
	output int
}

func TestFirstExample(t *testing.T) {
	testCase := TestCase{
		input: `3   4
4   3
2   5
1   3
3   9
3   3`,
		output: 11,
	}
	ans, err := SolveFirstPart(testCase.input)
	assert.Nil(t, err)
	assert.Equal(t, testCase.output, ans)
}

func TestSecondExample(t *testing.T) {
	testCase := TestCase{
		input: `3   4
4   3
2   5
1   3
3   9
3   3`,
		output: 31,
	}
	ans, err := SolveSecondPart(testCase.input)
	assert.Nil(t, err)
	assert.Equal(t, testCase.output, ans)
}
