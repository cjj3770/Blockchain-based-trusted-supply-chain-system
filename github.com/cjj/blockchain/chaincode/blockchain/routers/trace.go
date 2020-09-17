package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	//"time"

	"github.com/cjj/blockchain/chaincode/blockchain/lib"
	"github.com/cjj/blockchain/chaincode/blockchain/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//新建溯源,创世区块(生产商)
func CreateTrace(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 8 {
		return shim.Error("参数个数不满足")
	}
	traceId := args[0]
	userAccountId := args[1] //userAccountId用于验证是否为生产商
	productId := args[2]
	productName := args[3]
	productNumber := args[4]
	productPrice := args[5]
	productTime := args[6]
	rawIds := args[7] //原材料列表
	if userAccountId == "" || productId == "" || productNumber == "" || productName == "" || productPrice == "" || productTime == "" || rawIds == "" {
		return shim.Error("参数存在空值")
	}
	// 参数数据格式转换
	var formattedProductNumber int
	if val, err := strconv.Atoi(productNumber); err != nil {
		return shim.Error(fmt.Sprintf("productNumber参数格式转换出错: %s", err))
	} else {
		formattedProductNumber = val
	}
	var formattedProductPrice float64
	if val, err := strconv.ParseFloat(productPrice, 64); err != nil {
		return shim.Error(fmt.Sprintf("productPrice参数格式转换出错: %s", err))
	} else {
		formattedProductPrice = val
	}
	//判断是否生产商操作
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{userAccountId})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("操作人权限验证失败%s", err))
	}
	var userAccount lib.UserAccount
	if err = json.Unmarshal(resultsAccount[0], &userAccount); err != nil {
		return shim.Error(fmt.Sprintf("查询操作人信息-反序列化出错: %s", err))
	}
	if userAccount.UserType != "manufacturer" {
		return shim.Error(fmt.Sprintf("操作人权限不足%s", err))
	}
	//判断创世区块是否存在
	resultsTrace, err := utils.GetStateByPartialCompositeKeys2(stub, lib.TraceKey, []string{userAccountId, traceId})
	if len(resultsTrace) != 0 {
		return shim.Error(fmt.Sprintf("创世区块traceId存在%s", err))
	}

	trace := &lib.Trace{
		TraceId:       traceId,
		UserAccountId: userAccountId,
		ProductId:     productId,
		ProductNumber: formattedProductNumber,
		ProductName:   productName,
		ProductPrice:  formattedProductPrice,
		ProductTime:   productTime,
		RawIds:        rawIds,
		TraceStatus:   "created",
	}
	// 写入账本
	if err := utils.WriteLedger(trace, stub, lib.TraceKey, []string{trace.UserAccountId, trace.TraceId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	traceByte, err := json.Marshal(trace)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(traceByte)
}

//查询创世区块(可查询所有，也可根据所有人查询名下商品)
func QueryTraceList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var traceList []lib.Trace
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.TraceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var trace lib.Trace
			err := json.Unmarshal(v, &trace)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryTraceList-反序列化出错: %s", err))
			}
			traceList = append(traceList, trace)
		}
	}
	traceListByte, err := json.Marshal(traceList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryTraceList-序列化出错: %s", err))
	}
	return shim.Success(traceListByte)
}
