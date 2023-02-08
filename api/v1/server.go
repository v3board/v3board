package v1

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"v3board/global"
	"v3board/models"
	"v3board/service"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type serverApi struct{}

var Server = serverApi{}

// User 获取用户列表
func (*serverApi) User(c *gin.Context) {}

// Config 获取节点配置
func (*serverApi) Config(c *gin.Context) {}

// Push 上报节点信息
func (*serverApi) Push(c *gin.Context) {
	var req models.ServerUniversalQueryReq

	if err := c.ShouldBindQuery(&req); err != nil {
		global.Log.Err(err).Msg("node request failed")
		c.String(http.StatusInternalServerError, "query param error")
		return
	}

	// TODO 验证token
	body := c.Request.Body
	defer body.Close()

	// 读取数据
	bodyByte, err := io.ReadAll(body)
	if err != nil {
		global.Log.Err(err).Msg("request body read failed")
		c.String(http.StatusInternalServerError, "invalid data")
		return
	}

	// 检查内容格式
	if !gjson.ValidBytes(bodyByte) {
		global.Log.Err(err).Msg("json parse failed")
		c.String(http.StatusInternalServerError, "invalid json data")
		return
	}

	// 解析数据
	data, ok := gjson.ParseBytes(bodyByte).Value().(map[string]interface{})
	if !ok {
		global.Log.Err(err).Msg("content parse failed")
		c.String(http.StatusInternalServerError, "invalid content")
		return
	}

	// 数据类型转换
	newData := make(map[int64][]int64)
	var err1 error
	for k, v := range data {
		ki, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			err1 = err
		}
		temp := make([]int64, 2)
		for k1, v1 := range v.([]interface{}) {
			//temp = append(temp, int64(v1.(float64)))
			temp[k1] = int64(v1.(float64))
		}

		newData[ki] = temp
	}

	if err1 != nil {
		global.Log.Err(err).Msg("type convert failed")
		c.String(http.StatusInternalServerError, "invalid content")
		return
	}

	for u, d := range newData {
		if err := service.Server.StateUpdate(context.Background(), u, req.NodeId, req.NodeType, d[0], d[1]); err != nil {
			fmt.Println("update error: ", err)
		}
	}

	c.String(http.StatusOK, "success")
}
