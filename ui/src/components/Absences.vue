<template>
<v-card class="colorful-title">
  <v-card-text>
    <v-layout align-center justify-center row fill-height>
      <flex>
        <v-btn flat icon @click="previous()"><v-icon>arrow_back</v-icon></v-btn>
      </flex>
      <flex>{{ page.format('MMMM YYYY') }}</flex>
      <flex>
        <v-btn flat icon @click="next()"><v-icon>arrow_forward</v-icon></v-btn>
      </flex>
    </v-layout>
    <month :month="month" :year="year" :header="true"/>
    <month :month="month" :year="year" :data="absences"/>
  </v-card-text>
</v-card>
</template>

<script>
import axios from 'axios';
import moment from 'moment';
import Month from '@/components/ui/Month.vue';

export default {
  data: () => ({
    page: moment(),
    year: moment().year(),
    month: moment().month(),
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
    },
    previous: function() {
      this.page.subtract(1, 'month');
      this.year = this.page.year();
      this.month = this.page.month();
    },
    next: function() {
      this.page.add(1, 'month');
      this.year = this.page.year();
      this.month = this.page.month();
    }
  }
}
</script>

<style>
.absences-page-selector {
  display: inline-block;
  vertical-align: middle;
  text-align: justify;
  vertical-align: middle;
  line-height: 35px;
}
</style>
