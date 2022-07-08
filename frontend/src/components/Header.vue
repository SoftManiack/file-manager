<template>
  <div class="c-header">
    <div class="c-header_form">
      <div class="c-header_icon">
        <svg width="16" height="14"><use xlink:href="@/assets/icons/sprite.svg#search"></use></svg>
      </div>
      <input class="c-header_input" placeholder="Поиск">
    </div>
    <button class="c-header_menu " id="popover-1">
      <svg width="20" height="16"><use xlink:href="@/assets/icons/sprite.svg#dots-vertical"></use></svg>
    </button>

    <b-popover
      target="popover-1"
      placement="bottomleft"
      triggers="focus"
    >
      <div class="header-popover">
        <b-row class="header-popover_user px-1 mt-1">
          <b-col cols="4">
            <div class="header-popover_avatar">
                <h5>A</h5>
            </div>
          </b-col>
          <b-col cols="8">
            <h6>Артем Суходоев</h6>
            <span>fetelupercal@gamil.com</span>
          </b-col>
        </b-row>
        <b-row>
          <b-list-group>
            <router-link :to="{ name: 'Trash' }">
              <b-list-group-item button>Настройки</b-list-group-item>
            </router-link>
            <router-link :to="{ name: 'Login' }">
              <b-list-group-item button @click="logout">Выход</b-list-group-item>
            </router-link>
          </b-list-group>
        </b-row>
      </div>
    </b-popover>
    
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "Header",
  methods: {
    ...mapActions({ 
      logoutActions: "logout"
    }),
    logout(){
      this.logoutActions()
    }
  }
}
</script>

<style>
</style>