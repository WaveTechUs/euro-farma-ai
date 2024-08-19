<script setup>
import { ref, onMounted, watch } from 'vue';
import '@/assets/index.css';
import { useGoTo } from '@/composables/Utils';
import { listUser } from '@/composables/User';

onMounted(async()=>{
    await listUser()

})
import { getUser } from '@/composables/User';
</script>

<template>
    <div class="w-full h-full flex items-center justify-center  bg-gray-900">
        <div
            class="w-3/12 h-8/10 md:h-3/5 xl:h-4/6 md:w-4/12 xl:w-3/12 bg-slate-50 rounded-2xl p-5 flex flex-col justify-between ">
            <img class="w-2/3 text-center mx-auto" src="../assets/images/eurofarma-logo.png" alt="">
            <form class="flex flex-col  w-full  h-auto">
                <div class="flex flex-col  h-auto p-1   rounded-lg ">
                    <label for="email" class="block text-md font-semibold text-gray-700 ">Email</label>
                    <input type="email" id="email" name="email" required
                        class="block shadow-md w-full h-12 px-3 py-4 border border-gray-300 rounded-md  focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                </div>

                <div class="flex flex-col mb:2  h-auto p-1 mt-4 rounded-lg ">
                    <label for="password" class="block text-md font-semibold text-gray-700 mb-1">Senha</label>
                    <input type="password" id="password" name="password" required
                        class="block shadow-md w-full h-12 px-3 py-4 border border-gray-300 rounded-md  focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                    <a href="#"
                        class="text-sm w-3/5   align-middle justify-center pt-2 text-indigo-600  hover:text-indigo-700" >Esqueceu a senha?</a>
                </div>
                <div
                    class="flex flex-row mt-10 xl:mt-8 justify-between align-middle w-full   px-2  h-auto  rounded-lg ">
                    <a href="#"
                        class="text-sm w-3/5   align-middle justify-center pt-2 text-indigo-600  hover:text-indigo-700" @click="useGoTo('/criar-conta')">Não
                        possui conta?</a>
                    <button type="submit" @click="getUser()"
                        class="w-2/5 py-3 px-4  bg-indigo-600 text-white font-semibold rounded-md shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">

                        Entrar
                    </button>


                </div>
            </form>
            <div class="text-2xl w-full h-1 font-bold text-center "></div>

        </div>
    </div>


</template>



<style lang="scss" scoped></style>
