<script setup>
import MainLayout from './../layout/MainLayout.vue'
</script>

<template>
    <MainLayout>
        <div class="flex flex-col min-h-screen bg-slate-200">
            <div class="flex w-auto md:ml-10 flex-col flex-grow p-1">
                <div class="w-full flex-grow bg-slate-200 border flex flex-col">
                    <div class="h-12 bg-slate-200">
                        <h2 class="text-4xl">Chat</h2>
                        <p>DashBoard / custumers' List</p>
                    </div>
                    <div class="flex-grow mb-36  overflow-y-auto">
                        <div class="card text-left rounded-md my-12 md:w-2/3 rounded-tl-none p-2 bg-white h-auto shadow-md">
                            <p class="text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Sit, illo. Eaque
                                odio saepe quas eum numquam, rerum quia beatae, est, nam dignissimos sapiente perferendis
                                velit molestias commodi totam exercitationem corporis?</p>
                                <p class="mt-5 text-slate-500">chat</p>
                        </div>
                        <div class="card text-left rounded-md my-12 md:w-2/3 rounded-tr-none p-2 ml-auto md:mr-16 bg-slate-400 h-auto shadow-md">
                            <p class="text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Sit, illo. Eaque
                                odio saepe quas eum numquam, rerum quia beatae, est, nam dignissimos sapiente perferendis
                                velit molestias commodi totam exercitationem corporis?</p>
                                <p class="mt-5 text-slate-100">you</p>

                        </div>
                        <div class="card text-left rounded-md my-12 md:w-2/3 rounded-tl-none p-2 bg-white h-auto shadow-md">
                            <p class="text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Sit, illo. Eaque
                                odio saepe quas eum numquam, rerum quia beatae, est, nam dignissimos sapiente perferendis
                                velit molestias commodi totam exercitationem corporis?</p>
                                <p class="mt-5 text-slate-500">chat</p>

                        </div>
                        <div class="card text-left rounded-md my-12 md:w-2/3 rounded-tl-none p-2 bg-white h-auto shadow-md">
                            <p class="text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Sit, illo. Eaque
                                odio saepe quas eum numquam, rerum quia beatae, est, nam dignissimos sapiente perferendis
                                velit molestias commodi totam exercitationem corporis?</p>
                                <p class="mt-5 text-slate-500">chat</p>

                        </div>
                        <div class="card text-left rounded-md my-12 md:w-2/3 rounded-tl-none p-2 bg-white h-auto shadow-md">
                            <p class="text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Sit, illo. Eaque
                                odio saepe quas eum numquam, rerum quia beatae, est, nam dignissimos sapiente perferendis
                                velit molestias commodi totam exercitationem corporis?</p>
                                <p class="mt-5 text-slate-500">chat</p>

                        </div>
                        <div class="card text-left rounded-md my-12 md:w-2/3 rounded-tl-none p-2 bg-white h-auto shadow-md">
                            <p class="text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Sit, illo. Eaque
                                odio saepe quas eum numquam, rerum quia beatae, est, nam dignissimos sapiente perferendis
                                velit molestias commodi totam exercitationem corporis?</p>
                                <p class="mt-5 text-slate-500">chat</p>

                        </div>
                        <div class="card text-left rounded-md my-12 md:w-2/3 rounded-tl-none p-2 bg-white h-auto shadow-md">
                            <p class="text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Sit, illo. Eaque
                                odio saepe quas eum numquam, rerum quia beatae, est, nam dignissimos sapiente perferendis
                                velit molestias commodi totam exercitationem corporis?</p>
                                <p class="mt-5 text-slate-500">chat</p>

                        </div>
                        <div class="card text-left rounded-md my-12 md:w-2/3 rounded-tl-none p-2 bg-white h-auto shadow-md">
                            <p class="text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Sit, illo. Eaque
                                odio saepe quas eum numquam, rerum quia beatae, est, nam dignissimos sapiente perferendis
                                velit molestias commodi totam exercitationem corporis?</p>
                                <p class="mt-5 text-slate-500">chat</p>

                        </div>
                        <div class="card text-left rounded-md my-12 md:w-2/3 rounded-tl-none p-2 bg-white h-auto shadow-md">
                            <p class="text-sm">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Sit, illo. Eaque
                                odio saepe quas eum numquam, rerum quia beatae, est, nam dignissimos sapiente perferendis
                                velit molestias commodi totam exercitationem corporis?</p>
                                <p class="mt-5 text-slate-500">chat</p>

                        </div>
                  
            
                    </div>
                </div>
            </div>
            <div class="w-5/6 md:w-auto lg:w-11/12 lg:ml-10 h-auto bg-slate-200 fixed bottom-0  p-2">
                <div class="md:gap-4 md:grid-cols-5 hidden md:grid">
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">QUAL A PALAVRA MAIS
                        UTILIZADA?</button>
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">QUAL A MAIOR RECLAMAÇÃO?</button>
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">QUAL O MAIOR ELOGIO?</button>
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">ASSUNTO MAIS COMUM</button>
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">PRINCIPAIS IDEIAS</button>
                </div>
                <div class="grid grid-cols-12 gap-1 mt-5">
                    <input class="p-2 h-10 border col-span-11 rounded-3xl" type="text" placeholder="teste">
                    <div class="flex items-center justify-center border">
                        <button class="bg-slate-800 rounded-3xl hover:bg-slate-900 h-10 hidden md:block">
                            <i class="fa-solid fa-magnifying-glass" style="color: #ffffff;"></i>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </MainLayout>
</template>


<style lang="scss" scoped>
@import '../scss/global.scss';
</style>