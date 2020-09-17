package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cjj/blockchain/chaincode/blockchain/lib"
	"github.com/cjj/blockchain/chaincode/blockchain/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//QueryOrderList
func QueryOrderList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var orderList []lib.Order
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var order lib.Order
			err := json.Unmarshal(v, &order)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryOrderList-反序列化出错: %s", err))
			}
			orderList = append(orderList, order)
		}
	}
	orderListByte, err := json.Marshal(orderList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryOrderList-序列化出错: %s", err))
	}
	return shim.Success(orderListByte)
}

//创建订单
func CreateOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 10 {
		return shim.Error("参数个数不满足")
	}
	orderId := args[0]
	buyerAccountId := args[1]
	sellerAccountId := args[2]
	logisticAccountId := args[3]
	orderName := args[4]
	orderTime := args[5]
	orderPrice := args[6]
	traceId := args[7]
	departAddress := args[8]
	arrivalAddress := args[9]

	if orderId == "" || buyerAccountId == "" || sellerAccountId == "" || logisticAccountId == "" || orderName == "" || orderTime == "" || orderPrice == "" || traceId == "" || departAddress == "" || arrivalAddress == "" {
		return shim.Error("参数存在空值")
	}
	//判断用户账户是否具有权限,根据buyer获取买家信息
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{buyerAccountId})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("buyer买家信息验证失败%s", err))
	}
	var buyerAccount lib.UserAccount
	if err = json.Unmarshal(resultsAccount[0], &buyerAccount); err != nil {
		return shim.Error(fmt.Sprintf("查询buyer买家信息-反序列化出错: %s", err))
	}
	if buyerAccount.UserType != "store" {
		return shim.Error(fmt.Sprintf("非门店不能下订单%s", err))
	}

	//判断创世区块是否存在
	resultsTrace, err := utils.GetStateByPartialCompositeKeys2(stub, lib.TraceKey, []string{sellerAccountId, traceId})
	if err != nil || len(resultsTrace) != 1 {
		return shim.Error(fmt.Sprintf("创世区块traceId信息验证失败%s", err))
	}
	var trace lib.Trace
	if err = json.Unmarshal(resultsTrace[0], &trace); err != nil {
		return shim.Error(fmt.Sprintf("resultsTrace-反序列化出错: %s", err))
	}
	if trace.TraceStatus == "ordered" {
		return shim.Error(fmt.Sprintf("this trace is ordered"))
	}

	//判断订单是否存在
	resultsOrder, err := utils.GetStateByPartialCompositeKeys(stub, lib.OrderKey, []string{buyerAccountId, orderId, sellerAccountId})
	if len(resultsOrder) != 0 {
		return shim.Error(fmt.Sprintf("订单orderId存在%s", err))
	}
	// 参数数据格式转换
	var formattedOrderPrice float64
	if val, err := strconv.ParseFloat(orderPrice, 64); err != nil {
		return shim.Error(fmt.Sprintf("orderPrice参数格式转换出错: %s", err))
	} else {
		formattedOrderPrice = val
	}
	//根据logisticAccountId获取物流公司信息，确认存在该物流公司
	resultsLogisticAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{logisticAccountId})
	if err != nil || len(resultsTrace) != 1 {
		return shim.Error(fmt.Sprintf("物流公司logisticAccountId信息验证失败%s", err))
	}
	var logisticAccount lib.UserAccount
	if err = json.Unmarshal(resultsLogisticAccount[0], &logisticAccount); err != nil {
		return shim.Error(fmt.Sprintf("resultsLogisticAccount-反序列化出错: %s", err))
	}

	createOrder := &lib.Order{
		OrderId:                   orderId,
		StoreUserAccountId:        buyerAccountId,
		ManufacturerUserAccountId: sellerAccountId,
		LogisticUserAccountId:     logisticAccountId,
		OrderName:                 orderName,
		OrderPrice:                formattedOrderPrice,
		OrderTime:                 orderTime,
		TraceId:                   trace.TraceId,
		OrderStatus:               "orderStart",
		DepartAddress:             departAddress,
		ArrivalAddress:            arrivalAddress,
	}
	// 写入账本
	if err := utils.WriteLedger(createOrder, stub, lib.OrderKey, []string{createOrder.StoreUserAccountId, createOrder.OrderId, createOrder.ManufacturerUserAccountId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	trace.TraceStatus = "ordered"
	//更新的trace区块写入账本
	if err := utils.WriteLedger(trace, stub, lib.TraceKey, []string{trace.UserAccountId, trace.TraceId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	createOrderByte, err := json.Marshal(createOrder)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(createOrderByte)
}

// 更新订单状态（manufacturer确认、store取消）
func UpdateOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 4 {
		return shim.Error(fmt.Sprintf("参数个数不满足!!!"))
	}
	orderId := args[0]
	mUserAccountId := args[1]
	sUserAccountId := args[2]
	status := args[3]
	if orderId == "" || mUserAccountId == "" || sUserAccountId == "" || status == "" {
		return shim.Error("参数存在空值")
	}
	if mUserAccountId == sUserAccountId {
		return shim.Error("买家和卖家不能同一人")
	}
	//根据orderId, sUserAccountId, mUserAccountId，确认存在该订单
	resultsOrder, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey, []string{sUserAccountId, orderId, mUserAccountId})
	if err != nil || len(resultsOrder) != 1 {
		return shim.Error(fmt.Sprintf("根据%s获取订单信息失败: %s", orderId, err))
	}
	var order lib.Order
	if err = json.Unmarshal(resultsOrder[0], &order); err != nil {
		return shim.Error(fmt.Sprintf("UpdateOrder-反序列化出错: %s", err))
	}
	//------------------------------------------------------------------------
	//判断输入的商店账户ID是否和订单上的对应
	if order.StoreUserAccountId != sUserAccountId {
		return shim.Error("store userAccountId is different from that in order")
	}

	//判断输入的生产商账户ID是否和订单上的对应
	if order.ManufacturerUserAccountId != mUserAccountId {
		return shim.Error("manufacturer userAccountId is different from that in order")
	}

	//根据sUserAccountId获取store账户信息
	resultssUserAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{sUserAccountId})
	if err != nil || len(resultssUserAccount) != 1 {
		return shim.Error(fmt.Sprintf("store账户信息验证失败%s", err))
	}
	var accountStore lib.UserAccount
	if err = json.Unmarshal(resultssUserAccount[0], &accountStore); err != nil {
		return shim.Error(fmt.Sprintf("查询store账户信息-反序列化出错: %s", err))
	}
	//根据账户服务ID查询store
	resultsStore, err := utils.GetStateByPartialCompositeKeys(stub, lib.StoreKey, []string{accountStore.ServiceId})
	if err != nil || len(resultsStore) != 1 {
		return shim.Error(fmt.Sprintf("store买家信息验证失败%s", err))
	}
	var store lib.Store
	if err = json.Unmarshal(resultsStore[0], &store); err != nil {
		return shim.Error(fmt.Sprintf("查询store买家信息-反序列化出错: %s", err))
	}
	//------------------------------------------------------------------------
	//根据mUserAccountId获取manufacturer账户信息
	resultsmUserAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{mUserAccountId})
	if err != nil || len(resultsmUserAccount) != 1 {
		return shim.Error(fmt.Sprintf("manufacturer账户信息验证失败%s", err))
	}
	var accountManufacturer lib.UserAccount
	if err = json.Unmarshal(resultsmUserAccount[0], &accountManufacturer); err != nil {
		return shim.Error(fmt.Sprintf("查询manufacturer账户信息-反序列化出错: %s", err))
	}
	//根据账户服务ID查询manufacturer
	resultsManufacturer, err := utils.GetStateByPartialCompositeKeys(stub, lib.ManufacturerKey, []string{accountManufacturer.ServiceId})
	if err != nil || len(resultsManufacturer) != 1 {
		return shim.Error(fmt.Sprintf("manufacturer卖家信息验证失败%s,%s", err, accountManufacturer.ServiceId))
	}
	var manufacturer lib.Manufacturer
	if err = json.Unmarshal(resultsManufacturer[0], &manufacturer); err != nil {
		return shim.Error(fmt.Sprintf("查询manufacturer卖家信息-反序列化出错: %s", err))
	}
	//------------------------------------------------------------------------
	resultsLogisticUserAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{order.LogisticUserAccountId})
	if err != nil || len(resultsLogisticUserAccount) != 1 {
		return shim.Error(fmt.Sprintf("根据%s和%s获取订单信息失败: %s", orderId, sUserAccountId, err))
	}
	var accountLogistic lib.UserAccount
	if err = json.Unmarshal(resultsLogisticUserAccount[0], &accountLogistic); err != nil {
		return shim.Error(fmt.Sprintf("查询logistic信息-反序列化出错: %s", err))
	}
	//根据账户服务ID查询logisticCom
	resultsLogisticCom, err := utils.GetStateByPartialCompositeKeys(stub, lib.LogisticComKey, []string{accountLogistic.ServiceId})
	if err != nil || len(resultsLogisticCom) != 1 {
		return shim.Error(fmt.Sprintf("logisticCom信息验证失败%s,%s", err, accountLogistic.ServiceId))
	}
	var logisticCom lib.LogisticCom
	if err = json.Unmarshal(resultsLogisticCom[0], &logisticCom); err != nil {
		return shim.Error(fmt.Sprintf("查询logisticCom信息-反序列化出错: %s", err))
	}

	var data []byte
	//判断需要更新的订单状态
	switch status {
	case "arrived":
		//如果是manufacturer确认收款操作,必须确保订单处于交付状态
		if order.OrderStatus != "delivery" {
			return shim.Error(fmt.Sprintf("This order is not in delivery"))
		}
		//订单状态设置为完成
		order.OrderStatus = status
		//重新将订单写入账本
		if err := utils.WriteLedger(order, stub, lib.OrderKey, []string{order.StoreUserAccountId, order.OrderId, order.ManufacturerUserAccountId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		data, err = json.Marshal(order)
		if err != nil {
			return shim.Error(fmt.Sprintf("序列化订单状态的信息出错: %s", err))
		}
		break
	case "done":
		//如果是manufacturer确认收款操作,必须确保订单处于交付状态
		if order.OrderStatus != "arrived" {
			return shim.Error(fmt.Sprintf("此交易并不处于交付中，确认收款失败"))
		}
		//确认收款,将款项加入到manufacturer,get logistic fee
		var logistic_fee = order.OrderPrice * 0.05
		manufacturer.Funds = manufacturer.Funds + order.OrderPrice - logistic_fee
		store.Funds = store.Funds - order.OrderPrice + logistic_fee
		logisticCom.Funds += logistic_fee

		if err := utils.WriteLedger(manufacturer, stub, lib.ManufacturerKey, []string{manufacturer.ManufacturerId}); err != nil {
			return shim.Error(fmt.Sprintf("manufacturer确认接收资金失败%s", err))
		}
		if err := utils.WriteLedger(store, stub, lib.StoreKey, []string{store.StoreId}); err != nil {
			return shim.Error(fmt.Sprintf("store确认支出资金失败%s", err))
		}
		if err := utils.WriteLedger(logisticCom, stub, lib.LogisticComKey, []string{logisticCom.LogisticComId}); err != nil {
			return shim.Error(fmt.Sprintf("logisticCom确认接收资金失败%s", err))
		}
		//订单状态设置为完成
		order.OrderStatus = status
		// // 从账本中撤销订单
		// if err := utils.DelLedger(stub, lib.OrderKey, []string{order.OrderId}); err != nil {
		// 	return shim.Error(fmt.Sprintf("%s", err))
		// }
		//重新将订单写入账本
		if err := utils.WriteLedger(order, stub, lib.OrderKey, []string{order.StoreUserAccountId, order.OrderId, order.ManufacturerUserAccountId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		data, err = json.Marshal(order)
		if err != nil {
			return shim.Error(fmt.Sprintf("序列化订单状态的信息出错: %s", err))
		}
		break
	case "cancelled":
		if order.OrderStatus == "cancelled" {
			return shim.Error(fmt.Sprintf("Order already cancelled"))
		}
		//判断Store用户账户是否具有权限
		if accountStore.UserType != "store" {
			return shim.Error(fmt.Sprintf("非门店不能操作%s", err))
		}
		return closeOrder(status, order, logisticCom, store, manufacturer, stub)
		break
	case "expired":
		if order.OrderStatus == "expired" {
			return shim.Error(fmt.Sprintf("Order already expired"))
		}
		//判断Logistic用户账户是否具有权限
		if accountLogistic.UserType != "administrator" {
			return shim.Error(fmt.Sprintf("非物流公司管理员不能操作%s", err))
		}
		return closeOrder(status, order, logisticCom, store, manufacturer, stub)
		break
	default:
		return shim.Error(fmt.Sprintf("%s状态不支持", status))
	}
	return shim.Success(data)
}

func closeOrder(closeStart string, order lib.Order, logisticCom lib.LogisticCom, store lib.Store, manufacturer lib.Manufacturer, stub shim.ChaincodeStubInterface) pb.Response {
	//运输途中取消，要付违约金:100
	if order.OrderStatus == "delivery" && closeStart != "expired" {
		logisticCom.Funds += 100
	}

	order.OrderStatus = closeStart

	if err := utils.WriteLedger(logisticCom, stub, lib.LogisticComKey, []string{logisticCom.LogisticComId}); err != nil {
		return shim.Error(fmt.Sprintf("logisticCom确认接收违约金:100失败%s", err))
	}
	if err := utils.WriteLedger(order, stub, lib.OrderKey, []string{order.StoreUserAccountId, order.OrderId, order.ManufacturerUserAccountId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	data, err := json.Marshal(order)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(data)
}
