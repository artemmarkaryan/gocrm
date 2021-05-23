package service

import (
	"github.com/artemmarkaryan/gocrm/internal"
	"testing"
)

func TestUserService_GetAll(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	//db, err := domain.GetDB()
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//testUser := domain.User{
	//	Name: "test" + strconv.FormatInt(time.Now().Unix(), 10),
	//}
	//db.Create(&testUser).Commit()

	service := UserService{}
	result, err := service.GetAll()
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(string(result))
	}

	//db.Delete(&testUser).Commit()
}
