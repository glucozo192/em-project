{{define "default"}}
package postgres

import (
	"context"
	"encoding/json"

	"{{.Module}}/internal/models"
    "{{.Module}}/internal/repositories"
)

type {{.CamelCase}}Repository struct {
	db models.DBTX
}

func New{{.PascalCase}}Repository(db models.DBTX) repositories.{{.PascalCase}}Repository {
	return &{{.CamelCase}}Repository{
		db,
	}
}
{{end}}