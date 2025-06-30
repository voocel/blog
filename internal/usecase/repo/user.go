package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"
	"errors"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) usecase.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetUserByIdRepo(ctx context.Context, uid int64) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("id = ?", uid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByNameRepo(ctx context.Context, name string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("username = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByEmailRepo(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByNameExistRepo(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("username = ?", name).Count(&count).Error
	return count > 0, err
}

func (r *userRepo) GetUserByEmailExistRepo(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *userRepo) GetUsersRepo(ctx context.Context, page, pageSize int, search string) ([]*entity.User, int64, error) {
	var users []*entity.User
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.User{})

	if search != "" {
		query = query.Where("username LIKE ? OR email LIKE ? OR nickname LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// 获取总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepo) AddUserRepo(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) UpdateUserRepo(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepo) DeleteUserRepo(ctx context.Context, uid int64) error {
	return r.db.WithContext(ctx).Delete(&entity.User{}, uid).Error
}

func (r *userRepo) UpdateAvatarUserRepo(ctx context.Context, uid int64, avatar string) error {
	return r.db.Model(&entity.User{}).Where("id = ?", uid).Update("avatar", avatar).Error
}

func (r *userRepo) UpdateAddressUserRepo(ctx context.Context, uid int64, address string) error {
	return r.db.Model(&entity.User{}).Where("id = ?", uid).Update("address", address).Error
}
