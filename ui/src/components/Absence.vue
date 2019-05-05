<template>
<v-layout row justify-center v-if="entity">
  <v-dialog v-model="model" persistent max-width="425px">
    <v-card class="absence-form">
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-card-text>
          <v-layout wrap>
            <v-flex xs12>
              <!-- Absence range picker -->
              <vc-date-picker
                v-model="inputs.range" mode="range"
                is-inline locale="en"
                class="absence-picker"
              />
            </v-flex>
            <v-flex xs12>
              <v-layout
                align-center justify-center row fill-height
                :key="inputs.range.start + '-' + inputs.range.end"
                v-if="mode() == 'range'"
              >
                <!-- FROM details (am/pm) -->
                <div class="absence-details-from">
                  <v-layout align-center justify-center column fill-height>
                    <v-flex xs12>
                      <h3>FROM</h3>
                    </v-flex>
                    <v-flex xs12>
                      {{ format(inputs.range.start) }}
                    </v-flex>
                    <v-flex xs12>
                      <v-chip
                        text-color="white" color="accent"
                        :class="{'halfday-off': !inputs.start.am}"
                        @click="inputs.start.am = !inputs.start.am"
                      >AM</v-chip>
                      <v-chip text-color="white" color="accent">PM</v-chip>
                    </v-flex>
                  </v-layout>
                </div>
                <div class="absence-details-join">
                  <v-icon>arrow_forward</v-icon>
                </div>
                <!-- UNTIL details (am/pm) -->
                <div class="absence-details-to">
                  <v-layout align-center justify-center column fill-height>
                    <v-flex xs12>
                      <h3>UNTIL</h3>
                    </v-flex>
                    <v-flex xs12>
                      {{ format(inputs.range.end) }}
                    </v-flex>
                    <v-flex xs12>
                      <v-chip text-color="white" color="accent">AM</v-chip>
                      <v-chip
                        text-color="white" color="accent"
                        :class="{'halfday-off': !inputs.end.pm}"
                        @click="inputs.end.pm = !inputs.end.pm"
                      >PM</v-chip>
                    </v-flex>
                  </v-layout>
                </div>
              </v-layout>
              <v-layout
                align-center justify-center row fill-height
                :key="inputs.range.start + '-' + inputs.range.end"
                v-if="mode() == 'single'"
              >
                <!-- Signle day details (am/pm) -->
                <div class="absence-details-from">
                  <v-layout align-center justify-center column fill-height>
                    <v-flex xs12>
                      <h3>ABSENCE</h3>
                    </v-flex>
                    <v-flex xs12>
                      {{ format(inputs.range.start) }}
                    </v-flex>
                    <v-flex xs12>
                      <v-chip
                        text-color="white" color="accent"
                        :class="{'halfday-off': !inputs.start.am}"
                        @click="inputs.start.am = !inputs.end.pm ? true : !inputs.start.am"
                      >AM</v-chip>
                      <v-chip
                        text-color="white" color="accent"
                        :class="{'halfday-off': !inputs.end.pm}"
                        @click="inputs.end.pm = !inputs.start.am ? true : !inputs.end.pm"
                      >PM</v-chip>
                    </v-flex>
                  </v-layout>
                </div>
              </v-layout>
            </v-flex>
            <v-flex xs12
              :key="inputs.range.start + '-' + inputs.range.end"
              v-show="mode() != null"
            >
              <!-- Reason -->
              <v-select class="absence-details-reason"
                label="Reason"
                v-model="entity.reason"
                required
                :items="reasons"
                :rules="[v => !!v || 'Field required']"
                item-text="label" item-value="id"
              />
            </v-flex>
          </v-layout>
        </v-card-text>
        <v-card-actions>
          <v-btn v-if="!model.new" color="red darken-1" flat @click="remove()">Delete</v-btn>
          <v-spacer></v-spacer>
          <v-btn
            color="accent darken-1" flat
            :disabled="!valid"
            @click="model.new ? create() : update()"
          >{{ model.new ? 'Create' : 'Update'}}</v-btn>
          <v-btn color="accent darken-1" flat @click="cancel()">Cancel</v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </v-dialog>
  <v-snackbar v-model="error" color="error" bottom>
    Invalid form values
    <v-btn flat @click="error = null">ok</v-btn>
  </v-snackbar>
</v-layout>
</template>

<script>
import axios from 'axios';
import moment from 'moment';
import _ from 'lodash';

export default {
  props: ['model', 'reasons'],
  data: () => ({
    entity: null,
    inputs: {},
    valid: false,
    error: null
  }),
  watch: {
    'model.entity': function(e) {
      if (e != null) {
        this.init();
      }
    }
  },

  methods: {
    init: function() {
      this.entity = _.cloneDeep(this.model.entity);
      if (!this.model.new) {
        let from = moment(this.entity.from);
        let to = moment(this.entity.to);
        this.inputs = {
          range: {start: from.toDate(), end: to.toDate()},
          start: {am: from.hours() < 12, pm: from.hours() >= 12},
          end: {am: to.hours() <= 12, pm: to.hours() > 12},
        }
      } else {
        this.inputs = {
          range: {start: null, end: null},
          start: {am: true, pm: true},
          end: {am: true, pm: true},
        }
      }
    },
    format: function(day) {
      return moment(day).format("dddd, MMMM Do");
    },
    validate: function() {
      this.valid = this.$refs.form.validate();
      return this.valid;
    },
    mode: function() {
      if (!this.inputs.range || !this.inputs.range.start || !this.inputs.range.end) {
        return null;
      }
      let from = moment(this.inputs.range.start);
      let to = moment(this.inputs.range.end);
      return (
        from.year() == to.year() &&
        from.month() == to.month() &&
        from.date() == to.date()
      ) ? 'single' : 'range';
    },
    dto: function() {
      let from = moment(this.inputs.range.start);
      from.hours(this.inputs.start.am ? 0 : 12);
      let to = moment(this.inputs.range.end);
      to.hours(this.inputs.end.pm ? 23 : 12);
      return {
        from: from.format(),
        to: to.format(),
        reason: this.entity.reason
      }
    },
    create: async function() {
      try {
        if (!this.validate()) { throw "invalid form"; }
        await axios.post(
          this.$api + '/users/me/absences',
          this.dto(),
          this.$auth.header()
        );
      } catch(e) {
        this.error = e;
        return
      }
      this.$emit('done', true);
      this.reset();
    },
    update: async function() {
      try {
        if (!this.validate()) { throw "invalid form"; }
        await axios.patch(
          this.$api + '/users/me/absences/' + this.entity.id,
          this.dto(),
          this.$auth.header()
        );
      } catch(e) {
        this.error = e;
        return
      }
      this.$emit('done', true);
      this.reset();
    },
    remove: async function() {
      try {
        await axios.delete(
          this.$api + '/users/me/absences/' + this.entity.id,
          this.$auth.header()
        );
      } finally {
        this.$emit('done', true);
        this.reset();
      }
    },
    cancel: function() {
      this.$emit('done', false);
      this.reset();
    },
    reset: function() {
      this.entity = null;
      this.inputs = {};
      this.valid = true;
      this.error = null;
    }
  }
}
</script>

<style>
.absence-picker {
  display: block;
  margin-left: auto;
  margin-right: auto;
}
.absence-details-join {
  margin: 0 20px 10px 20px;
}
.absence-details-from h3,
.absence-details-to h3 {
  margin-bottom: 3px;
}
.absence-details-reason {
  width: 275px;
  margin: 40px auto 0 auto !important;
}
.vc-border {
  border: none;
}
.vc-highlight {
  background-color: var(--v-accent-base);
}
.vc-container, .vc-container * {
  border-color: var(--v-tertiary-base);
}
.vc-text-blue-900 {
  color: white;
}
.absence-form .v-chip .v-chip__content:hover {
  cursor: pointer;
}
.absence-form .v-chip .v-chip__content {
  height: 28px;
}
.absence-form .v-icon.am {
  transform: rotate(180deg) !important;
}
.absence-form .v-chip:focus:not(.v-chip--disabled) {
  -webkit-box-shadow: none;
  box-shadow: none;
}
.absence-form .v-chip:not(.halfday-off) {
  border: solid 1px var(--v-accent-base);
  padding: 0px;
}
.absence-form .v-chip.halfday-off {
  background-color: transparent !important;
  border: solid 1px var(--v-accent-base);
}
.absence-form .v-chip.halfday-off .v-chip__content {
  color: var(--v-primary-base);
}
</style>
