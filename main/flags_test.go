package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Parse(t *testing.T) {
	bf := BoolFlag{
		Default: true,
	}
	require.Equal(t, true, bf.Value())
	f, err := bf.Parse("false")
	require.NoError(t, err)
	require.Equal(t, false, f.Value())

	sf := StringFlag{
		Default: "default",
	}
	require.Equal(t, "default", sf.Value())
	f, err = sf.Parse("new")
	require.NoError(t, err)
	require.Equal(t, "new", f.Value())
}
