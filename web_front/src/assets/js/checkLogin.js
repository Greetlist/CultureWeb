import { Get, Post } from '@js/server';

export function checkLogin() {
  const res = await Get("/api/user/normal/checkLogin")
  console.log(res)
  return res
}
