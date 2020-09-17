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

type TraceRequestBody struct {
	TraceId       string  `json:"traceId"`       //创世ID
	UserAccountId string  `json:"userAccountId"` //用户账户
	ProductId     string  `json:"productId"`     //商品ID
	ProductName   string  `json:"productName"`   //商品名称
	ProductNumber int     `json:"productNumber"` //商品数量
	ProductPrice  float64 `json:"productPrice"`  //商品价格
	ProductTime   string  `json:"productTime"`   //生产日期
	RawIds        string  `json:"rawIds"`        //原材料
}
type TraceListQueryRequestBody struct {
	UserAccountId string `json:"userAccountId"` //用户账户
	TraceId       string `json:"traceId"`       //创世ID
}

// @Summary 发起Trace
// @Param Trace body TraceRequestBody true "Trace"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/createOrder [post]
func CreateTrace(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TraceRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	//format data
	formattedProductNumber := strconv.Itoa(body.ProductNumber)
	formattedProductPrice := strconv.FormatFloat(body.ProductPrice, 'E', -1, 64)
	if body.TraceId == "" || body.UserAccountId == "" || body.ProductId == "" || body.ProductName == "" || formattedProductNumber == "" || formattedProductPrice == "" || body.ProductTime == "" || body.RawIds == "" {
		appG.Response(http.StatusBadRequest, "失败", "不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.TraceId))
	bodyBytes = append(bodyBytes, []byte(body.UserAccountId))
	bodyBytes = append(bodyBytes, []byte(body.ProductId))
	bodyBytes = append(bodyBytes, []byte(body.ProductName))
	bodyBytes = append(bodyBytes, []byte(formattedProductNumber))
	bodyBytes = append(bodyBytes, []byte(formattedProductPrice))
	bodyBytes = append(bodyBytes, []byte(body.ProductTime))
	bodyBytes = append(bodyBytes, []byte(body.RawIds))
	//调用智能合约
	resp, err := bc.ChannelExecute("createTrace", bodyBytes)
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

// @Summary 获取Trace信息

func QueryTraceList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TraceListQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.UserAccountId != "" {
		bodyBytes = append(bodyBytes, []byte(body.UserAccountId))
	}
	if body.TraceId != "" {
		bodyBytes = append(bodyBytes, []byte(body.TraceId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryTraceList", bodyBytes)
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
