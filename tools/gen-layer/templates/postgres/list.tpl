{{define "list"}}
func (r *{{.CamelCase}}Repository) GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.{{.PascalCase}}, error) {
	q := models.New(r.db)
	result, err := q.GetList{{.PascalCase}}(ctx, models.GetList{{.PascalCase}}Params{
		Offset: int32(offset),
		Limit:  int32(limit),
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
{{end}}