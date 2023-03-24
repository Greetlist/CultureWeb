import axios from "axios";
import https from "https";
//axios.defaults.baseURL = process.env.VUE_APP_BASE_API;
export const baseURL = "https://121.5.100.186:8900";
axios.defaults.baseURL = baseURL;

const httpsAgent = new https.Agent({
  rejectUnauthorized: false,
})
axios.defaults.httpsAgent = httpsAgent
const header = {
  'Access-Control-Allow-Origin': '*',
}

export const Get = (url, params = {}) =>
  new Promise((resolve) => {
    axios
      .get(url, params, { header })
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
      .post(url, data, { header })
      .then((result) => {
        resolve(result);
      })
      .catch((err) => {
        resolve(err);
      });
  });
};
