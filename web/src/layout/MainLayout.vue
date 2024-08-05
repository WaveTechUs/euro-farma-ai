<script setup>
import Menu from '../components/Menu.vue'
</script>
<template>
    <main class="w-full h-auto  bg-slate-200 flex flex-row">
        <Menu />
       
        <div class="w-full h-full flex flex-col bg-slate-200 ml-14 ">
            <slot></slot>
          </div>

    </main>
</template>


<style lang="scss" scoped></style>