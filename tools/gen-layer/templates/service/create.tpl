{{define "create"}}
func (s *{{.CamelCase}}Service) Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}) error {
	{{.CamelCase}}.ID = uuid.NewString()
	if err := s.{{.CamelCase}}Repo.Create(ctx, {{.CamelCase}}); err != nil {
		return err
	}
	return nil
}
{{end}}