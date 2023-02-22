package gocleaninterface

import (
	"github.com/Namchee/go-clean-interface/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
