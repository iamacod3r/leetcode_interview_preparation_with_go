package sorting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelection(t *testing.T) {
	slice := []int{20, 12, 10, 15, 2}
	Selection(slice)
	fmt.Println(slice)
	expected := []int{2, 10, 12, 15, 20}
	require.Equal(t, slice, expected)
}
