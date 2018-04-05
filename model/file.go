package model

import "time"

type File struct {
	Id               int
	Name             string
	Description      string
	Uploaded         time.Time
	Rating           int
	OriginalFilename string
	Source           string
	Hashsum          string // must be unique
	Tags             []string
	Unsafe           bool
}
