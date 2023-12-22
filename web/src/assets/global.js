// store/index.js
import { createStore } from 'vuex'

const store = createStore({
  state() {
    return {
      counter: 0,
      user : {}
    }
  },

  mutations: {
    SET_USER(state, user) {
      state.user = user;
    }
  },
  actions: {

  },
  getters: {
    
  }
})

export default store
