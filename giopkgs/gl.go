// this file was generated by gomacro command: import _i "gioui.org/gpu/gl"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	gl "gioui.org/gpu/gl"
)

// reflection: allow interpreted code to import "gioui.org/gpu/gl"
func init() {
	imports.Packages["gioui.org/gpu/gl"] = imports.Package{
	Name: "gl",
	Binds: map[string]r.Value{
		"NewBackend":	r.ValueOf(gl.NewBackend),
	}, Types: map[string]r.Type{
		"Backend":	r.TypeOf((*gl.Backend)(nil)).Elem(),
		"Context":	r.TypeOf((*gl.Context)(nil)).Elem(),
	}, 
	}
}
