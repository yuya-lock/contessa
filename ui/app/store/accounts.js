export const state = () => ({
    name: "",
    password: "",
    age: "",
    job: "",
    sex: "",
})

export const getters = {
    name: (state) => state.name,
    password: (state) => state.password,
    age: (state) => state.age,
    job: (state) => state.job,
    sex: (state) => state.sex,
}

export const mutations = {
    setName(state, name) {
        state.name = name
    },
    setPassword(state, password) {
        state.password = password
    },
    setAge(state, age) {
        state.age = age
    },
    setJob(state, job) {
        state.job = job
    },
    setSex(state, sex) {
        state.sex = sex
    },
}

export const actions = {
    async signup(context) {
        const response = await this.$axios.$
    }
}
