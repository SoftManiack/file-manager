<script setup lang="ts">
  import '@/assets/styles/index.scss';
  import { RouterView } from 'vue-router'
  import { useRoute } from 'vue-router'
  import { computed } from 'vue'
  import MainLayout from '../src/layout/MainLayout.vue'
  import EmptyLayout from '../src/layout/EmptyLayout.vue'
  
  const route = useRoute()
  const isEmptyLayout = computed(() => route.meta.layout === 'empty')

</script>

<template>
  <div id="app" >
    {{isEmptyLayout}}
      <EmptyLayout v-if="isEmptyLayout" />
      <MainLayout v-else />
  </div>
</template>  