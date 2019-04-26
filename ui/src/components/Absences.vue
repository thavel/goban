<template>
<div>
  <month :year="2019" :month="3" :header="true"/>
  <month :year="2019" :month="3" :data="absences"/>
</div>
</template>

<script>
import axios from 'axios';
import Month from '@/components/ui/Month.vue';

export default {
  data: () => ({
    absences: []
  }),
  mounted() {
    this.fetch();
  },
  components: {
    Month
  },

  methods: {
    fetch: async function() {
      let res = await axios.get(this.$api + '/absences', this.$auth.header());
      this.absences = res.data;
    }
  }
}
</script>

<style>
</style>
