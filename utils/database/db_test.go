package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Actor struct {
	ID    string `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

func TestIsExistFieldInTable(t *testing.T) {

	got := IsExistFieldInTable(Actor{}, "id")
	assert.True(t, got)

	got = IsExistFieldInTable(Actor{}, "id_1q2312312")
	assert.False(t, got)
}
