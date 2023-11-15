import axios from 'axios';

const state = {
    username: null,
    token: null,
    loginMessage: null
};

const getters = {
    isAuthenticated: (state) => state.username != null && state.token != null,
    getUsername: (state) => state.username,
    getToken: (state) => state.token,
    getAllState: (state) => state,
    getLoginMessage: (state) => state.loginMessage
};

const actions = {
    async register({dispatch}, form) {
        // TODO
    },

    async login({ commit }, {username, password}) {
        // TODO more password validation:?
        return axios.post("/api/login", {
            "username": username,
            "password": password,
        }).then((response) => {
            commit("setUsername", username)
            commit("setToken", response.data.AuthToken)
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
        commit("setUsername", null)
        commit("setToken", null)
    }
};

const mutations = {
    setUsername(state, username) {
        state.username = username
    },
    setToken(state, token) {
        state.token = token
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