package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path/filepath"
)

// ParseGoDir parses all Go files in a directory into ASTs,
// respecting build constraints (build tags) exactly like `go build`.
func ParseGoDir(dir string) (map[string]*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()

	// Use go/build to handle build tags correctly
	ctx := build.Default
	pkg, err := ctx.ImportDir(dir, build.IgnoreVendor)
	if err != nil {
		return nil, nil, fmt.Errorf("ImportDir error: %w", err)
	}

	// Collect files filtered by build tags
	allGoFiles := append([]string{}, pkg.GoFiles...)
	allGoFiles = append(allGoFiles, pkg.CgoFiles...)
	allGoFiles = append(allGoFiles, pkg.TestGoFiles...)

	asts := make(map[string]*ast.File)

	for _, f := range allGoFiles {
		fullPath := filepath.Join(dir, f)

		fileAst, err := parser.ParseFile(fset, fullPath, nil, parser.ParseComments)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing %s: %w", f, err)
		}

		asts[f] = fileAst
	}

	return asts, fset, nil
}

func main() {
	asts, _, err := ParseGoDir("./wgpu")
	if err != nil {
		panic(err)
	}

	for name, root := range asts {
		fmt.Println("Writing wrapper for", name)

		for _, decl := range root.Decls {
			switch decl := decl.(type) {
			case *ast.FuncDecl:
				res := decl.Type.Results
				if res != nil && len(res.List) > 0 {
					// go through all params and remove
				}
			}
		}
	}
}
