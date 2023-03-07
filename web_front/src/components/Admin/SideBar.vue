<template>
  <div>
    <el-scrollbar class='page-component__scroll'>
      <transition :duration="{ enter: 800, leave: 100}" mode="out-in">
        <el-menu
          :collapse="false"
          :collapse-transition="true"
          :default-active="active"
          :default-openeds="overview"
          class="el-menu-vertical"
          unique-opened
          router
          background-color="#545c64"
          active-text-color="#ffd04b"
          text-color="#fff"
        >
          <el-submenu v-for="(singleMenu, idx) in menuList" :index="idx.toString()" :key="singleMenu.id" style="text-align: left">
            <template slot="title">
              <i :class="singleMenu.headIcon"></i>
              {{ singleMenu.name }}
            </template>
            <el-menu-item-group>
              <el-menu-item v-for="item in singleMenu.subMenuList" :index="item.route" :key="item.id">
              {{ item.name }}
              </el-menu-item>
            </el-menu-item-group>
          </el-submenu>
        </el-menu>
      </transition>
    </el-scrollbar>
  </div>
</template>

<script>
var menuList = [
  {
    name: '总览',
    id: 'overview',
    headIcon: 'el-icon-s-data',
    subMenuList: [
      {
        name: '用户总览',
        id: 'user overview',
        route: '/admin/userOverview'
      },
      {
        name: '访问总览',
        id: 'request overview',
        route: '/admin/requestOverview'
      },
      {
        name: '存储总览',
        id: 'storage overview',
        route: '/admin/storageOverview'
      }
    ]
  },
  {
    name: '系统设置',
    id: 'system config',
    headIcon: 'el-icon-cpu',
    subMenuList: [
      {
        name: '基本信息',
        id: 'baseInfo',
        route: '/admin/baseInfo'
      },
      {
        name: '参数设置',
        id: 'baseConfig',
        route: '/admin/baseConfig'
      }
    ]
  },
  {
    name: '用户管理',
    id: 'user-control',
    headIcon: 'el-icon-alarm-clock',
    subMenuList: [
      {
        name: '所有用户',
        id: 'total-user-list',
        route: '/admin/totalUserList'
      },
      {
        name: '新增用户',
        id: 'add-user',
        route: '/admin/addUser'
      }
    ]
  },
  {
    name: '文章管理',
    id: 'article',
    headIcon: 'el-icon-s-comment',
    subMenuList: [
      {
        name: '文章列表',
        id: 'article-list',
        route: '/admin/articleList'
      },
      {
        name: '新增文章',
        id: 'add-new-article',
        route: '/admin/newArticle'
      }
    ]
  },
  {
    name: '活动管理',
    id: 'activity',
    headIcon: 'el-icon-printer',
    subMenuList: [
      {
        name: '活动列表',
        id: 'activity-list',
        route: '/admin/activityList'
      },
      {
        name: '新增活动',
        id: 'new-activity',
        route: '/admin/newActivity'
      }
    ]
  },
  {
    name: '场馆预约管理',
    id: 'site',
    headIcon: 'el-icon-setting',
    subMenuList: [
      {
        name: '所有预约',
        id: 'total-site-management',
        route: '/admin/siteManagementList'
      },
      {
        name: '新增场地',
        id: 'add-new-site',
        route: '/admin/addNewSite'
      }
    ]
  },
  {
    name: 'Enjoy PornHub',
    id: 'pornhub',
    headIcon: 'el-icon-s-promotion',
    subMenuList: [
      {
        name: 'Asia',
        id: 'asia',
        route: 'cn.pornhub.com'
      },
      {
        name: 'Japan',
        id: 'japan',
        route: 'japan.pornhub.com'
      }
    ]
  }
]

export default {
  name: 'SideBar',
  data: function () {
    return {
      menuList: menuList
    }
  }
}
</script>
