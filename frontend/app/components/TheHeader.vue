<template>
    <v-app-bar
        fixed
        app
        color="dark"
    >
        <img
            src="/logo.png"
            alt="logo"
            :class="{ logo: !xs, 'ml-2': !xs, 'logo-xs': xs, 'ml-1': xs }"
        >
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

<style scoped>
.logo {
    height: 48px;
}
.logo-xs {
    height: 40px;
}
</style>
