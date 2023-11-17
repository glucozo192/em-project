package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFieldMap(t *testing.T) {

	// fields, values := FieldMap(Actor{})
	// assert.Contains(t, fields, "id")
	// assert.NotNil(t, values)
	fields, values := FieldMap(&Actor{})
	assert.Contains(t, fields, "id")
	assert.NotNil(t, values)
}

func TestGeneratePlaceHolderForBulkUpsert(t *testing.T) {
	got := GeneratePlaceHolderForBulkUpsert(2, 2)
	require.Equal(t, got, "($1, $2), ($3, $4)")

}
