package ioc

import (
	"github.com/labstack/gommon/log"

	"github.com/mfsyahrz/bareksaproject/internal/config"
	"github.com/mfsyahrz/bareksaproject/internal/db"
	postgresNews "github.com/mfsyahrz/bareksaproject/internal/infrastructure/postgres/news"
	postgresTag "github.com/mfsyahrz/bareksaproject/internal/infrastructure/postgres/tag"
	postgresTopic "github.com/mfsyahrz/bareksaproject/internal/infrastructure/postgres/topic"
	"github.com/mfsyahrz/bareksaproject/internal/infrastructure/redis/cache"
	serviceNews "github.com/mfsyahrz/bareksaproject/internal/service/news"
	serviceTag "github.com/mfsyahrz/bareksaproject/internal/service/tag"
	serviceTopic "github.com/mfsyahrz/bareksaproject/internal/service/topic"
)

type IOC struct {
	Config       *config.Config
	NewsService  serviceNews.NewsService
	TagService   serviceTag.TagService
	TopicService serviceTopic.TopicService
	Log          *log.Logger
}

func Setup() *IOC {
	cfg, err := config.New(".env")
	if err != nil {
		panic(err.Error())
	}

	postgresDB, err := db.NewPostgres(cfg.Postgres)
	if err != nil {
		panic(err.Error())
	}

	log := log.New("bareksapr")
	redis := cache.New(&cfg.Redis)
	// construct postgres repo
	newsRepoPG := postgresNews.NewNewsRepo(postgresDB.Conn, log, redis)
	tagRepoPG := postgresTag.NewTagRepo(postgresDB.Conn, log, redis)
	topicRepoPG := postgresTopic.NewTopicRepo(postgresDB.Conn, log)

	// construct services
	newsSvc := serviceNews.NewNewsService(newsRepoPG)
	tagSvc := serviceTag.NewTagService(tagRepoPG)
	topicSvc := serviceTopic.NewTopicService(topicRepoPG)

	return &IOC{
		Config:       cfg,
		NewsService:  newsSvc,
		TagService:   tagSvc,
		TopicService: topicSvc,
		Log:          log,
	}
}
