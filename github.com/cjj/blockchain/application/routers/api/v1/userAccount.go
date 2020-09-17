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

type UserAccountIdBody struct {
	UserAccountId string `json:"userAccountId"`
}

type UserAccountRequestBody struct {
	Args []UserAccountIdBody `json:"args"`
}

type CreateUserAccountRequestBody struct {
	AccountId       string `json:"userAccountId"`
	CreateAccountId string `json:"createUserAccountId"`
	UserName        string `json:"userName"`
	Password        string `json:"password"`
	UserType        string `json:"userType"`
	ServiceId       string `json:"serviceId"`
}

// @Summary 获取账户信息
// @Param account body UserAccountIdBody true "account"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/queryUserAccountList [post]
func QueryUserAccountList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserAccountRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	for _, val := range body.Args {
		bodyBytes = append(bodyBytes, []byte(val.UserAccountId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryUserAccountList", bodyBytes)
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

func QueryUserAccount(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserAccountIdBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.UserAccountId == "" {
		appG.Response(http.StatusBadRequest, "失败", "不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.UserAccountId))
	//调用智能合约
	resp, err := bc.ChannelQuery("queryUserAccountList", bodyBytes)
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

func CreateUserAccount(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CreateUserAccountRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.AccountId == "" || body.CreateAccountId == "" || body.Password == "" || body.UserType == "" || body.ServiceId == "" || body.UserName == "" {
		appG.Response(http.StatusBadRequest, "失败", "不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.AccountId))
	bodyBytes = append(bodyBytes, []byte(body.CreateAccountId))
	bodyBytes = append(bodyBytes, []byte(body.UserName))
	bodyBytes = append(bodyBytes, []byte(body.Password))
	bodyBytes = append(bodyBytes, []byte(body.UserType))
	bodyBytes = append(bodyBytes, []byte(body.ServiceId))
	//调用智能合约
	resp, err := bc.ChannelExecute("createUserAccount", bodyBytes)
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
