<template>
     <div id="app" class="app">
        <component :is="layout">
            <router-view />
            <loader object="#de2d34" color1="#ffffff" color2="#17fd3d" size="5" speed="2" bg="#ffffff" objectbg="#999793" opacity="100" disableScrolling="false" name="dots"></loader>
        </component>
        {{loadingStatus}}
    </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "App",
    components: {
        MainLayout: () => import("@/layouts/MainLayout"),
        EmptyLayout: () => import("@/layouts/EmptyLayout"),
        Loader: () => import("@/layouts/Loader"),
    },
    computed: {
        ...mapGetters({
            loadingStatus: 'loadingStatus',
        }),
        layout() {
            if(this.loadingStatus == "LOADING"){
                return "Loader";
            }else{  
                return this.$route.meta.layout || "EmptyLayout";
            }
        },
    },
}
</script>

<style lang="scss">
    @import 'assets/scss/style';
</style>
