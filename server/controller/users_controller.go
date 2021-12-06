package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/starship-cloud/starship-iac/server/core/db"
	"github.com/starship-cloud/starship-iac/server/events"
	"github.com/starship-cloud/starship-iac/server/events/models"
	"github.com/starship-cloud/starship-iac/server/logging"
	"github.com/starship-cloud/starship-iac/server/services"
)


type UserResp struct {
	StatusCode  uint
	Description string
	models.UserEntity
}

type userEntities []*models.UserEntity
type UsersResp struct {
	StatusCode  uint
	Description string
	userEntities
}

type UsersController struct {
	Logger  logging.SimpleLogging
	Drainer *events.Drainer
	DB      *db.MongoDB
}

type deleteUsersResp struct {
	StatusCode bool `json:"status_code"`
}

func (uc *UsersController) Get(ctx iris.Context) {
	var userReq models.UserEntity
	ctx.ReadJSON(&userReq)

	userId := ctx.Params().Get("userid")

	result, err := users_service.GetUser(userId, uc.DB)
	if err != nil {
		uc.Logger.Err(err.Error())
		ctx.JSON(&UserResp{
			StatusCode:  iris.StatusInternalServerError,
			Description: err.Error(),
			UserEntity:  models.UserEntity{},
		})
	} else {
		if result != nil {
			ctx.JSON(&UserResp{
				StatusCode:  iris.StatusOK,
				Description: "found",
				UserEntity:  *result,
			})
		}else{
			ctx.JSON(&UserResp{
				StatusCode:  iris.StatusNotFound,
				Description: "Not found",
				UserEntity:  models.UserEntity{Userid: userReq.Userid},
			})
		}
	}
}

func (uc *UsersController) Create(ctx iris.Context) {
	var userReq models.UserEntity
	ctx.ReadJSON(&userReq)
	result, err := users_service.CreateUser(&userReq, uc.DB)
	if err != nil {
		uc.Logger.Err(err.Error())
		ctx.JSON(&UserResp{
			StatusCode:  iris.StatusInternalServerError,
			Description: err.Error(),
			UserEntity:  models.UserEntity{},
		})
	} else {
		ctx.JSON(&UserResp{
			StatusCode:  iris.StatusOK,
			Description: "created",
			UserEntity:  *result,
		})
	}
}

func (uc *UsersController) Delete(ctx iris.Context) {
	var userReq models.UserEntity
	ctx.ReadJSON(&userReq)
	_, err := users_service.DeleteUser(&userReq, uc.DB)
	if err != nil {
		uc.Logger.Err(err.Error())
		ctx.JSON(&UserResp{
			StatusCode:  iris.StatusInternalServerError,
			Description: err.Error(),
			UserEntity:  models.UserEntity{Userid: userReq.Userid},
		})
	} else {
		ctx.JSON(&UserResp{
			StatusCode:  iris.StatusOK,
			Description: "deleted",
			UserEntity:  models.UserEntity{Userid: userReq.Userid},
		})
	}
}

func (uc *UsersController) Update(ctx iris.Context) {
	var userReq models.UserEntity
	ctx.ReadJSON(&userReq)
	_, err := users_service.UpdateUser(&userReq, uc.DB)
	if err != nil {
		uc.Logger.Err(err.Error())
		ctx.JSON(&UserResp{
			StatusCode:  iris.StatusInternalServerError,
			Description: err.Error(),
			UserEntity:  models.UserEntity{Userid: userReq.Userid},
		})
	} else {
		ctx.JSON(&UserResp{
			StatusCode:  iris.StatusOK,
			Description: "updated",
			UserEntity:  models.UserEntity{Userid: userReq.Userid},
		})
	}
}

func (uc *UsersController) Search(ctx iris.Context) {
	var userReq models.UserEntity
	ctx.ReadJSON(&userReq)

	userName := ctx.Params().Get("username")

	result, err := users_service.SearchUsers(userName, uc.DB)
	if err != nil {
		uc.Logger.Err(err.Error())
		ctx.JSON(&UserResp{
			StatusCode:  iris.StatusInternalServerError,
			Description: err.Error(),
			UserEntity:  models.UserEntity{},
		})
	} else {
		if result != nil {
			ctx.JSON(&UsersResp{
				StatusCode:  iris.StatusOK,
				Description: "found",
				userEntities:  result,
			})
		}else{
			ctx.JSON(&UserResp{
				StatusCode:  iris.StatusNotFound,
				Description: "Not found",
				UserEntity:  models.UserEntity{Userid: userReq.Username},
			})
		}
	}
}





