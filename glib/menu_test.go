package glib_test

import(
	"testing"

	"github.com/gotk3/gotk3/glib"
)

func TestGetSetAttributeValueCustomBool(t *testing.T) {
	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "false",
			value: false,
		},
	}

	menuItem := glib.MenuItemNew("unit_label", "unit_action")

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			variant := glib.VariantFromBoolean(tC.value)

			menuItem.SetAttributeValue("custom-bool-attribute", variant)
			actual := menuItem.GetAttributeValue("custom-bool-attribute", glib.VARIANT_TYPE_BOOLEAN)

			if !actual.IsType(glib.VARIANT_TYPE_BOOLEAN) {
				t.Error("Expected value of type", glib.VARIANT_TYPE_BOOLEAN, "got", actual.Type())
			}

			if tC.value != actual.GetBoolean() {
				t.Error("Expected", tC.value, "got", actual)
			}
		})
	}
}
