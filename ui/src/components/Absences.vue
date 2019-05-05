<template>
<v-card class="colorful-title absences" :class="{'absences-content': !isEmpty(absences)}">
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
    <v-data-table
      :headers="headers"
      :items="content"
      hide-actions
    >
      <template slot="items" slot-scope="props">
        <tr @click="select(absences[props.item.index])">
          <td class="text-xs-center">{{ props.item.reason }}</td>
          <td class="text-xs-center">{{ props.item.from }}</td>
          <td class="text-xs-center">{{ props.item.to }}</td>
        </tr>
      </template>
    </v-data-table>
  </v-card-text>
  <v-card-text v-if="isEmpty(absences)" class="absence-empty">
    <img src="../assets/mini-logo.png"/><br/>
    No absence has been registered for this month
  </v-card-text>
  <v-fab-transition>
    <v-btn dark color="accent" fab fixed bottom right @click="create()">
      <v-icon>add</v-icon>
    </v-btn>
  </v-fab-transition>
  <absence
    :model="selected"
    :reasons="reasons"
    @done="done($event)"
  />
</v-card>
</template>

<script>
import _ from 'lodash';
import axios from 'axios';
import moment from 'moment';
import auth from '@/common/auth';
import Absence from '@/components/Absence.vue';

export default {
  data: () => ({
    today: moment(),
    current: moment(),
    year: moment().year(),
    month: moment().month(),
    absences: [],
    reasons: [],
    selected: {
      new: false,
      entity: null
    },
    headers: [
      { text: 'Reason', value: 'reason', align: 'center' },
      { text: 'From', value: 'from', align: 'center' },
      { text: 'Until', value: 'to', align: 'center' },
    ]
  }),
  mounted() {
    this.fetchAbsences();
    this.fetchReasons();
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
    },
    content: function() {
      if (!this.reasons || !this.absences) {
        return [];
      }
      let reasons = _.chain(this.reasons).keyBy('id').mapValues('label').value();
      let absences = [];
      for (var i = 0; i < this.absences.length; i++) {
        let abs = this.absences[i];
        absences.push({
          index: i,
          reason: reasons[abs.reason],
          from: moment(abs.from).format("dddd, MMMM Do"),
          to: moment(abs.to).format("dddd, MMMM Do"),
        })
      }
      return absences;
    },
  },
  components: {
    Absence
  },

  methods: {
    uri: function(datetime) {
      return encodeURIComponent(datetime.format());
    },
    fetchAbsences: auth.wrap(async function() {
      var query = '?from=' + this.uri(this.from) + '&to=' + this.uri(this.to);
      let res = await axios.get(this.$api + '/users/me/absences' + query, this.$auth.header());
      this.absences = res.data;
    }),
    fetchReasons: auth.wrap(async function() {
      let res = await axios.get(this.$api + '/reasons', this.$auth.header());
      this.reasons = res.data;
    }),
    create: function() {
      this.selected.new = true;
      this.selected.entity = {};
    },
    select: function(absence) {
      this.selected.new = false;
      this.selected.entity = absence;
    },
    navigate: function(datetime) {
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
    done: function(reload) {
      this.selected.entity = null;
      if (reload) {
        this.fetchAbsences();
      }
    },
    isEmpty: _.isEmpty
  }
}
</script>

<style>
.absence-empty {
  width: 500px;
  min-width: 500px;
  max-width: 500px;
  text-align: center;
  color: #707070;
  font-size: 13px;
}
.absence-empty img {
  -webkit-filter: grayscale(100%);
  filter: grayscale(100%);
  opacity: 0.35;
}
.absences.absences-content {
  min-width: 1100px;
}
.v-btn.v-btn--floating.v-btn--fixed.v-btn--bottom.v-btn--right {
  bottom: 64px;
  right: 32px;
}
</style>
