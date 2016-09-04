Android binary manifest XML file parser library for Golang
=======

This is a golang port of [github.com/xgouchet/AXML](http://github.com/xgouchet/AXML) with modifications.

This package aims to parse an android binary mainfest xml file on an apk file, including references to resources.arsc.

Installation
------
```
go get github.com/mathuin/axmlParser
```

Usage
------

```Go
listener := new(AppNameListener)
_, err := ParseApk(apkfilepath, listener)
```
