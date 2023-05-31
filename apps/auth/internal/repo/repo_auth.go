package repo

import (
	"github.com/marsxingzhi/goim/domain/model"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type AuthRepo interface {
	TxCreate(tx *gorm.DB, user *model.User) (err error)
}

type authRepo struct {
}

func NewAuthRepo() AuthRepo {
	return &authRepo{}
}

func (ar *authRepo) TxCreate(tx *gorm.DB, user *model.User) (err error) {
	// 这里才生成uid和aid，前面如果生成了，但是在入库的过程中失败了，那么就是白生成，意义不大
	// TODO-xz 下面两个使用雪花算法
	user.Uid = time.Now().Unix()
	user.Aid = strconv.FormatInt(time.Now().Unix(), 10)

	return tx.Create(user).Error
}
