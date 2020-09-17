const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  // token: state => state.account.token,
  // accountId: state => state.account.accountId,
  // userName: state => state.account.userName,
  // balance: state => state.account.balance,
  // roles: state => state.account.roles,
  // permission_routes: state => state.permission.routes
  token: state => state.useraccount.token,
  accountId: state => state.useraccount.accountId,
  userName: state => state.useraccount.userName,
  userType: state => state.useraccount.userType,
  serviceId: state => state.useraccount.serviceId,
  balance: state => state.useraccount.balance,
  roles: state => state.useraccount.roles,
  permission_routes: state => state.permission.routes
}
export default getters
