import '~/plugins/composition-api';

import Vue from 'vue';
import Toasted from 'vue-toasted';
import App from '~/App.vue';
import store from '~/store';
import '~/assets/css/fonts.css';
import '~/assets/css/global.css';

Vue.use(Toasted);

const app = new Vue({
  store,
  render: (h) => h(App),
}).$mount('#app');

export default app;
