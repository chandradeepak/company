package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parsing(t *testing.T) {
	t.Run("quotes in record", func(t *testing.T) {
		emp, err := parseRecord(`Alan,"1",150`)
		require.NoError(t, err)
		require.NotNil(t, emp)
	})

	t.Run("more than 3 delimeters", func(t *testing.T) {
		emp, err := parseRecord(`Alan,1,150,`)
		require.Error(t, err)
		require.Nil(t, emp)
	})

	t.Run("empty ID", func(t *testing.T) {
		emp, err := parseRecord(`Alan,"",150,`)
		require.Error(t, err)
		require.Nil(t, emp)
	})

}

func Test_LoadData(t *testing.T) {
	t.Run("empty ID", func(t *testing.T) {
		emp, err := loadData("data_duplicate.csv")
		require.Error(t, err)
		require.Nil(t, emp)
	})
}
