// this file was generated by gomacro command: import _i "gioui.org/x/profiling"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package giopkgs

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	profiling "gioui.org/x/profiling"
)

// reflection: allow interpreted code to import "gioui.org/x/profiling"
func init() {
	imports.Packages["gioui.org/x/profiling"] = imports.Package{
	Name: "profiling",
	Binds: map[string]r.Value{
		"NewRecorder":	r.ValueOf(profiling.NewRecorder),
	}, Types: map[string]r.Type{
		"CSVTimingRecorder":	r.TypeOf((*profiling.CSVTimingRecorder)(nil)).Elem(),
		"Timings":	r.TypeOf((*profiling.Timings)(nil)).Elem(),
	}, 
	}
}
