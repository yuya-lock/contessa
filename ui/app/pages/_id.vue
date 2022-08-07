<template>
    <div>
        <v-row justify="center" class="mt-1">
            <v-col cols="12" sm="12" md="12">
                <v-card
                    class="pa-4"
                    elevation="0"
                    outlined
                >
                    <v-row class="mb-1" justify="end">
                        <div v-if="showLike">
                            <v-btn
                                elevation="0"
                                outlined
                                color="blue"
                                @click="createLike"
                            >
                                <v-icon left>mdi-thumb-up</v-icon>
                                MY COCKTAILに追加済み
                            </v-btn>
                        </div>
                        <div v-else>
                            <v-btn
                                elevation="0"
                                outlined
                                color="blue-grey lighten-2"
                                @click="deleteLike"
                            >
                                <v-icon left>mdi-thumb-up</v-icon>
                                MY COCKTAILに追加
                            </v-btn>
                        </div>
                    </v-row>
                    <v-card-subtitle class="py-0">{{ cocktail.cocktail_digest }}</v-card-subtitle>
                    <v-card-title class="pt-0">{{ cocktail.cocktail_name }}</v-card-title>
                    <v-card-subtitle>{{ cocktail.cocktail_name_english }}</v-card-subtitle>
                    <v-card-text>
                        <div class="pb-2">{{ cocktail.cocktail_desc }}</div>
                        <v-row
                            align="center"
                            class="mx-0 pt-3"
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
                        <v-row
                            align="center"
                            class="pb-3"
                        >
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
                        <v-row justify="center" class="mt-1">
                            <v-col cols="12" sm="5" md="4">
                                <v-img
                                    src="https://cdn.pixabay.com/photo/2017/08/03/21/48/drinks-2578446_1280.jpg"
                                ></v-img>
                            </v-col>
                            <v-col cols="12" sm="7" md="8">
                                <v-simple-table>
                                    <tbody>
                                        <tr
                                            v-for="recipe in cocktail.recipes"
                                            :key="recipe.ingredient_id"
                                        >
                                            <td>{{ recipe.ingredient_name }}</td>
                                            <td>{{ recipe.amount }}{{ recipe.unit }}</td>
                                        </tr>
                                    </tbody>
                                </v-simple-table>
                                {{ cocktail.recipe_desc }}
                            </v-col>
                        </v-row>
                    </v-card-text>
                </v-card>
            </v-col>
            <v-col cols="12" sm="5" md="4" class="mt-3">
                <v-card
                    class="pa-4"
                    elevation="0"
                    outlined
                >
                    <v-card-title>Review This Cocktail</v-card-title>
                    <v-textarea
                        v-model="comment"
                        label="このカクテルのレビューを書く"
                        outlined
                        auto-grow
                        class="pt-2"
                    ></v-textarea>
                    <v-row justify="center" class="mb-1">
                        <v-btn
                            width="88px"
                            elevation="1"
                            color="indigo lighten-3"
                            @click="createComment"
                        >SUBMIT</v-btn>
                    </v-row>
                </v-card>
                <v-card
                    class="pa-4 mt-5"
                    elevation="0"
                    outlined
                >
                    <v-card-title>Rate This Cocktail</v-card-title>
                    <v-card-text>
                        If you enjoy drinking this cocktail, please take a few seconds to rate your experience. It really helps!
                        <div class="text-center mt-6">
                            <v-rating
                                v-model="rating"
                                color="yellow darken-3"
                                background-color="grey darken-1"
                                empty-icon="$ratingFull"
                                hover
                            ></v-rating>
                        </div>
                    </v-card-text>
                    <v-divider></v-divider>
                    <v-card-actions class="justify-center">
                        <v-btn
                            color="primary"
                            @click="createRate"
                            text
                        >Rate Now</v-btn>
                    </v-card-actions>
                </v-card>
            </v-col>
            <v-col cols="12" sm="7" md="8">
                <v-card
                    class="my-3"
                    elevation="0"
                    outlined
                    v-for="comment in cocktail.comments"
                    :key="comment.ID"
                >
                    <v-card-text>{{ comment.body }}</v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </div>
</template>

<script>
import { mapGetters } from "vuex"

export default {
    data() {
        return {
            showLike: false,
            comment: "",
            rating: 0,
        }
    },
    computed: {
        ...mapGetters({
            cocktail: "cocktail_detail"
        }),
        ...mapGetters("accounts", [
            "favorite_cocktails", "token",
        ])
    },
    methods: {
        createComment() {
            if (!this.token) {
                alert('カクテルのレビューを書く場合は、先にログインをしてください。')
                return
            }
            this.$store.dispatch("createComment", {
                body: this.comment,
                cocktail_id: this.cocktail.cocktail_id,
                user_id: this.$store.getters["accounts/user"].ID
            })
            this.comment = "";
        },
        createLike() {
            if (!this.token) {
                alert('MY COCKTAILの追加をする場合は、先にログインをしてください。')
                return
            }
            this.$store.dispatch("createLike", {
                cocktail_id: this.cocktail.cocktail_id,
                user_id: this.$store.getters["accounts/user"].ID
            })
            this.showLike = true
        },
        deleteLike() {
            if (!this.token) {
                alert('MY COCKTAILの解除をする場合は、先にログインをしてください。')
                return
            }
            this.$store.dispatch("deleteLike", {
                cocktail_id: this.cocktail.cocktail_id,
                user_id: this.$store.getters["accounts/user"].ID
            })
            this.showLike = false
        },
        createRate() {
            if (!this.token) {
                alert('カクテルを評価する場合は、先にログインをしてください。')
                return
            }
            this.$store.dispatch("createRate", {
                rating: this.rating,
                cocktail_id: this.cocktail.cocktail_id,
                user_id: this.$store.getters["accounts/user"].ID
            })
            this.rating = 0
        },
    },
    created() {
        for (const favorite_cocktail of this.favorite_cocktails) {
            if (favorite_cocktail.cocktail_id == this.$route.params.id) {
                this.showLike = true
            }
        }
    },
}
</script>