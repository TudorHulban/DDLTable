package ddltable

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Person struct {
	Persons struct{} `hera:"tablename"`

	Name string

	ID uint `hera:"pk, order:0"`

	Age             int16
	AllowedToDrive  bool `hera:"default:false, columnname:driving, order:2"`
	skipNotExported bool //nolint:unused
	SkipExported    bool `hera:"-"`
	Birthdate       sql.NullString
}

type PersonsInGroups struct {
	IDPersons uint `hera:"index:ix_personsingroups"`
	IDGroups  uint `hera:"index:ix_personsingroups"`

	FUnique string `hera:"indexunique:ixunique"`

	Field1 uint `hera:"index:ix_fields_personsingroups, indexunique:ixunique"`
	Field2 uint `hera:"index:ix_fields_personsingroups, indexunique:ixunique"`
}

type SQLNullTypes struct {
	SomeInteger16 sql.NullInt16
	SomeInteger32 sql.NullInt32
	SomeInteger64 sql.NullInt64

	SomeFloat64 sql.NullFloat64
}

func TestPersonsTable(t *testing.T) {
	table, errNew := NewTable(
		RootTagName,
		&Person{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	fmt.Println(table.MigrationTable)
	fmt.Println(table.MigrationIndexes)
}

func TestTablePersonsInGroups(t *testing.T) {
	table, errNew := NewTable(
		RootTagName,
		&PersonsInGroups{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	fmt.Println(table.MigrationTable)
	fmt.Println(table.MigrationIndexes)
}

func TestTableWSQLNullTypes(t *testing.T) {
	table, errNew := NewTable(
		RootTagName,
		&SQLNullTypes{},
	)
	require.NoError(t, errNew)
	require.NotZero(t, table)

	fmt.Println(table.MigrationTable)
	fmt.Println(table.MigrationIndexes)
}
