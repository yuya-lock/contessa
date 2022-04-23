<template>
    <v-main>
        <v-container>
            <p>{{ message }}</p>
            <a @click="getMessage">webサーバと連携</a>
        </v-container>
    </v-main>
</template>

<script>
export default {
    data: function() {
        return {
            message: "初期値"
        }
    },
    methods: {
        async getMessage() {
            await this.$axios.$get('http://localhost')
                .then(response => {
                    console.log(response)
                })
        }
    }
}
</script>
