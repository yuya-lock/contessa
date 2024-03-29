export const state = () => ({
    total_pages: 1,
    current_page: 1,
    word: "",
    base: "",
    technique: "",
    taste: "",
    style: "",
    alcohol_from: 0,
    alcohol_to: 50,
    top: "",
    glass: "",
    page: 1,
    limit: 10,
    cocktails: [],
    cocktail_detail: [],
    showPagination: false,
    ave_rating: 0,
    ratings_count: 0,
})

export const getters = {
    total_pages: (state) => state.total_pages,
    current_page: (state) => state.current_page,
    word: (state) => state.word,
    base: (state) => state.base,
    technique: (state) => state.technique,
    taste: (state) => state.taste,
    style: (state) => state.style,
    alcohol_from: (state) => state.alcohol_from,
    alcohol_to: (state) => state.alcohol_to,
    top: (state) => state.top,
    glass: (state) => state.glass,
    page: (state) => state.page,
    limit: (state) => state.limit,
    cocktails: (state) => state.cocktails,
    cocktail_detail: (state) => state.cocktail_detail,
    showPagination: (state) => state.showPagination,
    ave_rating: (state) => state.ave_rating,
    ratings_count: (state) => state.ratings_count,
}

export const mutations = {
    setTotalPages(state, total_pages) {
        state.total_pages = total_pages
    },
    setCurrentPage(state, current_page) {
        state.current_page = current_page
    },
    setWord(state, word) {
        state.word = word
    },
    setBase(state, base) {
        state.base = base
    },
    setTechnique(state, technique) {
        state.technique = technique
    },
    setTaste(state, taste) {
        state.taste = taste
    },
    setStyle(state, style) {
        state.style = style
    },
    setAlcoholFrom(state, alcohol_from) {
        state.alcohol_from = alcohol_from
    },
    setAlcoholTo(state, alcohol_to) {
        state.alcohol_to = alcohol_to
    },
    setTop(state, top) {
        state.top = top
    },
    setGlass(state, top) {
        state.glass = top
    },
    setPage(state, page) {
        state.page = page
    },
    setLimit(state, limit) {
        state.limit = limit
    },
    setCocktails(state, cocktails) {
        state.cocktails = cocktails
    },
    setCocktailDetail(state, cocktail_detail) {
        state.cocktail_detail = cocktail_detail
    },
    setShowPagination(state, showPagination) {
        state.showPagination = showPagination
    },
    setAveRating(state, ave_rating) {
        state.ave_rating = ave_rating
    },
    setRatingsCount(state, ratings_count) {
        state.ratings_count = ratings_count
    },
}

export const actions = {
    async fetchCocktails({ state, commit }) {
        const params = {
            word: state.word,
            base: state.base,
            technique: state.technique,
            taste: state.taste,
            style: state.style,
            alcohol_from: state.alcohol_from,
            alcohol_to: state.alcohol_to,
            top: state.top,
            glass: state.glass,
            page: state.page,
            limit: state.limit
        }
        const response = await this.$axios.$get('/cocktails', { params })
        commit("setCocktails", response.cocktails)
        commit("setTotalPages", response.total_pages)
        commit("setCurrentPage", response.current_page)
        commit("setShowPagination", true)
    },
    fetchCocktailDetail({ state, commit }, cocktail_id) {
        this.$axios
            .$get('/cocktails/' + cocktail_id)
            .then(response => {
                commit("setCocktailDetail", response)

                let sum_ratings = 0
                for (const rate of state.cocktail_detail.rates) {
                    sum_ratings += rate.rating
                }
                commit("setRatingsCount", (state.cocktail_detail.rates).length)
                if (state.ratings_count === 0) {
                    commit("setAveRating", 0)
                } else {
                    commit("setAveRating", Math.round((sum_ratings / state.ratings_count) * 2)/2)
                }
            })
    },
    createComment({ dispatch }, payload) {
        this.$axios
            .$post('/comment', {
                body: payload.body,
                cocktail_id: payload.cocktail_id,
                user_id: payload.user_id
            })
            .then(response => {
                dispatch("accounts/getMyCommentCocktails", null, { root: true })
                this.$router.push("/accounts/mypage")
            })
    },
    createLike({ dispatch }, payload) {
        this.$axios
            .$post('/createlike', {
                cocktail_id: payload.cocktail_id,
                user_id: payload.user_id
            })
            .then(response => {
                dispatch("accounts/getFavoriteCocktails", null, { root: true })
            })
    },
    deleteLike({ dispatch }, payload) {
        this.$axios
            .$post('/deletelike', {
                cocktail_id: payload.cocktail_id,
                user_id: payload.user_id
            })
            .then(response => {
                dispatch("accounts/getFavoriteCocktails", null, { root: true })
            })
    },
    createRate({ dispatch }, payload) {
        this.$axios
            .$post('/rate', {
                rating: payload.rating,
                cocktail_id: payload.cocktail_id,
                user_id: payload.user_id
            })
            .then(response => {
                dispatch("accounts/getMyRateCocktails", null, { root: true })
                this.$router.push("/accounts/mypage")
            })
    },
}
