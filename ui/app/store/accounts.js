export const state = () => ({
    isLoggedIn: false,
    user: null,
    token: ""
})

export const getters = {
    isLoggedIn: (state) => state.isLoggedIn,
    user: (state) => state.user,
    token: (state) => state.token
}

export const mutations = {
    setUser(state, { user }) {
        state.user = user
    },
    setIsLoggedIn(state, isLoggedIn) {
        state.isLoggedIn = isLoggedIn
    },
    setToken(state, token) {
        state.token = token
    }
}

export const actions = {
    async signup({ commit }, { user }) {
        const response = await this.$axios.$post("/signup", {
            name: user.name,
            password: user.password,
            job: user.job,
            age: user.age,
            sex: user.sex
        }).catch(err => {
            return err.response
        })
        commit("setUser", { response })
        await this.$router.push("/accounts/login")
        console.log(response)
    },
    async login({ commit }, { user }) {
        const response = await this.$axios.$post("/login", {
            name: user.name,
            password: user.password
        }).catch(err => {
            return err.response
        })
        await this.$router.push("/")
        console.log(response)
    }
}
