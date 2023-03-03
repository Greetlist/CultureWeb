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
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
