import './assets/css/fonts.css'
import './assets/css/global.css'

import { createApp } from 'vue'
import Toasted from 'vue-toasted'
import { createPinia } from 'pinia'
import App from './App.vue'

export const app = createApp(App)

app.use(createPinia())
app.use(Toasted)

app.mount('#app')
