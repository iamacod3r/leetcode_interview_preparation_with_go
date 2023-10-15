package sorting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubble(t *testing.T) {

	slice1 := []int{3, 4, 5, 2, 1}
	Bubble(slice1)
	fmt.Println(slice1)
	expected1 := []int{1, 2, 3, 4, 5}
	require.Equal(t, slice1, expected1)

	slice2 := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	Bubble(slice2)
	fmt.Println(slice2)
	expected2 := []int{-3, -1, 1, 2, 3, 4, 5, 7, 8}
	require.Equal(t, slice2, expected2)
}
