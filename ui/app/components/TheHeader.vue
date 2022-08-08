<template>
    <div>
        <v-app-bar
            fixed
            app
        >
            <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>
            <nuxt-link to="/">
                <v-img
                    src="/logo.png"
                    alt="logo"
                    max-height="40"
                    max-width="150"
                ></v-img>
            </nuxt-link>
        </v-app-bar>

        <v-navigation-drawer
            v-model="drawer"
            absolute
            temporary
            style="position:fixed;"
        >
            <v-list
                nav
                dense
            >
                <v-list-item-group
                    v-model="group"
                    active-class="deep-purple--text text--accent-4"
                >
                    <v-list-item nuxt to="/">
                        <v-list-item-icon>
                            <v-icon>mdi-home</v-icon>
                        </v-list-item-icon>
                        <v-list-item-title>Home</v-list-item-title>
                    </v-list-item>

                    <v-list-item nuxt to="/accounts/signup">
                        <v-list-item-icon>
                            <v-icon>mdi-account-plus</v-icon>
                        </v-list-item-icon>
                        <v-list-item-title>Sign Up</v-list-item-title>
                    </v-list-item>

                    <v-list-item nuxt to="/accounts/login">
                        <v-list-item-icon>
                            <v-icon>mdi-account-lock-open</v-icon>
                        </v-list-item-icon>
                        <v-list-item-title>Log In</v-list-item-title>
                    </v-list-item>

                    <v-list-item @click="logout">
                        <v-list-item-icon>
                            <v-icon>mdi-account-lock</v-icon>
                        </v-list-item-icon>
                        <v-list-item-title>Log Out</v-list-item-title>
                    </v-list-item>

                    <v-list-item nuxt to="/accounts/mypage">
                        <v-list-item-icon>
                            <v-icon>mdi-account-cog</v-icon>
                        </v-list-item-icon>
                        <v-list-item-title>My Page</v-list-item-title>
                    </v-list-item>
                </v-list-item-group>
            </v-list>
        </v-navigation-drawer>
    </div>
</template>

<script>
export default {
    data() {
        return {
            drawer: false,
            group: null,
        }
    },
    computed: {
        xs () {
            return this.$vuetify.breakpoint.xs
        }
    },
    methods: {
        logout() {
            if (confirm("ログアウトしますか？")) {
                this.$store.dispatch("accounts/logout")
                this.drawer = false
            } else {
                this.drawer = false
            }
        }
    },
}
</script>
