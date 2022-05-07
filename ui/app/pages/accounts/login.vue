<template>
    <div>
        <v-row justify="center" class="mt-1">
            <v-col cols="12" sm="12" md="12">
                <v-card
                    class="pa-4"
                    elevation="0"
                    outlined
                >
                    <v-card-subtitle>
                        まだ登録がお済みでない方は、
                        <span @click="signUpForm" style="color: blue; text-decoration: underline;">こちら</span>
                        から新規登録を行なってください。
                    </v-card-subtitle>
                    <v-card-title>ログイン</v-card-title>
                    <v-form
                        ref="form"
                        v-model="valid"
                        @submit.prevent="login"
                        lazy-validation
                    >
                        <v-row justify="center">
                            <v-col cols="12" sm="5" md="5" class="mx-2">
                                <v-text-field
                                    v-model="name"
                                    :counter="20"
                                    label="ユーザー名"
                                    class="pt-2"
                                ></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="5" md="5" class="mx-2">
                                <v-text-field
                                    v-model="password"
                                    :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                                    :type="show ? 'text' : 'password'"
                                    hint="8文字以上で入力してください。"
                                    counter
                                    label="パスワード"
                                    @click:append="show = !show"
                                    class="pt-2"
                                ></v-text-field>
                            </v-col>
                        </v-row>
                        <v-row justify="center">
                            <v-btn
                                class="ma-4 px-4"
                                type="submit"
                                color="blue lighten-4"
                            >LOG IN</v-btn>
                        </v-row>
                    </v-form>
                </v-card>
            </v-col>
        </v-row>
    </div>
</template>

<script>
import { mapGetters, mapMutations, mapActions } from 'vuex'

export default {
    data() {
        return {
            valid: true,
            show: false,
            rules: {
                required: value => !!value || "この項目は必須です。",
                max: v => v.length <= 20 || "20文字以内で入力してください。",
                min: v => v.length >= 8 || "8文字以上で入力してください。",
            },
            name: "",
            password: "",
        }
    },
    methods: {
        signUpForm() {
            this.$router.push("/accounts/signup")
        },
        login() {
            let user = {
                name: this.name,
                password: this.password
            }
            this.$store.dispatch("accounts/login", { user })
        },
    }
}
</script>