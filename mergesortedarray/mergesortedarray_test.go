package mergesortedarray_test

import (
	"main/mergesortedarray"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	assert := assert.New(t)

	// []int{1, 2, 3, 0, 0, 0}, 3, []int{1, 2, 3, 0, 0, 0}, 2
	actual := mergesortedarray.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	expect := []int{1, 2, 2, 3, 5, 6}

	assert.Equal(expect, actual, "The two words should be the same.")

}
