// This file includes wrappers for symbols included since GTK 3.16, and
// and should not be included in a build intended to target any older GTK
// versions. To target an older build, such as 3.12, use
// 'go build -tags gtk_3_12'.  Otherwise, if no build tags are used, GTK 3.16
// is assumed and this file is built.

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_16

package gtk

// #include <gtk/gtk.h>
// #include "popover_menu_since_3_16.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_popover_menu_get_type()), marshalPopover},
	}

	glib.RegisterGValueMarshalers(tm)
	WrapMap["GtkPopoverMenu"] = wrapPopoverMenu
}

// PopoverMenu is a representation of GTK's GtkPopoverMenu.
type PopoverMenu struct {
	Bin
}

func (v *PopoverMenu) native() *C.GtkPopoverMenu {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkPopoverMenu(p)
}

func marshalPopoverMenu(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapPopoverMenu(glib.Take(unsafe.Pointer(c))), nil
}

func wrapPopoverMenu(obj *glib.Object) *PopoverMenu {
	return &PopoverMenu{Popover{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}

// PopoverMenuNew is a wrapper around gtk_popover_menu_new
func PopoverMenuNew() (*PopoverMenu, error) {
	c := C.gtk_popover_menu_new(nil)
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapPopoverMenu(glib.Take(unsafe.Pointer(c))), nil
}

// OpenSubmenu is a wrapper around gtk_popover_menu_open_submenu
func OpenSubmenu(name string) {
	cstr1 := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr1))

	C.gtk_popover_menu_open_submenu(cstr1)
}
