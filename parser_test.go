package axmlParser

import (
	"reflect"
	"testing"
)

func check(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Wanted \"%+v\", got \"%+v\" instead", expected, actual)
	}
}

var parserTests = []struct {
	name, version, label, icon, filename string
	err                                  error
}{
	{"org.twilley.android.firstapp", "1.0", "FirstApp", "res/drawable-mdpi-v4/ic_launcher.png", "a.apk", nil},
}

func TestParser(t *testing.T) {
	for _, tt := range parserTests {
		listener := new(AppNameListener)
		_, err := ParseApk(tt.filename, listener)
		check(t, tt.name, listener.PackageName)
		check(t, tt.version, listener.VersionName)
		check(t, tt.label, listener.ApplicationLabel)
		check(t, tt.icon, listener.ApplicationIcon)
		if !reflect.DeepEqual(tt.err, err) {
			t.Errorf("Wanted \"%+v\", got \"%+v\" instead", tt.err, err)
		}
	}
}
