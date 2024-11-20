package main

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_ParseFlagValue(t *testing.T) {
	cases := []struct {
		Flag     Flag
		SetTo    string
		Expected any
	}{
		{
			Flag:     BoolFlag{},
			SetTo:    "false",
			Expected: false,
		},
		{
			Flag:     StringFlag{Default: "default"},
			SetTo:    "new",
			Expected: "new",
		},
		{
			Flag:     IntFlag{Default: 10},
			SetTo:    "22",
			Expected: 22,
		},
	}

	for _, c := range cases {
		f := c.Flag
		require.Equal(t, reflect.ValueOf(f).FieldByName("Default").Interface(), f.Value())
		f, err := f.WithValue(c.SetTo)
		require.NoError(t, err)
		require.Equal(t, c.Expected, f.Value())
	}
}
