<template>
  <section class="container">
    <div>
      <h1 class="title"> MENU </h1>
      <div class="links">
        <nuxt-link to="/menus/gacha" class="btn btn-outline-success">
          TELL ME! 
        </nuxt-link>
        <nuxt-link to="/menus" class="btn btn-outline-primary">
          MENU LIST
        </nuxt-link>
      </div>
    </div>
  </section>
</template>

<script>
import menuApifrom '@/api/menu'
import { mapState } from 'vuex'
export default {
  computed: mapState(['menu_list']),
  async fetch({ store }) {
    const json = await menuApi.menus()
    store.commit('setMenuList', json)
  }
}
</script>

<style scoped>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  text-align: center;
}
.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 60px;
  color: #35495e;
  letter-spacing: 1px;
}
.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}
.links {
  padding-top: 15px;
}
</style>
