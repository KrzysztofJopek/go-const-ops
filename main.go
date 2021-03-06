package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"

	"github.com/go-const-utils/nodefinder"
)

type dupsFinder struct {
	mapper map[litVal][]string
}

type litVal struct {
	kind  token.Token
	value string
}

func (v *dupsFinder) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case *ast.BasicLit:
		gd := node.(*ast.BasicLit)
		lv := litVal{
			kind:  gd.Kind,
			value: gd.Value,
		}
		v.mapper[lv] = append(v.mapper[lv], "X")
	}
	return v
}

func printDups(mapper dupsFinder, r io.Writer) error {
	buff := bytes.Buffer{}
	for i, v := range mapper.mapper {
		if len(v) > 1 {
			buff.WriteString(fmt.Sprintf("Duplicated value: %s\n", i.value))
		}
	}
	_, err := io.Copy(r, &buff)
	if err != nil {
		return fmt.Errorf("Could not print duplicates, IO error: %w", err)
	}
	return nil
}

func getDups(w io.Writer, r io.Reader, pos int) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "const_finder.go", r, 0)
	if err != nil {
		return err
	}

	constNode, err := nodefinder.GetConstBlock(file, token.Pos(pos))
	if err != nil {
		return err
	}

	gd := &dupsFinder{mapper: make(map[litVal][]string)}
	ast.Walk(gd, constNode)
	err = printDups(*gd, w)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var (
		err error
		pos int
	)
	flag.IntVar(&pos, "pos", 0, "position")
	flag.Parse()

	err = getDups(os.Stdout, os.Stdin, pos)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
}
