import axios from 'axios';

const state = {
    user: null,  // {name, token}
    loginMessage: null
};

const getters = {
    isAuthenticated: (state) => !state.user,
    getUsername: (state) => state.user?.name,
    getToken: (state) => state.token?.token,
    getLoginMessage: (state) => state.loginMessage
};

const actions = {
    async register({dispatch}, form) {
        // TODO
    },

    async login({ commit }, {username, password}) {
        // TODO more password validation:?
        axios.post("/api/login", {
          "username": username,
          "password": password,
        }).then((response) => {
          commit("setUser", {
            "name": username,
            "token": response.data.Token
          })
          commit("setLoginMessage", "")  // Clear login message
        }).catch((error) => {
          if (error.response.data.Message) {
            commit("setLoginMessage", error.response.data.Message)
          } else {
            commit("setLoginMessage", "Unable to log in")
          }
        })
    },

    async logout({ commit }) {
        commit("setUser", null)
    }
};

const mutations = {
    setUser(state, user) {
        state.user = user
    },
    setLoginMessage(state, message) {
        state.loginMessage = message
    }
};

export default {
  state,
  getters,
  actions,
  mutations
};