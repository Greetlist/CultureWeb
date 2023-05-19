import { Get, Post } from "@js/server";
import { baseURL } from "@js/server";

export function addUser(params) {
  return Post("/api/admin/addUser", params);
}

export function getTotalUserInfo() {
  return Get("/api/admin/getTotalUserInfo", {});
}

export function getUserInfo(params) {
  return Get("/api/admin/getUserInfo", params);
}

export function submitArticle(params) {
  return Post("/api/admin/submitArticle", params)
}

export function getTotalArticle() {
  return Get("/api/admin/getTotalArticle", {});
}

export function getArticleInfo(params) {
  return Post("/api/user/normal/getArticleInfo", params);
}

export function searchArticle(params) {
  return Post("/api/user/normal/searchArticle", params);
}

export function batchDeleteArticle(params) {
  return Post("/api/admin/batchDeleteArticle", params)
}

export function batchModifyArticle(params) {
  return Post("/api/admin/batchModifyArticle", params)
}

export function getArticleContent(params) {
  return Post("/api/user/normal/getArticleContent", params)
}

export function getTotalLabel() {
  return Get("/api/user/normal/getTotalLabel")
}

export function addSingleLabel(params) {
  return Post("/api/admin/addSingleLabel", params)
}

export function deleteLabel(params) {
  return Post("/api/admin/deleteLabel", params)
}

export function batchModifyLabel(params) {
  return Post("/api/admin/modifyLabel", params)
}

export function submitActivity(params) {
  return Post("/api/admin/submitActivity", params)
}

export function getTotalActivity() {
  return Get("/api/admin/getTotalActivity", {});
}

export function getActivityInfo(params) {
  return Post("/api/user/normal/getActivityInfo", params);
}

export function searchActivity(params) {
  return Post("/api/user/normal/searchActivity", params);
}

export function batchDeleteActivity(params) {
  return Post("/api/admin/batchDeleteActivity", params)
}

export function batchModifyActivity(params) {
  return Post("/api/admin/batchModifyActivity", params)
}

export function getActivityContent(params) {
  return Post("/api/user/normal/getActivityContent", params)
}

export const uploadMediaURL = baseURL + "/api/admin/saveMedia";

export const adminApi = {
  addUser,
  getTotalUserInfo,
  getUserInfo,

  submitArticle,
  getTotalArticle,
  batchDeleteArticle,
  batchModifyArticle,
  getArticleInfo,
  searchArticle,
  getArticleContent,

  getTotalLabel,
  addSingleLabel,
  deleteLabel,
  batchModifyLabel,

  submitActivity,
  getTotalActivity,
  batchDeleteActivity,
  batchModifyActivity,
  getActivityInfo,
  searchActivity,
  getActivityContent
};
