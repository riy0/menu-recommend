<template>
  <div>
    <h1 class="title"> menu </h1> 
    <div class="menu-list">
      <div v-for="(item, i) in menu_list" :key="i">
        <nuxt-link ;to="{ path: 'menus/' + item.id }">
          <b-card
            :title="item.title"
            :img-src="item.imgae_url"
            img-top
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
          >
            <b-card-text>calorie: {{ item.calorie }} </b-card-text>
          </b-card>
        </nuxt-link>
      </div>
    </div>
  </div>
</template>

<script>
import menuApi from '@/api/menu'
import {mapState} from 'vuex'

export default{
  computed: mapState(['menu_list']),
  async fetch({ store }) {
    const json = await menuApi.menus()
    store.commit('setMenuList', json)
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
