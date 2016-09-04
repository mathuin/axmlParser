package axmlParser

import (
	"archive/zip"
	"io/ioutil"

	"github.com/mathuin/axmlParser/binres"
)

func ParseApk(apkpath string, listener Listener) (*Parser, error) {
	r, err := zip.OpenReader(apkpath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var xmlf *zip.File
	var arscf *zip.File

	for _, f := range r.File {
		if f.Name == "AndroidManifest.xml" {
			xmlf = f
		}
		if f.Name == "resources.arsc" {
			arscf = f
		}
		if xmlf != nil && arscf != nil {
			break
		}
	}

	if xmlf == nil || arscf == nil {
		return nil, err
	}

	arc, err := arscf.Open()
	if err != nil {
		return nil, err
	}
	defer arc.Close()

	abuf := make([]byte, 1)
	abuf, err = ioutil.ReadAll(arc)
	if err != nil {
		return nil, err
	}

	tbl := new(binres.Table)
	if err = tbl.UnmarshalBinary(abuf); err != nil {
		return nil, err
	}

	rc, err := xmlf.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	bs, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	parser := New(listener, tbl)
	err = parser.Parse(bs)
	if err != nil {
		return nil, err
	}
	return parser, nil
}

func ParseAxml(axmlpath string, listener Listener) (*Parser, error) {
	bs, err := ioutil.ReadFile(axmlpath)
	if err != nil {
		return nil, err
	}
	parser := New(listener, nil)
	err = parser.Parse(bs)
	if err != nil {
		return nil, err
	}
	return parser, nil
}
