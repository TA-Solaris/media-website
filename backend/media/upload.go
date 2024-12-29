package main

import "context"

type upload struct {
	// TODO - Add storage dependancy
}

func NewUpload() *upload {
	return &upload{}
}

func (u *upload) Upload(context.Context) error {
	return nil
}
