// this file was generated by gomacro command: import _i "gioui.org/font/gofont"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	gofont "gioui.org/font/gofont"
)

// reflection: allow interpreted code to import "gioui.org/font/gofont"
func init() {
	imports.Packages["gioui.org/font/gofont"] = imports.Package{
	Name: "gofont",
	Binds: map[string]r.Value{
		"Collection":	r.ValueOf(gofont.Collection),
	}, 
	}
}