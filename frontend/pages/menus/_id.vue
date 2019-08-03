<template>
  <div class="container">
    <b-card-group deck>
      <b-card :img-src="menu.image_url" img-alt="Card image" img-top>
        <b-card-text>
          {{ menu.title }}
        </b-card-text>
          <p>{{ menu.calorie }} kcal </p>
          <p>{{ menu.cooking_time}} min </p>
          <div v-for="(item, i) in ingredient_list" : key="i">
            <p>{{ item.name }} </P>
          </div>
      </b-card>
    </b-card-group>
  </div>
</template>

<script>
import menuApi from '@/api/menu'
import { mapState } from 'vuex'

export default{
  computed: mapState(['menu', 'ingredient_list']),
  async fetch({ store, params }) {
    const menuJson= await menuApi.menu(params.id)
    store.commit('setMenu', menuJson)
    const ingredientJson = await menuApi.ingredients(params.id)
    store.commit('setIngredientList', ingredientJson)
  },
  created(){
      console.log(this.ingredient_list)
  }
}
</script>
