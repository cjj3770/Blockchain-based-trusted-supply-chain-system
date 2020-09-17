package routers

import (
	"encoding/json"
	"fmt"

	//"strconv"
	//"time"

	"github.com/cjj/blockchain/chaincode/blockchain/lib"
	"github.com/cjj/blockchain/chaincode/blockchain/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//新建用户账户(管理员)
func CreateUserAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 6 {
		return shim.Error("参数个数不满足")
	}
	userAccountId := args[0] //userAccountId用于验证是否为管理员
	createUserAccountId := args[1]
	userName := args[2]
	password := args[3]
	userType := args[4]
	serviceId := args[5]
	// lastLoginTime := args[4]
	// status := args[5]
	// remark := args[6]

	if userAccountId == "" || createUserAccountId == "" || userName == "" || password == "" || userType == "" || serviceId == "" {
		return shim.Error("参数存在空值")
	}
	//判断是否管理员操作
	resultsUserAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{userAccountId})
	if err != nil || len(resultsUserAccount) != 1 {
		return shim.Error(fmt.Sprintf("操作人权限验证失败%s", err))
	}
	var userAccount lib.UserAccount
	if err = json.Unmarshal(resultsUserAccount[0], &userAccount); err != nil {
		return shim.Error(fmt.Sprintf("查询操作人信息-反序列化出错: %s", err))
	}
	if userAccount.UserType != "administrator" {
		return shim.Error(fmt.Sprintf("操作人权限不足%s,%s", err, userAccount.UserType))
	}
	//判断将创建的账户是否存在
	resultsCreateUserAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, []string{createUserAccountId})
	if len(resultsCreateUserAccount) != 0 {
		return shim.Error(fmt.Sprintf("将创建的账户存在%s,%s", err, createUserAccountId))
	}
	//check if usertype is valid
	if userType != "courier" && userType != "manufacturer" && userType != "store" {
		return shim.Error(fmt.Sprintf("userType is invalid%s,%s", err, userType))
	}
	createUserAccount := &lib.UserAccount{
		AccountId: createUserAccountId,
		UserName:  userName,
		Password:  password,
		UserType:  userType,
		ServiceId: serviceId,
	}
	// 写入账本
	if err := utils.WriteLedger(createUserAccount, stub, lib.UserAccountKey, []string{createUserAccount.AccountId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	createUserAccountByte, err := json.Marshal(createUserAccount)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(createUserAccountByte)
}

//查询账户(可查询所有，也可根据账户Id查询账户)
func QueryUserAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var userAccountList []lib.UserAccount
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserAccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var userAccount lib.UserAccount
			err := json.Unmarshal(v, &userAccount)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryUserAccountList-反序列化出错: %s", err))
			}
			userAccountList = append(userAccountList, userAccount)
		}
	}
	userAccountListByte, err := json.Marshal(userAccountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryUserAccountList-序列化出错: %s", err))
	}
	return shim.Success(userAccountListByte)
}
