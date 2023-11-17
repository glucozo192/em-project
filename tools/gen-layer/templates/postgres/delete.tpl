{{define "delete"}}
func (r *{{.CamelCase}}Repository) Delete(ctx context.Context, id string, opts ...repositories.Options) error {
	q := models.New(r.db)
	if err := q.Delete{{.PascalCase}}(ctx, id); err != nil {
		return err
	}

	return nil
}
{{end}}