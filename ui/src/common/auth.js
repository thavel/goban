export default {
  token: null,
  set(token) {
    this.token = token;
    localStorage.token = token;
  },
  retrive() {
    if (!('token' in localStorage)) {
      return null;
    }
    this.token = localStorage.token;
    return this.token;
  },
  clean() {
    this.token = null;
    localStorage.removeItem('token');
  },
  header() {
    return {headers: {'Authorization': 'Bearer ' + this.token}};
  },
  wrap(fun) {
    return async function() {
      try {
        await fun.apply(this, arguments);
      } catch(e) {
        if (e.response != undefined && e.response.status == 401) {
          this.$router.replace({path: '/login'});
        }
      }
    }
  }
};
