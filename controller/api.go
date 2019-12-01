package controller

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuzhikuan/user_manage/dao"
	"github.com/yuzhikuan/user_manage/dto"
	"github.com/yuzhikuan/user_manage/middleware"
)

type Api struct {
}

func (o *Api) Login(c *gin.Context) {
	api := &dto.LoginInput{}
	if err := api.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	// todo 这块应该从db验证
	if api.Username == "admin" && api.Password == "123456" {
		session := sessions.Default(c)
		session.Set("user", api.Username)
		session.Save()
		middleware.ResponseSuccess(c, api)
	} else {
		middleware.ResponseError(c, 2002, errors.New("账号或密码错误"))
	}
	return
}

func (o *Api) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	middleware.ResponseSuccess(c, nil)
	return
}

func (o *Api) ListPage(c *gin.Context) {
	listInput := &dto.ListPageInput{}
	if err := listInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	user := &dao.User{}
	pageInt, err1 := strconv.ParseInt(listInput.Page, 10, 64)
	sizeInt, err2 := strconv.ParseInt(listInput.Size, 10, 64)
	if err1 != nil || err2 != nil {
		middleware.ResponseError(c, 2004, errors.New("page or size error."))
		return
	}
	if userList, total, err := user.PageList(int(pageInt), int(sizeInt)); err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	} else {
		m := map[string]interface{}{
			"list":  userList,
			"total": total,
		}
		middleware.ResponseSuccess(c, m)
	}
	return
}

func (o *Api) AddUser(c *gin.Context) {
	addInput := &dto.AddUserInput{}
	if err := addInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2006, err)
		return
	}
	user := &dao.User{}
	user.Name = addInput.Name
	user.Sex = addInput.Sex
	user.Age = addInput.Age
	user.Birth = addInput.Birth
	user.Addr = addInput.Addr
	user.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	if err := user.Save(); err != nil {
		middleware.ResponseError(c, 2007, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (o *Api) EditUser(c *gin.Context) {
	editInput := &dto.EditUserInput{}
	if err := editInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2006, err)
		return
	}
	user := &dao.User{}
	if userDb, err := user.Find(int64(editInput.Id)); err != nil {
		middleware.ResponseError(c, 2006, err)
		return
	} else {
		user = userDb
	}
	user.Name = editInput.Name
	user.Sex = editInput.Sex
	user.Age = editInput.Age
	user.Birth = editInput.Birth
	user.Addr = editInput.Addr
	user.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	if err := user.Save(); err != nil {
		middleware.ResponseError(c, 2007, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (o *Api) RemoveUser(c *gin.Context) {
	removeInput := &dto.RemoveUserInput{}
	if err := removeInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2006, err)
		return
	}
	user := &dao.User{}
	if err := user.Del(strings.Split(removeInput.IDS, ",")); err != nil {
		middleware.ResponseError(c, 2007, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}
