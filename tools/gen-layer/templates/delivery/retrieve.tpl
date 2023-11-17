{{define "retrieve"}}
func (d *{{.CamelCase}}Delivery) Get{{.PascalCase}}ByID(ctx context.Context, req *pb.Get{{.PascalCase}}ByIDRequest) (*pb.Get{{.PascalCase}}ByIDResponse, error) {
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cant be nil!")
	}

	{{.CamelCase}}, err := d.{{.CamelCase}}Service.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to get {{.CamelCase}} by id: %v", err)
	}

	return &pb.Get{{.PascalCase}}ByIDResponse{
		Data: transform.{{.PascalCase}}ToPbPtr({{.CamelCase}}),
	}, nil
}
{{end}}