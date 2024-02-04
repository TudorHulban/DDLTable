package ddltable

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type column struct {
	Name         string
	PGType       string
	DefaultValue string
	IndexName    string
	IndexType    string

	OrderNumber uint

	IsPK          bool
	IsNullable    bool
	IsUnique      bool
	IsIndexed     bool
	IsToBeSkipped bool
}

func newColumns(object any) (columns, string, error) {
	result := make([]*column, 0)

	var alreadyHavePK bool

	var tableName string

	for i := 0; i < reflect.TypeOf(object).Elem().NumField(); i++ {
		fieldRoot := reflect.TypeOf(object).
			Elem().
			FieldByIndex([]int{i})

		column := column{
			Name:        fieldRoot.Name,
			OrderNumber: uint(i),
		}

		if valueTag, hasTag := fieldRoot.Tag.Lookup(_TagName); hasTag {
			errUpdate := column.UpdateWith(valueTag, alreadyHavePK)
			if errUpdate != nil {
				if errUpdate.Error() == errIsOverrideTableName.Error() {
					tableName = strings.ToLower(fieldRoot.Name)

					continue
				}

				return nil, "",
					errUpdate
			}
		}

		if !fieldRoot.IsExported() {
			continue
		}

		var isNullable bool

		column.PGType, isNullable = reflectToPG(fieldRoot.Type.String(), column.IsPK)
		if isNullable {
			column.IsNullable = true
		}

		if column.IsPK {
			alreadyHavePK = true
		}

		result = append(result, &column)
	}

	return result,
		tableName,
		nil
}

func (col *column) UpdateWith(tagValues string, alreadyHavePK bool) error {
	for _, tagValue := range strings.Split(
		tagValues, ",",
	) {
		tagClean := strings.ToLower(
			strings.TrimSpace(tagValue),
		)

		var compoundTagValue string

		if strings.Contains(tagClean, _TagSeparator) {
			tagCompound := strings.Split(tagClean, _TagSeparator)

			if len(tagCompound) != 2 {
				return fmt.Errorf(
					"malformed tag value: %s",
					tagClean,
				)
			}

			tagClean = tagCompound[0]
			compoundTagValue = tagCompound[1]
		}

		switch tagClean {
		case "":
			return nil

		case "-":
			col.IsToBeSkipped = true

			return nil

		case _TagOverrideTableName:
			return errIsOverrideTableName

		case _TagPK:
			if alreadyHavePK {
				return errors.New("more than one primary key field detected. max is 1")
			}

			col.IsPK = true

		case _TagOverrideOrder:
			order, errConv := strconv.Atoi(compoundTagValue)
			if errConv != nil {
				return fmt.Errorf(
					"override order tag: %w",
					errConv,
				)
			}

			col.OrderNumber = uint(order)

		case _TagIndexed:
			col.IsIndexed = true
			col.IndexName = compoundTagValue

		case _TagIndexUnique:
			col.IsIndexed = true
			col.IndexType = _indexUnique
			col.IndexName = compoundTagValue

		case _TagRequired:
			col.IsNullable = false

		case _TagDefault:
			col.DefaultValue = compoundTagValue

		case _TagOverrideColumnName:
			col.Name = compoundTagValue
		}
	}

	return nil
}

func (col *column) AsDDLPostgres() string {
	if col.IsToBeSkipped {
		return ""
	}

	result := []string{
		strings.ToLower(col.Name),
	}

	result = append(result,
		col.PGType,
	)

	if col.IsPK {
		result = append(result,
			"PRIMARY KEY",
		)
	}

	if col.IsUnique {
		result = append(result, "UNIQUE")
	}

	if !col.IsNullable {
		result = append(result, "NOT NULL")
	}

	if len(col.DefaultValue) > 0 {
		result = append(result, "DEFAULT")
		result = append(result, col.DefaultValue)
	}

	return strings.Join(result, " ")
}
