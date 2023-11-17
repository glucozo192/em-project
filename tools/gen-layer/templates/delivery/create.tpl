{{define "create"}}
func (d *{{.CamelCase}}Delivery) Create{{.PascalCase}}(ctx context.Context, req *pb.Create{{.PascalCase}}Request) (*pb.Create{{.PascalCase}}Response, error) {
	if req.GetData() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "the request body data cant be nil!")
	}

	if err := d.{{.CamelCase}}Service.Create(ctx, transform.PbTo{{.PascalCase}}Ptr(req.GetData())); err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create {{.CamelCase}}: %v", err)
	}

	return &pb.Create{{.PascalCase}}Response{}, nil
}
{{end}}