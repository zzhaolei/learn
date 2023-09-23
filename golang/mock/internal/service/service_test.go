package service

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func TestPeopleSay(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockPeople(ctrl)
	m.EXPECT().Say().Return("这是一个测试")

	SayAnything(m)
}

func TestServerAndDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbName := "测试DB"

	dbMock := NewMockDB(ctrl)
	dbMock.EXPECT().GetName().Return(dbName)

	srvMock := NewMockServer(ctrl)
	srvMock.EXPECT().Get().Return(dbMock)

	db := GetDB(srvMock)
	if db.GetName() != dbName {
		t.Errorf("Expect '%s', got '%s'", dbName, db.GetName())
	}
}
