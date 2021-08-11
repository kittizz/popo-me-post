<template>
  <div id="app">
    <Navbar />
    <router-view />
  </div>
</template>
<script>
import "./assets/css/timos-icons.css";
import Navbar from "./components/Navbar.vue";
export default {
  name: "App",
  components: { Navbar },
  methods: {},
};
</script>
<style>
#app {
  font-family: Segoe UI, Frutiger, Frutiger Linotype, Dejavu Sans,
    Helvetica Neue, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  /* color: #2c3e50; */
  margin-top: 80px;
}
/* 
 a.router-link-exact-active {
  color: #42b983;
} */
</style>
