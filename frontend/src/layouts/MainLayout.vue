<template>
  <div class="l-main-layout">
    <Navbar/>
    <main>
      <Header/>
      <router-view/>
    </main>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar";
import Header from "@/components/Header";

export default {
  name: "MainLayout",
  components: { Navbar, Header }
}
</script>

<style>
</style>