<template>
<v-card class="colorful-title">
  <v-card-title>
    <v-layout align-center justify-center row fill-height>
      <v-btn class="month-today" flat icon @click="reset()" :disabled="today.isSameOrAfter(current)">
          <v-icon color="white">today</v-icon>
      </v-btn>
      <v-btn flat icon @click="previous()"><v-icon color="white">arrow_back</v-icon></v-btn>
      <div class="month-picker">{{ current.format('MMMM YYYY') }}</div>
      <v-btn flat icon @click="next()"><v-icon color="white">arrow_forward</v-icon></v-btn>
      <v-btn class="month-today" flat icon @click="reset()" :disabled="today.isSameOrBefore(current)">
        <v-icon color="white">today</v-icon>
      </v-btn>
    </v-layout>
  </v-card-title>
  <v-card-text v-if="!isEmpty(absences)">
    <v-layout align-center justify-center row fill-height :key="year + '-' + month">
      <div class="team-picker"/>
      <month class="absence-header" :month="month" :year="year" :header="true"/>
    </v-layout>
    <v-layout
      align-center justify-center row fill-height
      v-for="(abs, usr) in absences"
      :key="year + '-' + month + '-' + usr"
    >
      <inline-user :id="Number(usr)"/>
      <month :month="month" :year="year" :data="abs"/>
    </v-layout>
  </v-card-text>
  <v-card-text v-if="isEmpty(absences)" class="absence-empty">
    No absence has been registered for this month
  </v-card-text>
</v-card>
</template>

<script>
import _ from 'lodash';
import axios from 'axios';
import moment from 'moment';
import Month from '@/components/ui/Month.vue';
import InlineUser from '@/components/ui/InlineUser.vue';

export default {
  data: () => ({
    today: moment(),
    current: moment(),
    year: moment().year(),
    month: moment().month(),
    absences: {}
  }),
  mounted() {
    this.fetch();
  },
  computed: {
    from() {
      return moment({year: this.year, month: this.month, day: 1});
    },
    to() {
      let to = this.from.clone();
      to.add(1, 'months');
      to.subtract(1, 'hours');
      return to;
    }
  },
  components: {
    Month,
    InlineUser
  },

  methods: {
    uri: function(datetime) {
      return encodeURIComponent(datetime.format());
    },
    fetch: async function() {
      this.absences = {};
      let query = '?from=' + this.uri(this.from) + '&to=' + this.uri(this.to);
      let res = await axios.get(this.$api + '/absences' + query, this.$auth.header());
      let absences = {};
      for (var ab of res.data) {
        if (!absences[ab.user]) {
          absences[ab.user] = [];
        }
        absences[ab.user].push(ab);
      }
      this.absences = absences;
    },
    navigate: function(datetime) {
      this.year = datetime.year();
      this.month = datetime.month();
      this.fetch();
    },
    reset: function() {
      this.current = this.today.clone();
      this.navigate(this.today);
    },
    previous: function() {
      this.current.subtract(1, 'month');
      this.navigate(this.current);
    },
    next: function() {
      this.current.add(1, 'month');
      this.navigate(this.current);
    },
    isEmpty: _.isEmpty,
  }
}
</script>

<style>
.team-picker {
  width: 175px;
  min-width: 175px;
  max-width: 175px;
}
.month-picker {
  color: white;
  font-weight: bold;
  width: 115px;
  min-width: 115px;
  max-width: 115px;
  text-align: center;
}
.month-today.v-btn.v-btn--disabled i.v-icon {
  color: transparent !important;
}
.absence-header {
  margin-bottom: 8px;
}
.absence-empty {
  width: 500px;
  min-width: 500px;
  max-width: 500px;
  text-align: center;
  color: #707070;
  font-size: 12px;
}
</style>
