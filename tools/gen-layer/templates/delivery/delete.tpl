{{define "delete"}}
func (d *{{.CamelCase}}Delivery) Delete{{.PascalCase}}(ctx context.Context, req *pb.Delete{{.PascalCase}}Request) (*pb.Delete{{.PascalCase}}Response, error) {
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cant be nil!")
	}

	if err := d.{{.CamelCase}}Service.Delete(ctx, req.GetId()); err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete {{.CamelCase}}: %v", err)
	}

	return &pb.Delete{{.PascalCase}}Response{
	}, nil
}
{{end}}