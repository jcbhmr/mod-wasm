//go:build generate

//go:generate go tool wkg wit fetch
//go:generate go tool wit-bindgen-go generate ./wit/ --world mod --out ./internal/

package main
