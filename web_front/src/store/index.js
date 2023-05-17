import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    webBasicInfo: {}
  },
  getters: {
    getWebBasicInfo(state) {
      return state.webBasicInfo
    }
  },
  mutations: {
    setWebBasicInfo(state, info) {
      state.webBasicInfo = info
    }
  },
  actions: {},
  modules: {},
});
