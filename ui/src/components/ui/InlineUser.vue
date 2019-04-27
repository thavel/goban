<template>
<div class="inline-user">
  <span>{{ displayName }}</span>
</div>
</template>

<script>
import axios from 'axios';

export default {
  props: {
    id: Number,
  },
  data: function() {
    return {
      displayName: null
    }
  },
  mounted() {
    this.fetch();
  },

  methods: {
    fetch: async function() {
      let res = await axios.get(this.$api + '/users/' + this.id, this.$auth.header());
      if (res.data.firstname && res.data.lastname) {
        this.displayName = res.data.firstname + ' ' + res.data.lastname;
      } else if (res.data.firstname || res.data.lastname) {
        this.displayName = res.data.firstname || res.data.lastname;
      } else {
        this.displayName = res.data.email;
      }
    },
  }
}
</script>

<style>
.inline-user {
  width: 175px;
  max-width: 175px;
  min-width: 175px;
  text-overflow: ellipsis;
}
</style>
