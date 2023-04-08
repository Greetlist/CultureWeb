import Vue from "vue";
import VueRouter from "vue-router";
import LoginView from "@views/Login.vue";
import Layout from "@/layout/index.vue";
import OverviewIndex from "@/views/Admin/Overview/index.vue";
import SystemIndex from "@/views/Admin/SystemConfig/index.vue";
import ArticalIndex from "@/views/Admin/Artical/index.vue";
import ActivityIndex from "@/views/Admin/Activities/index.vue";
import SiteIndex from "@/views/Admin/Site/index.vue";
import UserIndex from "@/views/Admin/UserManage/index.vue";

Vue.use(VueRouter);

const adminRoutes = {
  path: "/admin",
  name: "admin",
  redirect: "/admin/overview/userOverview",
  component: Layout,
  children: [
    {
      path: "/admin/overview",
      name: "overview",
      component: OverviewIndex,
      meta: {
        title: "总览",
        icon: "el-icon-pie-chart",
      },
      children: [
        {
          path: "userOverview",
          name: "userOverview",
          component: () => import("@/views/Admin/Overview/User.vue"),
          meta: {
            title: "用户总览",
          },
        },
        {
          path: "visitOverview",
          name: "visitOverview",
          component: () => import("@/views/Admin/Overview/Visit.vue"),
          meta: {
            title: "访问总览",
          },
        },
        {
          path: "storageOverview",
          name: "storageOverview",
          component: () => import("@/views/Admin/Overview/Storage.vue"),
          meta: {
            title: "存储总览",
          },
        },
      ],
    },
    {
      path: "/admin/systemConfig",
      name: "overview",
      component: SystemIndex,
      meta: {
        title: "系统设置",
        icon: "el-icon-setting",
      },
      children: [
        {
          path: "baseInfo",
          name: "baseInfo",
          component: () => import("@/views/Admin/SystemConfig/BaseInfo.vue"),
          meta: {
            title: "基本信息",
          },
        },
        {
          path: "baseConfig",
          name: "baseConfig",
          component: () => import("@/views/Admin/SystemConfig/BaseConfig.vue"),
          meta: {
            title: "参数设置",
          },
        },
      ],
    },
    {
      path: "/admin/userManage",
      name: "userManage",
      component: UserIndex,
      meta: {
        title: "用户设置",
        icon: "el-icon-user",
      },
      children: [
        {
          path: "totalUserList",
          name: "totalUserList",
          component: () => import("@/views/Admin/UserManage/TotalUserList.vue"),
          meta: {
            title: "所有用户",
          },
        },
        {
          path: "addUser",
          name: "addUser",
          component: () => import("@/views/Admin/UserManage/AddUser.vue"),
          meta: {
            title: "新增用户",
          },
        },
      ],
    },
    {
      path: "/admin/articalManage",
      name: "articalManage",
      component: ArticalIndex,
      meta: {
        title: "文章管理",
        icon: "el-icon-reading",
      },
      children: [
        {
          path: "articalList",
          name: "articalList",
          component: () => import("@/views/Admin/Artical/ArticalList.vue"),
          meta: {
            title: "文章列表",
          },
        },
        {
          path: "addArtical",
          name: "addArtical",
          component: () => import("@/views/Admin/Artical/AddArtical.vue"),
          meta: {
            title: "新增文章",
          },
        },
      ],
    },
    {
      path: "/admin/activityManage",
      name: "activityManage",
      component: ActivityIndex,
      meta: {
        title: "活动管理",
        icon: "el-icon-alarm-clock",
      },
      children: [
        {
          path: "activitiesList",
          name: "activitiesList",
          component: () =>
            import("@/views/Admin/Activities/ActivitiesList.vue"),
          meta: {
            title: "活动列表",
          },
        },
        {
          path: "addActivity",
          name: "addActivity",
          component: () => import("@/views/Admin/Activities/AddActivity.vue"),
          meta: {
            title: "新增活动",
          },
        },
      ],
    },
    {
      path: "/admin/siteManage",
      name: "siteManage",
      component: SiteIndex,
      meta: {
        title: "场馆预约管理",
        icon: "el-icon-location-outline",
      },
      children: [
        {
          path: "siteManagementLis",
          name: "siteManagementLis",
          component: () => import("@/views/Admin/Site/SiteList.vue"),
          meta: {
            title: "所有预约",
          },
        },
        {
          path: "siteManagementList",
          name: "siteManagementList",
          component: () => import("@/views/Admin/Site/AddNewSite.vue"),
          meta: {
            title: "新增场地",
          },
        },
      ],
    },
  ],
};

const routes = [
  {
    path: "/",
    name: "login",
    component: LoginView,
    hidden: true,
  },
  {
    path: "/home",
    name: "home",
    component: () => import("@/views/User/Home.vue"),
    children: [],
    hidden: true,
  },
  adminRoutes,
];

const router = new VueRouter({
  mode: "history",
  routes,
});

export default router;
export { adminRoutes };
