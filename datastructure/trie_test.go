package datastructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitTrie(t *testing.T) {

	result := InitTrie()
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.root)
	result.Insert("Cihan")
	require.True(t, result.Search("Cihan"))
	require.False(t, result.Search("Ciha"))
}
