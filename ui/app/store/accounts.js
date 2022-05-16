export const state = () => ({
    isLoggedIn: false,
    user: [],
})

export const getters = {
    isLoggedIn: (state) => state.isLoggedIn,
    user: (state) => state.user
}

export const mutations = {
    setUser(state, user) {
        state.user = user
        state.isLoggedIn = true
    }
}

export const actions = {
    login({ commit }, authData) {
        this.$axios
            .$post("/login", {
                name: authData.name,
                password: authData.password
            })
            .then(response => {
                commit("setUser", response)
                console.log(response)
                this.$router.push("/accounts/mypage")
            });
    },
    signup({ commit }, authData) {
        this.$axios
            .$post("/signup", {
                name: authData.name,
                password: authData.password,
                job: authData.job,
                age: authData.age,
                sex: authData.sex
            })
            .then(response => {
                console.log(response)
                this.$router.push("/accounts/login")
            });
    },
}
