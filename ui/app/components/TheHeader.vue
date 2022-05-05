<template>
    <v-app-bar
        fixed
        app
    >
        <v-app-bar-nav-icon></v-app-bar-nav-icon>
        <nuxt-link to="/">
            <v-img
                src="/logo.png"
                alt="logo"
                max-height="40"
                max-width="150"
            ></v-img>
        </nuxt-link>

        <v-spacer></v-spacer>

        <v-btn
            class="ma-2 px-2"
            outlined
            color="indigo"
            to="/accounts/signup"
            nuxt
        >
            SIGN UP
        </v-btn>
        <v-btn
            class="ma-2 px-2"
            outlined
            color="indigo"
            to="/accounts/login"
            nuxt
        >
            LOG IN
        </v-btn>

        <v-avatar
            color="indigo"
            size="32"
            class="ma-2"
        >
            <nuxt-link to="/accounts/mypage">
                <v-icon dark>mdi-account-circle</v-icon>
            </nuxt-link>
        </v-avatar>
    </v-app-bar>
</template>

<script>
export default {
    computed: {
        xs () {
            return this.$vuetify.breakpoint.xs
        }
    }
}
</script>
