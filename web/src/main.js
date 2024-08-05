import './assets/index.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from '@/routes/index'
import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'

// import './middlewares/interceptor'
// import 'video.js/dist/video-js.css'

const app = createApp(App)

app.use(router)

app.mount('#app')