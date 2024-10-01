package main

import goapitestgen "github.com/perfectogo/go-api-test-gen"

func main() {
	goapitestgen.Exec("swagger.json", "handlers")
}
