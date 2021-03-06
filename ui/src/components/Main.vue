<template>
<div id="main">
  <v-navigation-drawer v-model="drawer" fixed app dark>
    <!-- Menu -->
    <v-list>
      <!-- Menu header -->
      <v-list-tile class="menu-header" avatar id="main-logo">
        <v-list-tile-avatar>
          <img src="../assets/mini-logo.png"/>
        </v-list-tile-avatar>
        <v-list-tile-title class="title">
          Goban
        </v-list-tile-title>
      </v-list-tile>
      <v-divider id="main-title-divider"/>
      <!-- Menu content -->
      <div
        class="menu-content"
        v-for="(item, i) in menu"
        :key="i"
      >
        <!-- Menu item -->
        <v-list-tile
          v-if="item.children == undefined"
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
        <!-- Menu groups -->
        <v-list-group
          v-else
          v-model="item.open"
          :to="item.path"
        >
          <v-list-tile slot="activator">
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>{{ item.name }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
          <v-list-tile
            v-for="(child, j) in item.children"
            :key="i+''+j"
            :to="child.path"
            @click="setTitle"
          >
            <v-list-tile-action>
              <v-icon>{{ child.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>{{ child.name }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list-group>
      </div>
    </v-list>
  </v-navigation-drawer>
  <!-- Toolbar -->
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
                <v-list-tile-action>
                  <v-icon>person</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                  <v-list-tile-title>Profile</v-list-tile-title>
                </v-list-tile-content>
              </v-list-tile>
              <v-list-tile to="/login">
                <v-list-tile-action>
                  <v-icon>input</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                  <v-list-tile-title>logout</v-list-tile-title>
                </v-list-tile-content>
              </v-list-tile>
            </v-list>
          </v-card>
        </v-menu>
      </v-flex>
    </v-toolbar-items>
  </v-toolbar>
  <!-- Page content -->
  <v-content>
    <v-container fluid fill-height>
      <v-layout justify-center align-center>
        <v-flex shrink>
          <router-view/>
        </v-flex>
      </v-layout>
    </v-container>
  </v-content>
  <!-- Footer -->
  <v-footer inset fixed id="main-footer">
    <v-spacer/>
    <div id="footer-content">
      <a href="https://github.com/thavel/goban"><img src="../assets/github.png" width="16px" height="16px"/></a>
      - Made by
      <a href="https://github.com/thavel"><v-img v-if="author" :src="author" width="16px" height="16px"/></a>
      with
      <v-icon size="19px" color="error">favorite</v-icon>
    </div>
  </v-footer>
</div>
</template>

<script>
import axios from 'axios';

export default {
  data: () => ({
    drawer: null,
    mini: false,
    title: null,
    menu: [
      {
        name: 'Absences',
        icon: 'beach_access',
        open: false,
        children: [
          {
            name: 'Calendar',
            icon: 'date_range',
            path: '/calendar',
          },
          {
            name: 'My absences',
            icon: 'assignment',
            path: '/absences',
          }
        ]
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
      this.mini = window.innerWidth <= 1260;
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
    searchMenuItem: function(path) {
      for (var i of this.menu) {
        if (i.path === path) {
          return i;
        }
        for (var c of i.children || []) {
          if (c.path === path) {
            return c;
          }
        }
      }
      return null;
    },
    setTitle: function() {
      let item = this.searchMenuItem(this.$router.history.current.path);
      this.title = item.name;
    }
  }
}
</script>

<style>
#main #main-toolbar {
  background-color: transparent;
  color: var(--v-secondary-base);
}
#main #main-footer {
  background-color: transparent;
}
#main #main-footer .v-image,
#main #main-footer img {
  display: inline-block;
  vertical-align: middle;
}
#main #main-title-divider {
  margin-bottom: 35px;
}
#main .v-avatar img {
  margin-top: 10px;
  margin-left: 15px;
  border-radius: 0;
  height: 65px;
  width: auto;
}
#main #footer-content {
  margin-right: 10px;
}
#main .menu-header .title {
  margin-left: 15px;
}
#main .menu-content .v-list__group__header,
#main .menu-content .v-list__tile {
  border-radius: 8px;
}
#main .menu-content .v-list__group__header,
#main .menu-content .v-list__group__items,
#main .menu-content > div[role=listitem] {
  margin: 15px;
}
#main .menu-content .v-list__group__items div[role=listitem] {
  margin-bottom: 15px;
}
#main .menu-content .v-list__tile--active {
  background-color: var(--v-secondary-base);
  color: white !important;
}
#main .menu-content .v-list__group__items .v-list__tile--active {
  background-color: var(--v-tertiary-base);
  color: white !important;
  /* filter: brightness(200%); */
}
#main .menu-content .v-list__group--active:after,
#main .menu-content .v-list__group--active:before {
  background: none;
}
.tiny-menu .v-list__tile__action {
  min-width: 35px;
}
.v-menu__content.menuable__content__active {
  margin-top: 8px;
  border-radius: 8px;
}
.colorful-title.v-sheet {
  border-radius: 4px;
}
.colorful-title .v-card__title {
  display: inline-block;
  position: relative;
  padding: 8px;
  width: 94%;
  left: 3%;
  right: 3%;
  margin-top: -25px;
  margin-bottom: 15px;
  background-color: var(--v-secondary-base);
  border-radius: 8px !important;
  color: white;
  -webkit-box-shadow: 0 12px 20px -10px rgba(52,87,29,.28),
                      0 4px 20px 0 rgba(0,0,0,.12),
                      0 7px 8px -5px rgba(52,87,29,.2) !important;
  box-shadow: 0 12px 20px -10px rgba(52,87,29,.28),
              0 4px 20px 0 rgba(0,0,0,.12),
              0 7px 8px -5px rgba(52,87,29,.2) !important;
}
</style>
