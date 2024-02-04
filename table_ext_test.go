package ddltable_test

import (
	"fmt"
	"testing"

	ddltable "github.com/TudorHulban/DDLTable"
	"github.com/stretchr/testify/require"
)

func TestExtPersons(t *testing.T) {
	table, errNew := ddltable.NewTable(
		&ddltable.Person{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	fmt.Println(table.MigrationTable)
	fmt.Println(table.MigrationIndexes)
}
