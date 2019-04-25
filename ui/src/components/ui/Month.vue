<template>
<v-layout class="month" align-center justify-center row fill-height>
  <v-flex
    v-for="(day, i) in days"
    :key="i"
    :class="{
      day: true,
      header: header,
      weekend: (!header && day.day() > 4),
      today: (header && isToday(day))
    }"
  >
    {{ header ? day.format('dddd')[0] : day.date() }}
  </v-flex>
</v-layout>
</template>

<script>
import moment from 'moment';

export default {
  props: {
    year: Number,
    month: Number,
    header: Boolean
  },
  data: function() {
    return {
      days: [],
      today: moment()
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

  methods: {
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
  color: var(--v-secondary-base);
}
.month .day.weekend {
  background-color: #cccccc;
}
.month .day.today {
  border: solid 2px var(--v-accent-base);
  border-radius: 50%;
}

</style>
