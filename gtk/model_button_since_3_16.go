// This file includes wrappers for symbols included since GTK 3.16, and
// and should not be included in a build intended to target any older GTK
// versions. To target an older build, such as 3.12, use
// 'go build -tags gtk_3_12'.  Otherwise, if no build tags are used, GTK 3.16
// is assumed and this file is built.
// NOTE: property “use-markup” is only available since GTK 3.24

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_16

package gtk

// #include <gtk/gtk.h>
// #include "model_button_since_3_16.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_model_button_get_type()), marshalPopover},
	}

	glib.RegisterGValueMarshalers(tm)
	WrapMap["GtkModelButton"] = wrapModelButton
}

// ModelButton is a representation of GTK's GtkModelButton.
type ModelButton struct {
	Bin

	// Interfaces
	IActionable
}

func (v *ModelButton) native() *C.GtkModelButton {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkModelButton(p)
}

func marshalModelButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapModelButton(glib.Take(unsafe.Pointer(c))), nil
}

func wrapModelButton(obj *glib.Object) *ModelButton {
	actionable := &Actionable{obj}
	return &ModelButton{Button{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}, actionable}, actionable}
}

// ModelButtonNew is a wrapper around gtk_model_button_new
func ModelButtonNew() (*ModelButton, error) {
	c := C.gtk_popover_menu_new(nil)
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapModelButton(glib.Take(unsafe.Pointer(c))), nil
}
