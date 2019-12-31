// Same copyright and license as the rest of the files in this project

package glib_test

import (
	"testing"

	"github.com/gotk3/gotk3/glib"
)

func Test_AcceleratorParse(t *testing.T) {
	/*
		testVariant := &Variant{}
		t.Log("native: " + testVariant.Native())
	*/
}

func TestVariantBool(t *testing.T) {
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
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			variant := glib.VariantFromBoolean(tC.value)
			actual := variant.GetBoolean()
			if tC.value != actual {
				t.Error("Expected", tC.value, "got", actual)
			}
		})
	}
}

func TestVariantType(t *testing.T) {
	variant := glib.VariantFromBoolean(true)
	variantType := variant.Type()
	if !glib.VariantTypeEqual(glib.VARIANT_TYPE_BOOLEAN, variantType) {
		t.Error("Expected", glib.VARIANT_TYPE_BOOLEAN, "got", variantType)
	}
}
