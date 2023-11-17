{{define "update"}}
func (s *{{.CamelCase}}Service) UpdateByID(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}) error {
	if err := s.{{.CamelCase}}Repo.Update(ctx, &models.{{.PascalCase}}{ ID: id }, {{.CamelCase}}); err != nil {
		return err
	}
	return nil
}
{{end}}