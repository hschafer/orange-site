import axios from 'axios';

const state = {
    user: null
};

const getters = {
    isAuthenticated: (state) => !state.user,
    getUser: (state) => state.user
};

const actions = {
    async register({dispatch}, form) {
        // TODO
    },

    async login({ commit }, user) {
        await axios.post("/login", user)
        await commit("setUser", user.get("username"))
    },

    async logout({ commit }) {
        commit("logout")
    }

};

const mutations = {
    setUser(state, username) {
        state.user = username
    },
    logout(state) {
        state.user = null;
    }

};

export default {
  state,
  getters,
  actions,
  mutations
};