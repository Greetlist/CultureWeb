import Vue from "vue";
import VueCookies from 'vue-cookies'
import App from "./App.vue";
import router from "./router";
import store from "./store";
import ElementUI from "element-ui";
import "@style/element-variables.scss";

import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

Vue.config.productionTip = false;
Vue.use(ElementUI);
Vue.use(VueQuillEditor);
Vue.use(VueCookies);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
