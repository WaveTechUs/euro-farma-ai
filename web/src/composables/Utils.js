import router from '@/routes/index.js'
import {
  createToaster
} from "@meforma/vue-toaster";

const toaster = createToaster({
  position: 'bottom',
  duration: 4000
})

export const useGoTo = (path) => {
  router.push(path);
}

export const useToast = () => {
  return toaster
}

export const useCreateLocalStorage = (type, use) => {
  const userString = JSON.stringify(use);
  localStorage.setItem(type, userString);
}
