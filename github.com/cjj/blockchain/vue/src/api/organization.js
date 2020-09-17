import request from '@/utils/request'

// 查询LogisticCom
export function getLogisticCom(data) {
return request({
url: '/queryLogisticComInfo',
method: 'post',
data:data
})
}

//查询Manufacturer详细
export function getManufacturer (data) {
return request({
url: '/queryManufacturerInfo',
method: 'post',
data:data
})
}
//查询Store详细
export function getStore (data) {
return request({
url: '/queryStoreInfo',
method: 'post',
data:data
})
}