{{define "update"}}
func (d *{{.CamelCase}}Delivery) Update{{.PascalCase}}(ctx context.Context, req *pb.Update{{.PascalCase}}Request) (*pb.Update{{.PascalCase}}Response, error) {
	if req.GetData() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "the request body data cant be nil!")
	}

	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cant be nil!")
	}

	if err := d.{{.CamelCase}}Service.UpdateByID(ctx, req.GetId(), transform.PbTo{{.PascalCase}}Ptr(req.GetData())); err != nil {
		return nil, status.Errorf(codes.Internal, "unable to update {{.CamelCase}}: %v", err)
	}
	
	return &pb.Update{{.PascalCase}}Response{
	}, nil
}
{{end}}