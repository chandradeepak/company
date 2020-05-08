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
		empMap, err := loadData("test/data_duplicate.csv")
		require.Error(t, err)
		require.Nil(t, empMap)
	})
	t.Run("employee ID empty", func(t *testing.T) {
		empMap, err := loadData("test/empidempty.csv")
		require.Error(t, err)
		require.Nil(t, empMap)
	})
	t.Run("wrong data format ID", func(t *testing.T) {
		empMap, err := loadData("test/wrongdata.csv")
		require.Error(t, err)
		require.Nil(t, empMap)
	})
}

func Test_LinkRelationShip(t *testing.T) {
	t.Run("no ceo", func(t *testing.T) {
		empMap, err := loadData("test/noceo.csv")
		require.NoError(t, err)
		require.NotNil(t, empMap)
		ceo, err := linkRelationShip(empMap)
		require.Error(t, err)
		require.Nil(t, ceo)
	})

	t.Run("circular dependency", func(t *testing.T) {
		empMap, err := loadData("test/circular.csv")
		require.NoError(t, err)
		require.NotNil(t, empMap)
		ceo, err := linkRelationShip(empMap)
		require.Error(t, err)
		require.Nil(t, ceo)
	})
}
