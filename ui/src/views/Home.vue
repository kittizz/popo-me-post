<template>
  <div class="home">
    <HelloWorld msg="Welcome to Your PoPoMePost" />
    <tc-action name="app">
      <tc-action-item icon="heart" title="Like" />
      <tc-action-item icon="bookmark" title="Bookmark" />
      <tc-action-item icon="share" title="Share" />
    </tc-action>
  </div>
</template>

<script>
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "Home",
  components: {
    HelloWorld,
  },
};
</script>
