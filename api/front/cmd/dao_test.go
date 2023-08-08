package cmd

import (
	"bibirt-api/api/front/dao"
	"sync"
	"testing"
)

func TestNewTmpUser(t *testing.T) {
	dao.InitDB()
	user := dao.NewTmpUser()
	t.Error(user.ID)
}

var once = sync.Once{}

func DoOnce() {
	once.Do(dao.InitDB)
}

func BenchmarkNewTmpUser(b *testing.B) {
	DoOnce()
	for i := 0; i < 100; i++ {
		dao.NewTmpUser()
	}
}
