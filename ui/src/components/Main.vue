<template>
<div id="main">
  <v-navigation-drawer v-model="drawer" fixed app dark>
    <v-list>
      <v-list-tile avatar id="main-logo">
        <v-list-tile-avatar>
          <img src="../assets/logo.png" height="34"/>
        </v-list-tile-avatar>
        <v-list-tile-title class="title">
          Goban
        </v-list-tile-title>
      </v-list-tile>
      <v-divider id="main-title-divider"/>

      <v-list-tile
        v-for="(item, i) in menu"
        :key="i"
        :to="item.path"
        @click="setTitle"
      >
        <v-list-tile-action>
          <v-icon>{{ item.icon }}</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>{{ item.name }}</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </v-navigation-drawer>
  <v-toolbar app fixed flat id="main-toolbar">
    <v-toolbar-side-icon @click.stop="drawer = !drawer" v-if="mini"/>
    <v-toolbar-title v-if="title">{{ title }}</v-toolbar-title>
    <v-spacer/>
    <v-toolbar-items>
      <v-flex align-center layout>
        <v-menu bottom left offset-y transition="slide-y-transition">
          <v-btn flat icon slot="activator">
            <v-icon>account_circle</v-icon>
          </v-btn>
          <v-card class="tiny-menu">
            <v-list dense>
              <v-list-tile to="/profile">
                <v-list-tile-action><v-icon>person</v-icon></v-list-tile-action>
                <v-list-tile-content><v-list-tile-title>Profile</v-list-tile-title></v-list-tile-content>
              </v-list-tile>
              <v-list-tile to="/login">
                <v-list-tile-action><v-icon>input</v-icon></v-list-tile-action>
                <v-list-tile-content><v-list-tile-title>logout</v-list-tile-title></v-list-tile-content>
              </v-list-tile>
            </v-list>
          </v-card>
        </v-menu>
      </v-flex>
    </v-toolbar-items>
  </v-toolbar>
  <v-content>
    <v-container fluid fill-height>
      <v-layout justify-center align-center>
        <v-flex shrink>
          <router-view/>
        </v-flex>
      </v-layout>
    </v-container>
  </v-content>
  <v-footer inset fixed id="main-footer">
    <v-spacer/>
    <div id="footer-content">
      <a href="https://github.com/thavel/goban"><img src="../assets/github.png" width="16px" height="16px"/></a>
      - Made by
      <a href="https://github.com/thavel"><v-img v-if="author" :src="author" width="16px" height="16px"/></a>
      with
      <v-icon size="16px" color="error">favorite</v-icon>
    </div>
  </v-footer>
</div>
</template>

<script>
import _ from 'lodash';
import axios from 'axios';

export default {
  data: () => ({
    drawer: null,
    mini: false,
    title: null,
    menu: [
      {
        name: 'Absences',
        icon: 'date_range',
        path: '/absences'
      },
      {
        name: 'Settings',
        icon: 'settings',
        path: '/settings'
      }
    ],
    author: null
  }),
  mounted() {
    this.setTitle();
    this.getAuthor();
    this.responsive();
    window.addEventListener('resize', this.responsive);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.responsive);
  },

  methods: {
    responsive: function() {
      this.mini = window.innerWidth <= 1250;
      if (!this.mini) {
        this.drawer = true;
      }
    },
    getAuthor: async function() {
      try {
        let res = await axios.get('https://api.github.com/users/thavel');
        this.author = res.data.avatar_url;
      } catch(e) {
        this.author = '../assets/logo.png';
      }
    },
    setTitle: function() {
      let path = this.$router.history.current.path;
      let item = _.find(this.menu, {path: path});
      this.title = item.name;
    }
  }
}
</script>

<style>
#main #main-toolbar {
  background-color: transparent;
}
#main #main-footer .v-image,
#main #main-footer img {
  display: inline-block;
  vertical-align: middle;
}
#main #main-title-divider {
  margin-bottom: 15px;
}
#main .main-logo div {
  margin: 0 auto;
}
#main .v-avatar img {
  border-radius: 0;
}
#main #footer-content {
  margin-right: 20px;
}
.tiny-menu .v-list__tile__action {
  min-width: 35px;
}
.v-menu__content.menuable__content__active {
  margin-top: 8px;
  border-radius: 14px;
}

</style>
