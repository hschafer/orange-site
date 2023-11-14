import auth from './modules/auth';
import createPersistedState from 'vuex-persistedstate'
import vuex from 'vuex'

export default new vuex.Store({
    modules: {
        auth
    },
    plugins: [createPersistedState()]
})