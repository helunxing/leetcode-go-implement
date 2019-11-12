package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// [-1,0,1,2,-1,-4]
// [0,0,0]
// [-2, 0, 1, 1, 2]

func Test_1(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(threeSum([]int{-1, 0, 1, 2, -1, -4}), []int{})

}
