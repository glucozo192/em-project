{{define "retrieveByIDs"}}
func (r *{{.CamelCase}}Repository) GetByID(ctx context.Context, ids []string, opts ...repositories.Options) ([]*models.{{.PascalCase}}, error) {
  q := models.New(r.db)

	result, err := q.Get{{.PascalCase}}ByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return result, nil
}
{{end}}