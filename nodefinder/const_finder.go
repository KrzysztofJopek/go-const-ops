package nodefinder

import (
	"fmt"
	"go/ast"
	"go/token"
)

var (
	//ErrNodeNotFound signalise that node could not be found by finder
	ErrNodeNotFound = fmt.Errorf("Could not find node")
)

type constNodeFinder struct {
	pos   token.Pos
	found ast.Node
}

func (v *constNodeFinder) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *ast.GenDecl:
		gd := node.(*ast.GenDecl)
		if gd.Pos() <= v.pos && gd.End() > v.pos {
			v.found = gd
			return nil
		}
	}
	return v
}

//GetConstBlock returns first(outer) GenDecl node in the given position
//TODO check if node is const
func GetConstBlock(node ast.Node, pos token.Pos) (*ast.GenDecl, error) {
	nf := constNodeFinder{pos: pos}
	ast.Walk(&nf, node)
	if nf.found == nil {
		return nil, ErrNodeNotFound
	}
	found, ok := nf.found.(*ast.GenDecl)
	if !ok {
		return nil, fmt.Errorf("Could not convert ast.Node to ast.GenDecl")
	}
	return found, nil
}
