//go:build generate

//go:generate wkg wit fetch
//go:generate go tool rm-rf ./internal/wit-bindgen-go/
//go:generate go tool wit-bindgen-go generate ./wit/ --world mod --out ./internal/wit-bindgen-go/

package main
