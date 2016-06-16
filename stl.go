package main

import (
	"fmt"
	"path/filepath"

	"github.com/hschendel/stl"
)

type STL struct {
	stl.Solid
}

func OpenSTL(file string) *STL {
	file, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
	solid, err := stl.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return &STL{*solid}
}

func (s *STL) Vertices() (vertices []float32) {
	for _, triangle := range s.Triangles {
		for _, vertex := range triangle.Vertices {
			vertices = append(vertices, append(vertex[:], 0, 0)...)
		}
	}

	return
}
