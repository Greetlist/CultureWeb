import { Post } from "@js/server";

export function login(params) {
  return Post("/api/user/normal/login", params);
}

export function logout(params) {
  return Post("/api/user/normal/loginout", params);
}

export const normalApi = {
  login,
  logout,
};
