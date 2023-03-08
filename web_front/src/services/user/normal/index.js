import { Post } from "@js/server";

export function login(params) {
  return Post("/api/user/normal/login", params);
}

export function loginout(params) {
  return Post("/api/user/normal/loginout", params);
}

export const normalApi = {
  login,
  loginout,
};
