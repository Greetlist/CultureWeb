import { authApi } from "./user/auth";
import { normalApi } from "./user/normal";

export const adminApi = {};

export const userApi = {
  ...authApi,
  ...normalApi,
};
