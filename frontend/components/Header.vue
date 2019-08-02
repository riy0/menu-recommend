<template>
  <div>
    <b-navbar toggleable="lg" type="dark" class="nav-bar" small>
      <nuxt-link to="/">
        <b-navbar-brand>
          メニュー提案
        </b-navbar-brand>
      </nuxt-link>
    
      <b-navbar-toggle
        target="nav-collapse"
      ></b-navbar-toggle>
      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav class="nav">
          <nuxt-link to="/menus">
            おかず一覧
          </nuxt-link>
        </b-navbar-nav>
        <b-navbar-nav>
          <nuxt-link to="/contact">
            お問い合わせ
          </nuxt-link>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>


<style lang="scss">
.nav-bar{
  border: dashed gray;
  margin-bottom: 20px;
}

.nav{
  margin-right: 20px;
}

.brand{
    color #474949 !important;
}

</style>
