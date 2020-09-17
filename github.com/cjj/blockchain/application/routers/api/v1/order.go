package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	bc "github.com/cjj/blockchain/application/blockchain"
	"github.com/cjj/blockchain/application/pkg/app"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	OrderId                   string  `json:"orderId"`
	StoreUserAccountId        string  `json:"storeUserAccountId"`
	ManufacturerUserAccountId string  `json:"manufacturerUserAccountId"`
	LogisticUserAccountId     string  `json:"logisticUserAccountId"`
	OrderName                 string  `json:"orderName"`
	OrderTime                 string  `json:"orderTime"`
	OrderPrice                float64 `json:"orderPrice"`
	TraceId                   string  `json:"traceId"`
	DepartAddress             string  `json:"departAddress"`
	ArrivalAddress            string  `json:"arrivalAddress"`
}

type UpdateOrderRequestBody struct {
	OrderId        string `json:"orderId"`
	MUserAccountId string `json:"manufacturerUserAccountId"`
	SUserAccountId string `json:"storeUserAccountId"`
	Status         string `json:"status"`
}

type OrderListQueryRequestBody struct {
	SUserAccountId string `json:"storeUserAccountId"`
	OrderId        string `json:"orderId"`
}

// @Summary 发起Order
// @Param Order body CreateOrderRequestBody true "Order"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/createOrder [post]
func CreateOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CreateOrderRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	//format data
	formattedOrderPrice := strconv.FormatFloat(body.OrderPrice, 'E', -1, 64)
	if body.OrderId == "" || body.StoreUserAccountId == "" || body.ManufacturerUserAccountId == "" || body.LogisticUserAccountId == "" || body.OrderName == "" || body.OrderTime == "" || formattedOrderPrice == "" || body.TraceId == "" || body.DepartAddress == "" || body.ArrivalAddress == "" {
		appG.Response(http.StatusBadRequest, "失败", "不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.OrderId))
	bodyBytes = append(bodyBytes, []byte(body.StoreUserAccountId))
	bodyBytes = append(bodyBytes, []byte(body.ManufacturerUserAccountId))
	bodyBytes = append(bodyBytes, []byte(body.LogisticUserAccountId))
	bodyBytes = append(bodyBytes, []byte(body.OrderName))
	bodyBytes = append(bodyBytes, []byte(body.OrderTime))
	bodyBytes = append(bodyBytes, []byte(formattedOrderPrice))
	bodyBytes = append(bodyBytes, []byte(body.TraceId))
	bodyBytes = append(bodyBytes, []byte(body.DepartAddress))
	bodyBytes = append(bodyBytes, []byte(body.ArrivalAddress))
	//调用智能合约
	resp, err := bc.ChannelExecute("createOrder", bodyBytes)
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

// @Summary 获取订单信息

func QueryOrderList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(OrderListQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.SUserAccountId != "" {
		bodyBytes = append(bodyBytes, []byte(body.SUserAccountId))

	}
	if body.OrderId != "" {
		bodyBytes = append(bodyBytes, []byte(body.OrderId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryOrderList", bodyBytes)
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

// @Summary Update Order
// @Param Order body UpdateOrderRequestBody true "Order"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/createOrder [post]
func UpdateOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateOrderRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.OrderId == "" || body.SUserAccountId == "" || body.MUserAccountId == "" || body.Status == "" {
		appG.Response(http.StatusBadRequest, "失败", "不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.OrderId))
	bodyBytes = append(bodyBytes, []byte(body.MUserAccountId))
	bodyBytes = append(bodyBytes, []byte(body.SUserAccountId))
	bodyBytes = append(bodyBytes, []byte(body.Status))
	//调用智能合约
	resp, err := bc.ChannelExecute("updateOrder", bodyBytes)
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
