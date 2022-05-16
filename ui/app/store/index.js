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
    comment: [],
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
    comment: (state) => state.commnet,
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
    setComment(state, comment) {
        state.comment = comment
    },
}

export const actions = {
    async fetchCocktails(context) {
        const params = {
            word: context.state.word,
            base: context.state.base,
            technique: context.state.technique,
            taste: context.state.taste,
            style: context.state.style,
            alcohol_from: context.state.alcohol_from,
            alcohol_to: context.state.alcohol_to,
            top: context.state.top,
            glass: context.state.glass,
            page: context.state.page,
            limit: context.state.limit
        }
        const response = await this.$axios.$get('/cocktails', { params }).catch(err => {
            return err.response
        })
        context.commit("setCocktails", response.cocktails)
        context.commit("setTotalPages", response.total_pages)
        context.commit("setCurrentPage", response.current_page)
    },
    async fetchCocktailDetail({ commit }, cocktail_id) {
        const response = await this.$axios.$get('/cocktails/' + cocktail_id).catch(err => {
            return err.response
        })
        commit("setCocktailDetail", response)
        console.log(response)
    },
    createComment({ commit }, payload) {
        this.$axios
            .$post('/comment', {
                body: payload.body,
                cocktail_id: payload.cocktail_id,
                user_id: payload.user_id
            })
            .then(response => {
                console.log(response);
            })
            .catch(error => {
                console.log(error);
            });
    },
    createLike({ commit }, payload) {
        this.$axios
            .$post('/like', {
                cocktail_id: payload.cocktail_id,
                user_id: payload.user_id
            })
            .then(response => {
                console.log(response);
            })
            .catch(error => {
                console.log(error);
            });
    },
    createRate({ commit }, payload) {
        this.$axios
            .$post('/rate', {
                rating: payload.rating,
                cocktail_id: payload.cocktail_id,
                user_id: payload.user_id
            })
            .then(response => {
                console.log(response);
            })
            .catch(error => {
                console.log(error);
            });
    },
}
