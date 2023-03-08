export const adminMenuList = [
  {
    name: "总览",
    order: "1",
    path: "overview",
    headIcon: "el-icon-pie-chart",
    children: [
      {
        path: "userOverview",
        name: "用户总览",
      },
      {
        path: "visitOverview",
        name: "访问总览",
      },
      {
        path: "storageOverview",
        name: "存储总览",
      },
    ],
  },
  {
    name: "系统设置",
    order: "2",
    path: "systemConfig",
    headIcon: "el-icon-setting",
    children: [
      {
        path: "baseInfo",
        name: "基本信息",
      },
      {
        path: "baseConfig",
        name: "参数设置",
      },
    ],
  },
  {
    name: "用户管理",
    order: "3",
    path: "userManage",
    headIcon: "el-icon-user",
    children: [
      {
        path: "totalUserList",
        name: "所有用户",
      },
      {
        path: "addUser",
        name: "新增用户",
      },
    ],
  },
  {
    name: "文章管理",
    order: "4",
    path: "Artical",
    headIcon: "el-icon-reading",
    children: [
      {
        path: "articalList",
        name: "文章列表",
      },
      {
        path: "addArtical",
        name: "新增文章",
      },
    ],
  },
  {
    name: "活动管理",
    order: "5",
    path: "activities",
    headIcon: "el-icon-alarm-clock",
    children: [
      {
        path: "activitiesList",
        name: "活动列表",
      },
      {
        path: "addActivity",
        name: "新增活动",
      },
    ],
  },
  {
    name: "场馆预约管理",
    order: "6",
    path: "site",
    headIcon: "el-icon-location-outline",
    children: [
      {
        path: "siteManagementLis",
        name: "所有预约",
      },
      {
        path: "addNewSite",
        name: "新增场地",
      },
    ],
  },
];
