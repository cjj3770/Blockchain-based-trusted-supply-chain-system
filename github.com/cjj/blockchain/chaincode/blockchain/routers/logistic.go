package routers

import (
	"encoding/json"
	"fmt"

	"github.com/cjj/blockchain/chaincode/blockchain/lib"
	"github.com/cjj/blockchain/chaincode/blockchain/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//发起物流
func CreateLogistic(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 6 {
		return shim.Error("参数个数不满足")
	}
	orderId := args[0]
	storeUserAccountId := args[1]
	manufacturerUserAccountId := args[2]
	logisticId := args[3]
	courierAccountId := args[4]
	departTime := args[5]
	if orderId == "" || storeUserAccountId == "" || manufacturerUserAccountId == "" || logisticId == "" || courierAccountId == "" || departTime == "" {
		return shim.Error("参数存在空值")
	}
	//根据logisticId获取物流信息，确认不存在该物流
	resultsLogistic, err := utils.GetStateByPartialCompositeKeys(stub, lib.LogisticKey, []string{logisticId})
	if len(resultsLogistic) != 0 {
		return shim.Error(fmt.Sprintf("存在该物流: %s", logisticId))
	}
	//判断用户账户是否具有权限,根据courierAccountId获取账户信息
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{courierAccountId})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("courierAccountId获取账户信息失败%s", err))
	}
	var logisticAccount lib.UserAccount
	if err = json.Unmarshal(resultsAccount[0], &logisticAccount); err != nil {
		return shim.Error(fmt.Sprintf("查询courierAccountId获取物流公司信息-反序列化出错: %s", err))
	}
	// if logisticAccount.UserType != "courier" {
	// 	return shim.Error(fmt.Sprintf("非物流公司courier,administrator不能操作%s", err))
	// }
	//判断订单是否存在
	resultsOrder, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey, []string{storeUserAccountId, orderId, manufacturerUserAccountId})
	if err != nil || len(resultsOrder) != 1 {
		return shim.Error(fmt.Sprintf("订单信息验证失败%s", err))
	}
	var order lib.Order
	if err = json.Unmarshal(resultsOrder[0], &order); err != nil {
		return shim.Error(fmt.Sprintf("订单-反序列化出错: %s", err))
	}
	//判断订单状态是否已发单
	if order.OrderStatus != "orderStart" {
		return shim.Error("此订单已发单")
	}
	createLogistic := &lib.Logistic{
		OrderId:                   orderId,
		StoreUserAccountId:        storeUserAccountId,
		ManufacturerUserAccountId: manufacturerUserAccountId,
		LogisticId:                logisticId,
		CourierAccountId:          courierAccountId,
		DepartTime:                departTime,
		DepartAddress:             order.DepartAddress,
		ArrivalAddress:            order.ArrivalAddress,
		LogisticStatus:            "start",
	}
	// 写入账本
	if err := utils.WriteLedger(createLogistic, stub, lib.LogisticKey, []string{createLogistic.LogisticId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	createLogisticByte, err := json.Marshal(createLogistic)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	//更新订单状态
	order.OrderStatus = "delivery"
	//后将更新的订单写入账本
	if err := utils.WriteLedger(order, stub, lib.OrderKey, []string{order.StoreUserAccountId, order.OrderId, order.ManufacturerUserAccountId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err != nil {
		return shim.Error(fmt.Sprintf("将更新的订单写入账本的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(createLogisticByte)

}

//查询物流(可查询所有，也可根据物流Id查询)
func QueryLogisticList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var logisticList []lib.Logistic
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.LogisticKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var logistic lib.Logistic
			err := json.Unmarshal(v, &logistic)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryLogisticList-反序列化出错: %s", err))
			}
			logisticList = append(logisticList, logistic)
		}
	}
	logisticListByte, err := json.Marshal(logisticList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryLogisticList-序列化出错: %s", err))
	}
	return shim.Success(logisticListByte)
}

// 更新物流信息(delivery,done,"")
func UpdateLogistic(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 6 {
		return shim.Error("参数个数不满足")
	}
	courierAccountId := args[0]
	logisticId := args[1]
	viaAddress := args[2]
	viaTime := args[3]
	status := args[4]
	arrivalTime := args[5]
	if courierAccountId == "" || logisticId == "" {
		return shim.Error("courierAccountId,logisticId存在空值")
	}
	//判断用户账户是否具有权限,根据courierAccountId获取账户信息
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{courierAccountId})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("courierAccountId获取账户信息失败%s", err))
	}
	var logisticAccount lib.UserAccount
	if err = json.Unmarshal(resultsAccount[0], &logisticAccount); err != nil {
		return shim.Error(fmt.Sprintf("courierAccountId获取账户信息-反序列化出错: %s", err))
	}
	// if logisticAccount.UserType != "courier" {
	// 	return shim.Error(fmt.Sprintf("非物流公司courier不能操作%s", err))
	// }

	//根据物流ID判断物流是否存在
	resultsLogistic, err := utils.GetStateByPartialCompositeKeys(stub, lib.LogisticKey, []string{logisticId})
	if err != nil || len(resultsLogistic) != 1 {
		return shim.Error(fmt.Sprintf("物流ID信息验证失败%s", err))
	}
	var updateLogistic lib.Logistic
	if err = json.Unmarshal(resultsLogistic[0], &updateLogistic); err != nil {
		return shim.Error(fmt.Sprintf("获取物流信息-反序列化出错: %s", err))
	}

	//根据orderId判断订单是否存在
	resultsOrder, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey, []string{updateLogistic.StoreUserAccountId, updateLogistic.OrderId, updateLogistic.ManufacturerUserAccountId})
	if err != nil || len(resultsOrder) != 1 {
		return shim.Error(fmt.Sprintf("订单ID信息验证失败%s", err))
	}
	var updateOrder lib.Order
	if err = json.Unmarshal(resultsOrder[0], &updateOrder); err != nil {
		return shim.Error(fmt.Sprintf("获取订单信息-反序列化出错: %s", err))
	}
	switch status {
	case "":
		if updateLogistic.LogisticStatus != "delivery" {
			return shim.Error("物流状态不是delivery")
		}
		if viaAddress == "" || viaTime == "" {
			return shim.Error("参数存在空值")
		}
		//更新物流信息
		updateLogistic.CourierAccountId = courierAccountId
		updateLogistic.ViaAddress = viaAddress
		updateLogistic.ViaTime = viaTime
		//后重新将更新的物流信息写入账本
		if err := utils.WriteLedger(updateLogistic, stub, lib.LogisticKey, []string{updateLogistic.LogisticId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		if err != nil {
			return shim.Error(fmt.Sprintf("重新将更新的订单写入账本的信息出错: %s", err))
		}
		break
	//1.检查物流状态是否是start,2.更新物流状态
	case "delivery":
		if updateLogistic.LogisticStatus != "start" {
			return shim.Error("物流状态不是start")
		}
		//更新物流信息
		updateLogistic.CourierAccountId = courierAccountId
		updateLogistic.ViaAddress = viaAddress
		updateLogistic.ViaTime = viaTime
		updateLogistic.LogisticStatus = status
		//后重新将更新的物流信息写入账本
		if err := utils.WriteLedger(updateLogistic, stub, lib.LogisticKey, []string{updateLogistic.LogisticId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		if err != nil {
			return shim.Error(fmt.Sprintf("重新将更新的订单写入账本的信息出错: %s", err))
		}
		break
	//1.检查物流状态是否是delivery,2.更新物流状态,3.更新订单状态
	case "done":
		if updateLogistic.LogisticStatus != "delivery" {
			return shim.Error("物流状态不是delivery")
		}
		if arrivalTime == "" {
			return shim.Error("缺少到达时间")
		}
		//更新物流信息
		updateLogistic.CourierAccountId = courierAccountId
		updateLogistic.ArrivalTime = arrivalTime
		updateLogistic.LogisticStatus = status
		//后重新将更新的物流信息写入账本
		if err := utils.WriteLedger(updateLogistic, stub, lib.LogisticKey, []string{updateLogistic.LogisticId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		if err != nil {
			return shim.Error(fmt.Sprintf("重新将更新的订单写入账本的信息出错: %s", err))
		}
		//更新订单状态
		updateOrder.OrderStatus = "arrived"

		//将更新的订单写入账本
		if err := utils.WriteLedger(updateOrder, stub, lib.OrderKey, []string{updateOrder.StoreUserAccountId, updateOrder.OrderId, updateOrder.ManufacturerUserAccountId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		if err != nil {
			return shim.Error(fmt.Sprintf("重新将更新的订单写入账本的信息出错: %s", err))
		}
		break
	default:
		return shim.Error(fmt.Sprintf("%s状态不支持", status))
	}
	//将成功的信息返回
	updateLogisticByte, err := json.Marshal(updateLogistic)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化updateLogistic成功的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(updateLogisticByte)
}

//根据参与物流(订单OrderId)查询物流(参与的)(供生产商、门店查询)
// func QueryLogisticListByOther(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	if len(args) != 1 {
// 		return shim.Error(fmt.Sprintf("必须指定订单OrderId查询"))
// 	}
// 	var logisticOtherList []lib.Logistic
// 	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.LogisticKey, args)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("%s", err))
// 	}
// 	for _, v := range results {
// 		if v != nil {
// 			var logisticOther lib.Logistic
// 			err := json.Unmarshal(v, &logisticOther)
// 			if err != nil {
// 				return shim.Error(fmt.Sprintf("QueryLogisticOtherList-反序列化出错: %s", err))
// 			}
// 			logisticOtherList = append(logisticOtherList, logisticOther)
// 		}
// 	}
// 	logisticOtherListByte, err := json.Marshal(logisticOtherList)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("QueryLogisticOtherListByOther-序列化出错: %s", err))
// 	}
// 	return shim.Success(logisticOtherListByte)
// }
