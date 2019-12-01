package dto

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/yuzhikuan/user_manage/public"
	"gopkg.in/go-playground/validator.v9"
)

func commonPartWithBindValidParams(o interface{}, c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	if err := public.Validate.Struct(o); err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

type LoginInput struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func (o *LoginInput) BindingValidParams(c *gin.Context) error {
	return commonPartWithBindValidParams(o, c)
}

type ListPageInput struct {
	Page string `form:"page" json:"page" validate:"required"`
	Size string `form:"size" json:"size" validate:"required"`
}

func (o *ListPageInput) BindingValidParams(c *gin.Context) error {
	return commonPartWithBindValidParams(o, c)
}

type AddUserInput struct {
	Name  string `form:"name" json:"name" validate:"required"`
	Sex   int    `form:"sex" json:"sex" validate:""`
	Age   int    `form:"age" json:"age" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" validate:"required"`
	Addr  string `form:"addr" json:"addr" validate:"required"`
}

func (o *AddUserInput) BindingValidParams(c *gin.Context) error {
	return commonPartWithBindValidParams(o, c)
}

type EditUserInput struct {
	Id    int    `form:"id" json:"id" validate:"required"`
	Name  string `form:"name" json:"name" validate:"required"`
	Sex   int    `form:"sex" json:"sex" validate:""`
	Age   int    `form:"age" json:"age" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" validate:"required"`
	Addr  string `form:"addr" json:"addr" validate:"required"`
}

func (o *EditUserInput) BindingValidParams(c *gin.Context) error {
	return commonPartWithBindValidParams(o, c)
}

type RemoveUserInput struct {
	IDS string `form:"ids" json:"ids" validate:"required"`
}

func (o *RemoveUserInput) BindingValidParams(c *gin.Context) error {
	return commonPartWithBindValidParams(o, c)
}
