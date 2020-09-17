/**
 * @Author: Jiajie Chen
 * @Email: jc8n18@soton.ac.uk
 * @Date: 17-09-2020
 * @Description: define structurer
 */
package lib

//账户，虚拟管理员和若干业主账号
type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名
	Balance   float64 `json:"balance"`   //余额
}

//1.用户账户，虚拟管理员和若干业主账号
type UserAccount struct {
	//UserAccountId string  `json:"userAccountId"` //用户ID
	//AccountRole   AccountRole `json:"accountRole"`   //用户类型
	UserName      string  `json:"userName"`      //账户名
	Password      string  `json:"password"`      //密码
	UserType      string  `json:"userType"`      //用户类型
	LastLoginTime string  `json:"lastLoginTime"` //上一次登录时间
	Status        string  `json:"status"`        //状态
	Remark        string  `json:"remark"`        //备注
	AccountId     string  `json:"accountId"`
	Balance       float64 `json:"balance"`
	ServiceId     string  `json:"serviceId"` //服务机构ID
}

//2.角色权限
type RolePermission struct {
	RoleName   string `json:"roleName"`   //角色名字
	Permission string `json:"permission"` //角色权限
}

//3.用户角色
type AccountRole struct {
	RoleId         int            `json:"roleId"` //角色ID
	RolePermission RolePermission `json:"rolePermission"`
}

//4.商品信息
type Trace struct {
	TraceId       string  `json:"traceId"`       //创世ID
	UserAccountId string  `json:"userAccountId"` //用户账户
	ProductId     string  `json:"productId"`     //商品ID
	ProductName   string  `json:"productName"`   //商品名称
	ProductNumber int     `json:"productNumber"` //商品数量
	ProductPrice  float64 `json:"productPrice"`  //商品价格
	ProductTime   string  `json:"productTime"`   //生产日期
	RawIds        string  `json:"rawIds"`        //原材料
	TraceStatus   string  `json:"traceStatus"`   //商品状态
}

//5.原材料信息
type RawMaterial struct {
	RawId      int    `json:"rawId"`      //原材料ID
	RawName    string `json:"rawName"`    //原材料名称
	RawTime    string `json:"rawTime"`    //原材料生产日期
	RawComName string `json:"rawComName"` //raw company name
}

//6.logistic information
type Logistic struct {
	OrderId                   string `json:"orderId"` //Order ID
	StoreUserAccountId        string `json:"storeUserAccountId"`
	ManufacturerUserAccountId string `json:"manufacturerUserAccountId"`
	LogisticId                string `json:"logisticId"`       //logistic ID
	CourierAccountId          string `json:"courierAccountId"` //courier account ID
	ViaAddress                string `json:"viaAddress"`       //via address
	ViaTime                   string `json:"viaTime"`          //via time
	DepartTime                string `json:"departTime"`       //departure time
	ArrivalTime               string `json:"arrivalTime"`      //arrival time
	DepartAddress             string `json:"departAddress"`    //departure address
	ArrivalAddress            string `json:"arrivalAddress"`   //arrival address
	LogisticStatus            string `json:"logisticStatus"`   //logistic status
}

//7.logistic company information
type LogisticCom struct {
	LogisticComId        string  `json:"Id"`
	LogisticComName      string  `json:"Name"`
	LogisticComPrincipal string  `json:"Principal"`
	LogisticComAddress   string  `json:"Address"`
	Inventory            int     `json:"inventory"`
	Funds                float64 `json:"funds"`
	OrderId              string  `json:"orderId"`
}

//8.生产商信息
type Manufacturer struct {
	ManufacturerId        string  `json:"Id"`
	ManufacturerName      string  `json:"Name"`
	ManufacturerPrincipal string  `json:"Principal"`
	ManufacturerAddress   string  `json:"Address"`
	Inventory             int     `json:"inventory"`
	Funds                 float64 `json:"funds"`
	OrderId               string  `json:"orderId"`
}

//8.门店信息
type Store struct {
	StoreId        string  `json:"Id"`
	StoreName      string  `json:"Name"`
	StorePrincipal string  `json:"Principal"`
	StoreAddress   string  `json:"Address"`
	Inventory      int     `json:"inventory"`
	Funds          float64 `json:"funds"`
	OrderId        string  `json:"orderId"`
}

//9.订单查询
type Order struct {
	//ModifyAccountId           string  `json:"modifyAccountId"`
	OrderId                   string  `json:"orderId"`
	StoreUserAccountId        string  `json:"storeUserAccountId"`
	ManufacturerUserAccountId string  `json:"manufacturerUserAccountId"`
	LogisticUserAccountId     string  `json:"logisticUserAccountId"`
	OrderName                 string  `json:"orderName"`
	OrderTime                 string  `json:"orderTime"`
	OrderPrice                float64 `json:"orderPrice"`
	TraceId                   string  `json:"traceId"`
	OrderStatus               string  `json:"orderStatus"`
	DepartAddress             string  `json:"departAddress"`  //寄件地址
	ArrivalAddress            string  `json:"arrivalAddress"` //收件地址
}

const (
	AccountKey      = "account-key"
	OrderKey        = "order-key"
	UserAccountKey  = "userAccount-key"
	TraceKey        = "trace-key"
	LogisticKey     = "logistic-key"
	LogisticComKey  = "logisticCom-key"
	ManufacturerKey = "manufacturer-key"
	StoreKey        = "store-key"
)
