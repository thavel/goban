<template>
<v-card class="colorful-title calendar" :class="{'calendar-content': !isEmpty(absences)}">
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
  <v-card-text>
    <v-layout align-center justify-center row fill-height :key="year + '-' + month">
      <v-select class="team-picker"
        label="Team"
        v-model="filter"
        :items="teams"
        item-text="name" item-value="id"
        @change="fetchAbsences()"
      >
        <template v-slot:prepend-item>
          <v-list-tile ripple @click="filter = null; fetchAbsences()">Everybody</v-list-tile>
          <v-divider/>
        </template>
      </v-select>
      <month class="absence-header" :month="month" :year="year" :header="true"/>
    </v-layout>
    <div v-if="!isEmpty(absences)">
      <v-layout
        align-center justify-center row fill-height
        v-for="(abs, usr) in absences"
        :key="year + '-' + month + '-' + usr"
      >
        <inline-user :id="Number(usr)" class="absence-user"/>
        <month :month="month" :year="year" :data="abs"/>
      </v-layout>
    </div>
    <div v-else class="filter-empty">
      <img src="../assets/mini-logo.png"/><br/>
      No absence has been registered for this month
    </div>
  </v-card-text>
</v-card>
</template>

<script>
import _ from 'lodash';
import axios from 'axios';
import moment from 'moment';
import auth from '@/common/auth';
import Month from '@/components/ui/Month.vue';
import InlineUser from '@/components/ui/InlineUser.vue';

export default {
  data: () => ({
    today: moment(),
    current: moment(),
    year: moment().year(),
    month: moment().month(),
    absences: {},
    teams: [],
    filter: null
  }),
  mounted() {
    this.fetchTeams();
    this.fetchAbsences();
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
    fetchTeams: auth.wrap(async function() {
      let res = await axios.get(this.$api + '/teams', this.$auth.header());
      this.teams = res.data;
    }),
    fetchAbsences: auth.wrap(async function() {
      var query = '?from=' + this.uri(this.from) + '&to=' + this.uri(this.to);
      if (this.filter) {
        query += '&team=' + encodeURIComponent(this.filter);
      }
      let res = await axios.get(this.$api + '/absences' + query, this.$auth.header());
      let absences = {};
      for (var ab of res.data) {
        if (!absences[ab.user]) {
          absences[ab.user] = [];
        }
        absences[ab.user].push(ab);
      }
      this.absences = absences;
    }),
    navigate: function(datetime) {
      this.filter = null;
      this.year = datetime.year();
      this.month = datetime.month();
      this.fetchAbsences();
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
.calendar.calendar-content {
  min-width: 1100px;
}

.absence-user,
.team-picker {
  width: 175px;
  min-width: 175px;
  max-width: 175px;
  margin-right: 50px;
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
.filter-empty {
  margin: 10px auto 20px auto;
  text-align: center;
  color: #707070;
  font-size: 13px;
}
.filter-empty img {
  -webkit-filter: grayscale(100%);
  filter: grayscale(100%);
  opacity: 0.35;
}
</style>
