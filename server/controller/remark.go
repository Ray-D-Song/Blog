package controller

import (
	"server/ent"
	"server/service"
	"server/util/code"
	"server/util/rest"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRemarkList(c *gin.Context) {
	list, err := service.GetRemarkList(c)
	if err != nil {
		rest.Res(c, code.Error, nil)
		return
	}
	rest.Res(c, code.OK, list)
}

func AddRemark(c *gin.Context) {
	var remark ent.Remark
	if err := c.ShouldBindJSON(&remark); err != nil {
		rest.Res(c, code.Error, nil)
		return
	}
	if _, err := service.AddRemark(c, remark); err != nil {
		rest.Res(c, code.Error, nil)
		return
	}
	rest.Res(c, code.OK, nil)
}

func DelRemark(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		rest.Res(c, code.Error, nil)
	}
	if _, err := service.DelRemark(c, id); err != nil {
		rest.Res(c, code.Error, nil)
	}
	rest.Res(c, code.OK, nil)
}
