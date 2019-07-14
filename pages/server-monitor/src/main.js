import Vue from 'vue'
import App from './App.vue'
import VueUi from '@vue/ui'
import store from './store'
import '@vue/ui/dist/vue-ui.css'
import ECharts from 'vue-echarts'

Vue.use(VueUi)
Vue.config.productionTip = false

new Vue({
  store,
  render: h => h(App),
}).$mount('#app')
