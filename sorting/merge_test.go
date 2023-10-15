package sorting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	slice := []int{6, 5, 12, 10, 9, 1}
	Merge(slice)
	fmt.Println(slice)
	expected := []int{1, 5, 6, 9, 10, 12}
	require.Equal(t, slice, expected)
}
