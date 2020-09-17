package main

import (
	"fmt"
	"time"

	"github.com/cjj/blockchain/chaincode/blockchain/lib"
	"github.com/cjj/blockchain/chaincode/blockchain/routers"
	"github.com/cjj/blockchain/chaincode/blockchain/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type BlockChainRealEstate struct {
}

//链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("chaincode initialization")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal

	//chaincode initialization
	var accountIds = [6]string{"L000", "M001", "M002", "S001", "S002", "C001"}

	var serviceIds = [6]string{"logstic000", "manufacturer001", "manufacturer002", "store001", "store002", "logstic000"}
	var userName = [6]string{"Logistics company administrator", "Manufacturer①", "Manufacturer②", "Store①", "Store②", "Courier①"}
	var password = [6]string{"111", "222", "333", "444", "555", "666"}
	var userType = [6]string{"administrator", "manufacturer", "manufacturer", "store", "store", "courier"}
	var lastLoginTime = [6]string{"18/07/2020 14:47:25", "04/08/2020 13:44:25", "13/09/2019 13:44:25", "03/07/2020 13:44:25", "03/06/2020 13:44:25", "29/07/2020 23:44:25"}
	var status = [6]string{"0", "0", "0", "0", "0", "0"}
	var remark = [6]string{"login success", "login success", "login success", "login success", "login success", "login success"}

	var balances = [6]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000}

	var orderIds = [6]string{"s1m1-1", "s2m2-1", "s1m1-2", "s1m2-1", "s2m1-1", "s1m2-2"}
	var storeUserAccountId = [6]string{"S001", "S002", "S001", "S001", "S002", "S001"}
	var manufacturerUserAccountId = [6]string{"M001", "M002", "M001", "M002", "M001", "M002"}
	var logisticUserAccountId = [6]string{"L000", "L000", "L000", "L000", "L000", "L000"}
	var orderName = [6]string{"ORDER-IPHONE", "ORDER-SAMSUNG", "ORDER-SONY", "ORDER-HUAWEI", "ORDER-MIPHONE", "ORDER-OPPO"}
	var orderTime = [6]string{"2020-08-12 13:21:16", "2020-09-1 03:21:19", "2020-08-11 05:01:16", "2020-08-21 03:31:18", "2020-08-17 04:25:16", "2020-08-19 06:21:16"}
	var orderPrice = [6]float64{100000.00, 400000.00, 200000.00, 300000.00, 500000.00, 100000.00}
	var orderStatus = [6]string{"delivery", "delivery", "orderStart", "arrived", "expired", "orderStart"}
	var departAddress = [6]string{"Southampton", "Oxford", "Cambridge", "London", "Liverpool", "Bristol"}
	var arrivalAddress = [6]string{"London", "Cambridge", "Oxford", "Southampton", "Bristol", "Liverpool"}
	//initialize logistic
	var logisticId = [2]string{"l-000", "l-001"}
	var courierAccountId = [2]string{"C001", "C001"}
	var viaAddress = [2]string{"Hatfield", "Stevenage"}
	var viaTime = [2]string{"2020-08-11 03:21:16", "2020-08-11 03:21:16"}
	var departTime = [2]string{"2020-08-10 04:10:15", "2020-08-11 03:21:16"}
	var arrivalTime = [2]string{"", ""}
	var logisticStatus = [2]string{"start", "start"}
	for i, val := range logisticId {
		logistic := &lib.Logistic{
			OrderId:                   orderIds[i],
			StoreUserAccountId:        storeUserAccountId[i],
			ManufacturerUserAccountId: manufacturerUserAccountId[i],
			LogisticId:                val,
			CourierAccountId:          courierAccountId[i],
			ViaAddress:                viaAddress[i],
			ViaTime:                   viaTime[i],
			DepartTime:                departTime[i],
			ArrivalTime:               arrivalTime[i],
			LogisticStatus:            logisticStatus[i],
			ArrivalAddress:            arrivalAddress[i],
			DepartAddress:             departAddress[i],
		}
		// logistic写入账本
		if err := utils.WriteLedger(logistic, stub, lib.LogisticKey, []string{logistic.LogisticId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	//初始化manufacturer数据
	var manufacturerIds = [2]string{"manufacturer001", "manufacturer002"}
	var manufacturerName = [2]string{"Phone Manuf 1", "Phone Manuf 2"}
	var manufacturerPrincipal = [2]string{"Harry", "John"}
	var manufacturerAddress = [2]string{"Bournemouth", "Waymouth"}
	var inventory = [2]int{1000, 2000}
	var funds = [2]float64{10000, 10000}
	for i, val := range manufacturerIds {
		manufacturer := &lib.Manufacturer{
			ManufacturerId:        val,
			ManufacturerName:      manufacturerName[i],
			ManufacturerPrincipal: manufacturerPrincipal[i],
			ManufacturerAddress:   manufacturerAddress[i],
			Inventory:             inventory[i],
			Funds:                 funds[i],
			OrderId:               "",
		}
		// manufacturer写入账本
		if err := utils.WriteLedger(manufacturer, stub, lib.ManufacturerKey, []string{manufacturer.ManufacturerId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	//初始化logisticCom数据
	var logisticComIds = [1]string{"logstic000"}
	var logisticComName = [1]string{"Royal Mail"}
	var logisticComPrincipal = [1]string{"Harry Potter"}
	var logisticComAddress = [1]string{"London"}
	var logisticComInventory = [1]int{1000}
	var logisticComFunds = [1]float64{10000}
	for i, val := range logisticComIds {
		logisticCom := &lib.LogisticCom{
			LogisticComId:        val,
			LogisticComName:      logisticComName[i],
			LogisticComPrincipal: logisticComPrincipal[i],
			LogisticComAddress:   logisticComAddress[i],
			Inventory:            logisticComInventory[i],
			Funds:                logisticComFunds[i],
			OrderId:              "",
		}
		// logisticCom写入账本
		if err := utils.WriteLedger(logisticCom, stub, lib.LogisticComKey, []string{logisticCom.LogisticComId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	//初始化trace数据
	var manuAccountIds = [6]string{"M001", "M001", "M001", "M002", "M002", "M002"}
	var traceIds = [6]string{"B00", "B01", "B02", "A01", "A02", "A03"}
	var productId = [6]string{"a", "b", "c", "d", "e", "f"}
	var productName = [6]string{"IPHONEX", "IPHONE7", "IPHONE7-PLUS", "IPHONE6", "IPHONE8", "IPHONE10"}
	var productNumber = [6]int{10, 20, 30, 40, 50, 60}
	var productPrice = [6]float64{8000, 7000, 7500, 6000, 9000, 10000}
	var productTime = [6]string{"2020-08-12 04:10:15", "2020-08-11 03:21:16", "2020-08-11 03:21:16", "2020-08-11 03:21:16", "2020-08-11 03:21:16", "2020-08-11 03:21:16"}
	for i, val := range traceIds {
		trace := &lib.Trace{
			TraceId:       val,
			UserAccountId: manuAccountIds[i],
			ProductId:     productId[i],
			ProductName:   productName[i],
			ProductNumber: productNumber[i],
			ProductPrice:  productPrice[i],
			ProductTime:   productTime[i],
			TraceStatus:   "created",
			RawIds:        "r001,r002",
		}
		// trace写入账本
		if err := utils.WriteLedger(trace, stub, lib.TraceKey, []string{trace.UserAccountId, trace.TraceId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	//初始化order数据
	for i, val := range orderIds {
		order := &lib.Order{
			OrderId:                   val,
			StoreUserAccountId:        storeUserAccountId[i],
			ManufacturerUserAccountId: manufacturerUserAccountId[i],
			LogisticUserAccountId:     logisticUserAccountId[i],
			OrderName:                 orderName[i],
			OrderPrice:                orderPrice[i],
			OrderTime:                 orderTime[i],
			OrderStatus:               orderStatus[i],
			DepartAddress:             departAddress[i],
			ArrivalAddress:            arrivalAddress[i],
			TraceId:                   traceIds[i],
		}
		// order写入账本
		if err := utils.WriteLedger(order, stub, lib.OrderKey, []string{order.StoreUserAccountId, order.OrderId, order.ManufacturerUserAccountId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	//初始化store数据
	var storeIds = [2]string{"store001", "store002"}
	var storeName = [2]string{"iphone store 1", "iphone store 2"}
	var storePrincipal = [2]string{"Jack", "Lucy"}
	var storeAddress = [2]string{"Southampton", "Cambridge"}
	var storeInventory = [2]int{1000, 2000}
	var storeFunds = [2]float64{99999, 88888}
	for i, val := range storeIds {
		store := &lib.Store{
			StoreId:        val,
			StoreName:      storeName[i],
			StorePrincipal: storePrincipal[i],
			StoreAddress:   storeAddress[i],
			Inventory:      storeInventory[i],
			Funds:          storeFunds[i],
			OrderId:        "",
		}
		// store写入账本
		if err := utils.WriteLedger(store, stub, lib.StoreKey, []string{store.StoreId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	//初始化账号数据
	for i, val := range accountIds {
		account := &lib.Account{
			AccountId: val,
			UserName:  userName[i],
			Balance:   balances[i],
		}

		// 写入账本
		if err := utils.WriteLedger(account, stub, lib.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}

	}
	//初始化用户账号数据
	for i, val := range accountIds {
		userAccount := &lib.UserAccount{
			AccountId:     val,
			UserName:      userName[i],
			Password:      password[i],
			UserType:      userType[i],
			LastLoginTime: lastLoginTime[i],
			Status:        status[i],
			Remark:        remark[i],
			ServiceId:     serviceIds[i],
			Balance:       balances[i],
		}
		// 写入账本
		if err := utils.WriteLedger(userAccount, stub, lib.UserAccountKey, []string{userAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

//实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "createOrder":
		return routers.CreateOrder(stub, args)
	case "updateOrder":
		return routers.UpdateOrder(stub, args)
	case "queryOrderList":
		return routers.QueryOrderList(stub, args)
	case "queryUserAccountList":
		return routers.QueryUserAccountList(stub, args)
	case "createUserAccount":
		return routers.CreateUserAccount(stub, args)
	case "createTrace":
		return routers.CreateTrace(stub, args)
	case "queryTraceList":
		return routers.QueryTraceList(stub, args)
	case "createLogistic":
		return routers.CreateLogistic(stub, args)
	case "updateLogistic":
		return routers.UpdateLogistic(stub, args)
	case "queryLogisticList":
		return routers.QueryLogisticList(stub, args)
	case "queryLogisticComInfo":
		return routers.QueryLogisticComInfo(stub, args)
	case "queryManufacturerInfo":
		return routers.QueryManufacturerInfo(stub, args)
	case "queryStoreInfo":
		return routers.QueryStoreInfo(stub, args)

	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
