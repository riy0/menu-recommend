<template>
  <b-container class="menu-container">
    <b-row>
      <b-col>
        <b-button
          variant="outline-primary"
          @click="onSelected"
          class="menu-btn"
        >
          {{ this.decideBtn }}
        </b-button>
      </b-col>
    </b-row>
    <b-col class="menu-img">
      <nuxt-link :to="{ name: 'menu-id', params: { id: menu[value].id } }" >
          :img-src="menu[value].image_url"
          img-top
          tag="article"
            class="mb-2"
      </nuxt-link>
    </b-col>
    <b-row>
      <b-col>
        <div v-for="(item, i) in ingredients_list" :key="i">
          <p>  {{ item.name }} </p>
        </div>
      </b-col>
    </b-row>
  </b-container>
</template>


<script>
import menuApi from '@/api/menu'
import { mapState } from 'vuex' 

export default {
  data() {
    return {
      value: 0,
      canStart: false 
      decideBtn: 'decide',
      ingredients_list: []
    }
  },
  computed: mapState(['menu_list', 'ingredients']),
  created(){
    if(this.menu_list[0] === undefined){
      location.href = '/'
      onSelected() {
        if (this.canStart === false) {
          clearInterval(this.timer)
          this.canStart = false
          this.canStart = true
          this.decideBtn = 'again'
          this.fetchFuga()
      } else {
        this.start()
        this.canStart = false
        this.decideBtn = 'This ONE!'
      }

      async fetchFuga(){
          this.ingredients_list = await menuApi.ingredients(this.value)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.gacha-container {
  text-align: center;
}
.gacha-img {
  margin-top: 30px;
}
.gacha-btn {
  width: 100%;
}
</style>
