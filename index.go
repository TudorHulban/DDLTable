package ddltable

import (
	"fmt"
	"strings"
)

type index struct {
	Type        string
	ColumnNames []string
}

func (ix index) migrationUp() func(nameTable, nameIndex string) string {
	if len(ix.Type) == 0 {
		return func(nameTable, nameIndex string) string {
			return fmt.Sprintf(
				"create index if not exists %s on %s(\n"+strings.Join(ix.ColumnNames, ",\n")+"\n);",
				nameIndex,
				nameTable,
			)
		}
	}

	return func(nameTable, nameIndex string) string {
		return fmt.Sprintf(
			"create %s index if not exists %s on %s(\n"+strings.Join(ix.ColumnNames, ",\n")+"\n);",
			ix.Type,
			nameIndex,
			nameTable,
		)
	}
}

func (ix index) migrationDown() func(nameIndex string) string {
	return func(nameIndex string) string {
		return fmt.Sprintf(
			"drop index if exists %s;",
			nameIndex,
		)
	}
}
