package sys

import "flag"

var DSN string
var MODE string

func init() {
	flag.StringVar(&DSN, "DSN", "root:123456@tcp(127.0.0.1:3306)/gate", "DSN")
	flag.StringVar(&MODE, "mode", "main", "main / check, default is main")
	flag.Parse()
}
