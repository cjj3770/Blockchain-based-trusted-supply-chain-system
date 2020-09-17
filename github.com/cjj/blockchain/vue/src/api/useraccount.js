import request from '@/utils/request'


// 获取登录界面角色选择列表
export function queryUserAccountList() {
    return request({
      url: '/queryUserAccountList',
      method: 'post'
    })
  }
  
  // 登录
  export function login(data) {
    return request({
      url: '/queryUserAccountList',
      method: 'post',
      data
    })
  }



// 查询UserAccount列表
export function listUserAccount(query) {
return request({
url: '/queryUserAccountList',
method: 'post',
})
}

// 查询UserAccount详细
export function getUserAccount (data) {
return request({
url: '/queryUserAccount',
method: 'post',
data:data,
})
}


// 新增UserAccount
export function addUserAccount(data) {
return request({
url: '/createUserAccount',
method: 'post',
data: data
})
}

// 修改UserAccount
export function updateUserAccount(data) {
return request({
url: '/api/v1/useraccount',
method: 'put',
data: data
})
}

// 删除UserAccount
export function delUserAccount(userAccountId) {
return request({
url: '/api/v1/useraccount/' + userAccountId,
method: 'delete'
})
}