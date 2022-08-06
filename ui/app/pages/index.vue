<template>
    <div>
        <v-row justify="center" class="mt-1">
            <v-col cols="12" sm="5" md="4" class="mt-3">
                <v-form
                    ref="form"
                    v-model="valid"
                    lazy-validation
                >
                    <v-card
                        class="pa-4"
                        elevation="0"
                        outlined
                    >
                        <v-text-field
                            v-model="word"
                            :rules="wordRules"
                            :counter="20"
                            label="カクテル名, 材料, 説明文"
                            outlined
                            class="pt-2"
                        ></v-text-field>
                        <v-select
                            v-model="base"
                            :items="baseItems"
                            item-text="name"
                            item-value="number"
                            label="ベース"
                            return-object
                            outlined
                        ></v-select>
                        <v-subheader class="px-0">アルコール度数</v-subheader>
                        <v-row justify="space-between">
                            <v-col cols="6" sm="6">
                                <v-text-field
                                    v-model="alcohol_from"
                                    :rules="alcoholFromRules"
                                    label="MIN"
                                    type="number"
                                    outlined
                                    dense
                                ></v-text-field>
                            </v-col>
                            <v-col cols="6" sm="6">
                                <v-text-field
                                    v-model="alcohol_to"
                                    :rules="alcoholToRules"
                                    label="MAX"
                                    type="number"
                                    outlined
                                    dense
                                ></v-text-field>
                            </v-col>
                        </v-row>
                        <v-row justify="end" class="my-1 mr-1">
                            <v-btn @click="showDetail=!showDetail" text small class="px-1">
                                詳細条件
                                <v-icon>{{ showDetail ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
                            </v-btn>
                        </v-row>
                        <v-divider></v-divider>
                        <v-expand-transition>
                            <div v-show="showDetail" class="mt-3">
                                <v-select
                                    v-model="technique"
                                    :items="techniqueItems"
                                    item-text="name"
                                    item-value="number"
                                    label="技法"
                                    return-object
                                    dense
                                    outlined
                                    class="pt-2"
                                ></v-select>
                                <v-select
                                    v-model="taste"
                                    :items="tasteItems"
                                    item-text="name"
                                    item-value="number"
                                    label="味わい"
                                    return-object
                                    dense
                                    outlined
                                ></v-select>
                                <v-select
                                    v-model="style"
                                    :items="styleItems"
                                    item-text="name"
                                    item-value="number"
                                    label="スタイル"
                                    return-object
                                    dense
                                    outlined
                                ></v-select>
                                <v-select
                                    v-model="top"
                                    :items="topItems"
                                    item-text="name"
                                    item-value="number"
                                    label="時間帯"
                                    return-object
                                    dense
                                    outlined
                                ></v-select>
                                <v-select
                                    v-model="glass"
                                    :items="glassItems"
                                    item-text="name"
                                    item-value="number"
                                    label="グラス"
                                    return-object
                                    dense
                                    outlined
                                ></v-select>
                            </div>
                        </v-expand-transition>
                        <v-row justify="center" class="mb-1 mt-4">
                            <v-btn
                                width="88px"
                                @click="search"
                                elevation="1"
                                color="error"
                                class="mr-4"
                            >検索</v-btn>
                            <v-btn
                                width="88px"
                                @click="reset"
                                outlined
                                color="primary"
                                class="ml-4"
                                disabled
                            >クリア</v-btn>
                        </v-row>
                    </v-card>
                </v-form>
            </v-col>
            <v-col cols="12" sm="7" md="8">
                <v-card
                    v-show="!showPagination"
                    class="my-3"
                    elevation="0"
                    outlined
                >
                    <v-card-title style="color: #8E24AA;">LET'S ENJOY COCKTAIL LIFE</v-card-title>
                    <v-card-subtitle style="color: #8E24AA;">ー Contessa makes your life colorful ー</v-card-subtitle>
                    <v-card-text>
                        Contessaでは、現在の気分によって様々なカクテルの検索ができます。
                    </v-card-text>
                </v-card>
                <v-card
                    class="my-3"
                    elevation="0"
                    outlined
                    v-for="cocktail in cocktails"
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
                <div class="text-center">
                    <v-pagination
                        v-model="page"
                        class="my-4"
                        :length="total_pages"
                        @input="changePage"
                        :value="current_page"
                        total-visible="7"
                        circle
                        v-show="showPagination"
                    ></v-pagination>
                </div>
            </v-col>
        </v-row>
    </div>
</template>

<script>
import { mapGetters, mapMutations, mapActions } from "vuex"

export default {
    data() {
        return {
            inputs: {},
            temporaryInputs: {},
            valid: true,
            showDetail: false,
            wordRules: [
                v => v.length <= 20 || "20文字以内で入力してください。",
            ],
            alcoholFromRules: [
                v => v >= 0 || "0以上の値を入力してください。",
            ],
            alcoholToRules: [
                v => v <= 50 || "50以下の値を入力してください。",
            ],
            baseItems: [
                { name: "ジン", number: 1},
                { name: "ウォッカ", number: 2},
                { name: "テキーラ", number: 3},
                { name: "ラム", number: 4},
                { name: "ウイスキー", number: 5},
                { name: "ブランデー", number: 6},
                { name: "リキュール", number: 7},
                { name: "ワイン", number: 8},
                { name: "ビール", number: 9},
                { name: "日本酒", number: 10},
                { name: "ノンアルコール", number: 0},
            ],
            techniqueItems: [
                { name: "ビルド", number: 1},
                { name: "ステア", number: 2},
                { name: "シェイク", number: 3},
            ],
            tasteItems: [
                { name: "甘口", number: 1},
                { name: "中甘口", number: 2},
                { name: "中口", number: 3},
                { name: "中辛口", number: 4},
                { name: "辛口", number: 5},
            ],
            styleItems: [
                { name: "ショート", number: 1},
                { name: "ロング", number: 2},
            ],
            topItems: [
                { name: "食前", number: 1},
                { name: "食後", number: 2},
                { name: "オール", number: 3},
            ],
            glassItems: [
                { name: "カクテルグラス", number: 1},
                { name: "オールドファッションドグラス", number: 2},
                { name: "コリンズグラス", number: 3},
                { name: "タンブラー", number: 4},
                { name: "ワイングラス", number: 5},
                { name: "シャンパングラス", number: 6},
                { name: "ホットグラス", number: 7},
                { name: "ゴブレット", number: 8},
                { name: "リキュールグラス", number: 9},
                { name: "サワーグラス", number: 10},
                { name: "ピルスナーグラス", number: 11},
            ],
        }
    },
    computed: {
        ...mapGetters(["cocktails", "total_pages", "current_page", "showPagination"]),
        word: {
            get() {
                return this.$store.getters.word
            },
            set(value) {
                this.$store.commit("setWord", value);
            }
        },
        base: {
            get() {
                return this.$store.getters.base
            },
            set(value) {
                this.$store.commit("setBase", value.number)
            }
        },
        technique: {
            get() {
                return this.$store.getters.technique
            },
            set(value) {
                this.$store.commit("setTechnique", value.number)
            }
        },
        taste: {
            get() {
                return this.$store.getters.taste
            },
            set(value) {
                this.$store.commit("setTaste", value.number)
            }
        },
        style: {
            get() {
                return this.$store.getters.style
            },
            set(value) {
                this.$store.commit("setStyle", value.number)
            }
        },
        top: {
            get() {
                return this.$store.getters.top
            },
            set(value) {
                this.$store.commit("setTop", value.number)
            }
        },
        glass: {
            get() {
                return this.$store.getters.glass
            },
            set(value) {
                this.$store.commit("setGlass", value.number)
            }
        },
        alcohol_from: {
            get() {
                return this.$store.getters.alcohol_from
            },
            set(value) {
                this.$store.commit("setAlcoholFrom", value);
            }
        },
        alcohol_to: {
            get() {
                return this.$store.getters.alcohol_to
            },
            set(value) {
                this.$store.commit("setAlcoholTo", value);
            }
        },
        page: {
            get() {
                return this.$store.getters.page
            },
            set(value) {
                this.$store.commit("setPage", value)
            }
        },
    },
    methods: {
        initializeInputs() {
            let defaultInputs = {
                word: "",
                base: "",
                technique: "",
                taste: "",
                style: "",
                alcohol_from: 0,
                alcohol_to: 50,
                top: "",
                glass: "",
                page: "",
                limit: ""
            }
            this.inputs = JSON.parse(JSON.stringify(defaultInputs))
            this.temporaryInputs = JSON.parse(JSON.stringify(defaultInputs))
        },
        search() {
            this.$refs.form.validate()
            if (!this.valid) {
                return
            }
            this.inputs = JSON.parse(JSON.stringify(this.temporaryInputs))
            this.$store.commit("setPage")
            this.$store.dispatch("fetchCocktails")
        },
        cocktailDetail(cocktail_id) {
            this.$store.dispatch("fetchCocktailDetail", cocktail_id)
        },
        reset() {
            this.initializeInputs()
        },
        changePage(pageNumber) {
            if (pageNumber === this.current_page) return
            this.$store.dispatch("fetchCocktails")
        },
    }
}
</script>
