package service

import "gitlab.com/golang-gin-po/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type VideoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoservice) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoservice) FindAll() []entity.Video {
	return service.videos
}
