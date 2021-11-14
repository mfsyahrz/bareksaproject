package rest

import "github.com/mfsyahrz/bareksaproject/internal/interface/ioc"

type Handler struct {
	NewsHandler  *NewsHandler
	TagHandler   *TagHandler
	TopicHandler *TopicHandler
}

func SetupHandler(container *ioc.IOC) *Handler {
	return &Handler{
		NewsHandler:  NewNewsHandler(container.NewsService),
		TagHandler:   NewTagHandler(container.TagService),
		TopicHandler: NewTopicHandler(container.TopicService),
	}
}
