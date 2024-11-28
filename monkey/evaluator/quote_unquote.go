package evaluator

import (
	"github.com/taichi-kobaya/go-monkey-interpreter/monkey/ast"
	"github.com/taichi-kobaya/go-monkey-interpreter/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
