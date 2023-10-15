package sorting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuick(t *testing.T) {
	slice := []int{8, 7, 2, 1, 0, 9, 6}
	Quick(slice)
	fmt.Println(slice)
	expected := []int{0, 1, 2, 6, 7, 8, 9}
	require.Equal(t, slice, expected)
}
