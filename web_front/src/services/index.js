import { authApi } from "./admin/auth";
import { normalApi } from "./admin/normal";

export const adminApi = {};

export const userApi = {
  ...authApi,
  ...normalApi,
};
