<template>
    <div>
        <v-row justify="center" class="mt-1">
            <v-col cols="12" sm="12" md="12">
                <v-card
                    class="pa-4"
                    elevation="0"
                    outlined
                >
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
        </v-row>
    </div>
</template>

<script>
import { mapGetters } from "vuex"

export default {
    computed: {
        ...mapGetters({
            cocktail: "cocktail_detail"
        }),
    }
}
</script>