<template>
  <div>
    <h1>Accounts</h1>
    <p>manage your Facebook account with tokens.</p>

    <tl-grid>
      <tc-card :hover="true">
        <tc-table
          :selectable="true"
          :multiSelect="true"
          :border="false"
          :striped="true"
          v-model="selection"
          @sort="(s) => (sort = s)"
        >
          <template slot="head">
            <tc-th>#</tc-th>
            <tc-th>Facebook</tc-th>
            <tc-th>ID</tc-th>
            <tc-th>Note</tc-th>
            <tc-th>Status</tc-th>
            <tc-th></tc-th>
          </template>
          <tc-tr
            v-for="p in accounts"
            :key="p.id"
            :data="p.id"
            :tfbackground="p.status === false ? 'error' : ''"
          >
            <tc-td
              ><tc-avatar
                size="tiny"
                src="https://timos.design/assets/me.jpg"
              />
            </tc-td>
            <tc-td>{{ p.facebook }}</tc-td>
            <tc-td>{{ p.id }}</tc-td>
            <tc-td>{{ p.note }}</tc-td>
            <tc-td>{{ p.status }}</tc-td>
            <tc-td>
              <tc-action>
                <tc-action-item icon="pencil" title="Update" />
                <tc-action-item
                  icon="trashcan-duo"
                  title="Delete"
                /> </tc-action
            ></tc-td>
          </tc-tr>
        </tc-table>

        <tl-grid class="mt-1">
          <div class="ta-l">
            <tc-button
              tfbackground="error"
              variant="filled"
              name="Delete selection"
              icon="trashcan-duo"
            />
          </div>
          <div class="ta-r">
            <tc-button
              variant="filled"
              name="Add"
              icon="list-add"
              @click="modalAdd = !modalAdd"
            />
            <tc-button variant="filled" name="Refresh" icon="repeat" />
          </div>
        </tl-grid>
      </tc-card>
    </tl-grid>
    <tc-modal v-model="modalAdd" title="Add Account">
      <form>
        <br />
        <label class="ta-l p-30">
          <tc-link><i class="ti-facebook" /> click this</tc-link>
          for fetch facebook access_token.
        </label>
        <tc-textarea
          required="true"
          rows="6"
          placeholder="Enter your facebook access_token"
          class="p-30"
        />
        <tc-button
          slot="buttons"
          variant="filled"
          name="Submit"
          icon="checkmark"
        />
        <tc-button
          slot="buttons"
          variant="filled"
          tfbackground="error"
          name="Close"
          icon="cross"
          @click="modalAdd = !modalAdd"
        />
      </form>
    </tc-modal>
  </div>
</template>

<script>
export default {
  data() {
    return {
      modalAdd: false,
      query: "",
      selection: [],
      accounts: [
        {
          facebook: "Jonh CD",
          id: "10005123",
          note: "just do it",
          status: true,
        },
        {
          facebook: "Alex Json",
          id: "100021453",
          note: "n do it",
          status: false,
        },
        {
          facebook: "Jonh CD",
          id: "10005122",
          note: "just do it",
          status: true,
        },
        {
          facebook: "Jonh CD",
          id: "10005124",
          note: "just do it",
          status: true,
        },
        {
          facebook: "Jonh CD",
          id: "10005125",
          note: "just do it",
          status: false,
        },
      ],
    };
  },
  name: "Accounts",
  components: {},
};
</script>

<style></style>
