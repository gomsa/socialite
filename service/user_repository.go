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
	Exist(user *pb.User) bool
	Create(user *pb.User) (*pb.User, error)
	Get(user *pb.User) (*pb.User, error)
	List(req *pb.ListQuery) ([]*pb.User, error)
	Total(req *pb.ListQuery) (int64, error)
	Update(user *pb.User) (bool, error)
	Delete(user *pb.User) (bool, error)
}

// UserRepository 用户仓库
type UserRepository struct {
	DB *gorm.DB
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

// List 获取所有用户信息
func (repo *UserRepository) List(req *pb.ListQuery) (users []*pb.User, err error) {
	db := repo.DB
	// 分页
	var limit, offset int64
	if req.Limit > 0 {
		limit = req.Limit
	} else {
		limit = 10
	}
	if req.Page > 1 {
		offset = (req.Page - 1) * limit
	} else {
		offset = -1
	}

	// 排序
	var sort string
	if req.Sort != "" {
		sort = req.Sort
	} else {
		sort = "created_at desc"
	}
	// 查询条件
	if req.Id != "" {
		db = db.Where("id like ?", "%"+req.Id+"%")
	}
	if req.Openid != "" {
		db = db.Where("openid like ?", "%"+req.Openid+"%")
	}
	if req.Session != "" {
		db = db.Where("session like ?", "%"+req.Session+"%")
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return users, nil
}

// Total 获取所有用户查询总量
func (repo *UserRepository) Total(req *pb.ListQuery) (total int64, err error) {
	users := []pb.User{}
	db := repo.DB
	// 查询条件
	if req.Id != "" {
		db = db.Where("id like ?", "%"+req.Id+"%")
	}
	if req.Openid != "" {
		db = db.Where("openid like ?", "%"+req.Openid+"%")
	}
	if req.Session != "" {
		db = db.Where("session like ?", "%"+req.Session+"%")
	}
	if err := db.Find(&users).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
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
