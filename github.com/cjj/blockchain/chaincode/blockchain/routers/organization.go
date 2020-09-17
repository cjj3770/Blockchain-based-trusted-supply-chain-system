package routers

import (
	"encoding/json"
	"fmt"

	"github.com/cjj/blockchain/chaincode/blockchain/lib"
	"github.com/cjj/blockchain/chaincode/blockchain/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//查询物流公司
func QueryLogisticComInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var logisticComList []lib.LogisticCom
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.LogisticComKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var logisticCom lib.LogisticCom
			err := json.Unmarshal(v, &logisticCom)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryLogisticCom-反序列化出错: %s", err))
			}
			logisticComList = append(logisticComList, logisticCom)
		}
	}
	logisticComListByte, err := json.Marshal(logisticComList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryLogisticCom-序列化出错: %s", err))
	}
	return shim.Success(logisticComListByte)
}

//查询查询生产商列表
func QueryManufacturerInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var manufacturerList []lib.Manufacturer
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.ManufacturerKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var manufacturer lib.Manufacturer
			err := json.Unmarshal(v, &manufacturer)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryManufacturerList-反序列化出错: %s", err))
			}
			manufacturerList = append(manufacturerList, manufacturer)
		}
	}
	manufacturerListByte, err := json.Marshal(manufacturerList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryManufacturerList-序列化出错: %s", err))
	}
	return shim.Success(manufacturerListByte)
}

//查询查询门店列表
func QueryStoreInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var storeList []lib.Store
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.StoreKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var store lib.Store
			err := json.Unmarshal(v, &store)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryStoreList-反序列化出错: %s", err))
			}
			storeList = append(storeList, store)
		}
	}
	storeListByte, err := json.Marshal(storeList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryStoreList-序列化出错: %s", err))
	}
	return shim.Success(storeListByte)
}
