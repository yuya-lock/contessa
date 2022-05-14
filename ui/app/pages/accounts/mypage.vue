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
        </v-row>
    </div>
</template>

<script>
import { mapGetters, mapMutations, mapActions } from "vuex"

export default {
    data() {
        return {
            showDetail: false,
        }
    },
    computed: {
        ...mapGetters("accounts", [
            "user"
        ])
    },
}
</script>
