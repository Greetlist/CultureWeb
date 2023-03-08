import Vue from "vue";
import VueRouter from "vue-router";
import LoginView from "@views/Login.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "login",
    component: LoginView,
  },
  {
    path: "/home",
    name: "home",
    component: () => import("@/views/User/Home.vue"),
    children: [],
  },
  {
    path: "/admin",
    name: "admin",
    component: () => import("@/views/Admin/Home.vue"),
    children: [
      {
        path: "/userOverview",
        name: "userOverview", // 用户总览
        component: () => import("@/views/Admin/Overview/User.vue"),
      },
      {
        path: "/visitOverview",
        name: "visitOverview", // 访问总览
        component: () => import("@/views/Admin/Overview/Visit.vue"),
      },
      {
        path: "/storageOverview",
        name: "storageOverview", // 存储总览
        component: () => import("@/views/Admin/Overview/Storage.vue"),
      },
      {
        path: "/baseInfo",
        name: "baseInfo", // 基本信息
        component: () => import("@/views/Admin/SystemConfig/BaseInfo.vue"),
      },
      {
        path: "/baseConfig",
        name: "baseConfig", // 基本信息
        component: () => import("@/views/Admin/SystemConfig/BaseConfig.vue"),
      },
      {
        path: "/totalUserList",
        name: "totalUserList", // 所有用户
        component: () => import("@/views/Admin/UserManage/TotalUserList.vue"),
      },
      {
        path: "/addUser",
        name: "addUser", // 新增用户
        component: () => import("@/views/Admin/UserManage/AddUser.vue"),
      },
      {
        path: "/articalList",
        name: "articalList", // 活动列表
        component: () => import("@/views/Admin/Artical/ArticalList.vue"),
      },
      {
        path: "/addArtical",
        name: "addArtical", // 新增活动
        component: () => import("@/views/Admin/Artical/AddArtical.vue"),
      },
      {
        path: "/activitiesList",
        name: "activitiesList", // 活动列表
        component: () => import("@/views/Admin/Activities/ActivitiesList.vue"),
      },
      {
        path: "/addActivity",
        name: "addActivity", // 新增活动
        component: () => import("@/views/Admin/Activities/AddActivity.vue"),
      },
      {
        path: "/siteManagementLis",
        name: "siteManagementLis", // 所有预约
        component: () => import("@/views/Admin/Site/SiteList.vue"),
      },
      {
        path: "/addNewSite",
        name: "addNewSite", // 新增场地
        component: () => import("@/views/Admin/Site/AddNewSite.vue"),
      },
    ],
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
