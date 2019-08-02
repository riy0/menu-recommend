<template>
  <b-container class="menu-container">
    <b-row>
      <b-col class="menu-img">
        <nuxt-link
          :to="{ name: 'menu-id', params: { id: menu[value].id } }"
        >
            :img-src="menu[value].image_url"
            img-top
            tag="article"
            class="mb-2"
          />
        </nuxt-link>
      </b-col>
    </b-row>
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
  </b-container>
</template>

  data() {
    return {
      value: 0,
      canStart: false
      canStart: false,
      decideBtn: '決めた！'
    }
  },
  computed: mapState(['okazu_list']),
@@ -53,11 +61,28 @@ export default {
    onSelected() {
      if (this.canStart === false) {
        clearInterval(this.timer)
        this.canStart = false
        this.canStart = true
        this.decideBtn = 'もう一回...'
      } else {
        this.start()
        this.canStart = false
        this.decideBtn = '決めた！'
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
