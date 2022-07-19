<template>
     <div id="app" class="app">
        <component :is="layout">
            <router-view />
        </component>
    </div>
</template>

<script>
export default {
  name: "App",
    components: {
        MainLayout: () => import("@/layouts/MainLayout"),
        EmptyLayout: () => import("@/layouts/EmptyLayout"),
    },
    computed: {
        layout() { 
            return this.$route.meta.layout || "EmptyLayout";
        },
    },
}
</script>

<style lang="scss">
    @import 'assets/scss/style';
</style>
