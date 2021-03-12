package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

const (
	stringFuncMethodName = "String"
	stringFuncGenFile    = "./gen/stringfunc.go"
)

func getGenParseInfo() (*ParseInfo, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, stringFuncGenFile, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return &ParseInfo{
		Fset: fset,
		F:    f,
	}, nil
}

func GenerateStringFunc(target *ParseInfo) error {
	genParseInfo, err := getGenParseInfo()
	if err != nil {
		return err
	}

	// Add missing imports to the target file for the newly generate func
	importsToAdd, err := getMissingImports(target, genParseInfo)
	if err != nil {
		return err
	}

	for _, importToAdd := range importsToAdd {
		astutil.AddImport(target.Fset, target.F, importToAdd)
	}

	genStringFunc, err := getStringFuncASTNode(genParseInfo)
	if err != nil {
		return err
	}

	astutil.Apply(target.F, func(cr *astutil.Cursor) bool {
		funcDecal, ok := cr.Node().(*ast.FuncDecl)
		if !ok {
			return true
		}
		if funcDecal.Name.String() != stringFuncMethodName {
			return true
		}

		if len(funcDecal.Recv.List) != 1 {
			log.Fatal("invalid number of recievers")
		}

		genStringFunc.Recv = funcDecal.Recv
		// TODO: Allow for comments to be handled correctly
		genStringFunc.Doc = nil

		cr.Replace(genStringFunc)
		return false
	}, nil)

	return nil
}

func getMissingImports(target, genParseInfo *ParseInfo) ([]string, error) {
	genRequiredImports, err := getImports(genParseInfo)
	if err != nil {
		return nil, err
	}

	targetImports, err := getImports(target)
	if err != nil {
		return nil, err
	}

	importsToAdd := []string{}
	for genImport := range genRequiredImports {
		if !targetImports[genImport] {
			importsToAdd = append(importsToAdd, genImport)
		}
	}

	return importsToAdd, nil
}

func getStringFuncASTNode(genParseInfo *ParseInfo) (*ast.FuncDecl, error) {
	var stringFuncNode *ast.FuncDecl

	astutil.Apply(genParseInfo.F, func(cr *astutil.Cursor) bool {
		funcDecal, ok := cr.Node().(*ast.FuncDecl)
		if !ok {
			return true
		}
		if funcDecal.Name.String() == stringFuncMethodName {
			stringFuncNode = funcDecal
			return false
		}

		return true
	}, nil)

	if stringFuncNode == nil {
		return nil, errors.New("Failed to find String Func")
	}

	return stringFuncNode, nil
}

func getImports(target *ParseInfo) (map[string]bool, error) {
	requiredImports := map[string]bool{}

	importLists := astutil.Imports(target.Fset, target.F)
	for _, importList := range importLists {
		for _, importObj := range importList {
			path, err := strconv.Unquote(importObj.Path.Value)
			if err != nil {
				return nil, err
			}
			requiredImports[path] = true
		}
	}

	return requiredImports, nil
}
