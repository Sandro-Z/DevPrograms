import { createApp } from 'vue'
import App from './App.vue'
import router from './router/router.js'
import store from './store/index.js'
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'

import '../src/assets/styles/app.scss'

createApp(App).use(router).use(store).mount('#app')
