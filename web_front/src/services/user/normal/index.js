import { Post, Get } from "@js/server";

export function login(params) {
  return Post("/api/user/normal/login", params);
}

export function logout(params) {
  return Post("/api/user/normal/loginout", params);
}

export function getWebBasicInfo() {
  return Get("/api/user/normal/getWebBasicInfo", {})
}

export function checkLogin() {
  return Get("/api/user/normal/checkLogin", {})
}

export const userNormalApi = {
  login,
  logout,
  getWebBasicInfo,
  checkLogin
};
