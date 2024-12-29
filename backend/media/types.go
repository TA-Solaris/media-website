package main

import "context"

type MediaService interface {
	Upload(context.Context) error
}

type MediaUpload interface {
	Upload(context.Context) error
}
