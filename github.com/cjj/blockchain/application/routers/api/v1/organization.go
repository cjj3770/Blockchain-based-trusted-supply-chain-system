package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bc "github.com/cjj/blockchain/application/blockchain"
	"github.com/cjj/blockchain/application/pkg/app"
	"github.com/gin-gonic/gin"
)

//logistic company
type LogisticComIdBody struct {
	LogisticComId string `json:"Id"`
}

//manufacturer
type ManufacturerIdBody struct {
	ManufacturerId string `json:"Id"`
}

//store
type StoreIdBody struct {
	StoreId string `json:"Id"`
}

func QueryLogisticComInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(LogisticComIdBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.LogisticComId != "" {
		bodyBytes = append(bodyBytes, []byte(body.LogisticComId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryLogisticComInfo", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryManufacturerInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ManufacturerIdBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.ManufacturerId != "" {
		bodyBytes = append(bodyBytes, []byte(body.ManufacturerId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryManufacturerInfo", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryStoreInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(StoreIdBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.StoreId != "" {
		bodyBytes = append(bodyBytes, []byte(body.StoreId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryStoreInfo", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
