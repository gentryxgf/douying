package service

import "mime/multipart"

type VideoService struct{}

func (VideoService) UploadVideo(video *multipart.FileHeader) (err error) {

	return nil
}
