const { defineConfig } = require("@vue/cli-service");

module.exports = defineConfig({
  transpileDependencies: true,
  configureWebpack: {
    resolve: {
      extensions: [".js", ".css", ".vue"], // 自动补全识别后缀
      alias: {
        assets: "@/assets",
        components: "@/components",
        views: "@/views",
      },
    },
    module: {
      rules: [],
    },
  },
});
