<template>
    <div>
        <v-row justify="center" class="mt-1">
            <v-col cols="12" sm="12" md="12">
                <v-card
                    class="pa-4"
                    elevation="0"
                    outlined
                >
                    <v-card-text>
                        ＜採用担当者様へ＞ テストユーザーとしてログインいただけますので、
                        <span @click="logInForm" style="color: blue; text-decoration: underline;">こちら</span>
                        からログインへとお進み下さい。
                    </v-card-text>
                    <v-card-title>新規登録</v-card-title>
                    <v-form
                        ref="form"
                        v-model="valid"
                        @submit.prevent="signup"
                        lazy-validation
                    >
                        <v-row justify="center">
                            <v-col cols="12" sm="5" md="5" class="mx-2">
                                <v-text-field
                                    v-model="name"
                                    :rules="[rules.required, rules.max]"
                                    hint="20文字以内で入力してください。"
                                    :counter="20"
                                    label="ユーザー名"
                                    class="pt-2"
                                ></v-text-field>
                            </v-col>
                            <v-col cols="12" sm="5" md="5" class="mx-2">
                                <v-text-field
                                    v-model="password"
                                    :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                                    :rules="[rules.required, rules.min]"
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
                            <v-col cols="12" sm="3" md="3" class="mx-2">
                                <v-select
                                    v-model="job"
                                    :items="jobItems"
                                    label="職業"
                                ></v-select>
                            </v-col>
                            <v-col cols="12" sm="3" md="3" class="mx-2">
                                <v-select
                                    v-model="age"
                                    :items="ageItems"
                                    label="年齢"
                                ></v-select>
                            </v-col>
                            <v-col cols="12" sm="3" md="3" class="mx-2">
                                <v-select
                                    v-model="sex"
                                    :items="sexItems"
                                    label="性別"
                                ></v-select>
                            </v-col>
                        </v-row>
                        <v-row justify="center">
                            <v-btn
                                class="ma-4 px-4"
                                color="blue lighten-4"
                                type="submit"
                            >SIGN UP</v-btn>
                        </v-row>
                    </v-form>
                </v-card>
            </v-col>
        </v-row>
    </div>
</template>

<script>
import { mapGetters, mapMutations, mapActions } from "vuex"

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
            jobItems: [
                "会社員",
                "大学生",
                "専業主婦・主夫",
                "公務員",
                "自営業/個人事業",
                "会社役員",
                "自由業",
                "アルバイト/パート",
                "無職",
                "その他",
            ],
            ageItems: [
                "20代",
                "30代",
                "40代",
                "50代",
                "60歳以上",
            ],
            sexItems: [
                "男性",
                "女性",
                "その他",
            ],
            name: "",
            password: "",
            job: "",
            age: "",
            sex: "",
        }
    },
    methods: {
        logInForm() {
            this.$router.push("/accounts/login")
        },
        signup() {
            this.$store.dispatch("accounts/signup", {
                name: this.name,
                password: this.password,
                job: this.job,
                age: this.age,
                sex: this.sex
            })
        }
    }
}
</script>
