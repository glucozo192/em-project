package main

import (
	"github.com/glu/shopvui/tools/protoc-gen-custom/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	opt := protogen.Options{}
	internal.Run(opt)
}
