package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type LogstashRepo struct {
	db *gorm.DB
}

func NewLogstashRepo(db *gorm.DB) *LogstashRepo {
	return &LogstashRepo{db: db}
}

func (l LogstashRepo) AddLogstashRepo(ctx context.Context, logstash *entity.Logstash) error {
	return l.db.Create(logstash).Error
}

func (l LogstashRepo) GetLogstashByIdRepo(ctx context.Context, id int64) (*entity.Logstash, error) {
	var logstash = new(entity.Logstash)
	return logstash, l.db.Where("id = ?", id).First(logstash).Error
}

func (l LogstashRepo) GetLogstashRepo(ctx context.Context) ([]*entity.Logstash, error) {
	var logstash []*entity.Logstash
	return logstash, l.db.Find(&logstash).Error
}

func (l LogstashRepo) DeleteLogstashRepo(ctx context.Context, id int64) error {
	return l.db.Delete(&entity.Logstash{}, id).Error
}

func (l LogstashRepo) DeleteLogstashBatchRepo(ctx context.Context, ids []int64) error {
	return l.db.Delete(&entity.Logstash{}, ids).Error
}
