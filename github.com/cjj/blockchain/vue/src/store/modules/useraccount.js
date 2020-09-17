import {
    login
  } from '@/api/useraccount'
  import {
    getToken,
    setToken,
    removeToken
  } from '@/utils/auth'
  import {
    resetRouter
  } from '@/router'
  
  const getDefaultState = () => {
    return {
      token: getToken(),
      accountId: '',
      userName: '',
      userType:'',
      serviceId:'',
      balance: 0,
      roles: []
    }
  }
  
  const state = getDefaultState()
  
  const mutations = {
    RESET_STATE: (state) => {
      Object.assign(state, getDefaultState())
    },
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_ACCOUNTID: (state, accountId) => {
      state.accountId = accountId
    },
    SET_USERNAME: (state, userName) => {
      state.userName = userName
    },
    SET_USERTYPE: (state, userType) => {
      state.userType = userType
    },
    SET_SERVICEID: (state, serviceId) => {
      state.serviceId = serviceId
    },
    SET_BALANCE: (state, balance) => {
      state.balance = balance
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    }
  }
  
  const actions = {
    login({
      commit
    }, accountId) {
      return new Promise((resolve, reject) => {
        login({
          args: [{
            userAccountId: accountId
          }]
        }).then(response => {
          commit('SET_TOKEN', response[0].accountId)
          setToken(response[0].accountId)
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },
    // get user info
    getInfo({
      commit,
      state
    }) {
      return new Promise((resolve, reject) => {
        login({
          args: [{
            userAccountId: state.token
          }]
        }).then(response => {
          var roles
          switch (response[0].userType) {
            case 'administrator':
              roles = ['admin']
              break;
            case 'manufacturer':
              roles = ['manuf']
              break;
            case 'store':
              roles = ['store']
              break;
            case 'courier':
              roles=  ['courier']
              break;
          
            default:
              break;
          }
          // if (response[0].userType === 'administrator') {
          //   roles = ['admin']
          // } 
          // else{
          //   roles = ['editor']
          // }
          console.log("roles>>>>>>>>"+roles)
          console.log("roles>>>>>>>>"+response[0].accountId)
          console.log("roles>>>>>>>>"+response[0].userName)
          console.log("roles>>>>>>>>"+response[0].balance)
          console.log("roles>>>>>>>>"+response[0].serviceId)
          commit('SET_ROLES', roles)
          commit('SET_ACCOUNTID', response[0].accountId)
          commit('SET_USERNAME', response[0].userName)
          commit('SET_BALANCE', response[0].balance)
          commit('SET_SERVICEID',response[0].serviceId)
          resolve(roles)
        }).catch(error => {
          reject(error)
        })
      })
    },
    logout({
      commit
    }) {
      return new Promise(resolve => {
        removeToken()
        resetRouter()
        commit('RESET_STATE')
        resolve()
      })
    },
  
    resetToken({
      commit
    }) {
      return new Promise(resolve => {
        removeToken()
        commit('RESET_STATE')
        resolve()
      })
    }
  }
  
  export default {
    namespaced: true,
    state,
    mutations,
    actions
  }