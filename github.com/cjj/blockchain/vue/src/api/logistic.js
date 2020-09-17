import request from '@/utils/request'

// 查询Logistic列表
export function listLogistic() {
return request({
url: '/queryLogisticList',
method: 'post',
})
}

// // 查询Logistic详细
export function getLogistic (data) {
return request({
url: '/queryLogisticList',
method: 'post',
data:data
})
}


// 新增Logistic
export function addLogistic(data) {
return request({
url: '/createLogistic',
method: 'post',
data: data
})
}

// 修改Logistic
export function updateLogistic(data) {
return request({
url: '/updateLogistic',
method: 'post',
data: data
})
}

// 删除Logistic
export function delLogistic(logisticId) {
return request({
url: '/api/v1/logistic/' + logisticId,
method: 'delete'
})
}
