<template>
  <div class="container">
    <b-card-group deck>
      <b-card :img-src="menu.image_url" img-alt="Card image" img-top>
        <b-card-text>
          {{ menu.title }}
        </b-card-text>
      </b-card>
    </b-card-group>
  </div>
</template>

<script>
import menuApi from '@/api/menu'
import { mapState } from 'vuex'

export default{
  computed: mapState(['menu']),
  async fetch({ store, params }) {
    const json = await menuApi.menu(params.id)
    store.commit('setMenu', json)
  }
}
</script>
