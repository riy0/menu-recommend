<template>
  <section class="container">
    <div>
      <logo />
      <h1 class="title">
        料理たち
      </h1>
      <h2 class="menu-list">
      <div v-for="(item, i) in menu_list" :key="i">
        <a :href="item.url">
          <b-card
            :title="item.image_url"
            :img-src="item.image_url" 
            img-top 
            tag="article"
            style="max-width:30rem;"
            class="mb-2"
          >
            <b-card-text>calorie: {{ item.calorie }}</b-card-text>
          </b-card>
        </a>
      </div> 
    </div>
  </div>
  </section>
</template>

<script>
import menuApi from '@/api/menu'
import {mapState} from 'vuex'

export default{
  computed: mapState(['menu_list']),
  async fetch({ store }) {
    const json = await menuApi.menus()
    store.commit('menu_list_update', json)
  }
}
</script>

<style>
.menu-list{
  text-align: center;
  align-content: center;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items:center;
}
</style>

