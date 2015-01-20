package main

import (
	"flag"
	"strings"

	"github.com/lunny/tango"
	"github.com/tango-contrib/basicauth"
)

var (
	dir     = flag.String("dir", "./public", "static dir path")
	listen  = flag.String("listen", ":8000", "listen port")
	mode    = flag.Int("mode", tango.Dev, "run mode, 0: dev, 1: product")
	user    = flag.String("user", "", "basic auth user name")
	pass    = flag.String("pass", "", "basic auth user password")
	listDir = flag.Bool("listDir", false, "if list dir files")
	exts    = flag.String("exts", "", "filtered ext files will be supplied")
)

func main() {
	flag.Parse()

	t := tango.New()
	if *user != "" {
		t.Use(basicauth.New(*user, *pass))
		t.Logger().Info("Basic auth module loaded")
	}
	var filterExts []string
	if len(*exts) > 0 {
		filterExts = strings.Split(*exts, ",")
	}
	t.Use(tango.Logging())
	t.Use(tango.Static(tango.StaticOptions{
		RootPath:   *dir,
		ListDir:    *listDir,
		FilterExts: filterExts,
	}))

	t.Mode = *mode
	t.Run(*listen)
}
