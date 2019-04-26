<template>
<v-layout class="month" align-center justify-center row fill-height>
  <v-flex
    v-for="(day, i) in days"
    :key="i"
    :class="{
      day: true,
      header: header,
      weekend: (!header && day.day() > 4),
      today: (header && isToday(day)),
      ...absence(day)
    }"
  >
    {{ header ? day.format('dddd')[0] : day.date() }}
  </v-flex>
</v-layout>
</template>

<script>
import Moment from 'moment';
import { extendMoment } from 'moment-range';
const moment = extendMoment(Moment);

export default {
  props: {
    year: Number,
    month: Number,
    header: Boolean,
    data: Array
  },
  data: function() {
    return {
      days: [],
      today: moment(),
      absences: []
    }
  },
  mounted() {
    let ref = moment({year: this.year, month: this.month, day: 1});
    let size = ref.daysInMonth();
    for(var i = 1; i <= size; i++) {
      this.days.push(ref.clone());
      ref.add(1, 'days');
    }
  },
  watch: {
    'data': function(d) {
      this.parse(d);
    }
  },

  methods: {
    parse: function(data) {
      if (!data || data.length < 1) {
        return;
      }
      for (var i = 0; i < data.length; i++) {
        let from = moment(data[i].from);
        let to = moment(data[i].to);
        this.absences.push(moment.range(from, to));
      }
    },
    absence: function(morning) {
      let afternoon = morning.clone();
      afternoon.add(14, 'hours');

      for (var i = 0; i < this.absences.length; i++) {
        let absence = this.absences[i];
        let am = absence.contains(morning);
        let pm = absence.contains(afternoon);
        if (am || pm) {
          return {
            'absence': am && pm,
            'absence-am': am && !pm,
            'absence-pm': !am && pm
          };
        }
      }
      return {}
    },
    isToday: function(day) {
      return (
        day.year() == this.today.year() &&
        day.month() == this.today.month() &&
        day.date() == this.today.date()
      )
    }
  }
}
</script>

<style>
.month .day {
  margin: 2px 0;
  width: 35px;
  min-width: 30px;
  height: 35px;
  text-align: center;
  vertical-align: middle;
  line-height: 35px;
}
.month .day.header {
  color: var(--v-primary-base);
  font-weight: bold;
}
.month .day.today {
  border: solid 1px var(--v-primary-base);
  border-radius: 50%;
}
.month .day.absence {
  background-color: var(--v-accent-base);
  color: var(--v-secondary-base);
}
.month .day.absence-am {
  /* background: linear-gradient(to top left, transparent 0%, transparent 50%, var(--v-accent-base) 50%, var(--v-accent-base) 100%); */
  background: linear-gradient(90deg, var(--v-accent-base) 50%, transparent 50%);
  /* border: solid 1px var(--v-accent-base); */
  outline: 2px dashed var(--v-accent-base);
  outline-offset: -2px;
  color: var(--v-secondary-base);
}
.month .day.absence-pm {
  /* background: linear-gradient(to bottom right, transparent 0%, transparent 50%, var(--v-accent-base) 50%, var(--v-accent-base) 100%); */
  background: linear-gradient(90deg, transparent 50%, var(--v-accent-base) 50%);
  outline: 2px dashed var(--v-accent-base);
  outline-offset: -2px;
  color: var(--v-secondary-base);
}
.month .day.weekend {
  background-color: #cccccc;
  color: #707070;
}
</style>
