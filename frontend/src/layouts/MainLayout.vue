<template>
  <div>
    <div class="l-main-layout" v-if="loadingStatus != 'LOADING'">
      <Navbar/>
      <main>
        <Header/>
        <router-view/>
      </main>
    </div>
    1
  </div>
</template>

<script>
import Navbar from "@/components/Navbar";
import Header from "@/components/Header";
import { mapGetters, mapActions } from "vuex";

export default {
  name: "MainLayout",
  components: { Navbar, Header },
  computed: {
    ...mapGetters({ 
      loadingStatus: "getLoadingStatus", 
      uidUser: "uidUser"
    }),
  },
  methods: {
    ...mapActions({ elements: "getElements"})
  },
  created() { 
    this.elements(this.uidUser); 
  },
}

</script>

<style>

</style>