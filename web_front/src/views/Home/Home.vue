<template>
  <div class="home">
    <Navbar title="XXX文化馆"></Navbar>
    <Sidebar></Sidebar>
    <h1>HomePage</h1>
    <router-view></router-view>
  </div>
</template>

<script>
import { Navbar, Sidebar } from "@/components";

export default {
  name: "HomeView",
  components: { Navbar, Sidebar },
  data() {
    return {};
  },
};
</script>
