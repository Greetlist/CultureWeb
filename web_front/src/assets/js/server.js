import axios from "axios";
axios.defaults.baseURL = process.env.VUE_APP_BASE_API;

export const Get = (url, params = {}) =>
  new Promise((resolve) => {
    axios
      .get(url, params)
      .then((result) => {
        resolve(result);
      })
      .catch((err) => {
        resolve(err);
      });
  });

export const Post = (url, data = {}) => {
  return new Promise((resolve) => {
    axios
      .post(url, data)
      .then((result) => {
        resolve(result);
      })
      .catch((err) => {
        resolve(err);
      });
  });
};
