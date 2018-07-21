package service_transfer

import (
	"FKTrojan/dao"
	"testing"
)

func TestServerHandler(t *testing.T) {
	go ListenAndServe()
	TransFile("127.0.0.1", 7779, dao.TestTransferCmdPositive(dao.NewClient().GUID))
}
