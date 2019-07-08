package service

import (
	"fmt"
	// 公共引入
	"github.com/micro/go-log"

	pb "github.com/gomsa/socialite/proto/user"

	"github.com/jinzhu/gorm"
)

//URepository 仓库接口
type URepository interface {
	GetByID(req *pb.User) ([]*pb.User, error)
	Exist(user *pb.User) bool
	Create(user *pb.User) (*pb.User, error)
	Get(user *pb.User) (*pb.User, error)
	Update(user *pb.User) (bool, error)
	Delete(user *pb.User) (bool, error)
}

// UserRepository 用户仓库
type UserRepository struct {
	DB *gorm.DB
}

// GetByID 根据 id 获取绑定信息
func (repo *UserRepository) GetByID(req *pb.User) (User []*pb.User, err error) {
	if req.Id != "" {
		if err := repo.DB.Model(&User).Where("id = ?", req.Id).Find(&User).Error; err != nil {
			return nil, err
		}
	}
	return User, nil
}

// Exist 检测用户是否已经存在
func (repo *UserRepository) Exist(user *pb.User) bool {
	var count int
	if user.Id != "" {
		repo.DB.Model(&user).Where("id = ?", user.Id).Count(&count)
		if count > 0 {
			return true
		}
	}
	if user.Origin != "" {
		repo.DB.Model(&user).Where("origin = ?", user.Origin).Where("openid = ?", user.Openid).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// Get 获取用户信息
func (repo *UserRepository) Get(user *pb.User) (*pb.User, error) {
	if user.Id != "" {
		if err := repo.DB.Model(&user).Where("id = ?", user.Id).Find(&user).Error; err != nil {
			return nil, err
		}
	}
	if user.Openid != "" {
		if err := repo.DB.Model(&user).Where("origin = ?", user.Origin).Where("openid = ?", user.Openid).Find(&user).Error; err != nil {
			return nil, err
		}
	}
	return user, nil
}

// Create 创建用户
// bug 无用户名创建用户可能引起 bug
func (repo *UserRepository) Create(user *pb.User) (*pb.User, error) {
	if exist := repo.Exist(user); exist == true {
		return user, fmt.Errorf("注册用户已存在")
	}
	err := repo.DB.Create(user).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return user, fmt.Errorf("注册用户失败")
	}
	return user, nil
}

// Update 更新用户
func (repo *UserRepository) Update(user *pb.User) (bool, error) {
	if user.Id == "" {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.User{
		Id: user.Id,
	}
	err := repo.DB.Model(id).Updates(user).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除用户
func (repo *UserRepository) Delete(user *pb.User) (bool, error) {
	id := &pb.User{
		Id: user.Id,
	}
	err := repo.DB.Delete(id).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
