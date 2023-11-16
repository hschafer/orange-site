import axios from 'axios';

const state = {
    username: null,
    token: null,
};

const getters = {
    isAuthenticated: (state) => state.username != null && state.token != null,
    getUsername: (state) => state.username,
    getToken: (state) => state.token,
    getAllState: (state) => state,
};

async function loginOrRegister(route, commit, username, password) {
    return axios.post(route, {
        "username": username,
        "password": password,
    }).then((response) => {
        commit("setUsername", username)
        commit("setToken", response.data.AuthToken)
    }).catch((error) => {
        // Clear session info just in case
        commit("setUsername", null)
        commit("setToken", null)
        throw error;
    });
}

const actions = {
    async register({commit}, {username, password}) {
        return loginOrRegister("/api/register", commit, username, password)
    },

    async login({ commit }, {username, password}) {
        // TODO more password validation?
        return loginOrRegister("/api/login", commit, username, password)
    },

    logout({ commit }) {
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
    }
};

export default {
  state,
  getters,
  actions,
  mutations
};