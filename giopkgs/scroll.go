// this file was generated by gomacro command: import _i "gioui.org/x/scroll"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	scroll "gioui.org/x/scroll"
)

// reflection: allow interpreted code to import "gioui.org/x/scroll"
func init() {
	imports.Packages["gioui.org/x/scroll"] = imports.Package{
	Name: "scroll",
	Binds: map[string]r.Value{
		"DefaultBar":	r.ValueOf(scroll.DefaultBar),
		"Horizontal":	r.ValueOf(scroll.Horizontal),
		"Vertical":	r.ValueOf(scroll.Vertical),
	}, Types: map[string]r.Type{
		"Axis":	r.TypeOf((*scroll.Axis)(nil)).Elem(),
		"Bar":	r.TypeOf((*scroll.Bar)(nil)).Elem(),
		"C":	r.TypeOf((*scroll.C)(nil)).Elem(),
		"D":	r.TypeOf((*scroll.D)(nil)).Elem(),
		"Scrollable":	r.TypeOf((*scroll.Scrollable)(nil)).Elem(),
	}, Untypeds: map[string]string{
		"Horizontal":	"int:1",
		"Vertical":	"int:0",
	}, Wrappers: map[string][]string{
		"Bar":	[]string{"Scrolled","Update",},
		"C":	[]string{"Data","Refs","Reset","Version","Write","Write1","Write2",},
	}, 
	}
}
