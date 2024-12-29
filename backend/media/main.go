package main

import "context"

func main() {
	upload := NewUpload()
	svc := NewService(upload)

	svc.Upload(context.Background())
}
