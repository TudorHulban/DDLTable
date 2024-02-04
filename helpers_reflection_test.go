package ddltable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestObjectName(t *testing.T) {
	require.Equal(t,
		"*Person",
		getObjectName(
			&Person{},
		),
	)
}
