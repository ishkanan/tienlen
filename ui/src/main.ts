import './assets/css/fonts.css'
import './assets/css/global.css'

import { createApp, provide } from 'vue'
import Toasted from 'vue-toasted'
import { createPinia } from 'pinia'
import App from './App.vue'
import { init } from './lib/socket'

export const app = createApp(App)

app.use(createPinia())
app.use(Toasted)

provide('socket', init())

app.mount('#app')
