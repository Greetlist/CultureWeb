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
    component: () => import("@/views/Home/Home.vue"),
    children: [
      {
        path: "/userManage",
        name: "userManage",
        component: () => import("@/views/Home/UserManage.vue"),
      },
    ],
  },
  {
    path: "/admin",
    name: "admin",
    component: () => import("@/views/Admin/AdminHome.vue"),
    children: [
      {
        path: "/addUser",
        name: "addUser",
        component: () => import("@/views/Admin/AddUser.vue"),
      }
    ]
  },
  {
    path: "/addUserTest",
    name: "addUserTest",
    component: () => import("@/views/Admin/AddUser.vue"),
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
