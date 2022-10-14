package sys

import "flag"

var DSN string

func init() {
	flag.StringVar(&DSN, "DSN", "root:123456@tcp(127.0.0.1:3306)/gate", "DSN")
	flag.Parse()
}
