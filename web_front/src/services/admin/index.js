import { Get, Post } from "@js/server";
import { baseURL } from "@js/server";

export function addUser(params) {
  return Post("/api/admin/addUser", params);
}

export function submitArticle(params) {
  return Post("/api/admin/submitArticle", params)
}

export function getTotalUserInfo() {
  return Get("/api/admin/getTotalUserInfo", {});
}

export function getUserInfo(params) {
  return Get("/api/admin/getUserInfo", params);
}

export function getTotalArticle() {
  return Get("/api/admin/getTotalArticle", {});
}

export function batchModifyArticle(params) {
  return Post("/api/admin/batchModifyArticle", params)
}

export function getTotalLabel() {
  return Get("/api/user/normal/getTotalLabel")
}

export function addSingleLabel(params) {
  return Post("/api/admin/addSingleLabel", params)
}

export const uploadMediaURL = baseURL + "/api/admin/saveMedia";

export const adminApi = {
  addUser,
  submitArticle,
  getTotalUserInfo,
  getUserInfo,
  getTotalArticle,
  batchModifyArticle,
  getTotalLabel,
  addSingleLabel
};
