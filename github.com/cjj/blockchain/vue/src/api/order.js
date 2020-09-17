import request from '@/utils/request'

// 查询Order列表
export function listOrder(query) {
    return request({
        url: '/queryOrderList',
        method: 'post',
        })
}

// 查询Order详细
export function getOrder (data) {
return request({
url: '/queryOrderList',
method: 'post',
data:data
})
}


// 新增Order
export function addOrder(data) {
return request({
url: '/createOrder',
method: 'post',
data: data
})
}

// 确认Order
export function confirmOrder(data) {
return request({
url: 'updateOrder',
method: 'post',
data:data
})
}

// 取消Order
export function cancelOrder(data) {
return request({
url: 'updateOrder',
method: 'post',
data:data
})
}