<script setup>
import { ref, onMounted, watch } from 'vue';
import '@/assets/index.css';
import { useGoTo } from '@/composables/Utils';
import { postUser } from '@/composables/User';
import { authUserPost } from '@/composables/Login';

let nome = ref('')
let email = ref('')
let password = ref('')
let confirm = ref('')
const confirCreate = async (nome, email, password, confirm)=>{
    if(confirm == password){
        let post = await authUserPost(email, password, nome)
    }
}

</script>

<template>
    <div class="w-full h-full flex items-center justify-center  bg-gray-900">
        <div
            class="w-3/12 h-4/5 md:h-3/5 xl:h-5/10 md:w-7/12 xl:w-6/12 bg-slate-50 rounded-2xl p-5 flex flex-col justify-between ">
            <img class="w-1/3 text-center my-0 mx-auto" src="../assets/images/eurofarma-logo.png" alt="">
            <form class="flex flex-row  w-full my-10  h-auto ">
                <div class="flex flex-col  w-full  h-auto">

                    <div class="flex flex-col  h-auto p-1   rounded-lg ">
                        <label for="nome" class="block text-md font-semibold text-gray-700 ">Nome</label>
                        <input type="nome" id="nome" v-model="nome" name="nome" required
                            class="block shadow-md w-full h-12 px-3 py-4 border border-gray-300 rounded-md  focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                    </div>

                    <div class="flex flex-col mb:2  h-auto p-1 mt-4 rounded-lg ">
                        <label for="password" class="block text-md font-semibold text-gray-700 mb-1">Senha</label>
                        <input type="password" id="password" v-model="password" name="password" required
                            class="block shadow-md w-full h-12 px-3 py-4 border border-gray-300 rounded-md  focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />

                    </div>

                </div>
                <div class="flex flex-col  w-full  h-auto">

                    <div class="flex flex-col  h-auto p-1   rounded-lg ">
                        <label for="email" class="block text-md font-semibold text-gray-700 ">Email</label>
                        <input type="email" id="email" v-model="email" name="email" required
                            class="block shadow-md w-full h-12 px-3 py-4 border border-gray-300 rounded-md  focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                    </div>

                    <div class="flex flex-col mb:2  h-auto p-1 mt-4 rounded-lg ">
                        <label for="confirm-password" class="block text-md font-semibold text-gray-700 mb-1">Confirmar senha</label>
                        <input type="password" id="confirm-password" v-model="confirm" name="confirm-password" required
                            class="block shadow-md w-full h-12 px-3 py-4 border border-gray-300 rounded-md  focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />

                    </div>
                    <div
                        class="flex   flex-row mt-10 xl:mt-8 justify-end align-middle w-full   px-2  h-auto  rounded-lg ">

                        <button type="submit" @click.prevent="confirCreate(nome, email, password, confirm)"
                            class="w-2/5 py-3 px-4  bg-indigo-600 text-white font-semibold rounded-md shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">

                            Entrar
                        </button>


                    </div>
                </div>
            </form>
            <div class="text-2xl w-full h-1 font-bold text-center "></div>

        </div>
    </div>


</template>



<style lang="scss" scoped></style>
