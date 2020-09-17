package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cjj/blockchain/chaincode/blockchain/lib"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func initTest(t *testing.T) *shim.MockStub {
	scc := new(BlockChainRealEstate)
	stub := shim.NewMockStub("ex01", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	return stub
}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) pb.Response {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
	return res
}

// 测试链码初始化
func TestBlockChainRealEstate_Init(t *testing.T) {
	initTest(t)
}

// 测试获取组织信息
func Test_QueryLogisticComInfo(t *testing.T) {
	stub := initTest(t)
	fmt.Println(fmt.Sprintf("1、测试获取LogisticCom数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryLogisticComInfo"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取指定Manufacturer数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryManufacturerInfo"),
			[]byte("manufacturer002"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取Store数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryStoreInfo"),
			[]byte("store002"),
		}).Payload)))
}

//手动发起物流
func checkCreateLogistic(stub *shim.MockStub, t *testing.T) []lib.Logistic {
	var createLogisticList []lib.Logistic
	var createLogistic lib.Logistic
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createLogistic"),
		[]byte("3"),
		[]byte("S001"),
		[]byte("M001"),
		[]byte("l-003"),
		[]byte("C001"),
		[]byte("6/8/2020"),
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &createLogistic)
	createLogisticList = append(createLogisticList, createLogistic)
	return createLogisticList
}

// 测试获取物流信息
func Test_QueryLogisticList(t *testing.T) {
	stub := initTest(t)
	logisticList := checkCreateLogistic(stub, t)

	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryLogisticList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取指定数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryLogisticList"),
			[]byte(logisticList[0].LogisticId),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryLogisticList"),
			[]byte("0"),
		}).Payload)))
}

//手动更新物流
func checkUpdateLogistic(stub *shim.MockStub, t *testing.T) []lib.Logistic {
	var updateLogisticList []lib.Logistic
	var updateLogistic lib.Logistic
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("updateLogistic"),
		[]byte("C001"),
		[]byte("l-000"),
		[]byte("Baldock"),
		[]byte("2020-08-10 04:10:15"),
		[]byte("delivery"),
		[]byte(""),
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &updateLogistic)
	updateLogisticList = append(updateLogisticList, updateLogistic)
	return updateLogisticList
}

// 测试获取物流信息
func Test_QueryLogisticList2(t *testing.T) {
	stub := initTest(t)
	logisticList := checkUpdateLogistic(stub, t)

	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryLogisticList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取指定数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryLogisticList"),
			[]byte(logisticList[0].LogisticId),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryLogisticList"),
			[]byte("0"),
		}).Payload)))
}

// 测试查询账户
func Test_QueryUserAccountList(t *testing.T) {
	stub := initTest(t)
	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryUserAccountList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取多个数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryUserAccountList"),
			[]byte("L000"),
			[]byte("S003"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取单个数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryUserAccountList"),
			[]byte("M002"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("4、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryUserAccountList"),
			[]byte("0"),
		}).Payload)))
}

//手动创建一些用户账户
func checkCreateUserAccount(stub *shim.MockStub, t *testing.T) []lib.UserAccount {
	var userAccountList []lib.UserAccount
	var userAccount lib.UserAccount
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createUserAccount"),
		[]byte("L000"),     //管理员操作
		[]byte("S010"),     //创建账户的ID
		[]byte("cjj3779"),  //账户名
		[]byte("123123"),   //密码
		[]byte("store"),    //创建账户的用户类型
		[]byte("store003"), //创建账户所服务的机构ID
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &userAccount)
	userAccountList = append(userAccountList, userAccount)
	return userAccountList
}

// 测试获取创建的用户账户信息
func Test_QueryUserAccountList2(t *testing.T) {
	stub := initTest(t)
	userAccountList := checkCreateUserAccount(stub, t)

	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryUserAccountList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取指定数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryUserAccountList"),
			[]byte(userAccountList[0].AccountId),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryUserAccountList"),
			[]byte("0"),
		}).Payload)))
}

// 测试创建创世区块
func Test_CreateTrace(t *testing.T) {
	stub := initTest(t)
	//成功
	checkInvoke(t, stub, [][]byte{
		[]byte("createTrace"),
		[]byte("A00"),                 //traceId
		[]byte("M001"),                //userAccountId操作人
		[]byte("0"),                   //productId
		[]byte("iphoneX"),             //productName
		[]byte("100"),                 //productNumber
		[]byte("800"),                 //productPrice
		[]byte("04/08/2022 14:47:25"), //productTime
		[]byte("r001,r002"),           //rawIds
	})
}

//手动创建创世区块
func checkCreateTrace(stub *shim.MockStub, t *testing.T) []lib.Trace {
	var traceList []lib.Trace
	var trace lib.Trace
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createTrace"),
		[]byte("A00"),                 //traceId
		[]byte("M001"),                //userAccountId操作人
		[]byte("0"),                   //productId
		[]byte("iphoneX"),             //productName
		[]byte("100"),                 //productNumber
		[]byte("800"),                 //productPrice
		[]byte("04/08/2022 14:47:25"), //productTime
		[]byte("r001,r002"),           //rawIds
	})
	resp2 := checkInvoke(t, stub, [][]byte{
		[]byte("createTrace"),
		[]byte("A01"),                 //traceId
		[]byte("M001"),                //userAccountId操作人
		[]byte("0"),                   //productId
		[]byte("iphoneX"),             //productName
		[]byte("100"),                 //productNumber
		[]byte("800"),                 //productPrice
		[]byte("04/08/2022 20:47:25"), //productTime
		[]byte("r001,r002"),           //rawIds
	})
	resp3 := checkInvoke(t, stub, [][]byte{
		[]byte("createTrace"),
		[]byte("A01"),                 //traceId
		[]byte("M002"),                //userAccountId操作人
		[]byte("1"),                   //productId
		[]byte("iphoneX"),             //productName
		[]byte("100"),                 //productNumber
		[]byte("800"),                 //productPrice
		[]byte("05/08/2022 20:47:25"), //productTime
		[]byte("r001,r002"),           //rawIds
	})
	resp4 := checkInvoke(t, stub, [][]byte{
		[]byte("createTrace"),
		[]byte("A03"),                 //traceId
		[]byte("M002"),                //userAccountId操作人
		[]byte("1"),                   //productId
		[]byte("iphoneX"),             //productName
		[]byte("100"),                 //productNumber
		[]byte("800"),                 //productPrice
		[]byte("05/08/2022 22:47:25"), //productTime
		[]byte("r001,r002"),           //rawIds
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &trace)
	traceList = append(traceList, trace)
	json.Unmarshal(bytes.NewBuffer(resp2.Payload).Bytes(), &trace)
	traceList = append(traceList, trace)
	json.Unmarshal(bytes.NewBuffer(resp3.Payload).Bytes(), &trace)
	traceList = append(traceList, trace)
	json.Unmarshal(bytes.NewBuffer(resp4.Payload).Bytes(), &trace)
	traceList = append(traceList, trace)
	return traceList
}

// 测试获取创世区块信息
func Test_QueryTraceList(t *testing.T) {
	stub := initTest(t)
	//traceList := checkCreateTrace(stub, t)

	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryTraceList"),
			[]byte("M001"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取指定数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryTraceList"),
			[]byte("M001"),
			[]byte("B01"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryTraceList"),
			[]byte("0"),
		}).Payload)))
}

//手动创建一些订单
func checkCreateOrder(stub *shim.MockStub, t *testing.T) []lib.Order {
	var orderList []lib.Order
	var order lib.Order
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createOrder"),
		[]byte("7"),                   //orderId
		[]byte("S001"),                //storeAccountId
		[]byte("M001"),                //manufacturerAccountId
		[]byte("L000"),                //logisticAccountId
		[]byte("iphone-list1"),        //orderName
		[]byte("03/06/2020 00:44:25"), //orderTime
		[]byte("60000"),               //orderPrice
		[]byte("B00"),                 //traceId
		[]byte("London"),              //departureAddress
		[]byte("Southampton"),         //arrivalAddress
	})
	resp2 := checkInvoke(t, stub, [][]byte{
		[]byte("createOrder"),
		[]byte("8"),                   //orderId
		[]byte("S002"),                //storeAccountId
		[]byte("M001"),                //manufacturerAccountId
		[]byte("L000"),                //logisticAccountId
		[]byte("iphone-list2"),        //orderName
		[]byte("03/08/2020 00:44:30"), //orderTime
		[]byte("70000"),               //orderPrice
		[]byte("B01"),                 //traceId
		[]byte("Cambridge"),           //departureAddress
		[]byte("Southampton"),         //arrivalAddress
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &order)
	orderList = append(orderList, order)
	json.Unmarshal(bytes.NewBuffer(resp2.Payload).Bytes(), &order)
	orderList = append(orderList, order)

	return orderList
}

// 测试获取订单信息
func Test_QueryOrderList(t *testing.T) {
	stub := initTest(t)
	orderList := checkCreateOrder(stub, t)

	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryOrderList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取指定数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryOrderList"),
			[]byte(orderList[1].OrderId),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryOrderList"),
			[]byte("S002"),
			[]byte("2"),
		}).Payload)))
}

//手动更新一些订单
func checkUpdateOrder(stub *shim.MockStub, t *testing.T) []lib.Order {
	var orderList []lib.Order
	var order lib.Order
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("updateOrder"),
		[]byte("3"),    //orderId
		[]byte("M001"), //manufacturerAccountId
		[]byte("S001"), //storeAccountId
		[]byte("done"), //status
	})
	resp2 := checkInvoke(t, stub, [][]byte{
		[]byte("updateOrder"),
		[]byte("2"),    //orderId
		[]byte("M002"), //manufacturerAccountId
		[]byte("S002"), //storeAccountId
		[]byte("done"), //status
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &order)
	orderList = append(orderList, order)
	json.Unmarshal(bytes.NewBuffer(resp2.Payload).Bytes(), &order)
	orderList = append(orderList, order)
	return orderList
}

// 测试更新订单信息
func Test_UpdateOrder(t *testing.T) {
	stub := initTest(t)
	//orderList := checkUpdateOrder(stub, t)
	fmt.Println(fmt.Sprintf("1、测试更新指定订单数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("updateOrder"),
			[]byte("s2m2-1"),  //orderId
			[]byte("M002"),    //manufacturerAccountId
			[]byte("S002"),    //storeAccountId
			[]byte("arrived"), //status
		}).Payload)))

}

// //手动创建一些物流
// func checkCreateLogistic(stub *shim.MockStub, t *testing.T) []lib.Logistic {
// 	var logisticList []lib.Logistic
// 	var logistic lib.Logistic
// 	//成功
// 	resp1 := checkInvoke(t, stub, [][]byte{
// 		[]byte("createLogistic"),
// 		[]byte("1"),                   //orderId
// 		[]byte("l1"),                  //logisticId
// 		[]byte("C001"),                //courierAccountId
// 		[]byte("03/06/2020 00:44:25"), //departTime
// 	})
// 	resp2 := checkInvoke(t, stub, [][]byte{
// 		[]byte("createOrder"),
// 		[]byte("2"),                   //orderId
// 		[]byte("l2"),                  //logisticId
// 		[]byte("C001"),                //courierAccountId
// 		[]byte("03/06/2020 00:44:25"), //departTime
// 	})

// 	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &logistic)
// 	logisticList = append(logisticList, logistic)
// 	json.Unmarshal(bytes.NewBuffer(resp2.Payload).Bytes(), &logistic)
// 	logisticList = append(logisticList, logistic)
// 	return logisticList
// }

// // 测试获取物流信息
// func Test_QueryLogisticList(t *testing.T) {
// 	stub := initTest(t)
// 	logisticList := checkCreateLogistic(stub, t)

// 	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
// 		string(checkInvoke(t, stub, [][]byte{
// 			[]byte("queryLogisticList"),
// 		}).Payload)))
// 	fmt.Println(fmt.Sprintf("2、测试获取指定数据\n%s",
// 		string(checkInvoke(t, stub, [][]byte{
// 			[]byte("queryLogisticList"),
// 			[]byte(logisticList[1].LogisticId),
// 		}).Payload)))
// 	fmt.Println(fmt.Sprintf("3、测试获取无效数据\n%s",
// 		string(checkInvoke(t, stub, [][]byte{
// 			[]byte("queryLogisticList"),
// 			[]byte("3"),
// 		}).Payload)))
// }
