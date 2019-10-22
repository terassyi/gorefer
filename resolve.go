package gorefer

import (
	"errors"
	"fmt"
	"log"
)

//this file contain resolvers from *ast.FuncDecl or function name(string) to FunctionInfo

//to do
//resolve method list
//


func (pkg *PackageInfo) ResolveMethodList() {
	for _, typ := range pkg.Struct {
		var lis []FunctionInfo
		for _, f := range pkg.Function {
			if typ.Name == f.FuncInfo.ReceiverType {
				fmt.Printf("function:%v is methods of %v type.\n", f.FuncInfo.Name, typ.Name)
				lis = append(lis, f)
			}
		}
		typ.MethodList = lis
		fmt.Printf("methods list(type %v)-> %v\n", typ.Name, typ.MethodList)
	}
}

//util function
func FindFunctionFromId(list []FunctionInfo, id int) (interface{}, error) {
	for _, f := range list {
		if f.Id.Id == id {
			return f, nil
		}
	}
	log.Printf("Not found such function.")
	err := errors.New("Not found such function.")
	return nil, err
}

func FindStructFromId(list []StructInfo, id int) (interface{}, error) {
	for _, s := range list {
		if s.Id.Id == id {
			return s, nil
		}
	}
	log.Printf("Not found such struct.")
	err := errors.New("Not found such struct.")
	return nil, err
}

func FindInterfaceFromId(list []InterfaceInfo, id int) (interface{}, error) {
	for _, i := range list {
		if i.Id.Id == id {
			return i, nil
		}
	}
	log.Printf("Not found such interface.")
	err := errors.New("Not found such interface.")
	return nil, err
}

func FindVarFromId(list []VarInfo, id int) (interface{}, error) {
	for _, v := range list {
		if v.Id.Id == id {
			return v, nil
		}
	}
	log.Printf("Not found such valiable.")
	err := errors.New("Not found such variable.")
	return nil, err
}


//TODO
//func ResolveUnderLayingType(pkg PackageInfo) {
//
//}
type CallRelationship map[string][]string

func ResolveFuncRelationship(funcList []FunctionInfo) CallRelationship {
	relation := make(CallRelationship)
	var fList []string
	for _, f := range funcList {
		fList = append(fList, f.FuncInfo.Name)
	}
	for _, f := range funcList {
		var callList []string
		for _, call := range f.Call {
			if call.Package == "" {
				if name, exist := contains(fList, call.Name); exist {
					callList = append(callList, name)
				}
			}
		}
		relation[f.FuncInfo.Name] = callList
	}
	return relation
}

func contains(list []string, want string) (string, bool) {
	for _, s := range list {
		if s == want {
			return s, true
		}
	}
	return "", false
}