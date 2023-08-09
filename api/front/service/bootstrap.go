package service

import "flag"

var (
	sign_key = ""
)

func Bootstrap() {
	addFlag()
}

func addFlag() {
	flag.StringVar(&sign_key, "auth_sign_key", "", "authentic key")
}
