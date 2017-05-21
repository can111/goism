package sexpconv

import (
	"fmt"
	"go/types"
	"lisp"
	"lisp/function"
	"sexp"
)

func ZeroValue(typ types.Type) sexp.Form {
	if types.Identical(typ, lisp.Types.Symbol) {
		return sexp.Symbol{Val: "nil"}
	}

	if typ, ok := typ.(*types.Basic); ok {
		switch typ.Kind() {
		case types.String:
			return sexp.String{}
		case types.Bool:
			return sexp.Bool{}

		default:
			info := typ.Info()

			if info&types.IsFloat != 0 {
				return sexp.Float{}
			}
			if info&types.IsInteger != 0 {
				return sexp.Int{}
			}
		}
	}

	switch typ := typ.(type) {
	case *types.Basic:
		switch typ.Kind() {
		case types.String:
			return sexp.String{}
		case types.Bool:
			return sexp.Bool{}

		default:
			info := typ.Info()

			if info&types.IsFloat != 0 {
				return sexp.Float{}
			}
			if info&types.IsInteger != 0 {
				return sexp.Int{}
			}
		}

	case *types.Array:
		return &sexp.Call{
			Fn: function.MakeVector,
			Args: []sexp.Form{
				sexp.Int{Val: typ.Len()},
				ZeroValue(typ.Elem()),
			},
		}

	case *types.Map:
		return sexp.Var{
			Name: "Go--nil-map",
			Typ:  types.NewMap(lisp.Types.Object, lisp.Types.Object),
		}
	}

	panic(fmt.Sprintf("can not provide zero value for %#v", typ))
}
