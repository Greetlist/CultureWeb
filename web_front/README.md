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

### Quill 富文本编辑器的各种司马问题
1. 当文章过长的时候，滚动条是父组件的滚动条，不是编辑器的滚动条
这个问题不能直接在自己的代码CSS里面写,写了没逼用
```
.ql-container {
  overflow: auto;
  max-height: 33rem;
}
```
只能改源码里面的css:

> ./node_modules/quill/dist/quill.bubble.css

还他妈只能改这个文件里的```.ql-container```

2. 往编辑器里面粘贴的时候，滚动条直接回到顶部
也是直接改源码:

> ./node_modules/quill/dist/quill.bubble.css

```
.ql-clipboard {
  position: fixed;
  display: none;
}
```

3. 粘贴的时候，光标停止在当前位置
可以监听change事件，手动把光标定位到最后

```
this.$refs.article.quill.setSelection(event.text.length+1)
setTimeout(() => this.$refs.article.quill.setSelection(event.text.length+1), 0)
```
然后上面的代码第一行不管用，必须用第二种写法，嘻嘻
