package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	//"strconv"

	bc "github.com/cjj/blockchain/application/blockchain"
	"github.com/cjj/blockchain/application/pkg/app"
	"github.com/gin-gonic/gin"
)

type CreateLogisticRequestBody struct {
	OrderId                   string `json:"orderId"`
	StoreUserAccountId        string `json:"storeUserAccountId"`
	ManufacturerUserAccountId string `json:"manufacturerUserAccountId"`
	LogisticId                string `json:"logisticId"`
	CourierAccountId          string `json:"courierAccountId"`
	DepartTime                string `json:"departTime"`
	LogisticStatus            string `json:"logisticStatus"`
}

type UpdateLogisticRequestBody struct {
	CourierAccountId string `json:"courierAccountId"`
	LogisticId       string `json:"logisticId"`
	ViaAddress       string `json:"viaAddress"`
	ViaTime          string `json:"viaTime"`
	Status           string `json:"status"`
	ArrivalTime      string `json:"arrivalTime"`
}

type LogisticListQueryRequestBody struct {
	LogisticId string `json:"logisticId"`
}

// @Summary Create Logistic
// @Param Logistic body CreateLogisticRequestBody true "Logistic"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/createLogistic [post]
func CreateLogistic(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CreateLogisticRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.OrderId == "" || body.StoreUserAccountId == "" || body.ManufacturerUserAccountId == "" || body.LogisticId == "" || body.CourierAccountId == "" || body.DepartTime == "" {
		appG.Response(http.StatusBadRequest, "失败", "不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.OrderId))
	bodyBytes = append(bodyBytes, []byte(body.StoreUserAccountId))
	bodyBytes = append(bodyBytes, []byte(body.ManufacturerUserAccountId))
	bodyBytes = append(bodyBytes, []byte(body.LogisticId))
	bodyBytes = append(bodyBytes, []byte(body.CourierAccountId))
	bodyBytes = append(bodyBytes, []byte(body.DepartTime))
	//调用智能合约
	resp, err := bc.ChannelExecute("createLogistic", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 获取物流信息

func QueryLogisticList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(LogisticListQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.LogisticId != "" {
		bodyBytes = append(bodyBytes, []byte(body.LogisticId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryLogisticList", bodyBytes)
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

// @Summary Update Logistic
// @Param Logistic body UpdateLogisticRequestBody true "Logistic"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/updateLogistic [post]
func UpdateLogistic(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateLogisticRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.CourierAccountId == "" || body.LogisticId == "" {
		appG.Response(http.StatusBadRequest, "失败", "不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.CourierAccountId))
	bodyBytes = append(bodyBytes, []byte(body.LogisticId))
	bodyBytes = append(bodyBytes, []byte(body.ViaAddress))
	bodyBytes = append(bodyBytes, []byte(body.ViaTime))
	bodyBytes = append(bodyBytes, []byte(body.Status))
	bodyBytes = append(bodyBytes, []byte(body.ArrivalTime))
	//调用智能合约
	resp, err := bc.ChannelExecute("updateLogistic", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
