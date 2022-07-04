<template>
  <div>
    <div class="l-main-layout" v-if="loadingStatus != 'LOADING'">
      <Navbar/>
      <main>
        <Header/>
        <router-view/>
      </main>
    </div>
    <span class="loader d-flex justify-content-center align-items-center"></span>
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
  .loader {
      width: 48px;
      height: 48px;
      border: 5px solid #FFF;
      border-bottom-color: #FF3D00;
      border-radius: 50%;
      display: block;
      margin: auto;
      margin-top: 40vh;
      box-sizing: border-box;
      animation: rotation 1s linear infinite;
    }

    @keyframes rotation {
      0% {
          transform: rotate(0deg);
      }
      100% {
          transform: rotate(360deg);
      }
    } 
</style>