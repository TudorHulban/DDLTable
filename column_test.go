package ddltable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPersonNewColumns(t *testing.T) {
	columns, _, errParse := newColumns(
		RootTagName,
		&Person{},
	)
	require.NoError(t, errParse)
	require.NotZero(t, columns)

	fmt.Println(columns)
}

func TestPersonsInGroupsNewColumns(t *testing.T) {
	columns, _, errParse := newColumns(
		RootTagName,
		&PersonsInGroups{},
	)
	require.NoError(t, errParse)
	require.NotZero(t, columns)

	fmt.Println(columns)
}
