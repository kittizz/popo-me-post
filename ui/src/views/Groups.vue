<template>
  <div>
    <h1>Groups</h1>
    <p>prepare the target groups.</p>
    <tl-grid arrangement="auto-fill">
      <tc-card title="All Groups" :hover="true">
        <div class="mt-1">
          <p>choose one account for fetch group list.</p>
          <tc-select placeholder="Choose account">
            <tc-select-item title="Item #1" />
            <tc-select-item title="Item #2" />
            <tc-select-item title="Item #3" />
            <tc-select-item title="Item #4" />
            <tc-select-item title="Item #5" />
          </tc-select>
          <div class="mt-1">
            <tc-button slot="buttons" name="Fetch" icon="lens" />
          </div>
        </div>
      </tc-card>
      <tc-card title="Lists" :hover="true">
        <p slot="media">manage your group list.</p>
        <div class="mt-1">
          <tc-button slot="buttons" name="Manage" icon="tools" />
        </div>
      </tc-card>
    </tl-grid>
    <h1>Lists</h1>

    <tl-grid arrangement="auto-fill">
      <tc-card title="FiveM Set01" subtitle="342 Group" :hover="true">
        <p slot="media">note me.</p>
        <div>
          <tc-button tfbackground="success" name="Use" icon="newspaper" />

          <tc-button slot="buttons" name="Edit" icon="pencil" />
          <tc-button tfbackground="error" name="Delete" icon="trashcan-duo" />
        </div>
      </tc-card>
      <tc-card title="Minecraft PR" subtitle="42 Group" :hover="true">
        <p slot="media">note me.</p>
        <div>
          <tc-button tfbackground="success" name="Use" icon="newspaper" />
          <tc-button slot="buttons" name="Edit" icon="pencil" />
          <tc-button tfbackground="error" name="Delete" icon="trashcan-duo" />
        </div>
      </tc-card>
      <tc-card title="Minecraft TEST" subtitle="1,302 Group" :hover="true">
        <p slot="media">note me.</p>
        <div>
          <tc-button tfbackground="success" name="Use" icon="newspaper" />
          <tc-button slot="buttons" name="Edit" icon="pencil" />
          <tc-button tfbackground="error" name="Delete" icon="trashcan-duo" />
        </div>
      </tc-card>
      <tc-card title="TEST" subtitle="2 Group" :hover="true">
        <p slot="media">note me.</p>
        <div>
          <tc-button tfbackground="success" name="Use" icon="newspaper" />
          <tc-button slot="buttons" name="Edit" icon="pencil" />
          <tc-button tfbackground="error" name="Delete" icon="trashcan-duo" />
        </div>
      </tc-card>
    </tl-grid>
  </div>
</template>

<script>
export default {
  name: "Groups",
  components: {},
};
</script>

<style >
</style>
