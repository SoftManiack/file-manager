<template>
    <div id="contextmenu-1" class="l-main_contextmenu-main contextmenu-main">
        <div v-b-modal.modal-1  class="contextmenu-main_block" @click.exact="$emit('closeMenu')">
          <div class="contextmenu-main_selection">
            <svg class="c-nav_icon mr-2" width="19" height="15" fill ="#707070"><use xlink:href="@/assets/icons/sprite.svg#folder-add"></use></svg>
            Новоя папка
          </div>
        </div>
        <div v-b-modal.modal-2 class="contextmenu-main_block" >
          <div class="contextmenu-main_selection">
            <svg class="c-nav_icon mr-2" width="19" height="17" fill ="#707070"><use xlink:href="@/assets/icons/sprite.svg#file-add"></use></svg>
            Новый текстовый файл
          </div>
        </div>
        <div class="contextmenu-main_block">
          <div class="contextmenu-main_selection">
            <svg class="c-nav_icon mr-2" width="19" height="17" fill ="#707070"><use xlink:href="@/assets/icons/sprite.svg#file-donwload"></use></svg>
            Загрузить файл
          </div>
          <div class="contextmenu-main_selection mt-1">
            <svg class="c-nav_icon mr-2" width="19" height="16" fill ="#707070"><use xlink:href="@/assets/icons/sprite.svg#folder-donwload"></use></svg>
            Загрузить папку
          </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "MainContextmenu.vue",
}
</script>

<style>

</style>