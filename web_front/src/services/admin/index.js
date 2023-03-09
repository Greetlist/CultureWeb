import { Get, Post } from "@js/server";

export function addUser(params) {
  return Post("/api/admin/addUser", params);
}

export function getTotalUserInfo() {
  return Get("/api/admin/getTotalUserInfo", {});
}

export function getUserInfo(params) {
  return Get("/api/admin/getUserInfo", paramas);
}

export const adminApi = {
  addUser,
  getTotalUserInfo,
  getUserInfo,
};
