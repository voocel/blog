package usecase

import (
	"blog/internal/entity"
	"context"
)

type LogstashUseCase struct {
	repo LogstashRepo
}

func NewLogstashUseCase(repo LogstashRepo) *LogstashUseCase {
	return &LogstashUseCase{repo: repo}
}

func (c *LogstashUseCase) AddLogstash(ctx context.Context, req *entity.Logstash) error {
	return c.repo.AddLogstashRepo(ctx, req)
}

func (c *LogstashUseCase) Detail(ctx context.Context, id int64) (*entity.Logstash, error) {
	return c.repo.GetLogstashByIdRepo(ctx, id)
}

func (c *LogstashUseCase) List(ctx context.Context) ([]*entity.Logstash, error) {
	return c.repo.GetLogstashRepo(ctx)
}

func (c *LogstashUseCase) DeleteLogstashBatch(ctx context.Context, ids []int64) error {
	return c.repo.DeleteLogstashBatchRepo(ctx, ids)
}

func (c *LogstashUseCase) DeleteLogstash(ctx context.Context, id int64) error {
	return c.repo.DeleteLogstashRepo(ctx, id)
}
