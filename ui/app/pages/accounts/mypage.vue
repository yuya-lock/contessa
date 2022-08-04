<template>
    <div>
        <v-row justify="center" class="mt-1">
            <v-col cols="12" sm="5" md="4" class="mt-3">
                <v-card
                    class="pa-4"
                    elevation="0"
                    outlined
                >
                    <v-card-title>プロフィール</v-card-title>
                    <v-card-title>{{ user.name }}</v-card-title>
                    <v-row justify="start" class="my-1 ml-1">
                        <v-btn @click="showDetail=!showDetail" text small class="px-1">
                            <v-icon>{{ showDetail ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
                            詳細情報
                        </v-btn>
                    </v-row>
                    <v-divider></v-divider>
                    <v-expand-transition>
                        <div v-show="showDetail" class="mt-3">
                            <v-simple-table>
                                <tbody>
                                    <tr>
                                        <td>職業</td>
                                        <td>{{ user.job }}</td>
                                    </tr>
                                    <tr>
                                        <td>年齢</td>
                                        <td>{{ user.age }}</td>
                                    </tr>
                                    <tr>
                                        <td>性別</td>
                                        <td>{{ user.sex }}</td>
                                    </tr>
                                </tbody>
                            </v-simple-table>
                        </div>
                    </v-expand-transition>
                </v-card>
            </v-col>
            <v-col cols="12" sm="7" md="8">
                <v-card
                    class="my-3"
                    elevation="0"
                    outlined
                >
                    
                    <v-tabs>
                        <v-tab
                            @change="changeFavoriteCocktails"
                        >お気に入りカクテル</v-tab>
                        <v-tab
                            @change="changeMyCommentCocktails"
                        >コメント一覧</v-tab>
                        <v-tab
                            @change="changeMyRateCocktails"
                        >評価一覧</v-tab>
                    </v-tabs>
                </v-card>
                <v-expand-transition>
                    <div v-show="showFavoriteCocktails">
                        <v-card
                            class="my-3"
                            elevation="0"
                            outlined
                            v-for="cocktail in favorite_cocktails"
                            :key="cocktail.cocktail_id"
                        >
                            <v-card-subtitle>{{ cocktail.cocktail_digest }}</v-card-subtitle>
                            <v-card-title @click="cocktailDetail(cocktail.cocktail_id)">
                                <nuxt-link
                                    :to="'/' + cocktail.cocktail_id"
                                    class="grey--text text--darken-3"
                                >
                                    {{ cocktail.cocktail_name }}
                                </nuxt-link>
                            </v-card-title>
                            <v-card-subtitle>{{ cocktail.cocktail_name_english }}</v-card-subtitle>
                            <v-card-text>
                                <v-row
                                    align="center"
                                    class="mx-0 pb-3"
                                >
                                    <v-rating
                                        :value="4.5"
                                        color="amber"
                                        dense
                                        half-increments
                                        readonly
                                        size="14"
                                    ></v-rating>
                                    <div class="grey--text ms-4">
                                        4.5 (413)
                                    </div>
                                </v-row>
                                <v-row>
                                    <v-chip
                                        class="ma-2"
                                        color="blue darken-1"
                                        outlined
                                    >
                                        #{{cocktail.base_name}}
                                    </v-chip>
                                    <v-chip
                                        class="ma-2"
                                        color="blue darken-1"
                                        outlined
                                    >
                                        #{{cocktail.alcohol}}%
                                    </v-chip>
                                    <v-chip
                                        class="ma-2"
                                        color="blue-grey lighten-2"
                                        outlined
                                    >
                                        #{{cocktail.glass_name}}
                                    </v-chip>
                                    <v-chip
                                        class="ma-2"
                                        color="blue-grey lighten-2"
                                        outlined
                                    >
                                        #{{cocktail.taste_name}}
                                    </v-chip>
                                    <v-chip
                                        class="ma-2"
                                        color="blue-grey lighten-2"
                                        outlined
                                    >
                                        #{{cocktail.technique_name}}
                                    </v-chip>
                                    <v-chip
                                        class="ma-2"
                                        color="blue-grey lighten-2"
                                        outlined
                                    >
                                        #{{cocktail.style_name}}
                                    </v-chip>
                                    <v-chip
                                        class="ma-2"
                                        color="blue-grey lighten-2"
                                        outlined
                                    >
                                        #{{cocktail.top_name}}
                                    </v-chip>
                                </v-row>
                            </v-card-text>
                        </v-card>
                    </div>
                </v-expand-transition>
                <v-expand-transition>
                    <div v-show="showMyCommentCocktails">
                        <v-card
                            class="my-3"
                            elevation="0"
                            outlined
                            v-for="cocktail in mycomment_cocktails"
                            :key="cocktail.cocktail_id"
                        >
                            <v-card-subtitle>{{ cocktail.cocktail_digest }}</v-card-subtitle>
                            <v-card-title @click="cocktailDetail(cocktail.cocktail_id)">
                                <nuxt-link
                                    :to="'/' + cocktail.cocktail_id"
                                    class="grey--text text--darken-3"
                                >
                                    {{ cocktail.cocktail_name }}
                                </nuxt-link>
                            </v-card-title>
                            <v-card-subtitle>{{ cocktail.cocktail_name_english }}</v-card-subtitle>
                            <v-card-text>
                                <v-row
                                    align="center"
                                    class="mx-0 pb-3"
                                >
                                    
                                </v-row>
                            </v-card-text>
                        </v-card>
                    </div>
                </v-expand-transition>
                <v-expand-transition>
                    <div v-show="showMyRateCocktails">
                        <v-card
                            class="my-3"
                            elevation="0"
                            outlined
                            v-for="cocktail in myrate_cocktails"
                            :key="cocktail.cocktail_id"
                        >
                            <v-card-subtitle>{{ cocktail.cocktail_digest }}</v-card-subtitle>
                            <v-card-title @click="cocktailDetail(cocktail.cocktail_id)">
                                <nuxt-link
                                    :to="'/' + cocktail.cocktail_id"
                                    class="grey--text text--darken-3"
                                >
                                    {{ cocktail.cocktail_name }}
                                </nuxt-link>
                            </v-card-title>
                            <v-card-subtitle>{{ cocktail.cocktail_name_english }}</v-card-subtitle>
                            <v-card-text>
                                <v-row
                                    align="center"
                                    class="mx-0 pb-3"
                                >
                                    <v-rating
                                        :value="4.5"
                                        color="amber"
                                        dense
                                        half-increments
                                        readonly
                                        size="14"
                                    ></v-rating>
                                    <div class="grey--text ms-4">
                                        4.5 (413)
                                    </div>
                                </v-row>
                            </v-card-text>
                        </v-card>
                    </div>
                </v-expand-transition>
            </v-col>
        </v-row>
    </div>
</template>

<script>
import { mapGetters, mapMutations, mapActions } from "vuex"

export default {
    data() {
        return {
            showDetail: false,
            showFavoriteCocktails: true,
            showMyCommentCocktails: false,
            showMyRateCocktails: false,
        }
    },
    computed: {
        ...mapGetters("accounts", [
            "user", "favorite_cocktails", "mycomment_cocktails", "myrate_cocktails"
        ]),
    },
    methods: {
        cocktailDetail(cocktail_id) {
            this.$store.dispatch("fetchCocktailDetail", cocktail_id)
        },
        changeFavoriteCocktails() {
            this.showFavoriteCocktails = true
            this.showMyCommentCocktails = false
            this.showMyRateCocktails = false
        },
        changeMyCommentCocktails() {
            this.showFavoriteCocktails = false
            this.showMyCommentCocktails = true
            this.showMyRateCocktails = false
        },
        changeMyRateCocktails() {
            this.showFavoriteCocktails = false
            this.showMyCommentCocktails = false
            this.showMyRateCocktails = true
        }
    },
    created() {
        this.$store.dispatch("accounts/mypage")
        this.$store.dispatch("accounts/getFavoriteCocktails")
        this.$store.dispatch("accounts/getMyCommentCocktails")
        this.$store.dispatch("accounts/getMyRateCocktails")
    }
}
</script>
