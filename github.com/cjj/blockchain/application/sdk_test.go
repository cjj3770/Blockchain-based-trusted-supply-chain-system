package main_test

import (
	"fmt"
	"testing"

	"github.com/togettoyou/blockchain-real-estate/application/blockchain"

	"bytes"
	"encoding/json"

	//"net/http"

	//"github.com/gin-gonic/gin"
	bc "github.com/togettoyou/blockchain-real-estate/application/blockchain"
	//"github.com/togettoyou/blockchain-real-estate/application/pkg/app"
)

// func TestInvoke_QueryAccountList(t *testing.T) {
// 	blockchain.Init()
// 	response, e := blockchain.ChannelQuery("queryAccountList", [][]byte{})
// 	if e != nil {
// 		fmt.Println(e.Error())
// 		t.FailNow()
// 	}
// 	fmt.Println(string(response.Payload))
// }

type AccountIdBody struct {
	AccountId string `json:"accountId"`
}

type AccountRequestBody struct {
	Args []AccountIdBody `json:"args"`
}

// @Summary 获取账户信息
// @Param account body AccountRequestBody true "account"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/queryAccountList [post]
func TestInvoke_QueryAccountList(t *testing.T) {
	blockchain.Init()
	// 	response, e := blockchain.ChannelQuery("queryAccountList", [][]byte{})
	//appG := app.Gin{C: c}
	//body := new(AccountRequestBody)

	// accountId := "5feceb66ffc8"
	// accountIdBody := AccountIdBody{accountId}

	// args := []AccountIdBody{accountIdBody}

	// body := AccountRequestBody{Args: args}

	// //解析Body参数
	// if err := t.ShouldBind(body); err != nil {
	// 	//appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
	// 	fmt.Sprintf("参数出错%s", err.Error())
	// 	return
	// }
	// var bodyBytes [][]byte
	// for _, val := range body.Args {
	// 	bodyBytes = append(bodyBytes, []byte(val.AccountId))
	// }

	//1.创建json数据
	b := []byte(`{"args":[{"accountId":"5feceb66ffc8"}]}`)
	//2.声明结构体
	var arb AccountRequestBody
	//3.json解析到结构体
	err := json.Unmarshal(b, &arb)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(arb)
	var bodyBytes [][]byte
	for _, val := range arb.Args {
		bodyBytes = append(bodyBytes, []byte(val.AccountId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryAccountList", bodyBytes)
	if err != nil {
		//appG.Response(http.StatusInternalServerError, "失败", err.Error())
		fmt.Println(err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		//appG.Response(http.StatusInternalServerError, "失败", err.Error())
		fmt.Println(err.Error())
		return
	}
	//appG.Response(http.StatusOK, "成功", data)
	fmt.Println(string(resp.Payload))
}
