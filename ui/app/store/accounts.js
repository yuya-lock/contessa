export const state = () => ({
    user: [],
    token: null,
    favorite_cocktails: [],
    mycomment_cocktails: [],
    myrate_cocktails: []
})

export const getters = {
    user: (state) => state.user,
    token: (state) => state.token,
    favorite_cocktails: (state) => state.favorite_cocktails,
    mycomment_cocktails: (state) => state.mycomment_cocktails,
    myrate_cocktails: (state) => state.myrate_cocktails,
}

export const mutations = {
    setUser(state, user) {
        state.user = user
    },
    updateToken(state, token) {
        state.token = token
    },
    setFavoriteCocktails(state, favorite_cocktails) {
        state.favorite_cocktails = favorite_cocktails
    },
    setMyCommentCocktails(state, mycomment_cocktails) {
        state.mycomment_cocktails = mycomment_cocktails
    },
    setMyRateCocktails(state, myrate_cocktails) {
        state.myrate_cocktails = myrate_cocktails
    },
}

export const actions = {
    login({ commit }, authData) {
        this.$axios
            .$post("/login", {
                name: authData.name,
                password: authData.password
            })
            .then(response => {
                commit("updateToken", response.token)
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
                this.$router.push("/accounts/login")
            });
    },
    mypage({ state, commit }) {
        this.$axios
            .$get("/restricted/mypage", {
                headers: {
                    Authorization: `Bearer ${state.token}`
                }
            })
            .then(response => {
                commit("setUser", response.user)
            })
    },
    getFavoriteCocktails({ state, commit }) {
        this.$axios
            .$get("/restricted/favoritecocktails", {
                headers: {
                    Authorization: `Bearer ${state.token}`
                }
            })
            .then(response => {
                commit("setFavoriteCocktails", response.favoriteCocktails)
            })
    },
    getMyCommentCocktails({ state, commit }) {
        this.$axios
            .$get("/restricted/mycomments", {
                headers: {
                    Authorization: `Bearer ${state.token}`
                }
            })
            .then(response => {
                commit("setMyCommentCocktails", response.myCommentCocktails)
                for (const comment of state.user.Comments) {
                    for (const cocktail of state.mycomment_cocktails) {
                        if (comment.cocktail_id == cocktail.cocktail_id) {
                            cocktail["comment_body"] = comment.body
                        }
                    }
                }
            })
    },
    getMyRateCocktails({ state, commit }) {
        this.$axios
            .$get("/restricted/myrates", {
                headers: {
                    Authorization: `Bearer ${state.token}`
                }
            })
            .then(response => {
                commit("setMyRateCocktails", response.myRateCocktails)
                for (const rate of state.user.Rates) {
                    for (const cocktail of state.myrate_cocktails) {
                        if (rate.cocktail_id == cocktail.cocktail_id) {
                            cocktail["rating"] = rate.rating
                        }
                    }
                }
            })
    },
}
