import request from '@/utils/request'

// 查询Trace列表
export function listTrace(data) {
return request({
url: '/queryTraceList',
method: 'post',
data: data
})
}

// 查询Trace详细
export function getTrace (data) {
return request({
    url: '/queryTraceList',
    method: 'post',
    data: data
})
}


// 新增Trace
export function addTrace(data) {
return request({
url: '/createTrace',
method: 'post',
data: data
})
}

// 修改Trace
export function updateTrace(data) {
return request({
url: '/api/v1/trace',
method: 'put',
data: data
})
}

// 删除Trace
export function delTrace(traceId) {
return request({
url: '/api/v1/trace/' + traceId,
method: 'delete'
})
}