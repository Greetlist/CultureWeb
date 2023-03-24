# web_front

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Directory Structure
```
 ├─ web_front
     ├─ main.js         // 项目入口文件
     ├─ .env.development    // 开发环境配置文件
     ├─ .env.production     // 生产环境配置文件
     ├─ vue.config.js       // 项目配置文件
     ├─ src             // 开发目录
     │  ├─ assets           // 公用静态资源
     │  │  ├─ images           // 图片
     │  │  ├─ js               // 工具类
     │  │  ├─ style            // scss全局配置
     │  │  ├─ constant            // 常量文件    
     │  ├─ components	// 公用组件
     │  ├─ router	    // router
     │  ├─ store	    // store
     │  ├─ services	    // 接口定义(文件结构与后端保持一致)
     │  │  ├─ user      // 用户侧接口
     │  │  ├─ admin     // 管理测接口
     │  ├─ views        // 包含所有单页应用的文件夹
     │  │  ├─ User      // 用户侧页面
     │  │  ├─ Admin     // 管理侧页面    
```
### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
