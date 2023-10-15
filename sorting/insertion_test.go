package sorting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertion(t *testing.T) {
	slice := []int{9, 5, 1, 4, 3}
	Insertion(slice)
	fmt.Println(slice)
	expected := []int{1, 3, 4, 5, 9}
	require.Equal(t, slice, expected)
}
