<template>
  <div>
    <Header />
    <nuxt />
  </div>
</template>

<script>

import Header from '@/components/Header.vue'

export default {
    components: {
        Header
    }
}

<style lang="scss">
body {
  font-family: 'Sawarabi Gothic';
}

a {
  color: gray;
    :hover {
      color: black;
    }
}

</style>
