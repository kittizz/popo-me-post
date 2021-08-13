<template>
  <div id="app">
    <Navbar />
    <router-view />
  </div>
</template>
<script>
import "./assets/css/timos-icons.css";
import "./assets/scss/app.scss";
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
  /* color: #2c3e50; */
  padding: 2.5vw 5vw;
}
/* 
 a.router-link-exact-active {
  color: #42b983;
} */
</style>
