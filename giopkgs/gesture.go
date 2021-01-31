// this file was generated by gomacro command: import _i "gioui.org/gesture"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	gesture "gioui.org/gesture"
)

// reflection: allow interpreted code to import "gioui.org/gesture"
func init() {
	imports.Packages["gioui.org/gesture"] = imports.Package{
	Name: "gesture",
	Binds: map[string]r.Value{
		"Both":	r.ValueOf(gesture.Both),
		"Horizontal":	r.ValueOf(gesture.Horizontal),
		"StateDragging":	r.ValueOf(gesture.StateDragging),
		"StateFlinging":	r.ValueOf(gesture.StateFlinging),
		"StateIdle":	r.ValueOf(gesture.StateIdle),
		"TypeCancel":	r.ValueOf(gesture.TypeCancel),
		"TypeClick":	r.ValueOf(gesture.TypeClick),
		"TypePress":	r.ValueOf(gesture.TypePress),
		"Vertical":	r.ValueOf(gesture.Vertical),
	}, Types: map[string]r.Type{
		"Axis":	r.TypeOf((*gesture.Axis)(nil)).Elem(),
		"Click":	r.TypeOf((*gesture.Click)(nil)).Elem(),
		"ClickEvent":	r.TypeOf((*gesture.ClickEvent)(nil)).Elem(),
		"ClickState":	r.TypeOf((*gesture.ClickState)(nil)).Elem(),
		"ClickType":	r.TypeOf((*gesture.ClickType)(nil)).Elem(),
		"Drag":	r.TypeOf((*gesture.Drag)(nil)).Elem(),
		"Scroll":	r.TypeOf((*gesture.Scroll)(nil)).Elem(),
		"ScrollState":	r.TypeOf((*gesture.ScrollState)(nil)).Elem(),
	}, 
	}
}
