<template>
  <div 
    class="c-block"
    @click.ctrl.exact="$emit('selectManyBlock')"
    @click.exact="$emit('selectBlock')" 
    @contextmenu="$emit('switchContextmenu', $event), $emit('selectBlock')"
    @dblclick="$emit('openDirectory')"
    :class="{'c-block--select': selected}">
    <div class="d-flex justify-content-end c-block_favorites px-2 mt-2">
      <div v-if="isFavorites == true" >
        <svg width="15" height="15"><use xlink:href="@/assets/icons/sprite.svg#heart"></use></svg>
      </div>
    </div>
    <div class="d-flex justify-content-center mb-4">
      <div v-if="type == 'Directory'">
        <svg width="100" height="100"><use xlink:href="@/assets/icons/sprite.svg#folder-large"></use></svg>
      </div>
      <div class="mt-3 mb-2" v-if="type == 'File'">
        <svg width="75" height="75"><use xlink:href="@/assets/icons/sprite.svg#empry-file"></use></svg>
      </div>
    </div>
    <div class="d-flex">
      <div class="align-self-center c-block_name mt-1">{{name}}</div>
      <div class="c-block_menu mt-1" id="popover-block" :class="{'c-block_menu--select': selected}">
        <svg width="16" height="25"><use xlink:href="@/assets/icons/sprite.svg#horizontal-menu"></use></svg>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
      name: 'Block',
      props: {
        uid: { type: String, default: "" },
        name: { type: String, default: "" },
        type: { type: String, default: "" },
        isFavorites: { type: Boolean, default: false },
        selection: { type: Array, default: function () {} },
      },
      emits: ['selectBlock','selectManyBlock','switchContextmenu'],
      computed: {
        selected(){
          return  this.selection.find(item => item.uid == this.uid)
        }
      }
  }
</script>

<style>
</style>