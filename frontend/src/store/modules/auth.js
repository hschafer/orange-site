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

async function loginOrRegister(route, commit, username, password) {
    return axios.post(route, {
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
    });
    // TODO handle login message
}

const actions = {
    async register({commit}, {username, password}) {
        return loginOrRegister("/api/register", commit, username, password)
    },

    async login({ commit }, {username, password}) {
        // TODO more password validation:?
        return loginOrRegister("/api/login", commit, username, password)
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