//go:build !gen
// +build !gen

package embed

//go:embed www

//go:generate go run -tags generate make.go
