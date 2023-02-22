package analyzer

import (
	"flag"
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var (
	noPrefix bool
	noSuffix bool

	flags flag.FlagSet
)

var Analyzer = &analysis.Analyzer{
	Name:     "go-clean-interface",
	Doc:      "Checks for unnecessary interface prefix or suffixes",
	Run:      run,
	Flags:    flags,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func init() {
	flags.BoolVar(&noPrefix, "no-prefix", true, "disallow interfaces to be prefixed with `I`")
	flags.BoolVar(&noSuffix, "no-suffix", true, "disallow interfaces to be suffixed with `Itf`")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		decl, _ := node.(*ast.TypeSpec)

		if _, isItf := decl.Type.(*ast.InterfaceType); !isItf {
			return
		}

		name := decl.Name.Name

		if noPrefix && strings.HasPrefix(name, "I") {
			msg := fmt.Sprintf("Interface %s has unnecessary prefix `I`", name)
			pass.Reportf(node.Pos(), msg)
		}

		if noSuffix && strings.HasSuffix(name, "Itf") {
			msg := fmt.Sprintf("Interface %s has unnecessary suffix `Itf`", name)
			pass.Reportf(node.Pos(), msg)
		}
	})

	return nil, nil
}
