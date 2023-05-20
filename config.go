package main

type config struct {
	showHiddenFiles bool
}

func newConfig() *config {
	return &config{}
}
