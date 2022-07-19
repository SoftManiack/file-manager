<template>
  <div class="c-navbar">
    <div class="c-navbar_wrapper">
      <div class="c-navbar_nav .c-nav">
        <h5 class="pl-2 mt-3">Мои файлы</h5>
        <div class="c-nav mt-4 px-2">
          <TreeView />
        </div>
      </div>
      <div class="c-navbar_memory">
        <b-progress :max="11" class="с-memory_progress">
          <b-progress-bar class="с-memory_progressbar" :value="1">
          </b-progress-bar>
        </b-progress>
        <p class="с-memory_text"><span>50 МБ </span>из 1 ГБ</p>
      </div>
    </div>
    <div class="c-navbar_border" @mousedown="resizeNav"></div>
  </div>
</template>

<script>
import TreeView from '@/components/TreeView.vue'

export default {
  name: "Navbar",
  data() {
    return {}
  },
  components: {
    TreeView
  },
  methods: {
    resizeNav: () => {
      let downed = true;
      document.onmousemove = () => {
        if(Number(window.getComputedStyle(document.getElementsByClassName('c-navbar')[0]).minWidth.split('p')[0]) > 420){
          downed = false;
          document.getElementsByClassName('c-navbar')[0].style.minWidth = "415px"
        }
        if( downed == true ){
          console.log(window.event.clientX)
          document.getElementsByClassName('c-navbar')[0].style.minWidth = `${window.event.clientX}px`
        }
      }
      document.onmouseup = () => {
        downed=false;
      }
    }
  },
}
</script>

<style>
</style>