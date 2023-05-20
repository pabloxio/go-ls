package main

type config struct {
	showHiddenFiles   bool
	longListingFormat bool
}

func newConfig() *config {
	return &config{}
}
