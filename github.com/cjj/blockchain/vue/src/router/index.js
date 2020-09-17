import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [{
  path: '/login',
  component: () => import('@/views/login/index'),
  hidden: true
},

{
  path: '/404',
  component: () => import('@/views/404'),
  hidden: true
},

{
  path: '/',
  component: Layout,
  redirect: '/accountinfo',
  children: [{
    path: 'accountinfo',
    name: 'accountinfo',
    component: () => import('@/views/accountinfo/index'),
    meta: {
      title: 'Account Information',
      icon: 'realestate'
    }
  }]
},
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [

  {
    path: '/useraccount',
    component: Layout,
    meta: {
      roles: ['admin']
    },
    children: [{
      path: '/useraccount',
      name: 'useraccount',
      component: () => import('@/views/useraccount/index'),
      meta: {
        title: 'User Account',
        icon: 'donatingAll'
      }
    }]
  },

  {
    path: '/order_admin',
    component: Layout,
    meta: {
      roles: ['admin']
    },
    children: [{
      path: '/order_admin',
      name: 'order_admin',
      component: () => import('@/views/order_admin/index'),
      meta: {
        title: 'Order',
        icon: 'sellingBuy'
      }
    }]
  },
  {
    path: '/order_store',
    component: Layout,
    meta: {
      roles: ['store']
    },
    children: [{
      path: '/order_store',
      name: 'order_store',
      component: () => import('@/views/order_store/index'),
      meta: {
        title: 'Order',
        icon: 'sellingBuy'
      }
    }]
  },
  {
    path: '/trace',
    component: Layout,
    meta: {
      roles: ['manuf']
    },
    children: [{
      path: '/trace',
      name: 'trace',
      component: () => import('@/views/trace/index'),
      meta: {
        title: 'Trace',
        icon: 'donating'
      }
    }]
  },
  {
    path: '/order_manuf',
    component: Layout,
    meta: {
      roles: ['manuf']
    },
    children: [{
      path: '/order_manuf',
      name: 'order_manuf',
      component: () => import('@/views/order_manuf/index'),
      meta: {
        title: 'Order',
        icon: 'sellingBuy'
      }
    }]
  },
  {
    path: '/order_courier',
    component: Layout,
    meta: {
      roles: ['courier']
    },
    children: [{
      path: '/order_courier',
      name: 'order_courier',
      component: () => import('@/views/order_courier/index'),
      meta: {
        title: 'Order',
        icon: 'sellingBuy'
      }
    }]
  },
  {
    path: '/logistic',
    component: Layout,
    meta: {
      roles: ['courier']
    },
    children: [{
      path: '/logistic',
      name: 'logistic',
      component: () => import('@/views/logistic/index'),
      meta: {
        title: 'Logistic',
        icon: 'donatingGrantee'
      }
    }]
  },
  {
    path: '/logistic',
    component: Layout,
    meta: {
      roles: ['admin']
    },
    children: [{
      path: '/logistic',
      name: 'logistic',
      component: () => import('@/views/logistic/index'),
      meta: {
        title: 'Logistic',
        icon: 'donatingGrantee'
      }
    }]
  },


  // 404 page must be placed at the end !!!
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

const createRouter = () => new Router({
  base: '/web',
  // mode: 'history', // require service support
  scrollBehavior: () => ({
    y: 0
  }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
