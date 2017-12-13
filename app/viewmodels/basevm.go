package viewmodels

import (
	"github.com/go-gorp/gorp"
	"time"
)

type BaseVM struct {
	Txn *gorp.Transaction
}

type BaseVMResponse struct {
	IsSuccess bool
	Message string
}

func (vm BaseVM) GetMysqlTime() string{
	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	return datetime.Format("2006-01-02 15:04:05")
}

