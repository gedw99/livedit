// this file was generated by gomacro command: import _i "gioui.org/x/notify"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	notify "gioui.org/x/notify"
)

// reflection: allow interpreted code to import "gioui.org/x/notify"
func init() {
	imports.Packages["gioui.org/x/notify"] = imports.Package{
	Name: "notify",
	Binds: map[string]r.Value{
		"NewManager":	r.ValueOf(notify.NewManager),
	}, Types: map[string]r.Type{
		"Manager":	r.TypeOf((*notify.Manager)(nil)).Elem(),
		"Notification":	r.TypeOf((*notify.Notification)(nil)).Elem(),
	}, 
	}
}
