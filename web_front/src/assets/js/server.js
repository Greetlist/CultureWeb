import axios from "axios";
import https from "https";
//axios.defaults.baseURL = process.env.VUE_APP_BASE_API;
//export const baseURL = "https://121.5.100.186:8900";
//const CA = "-----BEGIN CERTIFICATE-----MIIDqzCCApOgAwIBAgIUKGU7VvJw6jv4O86PP9bBjEVaG6swDQYJKoZIhvcNAQELBQAwZTELMAkGA1UEBhMCQ04xETAPBgNVBAgMCHNoYW5naGFpMREwDwYDVQQHDAhzaGFuZ2hhaTESMBAGA1UECgwJZ3Jwcml2YXRlMQswCQYDVQQLDAJJVDEPMA0GA1UEAwwGZG9tYWluMB4XDTIzMDMwMjAzNDU1MloXDTMzMDIyNzAzNDU1MlowZTELMAkGA1UEBhMCQ04xETAPBgNVBAgMCHNoYW5naGFpMREwDwYDVQQHDAhzaGFuZ2hhaTESMBAGA1UECgwJZ3Jwcml2YXRlMQswCQYDVQQLDAJJVDEPMA0GA1UEAwwGZG9tYWluMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArr7PfH/0RTvK7YXsA9LxLFZBXuUCq7ewZvpMF233C32K283PGKkzHEaIco0gqOz2vgQ+bd5ccTh5vTMeb0EPZt+T33hQz94190zq/l9qDKX0RRYzglmnSQX3DlHT5/L18RtmdtzqnOW9QmxBnnsYt7nq0dCaTJDdxGy5R7yh6cTi5np99PX2SwoFVjC62O2oFA+RWNIN0ykVw07HXI6Dv84Rgh/pKzphm8X20vJoJizxlxKuJtk8B6GvV8fyednrIAEhW5Rh8Gufdk+odtvx0VmkS7mTnPMPaZf4CYV19q9KXQeBmsVUQZBK1zHof5W23W3fLOKiCHEfl53AoBodoQIDAQABo1MwUTAdBgNVHQ4EFgQUOoxWYtTqbqOsmH28hJcjvIseLwwwHwYDVR0jBBgwFoAUOoxWYtTqbqOsmH28hJcjvIseLwwwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAX7bZ6TZflFHjluUkaxLyYGXsWJBng/eUKAQaMM8c/vmIxuy4SeDMWPlxpbSkZFGVifSCuCfJGwLMuxZfsUjlij8WshgQHYu5PMzmm3vl6lI22uBkexclwlxTbDQ5h0r/uLIY9HkUcQ3ycroddYGzljTGjQSqL5uFL2qGMSXFFCvjrT1Eyo/Vz8m/C3sNO8ZmDPtHab6sqIhEXHGZVY9juMkUfG7IgCgjPYLkM0kjY6jD8IyKvYthYyWswdYdxmVriyj9IvXhGqyDhaxHPUXETcoL6LtnjmRDuTJwkvwVJu7J3b+V53ccOnkkJDd09+5yTkmuIqVdpU9oHGdjXbbO3g==-----END CERTIFICATE-----"
//const httpsAgent = new https.Agent({
//  ca: CA,
//  rejectUnauthorized: false,
//})
//axios.defaults.httpsAgent = httpsAgent

export const baseURL = "http://121.5.100.186:8082";
axios.defaults.baseURL = baseURL;
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
