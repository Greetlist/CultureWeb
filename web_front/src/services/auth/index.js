import { Post } from "@js/server";

export function login(params) {
  return Post("/api/auth/login", params);
}

export function loginout(params) {
  return Post("/api/auth/loginout", params);
}

export const authApi = {
  login,
};
