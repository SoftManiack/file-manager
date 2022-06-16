<template>
    <div id="contextmenu-3" class="l-main_contextmenu-main contextmenu-main" style="width: 225px" @click.exact="$emit('closeMenu')">
        <div class="contextmenu-main_block">
          <div class="contextmenu-main_selection">
             <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#info"></use></svg>
            Информация
          </div>
          <div class="contextmenu-main_selection">
            <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#heart"></use></svg>
            Избранное
          </div>
          <div class="contextmenu-main_selection">
            <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#donwload-3"></use></svg>
            Скачать
          </div>
          <div class="contextmenu-main_selection">
            <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#link"></use></svg>
            Получить ссылку
          </div>
        </div>
        <div class="contextmenu-main_block">
          <div class="contextmenu-main_selection">
            <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#edit"></use></svg>
            Переименовать
          </div>
          <div class="contextmenu-main_selection">
            <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#move"></use></svg>
            Переместить
          </div>
          <div class="contextmenu-main_selection">
            <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#copy"></use></svg>
            Копировать
          </div>
        </div>
        <div class="contextmenu-main_block">
          <div class="contextmenu-main_selection">
              <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#trash-2"></use></svg>
              Удалить
          </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "FileContextmenu",
    emits: ['closeMenu'],
}
</script>

<style>

</style>