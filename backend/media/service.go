package main

import "context"

type service struct {
	upload MediaUpload
}

func NewService(upload MediaUpload) *service {
	return &service{upload}
}

func (s *service) Upload(context.Context) error {
	return nil
}
