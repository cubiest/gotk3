// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

package gtk

import (
	"os"
	"testing"

	"github.com/gotk3/gotk3/glib"
)

func TestApplicationWindowNew(t *testing.T) {

	app, err := ApplicationNew("com.github.gotk3.unit", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		t.Error("ApplicationNew returned an error:", err)
	}
	if app == nil {
		t.Error("ApplicationNew returned nil")
	}

	app.Connect("activate", func() {
		defer app.Quit()

		win, err := ApplicationWindowNew(app)
		if err != nil {
			t.Error("ApplicationWindowNew returned an error:", err)
		}
		if app == nil {
			t.Error("ApplicationWindowNew returned nil")
		}

		if win.IActionGroup == nil {
			t.Error("Embedded IActionGroup is nil")
		}
		if win.IActionMap == nil {
			t.Error("Embedded IActionGroup is nil")
		}
	})

	rc := app.Run([]string{os.Args[0]})
	if rc != 0 {
		t.Errorf("Application return code was %v, expected 0", rc)
	}
}

func TestApplicationWindow_SetShowMenubar_GetShowMenubar(t *testing.T) {
	app, err := ApplicationNew("com.github.gotk3.unit", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		t.Error("ApplicationNew returned an error:", err)
	}
	if app == nil {
		t.Error("ApplicationNew returned nil")
	}

	app.Connect("activate", func() {
		defer app.Quit()

		win, err := ApplicationWindowNew(app)
		if err != nil {
			t.Error("ApplicationWindowNew returned an error:", err)
		}
		if app == nil {
			t.Error("ApplicationWindowNew returned nil")
		}

		win.SetShowMenubar(true)
		ret := win.GetShowMenubar()
		if !ret {
			t.Errorf("GetShowMenubar() = %v, want %v", ret, true)
		}

		win.SetShowMenubar(false)
		ret = win.GetShowMenubar()
		if ret {
			t.Errorf("GetShowMenubar() = %v, want %v", ret, false)
		}

	})

	rc := app.Run([]string{os.Args[0]})
	if rc != 0 {
		t.Errorf("Application return code was %v, expected 0", rc)
	}
}
