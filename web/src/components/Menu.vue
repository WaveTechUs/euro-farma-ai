<script setup>
import { ref, computed, onMounted } from 'vue';
import { useGoTo } from '@/composables/Utils';
import { useRoute } from 'vue-router';

const route = useRoute();
const x = ref(false);
let page = ref('')


function teste() {
  x.value = !x.value;
}

onMounted(()=>{
  if(route.path == '/chat') page = true
  else page = false

})

</script>

<template>
  <div
    :class="{
      'w-4/6 md:w-3/12 z-10 bg-gray-900 transition-all duration-300 flex flex-col': x,
      'w-12 z-10 bg-gray-800 border-none transition-all duration-300': !x
    }"
    class="h-full fixed top-0  left-0"
  >
    <div
      class="w-full h-20 mt-5 text-center"
      :class="{
        'w-full h-20 mt-5 text-center transition-all duration-300': x,
        'hidden transition-all duration-300': !x
      }"
    >
      <h1 class="text-4xl text-white">Brand.</h1>
      <ul class="h-auto flex justify-center flex-col items-center">
        <li
          @click="useGoTo('/dashboard')"
          :class="{'h-10 flex justify-center items-center rounded-xl cursor-pointer mb-5 px-4 bg-gray-800 w-4/5 bg-opacity-95' : !page, 
        'h-10 flex justify-center items-center rounded-xl cursor-pointer mb-2 px-4 w-4/5' : page }"
        >
          <i
            class="fa-solid fa-chart-pie w-2/12 pr-2 my-auto"
            style="color: #ffffff"
          ></i>
          <p class="text-white text-xl flex w-10/12 items-center">Dashboard</p>
        </li>
        <li
          @click="useGoTo('/chat')"
           :class="{'h-10 flex justify-center items-center rounded-xl cursor-pointer mb-5 px-4 bg-gray-800 w-4/5 bg-opacity-95' : page, 
        'h-10 flex justify-center items-center rounded-xl cursor-pointer mb-2 px-4 w-4/5' : !page }"
        >
          <i
            class="fa-regular fa-message w-2/12 pr-2 my-auto"
            style="color: #ffffff"
          ></i>
          <p class="text-white text-xl flex w-10/12 items-center">Chat</p>
        </li>
        <li
          @click="useGoTo('/Informacoes')"
           class=" h-10 flex justify-center items-center rounded-xl cursor-pointer mb-2 px-4 w-4/5"
        >
        <i class="fa-solid fa-circle-info w-2/12 pr-2 my-auto" style="color: yellow;"></i>

          <p style="color: yellow;" class=" text-xl flex w-10/12 items-center">Informações</p>
        </li>
        <li
          @click="useGoTo('/')"
           class=" h-10 flex justify-center items-center rounded-xl cursor-pointer mb-2 px-4 w-4/5"
        >
        <i class="fa-solid fa-right-from-bracket w-2/12 pr-2 my-auto" style="color: #e00b0b;"></i>
          <p style="color: #e00b0b;" class=" text-xl flex w-10/12 items-center">Sair</p>
        </li>
   
      </ul>
    </div>
    
    <div
      class="absolute -rotate-90 w-auto h-8 top-5 transition-all duration-300"
      :class="{ '-right-5': x, '-right-4': !x }"
    >
      <button @click="teste" :class="{ 'swallow__icon--active': x }" class="swallow__icon">
        <span></span>
      </button>
    </div>
  </div>
</template>


<style scoped>
/* <reset-style> ============================ */
button {
  border: none;
  background: none;
  padding: 0;
  margin: 0;
  cursor: pointer;
  font-family: inherit;
}

/* <main-style> ============================ */
.swallow__icon {
  width: 2rem;
  height: 2rem;
  padding: .25rem;
  background-color: #fff;
  backdrop-filter: saturate(180%) blur(20px);
  border: none;
  border-radius: 0.5em;
  position: relative;
}

.swallow__icon span {
  width: 1.5rem;
  height: .563rem;
  position: absolute;
  top: calc(.25rem + .563rem - 1px);
  left: calc(.25rem - 1px);
  transition: transform 1s cubic-bezier(.86, 0, .07, 1),
    transform-origin 1s cubic-bezier(.86, 0, .07, 1);
}

.swallow__icon span:before,
.swallow__icon span:after {
  content: "";
  width: .75rem;
  height: .125rem;
  background-color: #252525;
  position: absolute;
  bottom: 0;
  transition: transform 1s cubic-bezier(.86, 0, .07, 1),
    transform-origin 1s cubic-bezier(.86, 0, .07, 1);
}

.swallow__icon span:before {
  right: 50%;
  border-radius: 2px 0 0 2px;
  transform-origin: 100% 100%;
  transform: rotate(40deg);
}

.swallow__icon span:after {
  left: 50%;
  border-radius: 0 2px 2px 0;
  transform-origin: 0 100%;
  transform: rotate(-40deg);
}

.swallow__icon--active span {
  transform: translateY(-8px);
}

.swallow__icon--active span:before {
  transform-origin: 100% 0;
  transform: rotate(-40deg);
}

.swallow__icon--active span:after {
  transform-origin: 0 0;
  transform: rotate(40deg);
}
</style>
