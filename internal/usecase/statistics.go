package usecase

import (
	"blog/internal/entity"
	"context"
)

type StatisticsUseCase struct {
	repo StatisticsRepo
}

func NewStatisticsUseCase(repo StatisticsRepo) *StatisticsUseCase {
	return &StatisticsUseCase{
		repo: repo,
	}
}

// GetDashboardStatistics 获取仪表盘统计数据
func (uc *StatisticsUseCase) GetDashboardStatistics(ctx context.Context) (*entity.DashboardStatistics, error) {
	users, err := uc.repo.GetUsersCountRepo(ctx)
	if err != nil {
		return nil, err
	}

	articles, err := uc.repo.GetArticlesCountRepo(ctx)
	if err != nil {
		return nil, err
	}

	comments, err := uc.repo.GetCommentsCountRepo(ctx)
	if err != nil {
		return nil, err
	}

	visits, err := uc.repo.GetVisitsCountRepo(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.DashboardStatistics{
		Users:    users,
		Articles: articles,
		Comments: comments,
		Visits:   visits,
	}, nil
}
