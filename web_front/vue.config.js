const { defineConfig } = require("@vue/cli-service");
const path = require("path");

module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  configureWebpack: {
    resolve: {
      extensions: [".js", ".css", ".vue"], // 自动补全识别后缀
      alias: {
        "@": path.resolve(__dirname, "src"),
        "@js": path.resolve(__dirname, "src/assets/js"),
        "@style": path.resolve(__dirname, "src/assets/style"),
        "@img": path.resolve(__dirname, "src/assets/images"),
        "@c": path.resolve(__dirname, "src/components"),
        "@views": path.resolve(__dirname, "src/views"),
        "@services": path.resolve(__dirname, "src/services"),
      },
    },
    module: {
      rules: [],
    },
  },
});
