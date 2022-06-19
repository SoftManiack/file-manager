<template>
    <div id="contextmenu-2" class="l-main_contextmenu-main contextmenu-main" style="width: 225px" >
        <div  class="contextmenu-main_block" @click.exact="$emit('closeMenu')">
          <div class="contextmenu-main_selection">
             <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#folder-open-2"></use></svg>
            Открыть
          </div>
          <div v-b-modal.modal-infofolder class="contextmenu-main_selection">
            <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#info"></use></svg>
            Информация
          </div>
          <div class="contextmenu-main_selection">
            <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#hearth-2"></use></svg>
            Избранное
          </div>
        </div>
        <div class="contextmenu-main_block" @click.exact="$emit('closeMenu')">
           <div class="contextmenu-main_selection">
              <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#info"></use></svg>
              Удалить ссылку
          </div>
          <div class="contextmenu-main_selection">
              <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#link"></use></svg>
              Получить ссылку
          </div>
          <div class="contextmenu-main_selection">
              <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#donwload-3"></use></svg>
              Скачать
          </div>
        </div>
        <div class="contextmenu-main_block" @click.exact="$emit('closeMenu')">
          <div v-b-modal.modal-4 class="contextmenu-main_selection">
            <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#edit"></use></svg>
            Переиминовать
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
        <div class="contextmenu-main_block" @click.exact="$emit('closeMenu')">
           <div class="contextmenu-main_selection">
              <svg class="mr-2" width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#trash-2"></use></svg>
              Удалить
          </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "FolderContextmenu.vue",
}
</script>

<style>

</style>