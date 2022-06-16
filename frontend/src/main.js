import Vue from 'vue'
import loader from "vue-ui-preloader";
import { BootstrapVue } from 'bootstrap-vue';

import App from '@/App.vue'
import router from '@/router/index';
import store from '@/store/index';

Vue.config.productionTip = false

Vue.use(BootstrapVue);
Vue.use(loader);

new Vue({
  render: h => h(App),
  store,
  router,
}).$mount('#app')
