<script setup lang="ts">
  import '@/app/styles/index.scss';
  import { RouterView } from 'vue-router';
  import { useRoute } from 'vue-router'
  import { computed } from 'vue'
  import { MainLayout, EmptyLayout } from '@/shared/ui/layouts'
  
  const route = useRoute()
  const isEmptyLayout = computed(() => route.meta.layout === 'empty')

</script>

<template>
  <div id="app" >
      <EmptyLayout v-if="isEmptyLayout" />
      <MainLayout v-else />
  </div>
</template>