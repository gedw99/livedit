// this file was generated by gomacro command: import _i "gioui.org/op"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	op "gioui.org/op"
)

// reflection: allow interpreted code to import "gioui.org/op"
func init() {
	imports.Packages["gioui.org/op"] = imports.Package{
	Name: "op",
	Binds: map[string]r.Value{
		"Affine":	r.ValueOf(op.Affine),
		"Defer":	r.ValueOf(op.Defer),
		"Offset":	r.ValueOf(op.Offset),
		"Record":	r.ValueOf(op.Record),
		"Save":	r.ValueOf(op.Save),
	}, Types: map[string]r.Type{
		"CallOp":	r.TypeOf((*op.CallOp)(nil)).Elem(),
		"InvalidateOp":	r.TypeOf((*op.InvalidateOp)(nil)).Elem(),
		"MacroOp":	r.TypeOf((*op.MacroOp)(nil)).Elem(),
		"Ops":	r.TypeOf((*op.Ops)(nil)).Elem(),
		"StateOp":	r.TypeOf((*op.StateOp)(nil)).Elem(),
		"TransformOp":	r.TypeOf((*op.TransformOp)(nil)).Elem(),
	}, 
	}
}
