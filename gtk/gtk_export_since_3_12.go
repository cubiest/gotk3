// Same copyright and license as the rest of the files in this project
//go:build !gtk_3_6 && !gtk_3_8 && !gtk_3_10
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10

package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/internal/callback"
)

//export goFlowBoxSortFuncs
func goFlowBoxSortFuncs(child1 *C.GtkFlowBoxChild, child2 *C.GtkFlowBoxChild, userData C.gpointer) C.gint {
	fn := callback.Get(uintptr(userData)).(FlowBoxSortFunc)
	return C.gint(fn(
		wrapFlowBoxChild(glib.Take(unsafe.Pointer(child1))),
		wrapFlowBoxChild(glib.Take(unsafe.Pointer(child2))),
	))
}
