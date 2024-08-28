<script setup>
import { nextTick, onMounted, ref } from 'vue';
import MainLayout from './../layout/MainLayout.vue'
import { PostAsk } from '@/composables/Chat';
let asking = ref('')
// let history = ref([{OldText: "teste", Role: "user"}, {OldText: "Lorem ipsum dolor sit suscipit nisi perspiciatis eligendi adipisci a incidunt atque corrupti tempora ducimus praesentium dolorum! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem minima veritatis enim error in, consequuntur quaerat, ex eum voluptate quod maxime. Consectetur consequatur quo illum accusamus itaque! Ad, voluptates nulla.Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti ab eos molestias beatae molestiae voluptate reprehenderit possimus reiciendis sapiente voluptas dolorem laudantium debitis quis, qui architecto repellat delectus. Alias, labore! Lorem ipsum dolor sit amet consectetur adipisicing elit. Vel facilis architecto doloremque quas, mollitia animi quibusdam nobis, suscipit nisi perspiciatis eligendi adipisci a incidunt atque corrupti tempora ducimus praesentium dolorum! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem minima veritatis enim error in, consequuntur quaerat, ex eum voluptate quod maxime. Consectetur consequatur quo illum accusamus itaque! Ad, voluptates nulla.", Role: "model"},{OldText: "teste", Role: "user"}, {OldText: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti ab eos molestias beatae molestiae voluptate reprehenderit possimus reiciendis sapiente voluptas dolorem laudantium debitis quis, qui architecto repellat delectus. Alias, labore! Lorem ipsum dolor sit amet consectetur adipisicing elit. Vel facilis architecto doloremque quas, mollitia animi quibusdam nobis, suscipit nisi perspiciatis eligendi adipisci a incidunt atque corrupti tempora ducimus praesentium dolorum! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem minima veritatis enim error in, consequuntur quaerat, ex eum voluptate quod maxime. Consectetur consequatur quo illum accusamus itaque! Ad, voluptates nulla.Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti ab eos molestias beatae molestiae voluptate reprehenderit possimus reiciendis sapiente voluptas dolorem laudantium debitis quis, qui architecto repellat delectus. Alias, labore! Lorem ipsum dolor sit amet consectetur adipisicing elit. Vel facilis architecto doloremque quas, mollitia animi quibusdam nobis, suscipit nisi perspiciatis eligendi adipisci a incidunt atque corrupti tempora ducimus praesentium dolorum! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem minima veritatis enim error in, consequuntur quaerat, ex eum voluptate quod maxime. Consectetur consequatur quo illum accusamus itaque! Ad, voluptates nulla.Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti ab eos molestias beatae molestiae voluptate reprehenderit possimus reiciendis sapiente voluptas dolorem laudantium debitis quis, qui architecto repellat delectus. Alias, labore! Lorem ipsum dolor sit amet consectetur adipisicing elit. Vel facilis architecto doloremque quas, mollitia animi quibusdam nobis, suscipit nisi perspiciatis eligendi adipisci a incidunt atque corrupti tempora ducimus praesentium dolorum! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem minima veritatis enim error in, consequuntur quaerat, ex eum voluptate quod maxime. Consectetur consequatur quo illum accusamus itaque! Ad, voluptates nulla.Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti ab eos molestias beatae molestiae voluptate reprehenderit possimus reiciendis sapiente voluptas dolorem laudantium debitis quis, qui architecto repellat delectus. Alias, labore! Lorem ipsum dolor sit amet consectetur adipisicing elit. Vel facilis architecto doloremque quas, mollitia animi quibusdam nobis, suscipit nisi perspiciatis eligendi adipisci a incidunt atque corrupti tempora ducimus praesentium dolorum! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem minima veritatis enim error in, consequuntur quaerat, ex eum voluptate quod maxime. Consectetur consequatur quo illum accusamus itaque! Ad, voluptates nulla.Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti ab eos molestias beatae molestiae voluptate reprehenderit possimus reiciendis sapiente voluptas dolorem laudantium debitis quis, qui architecto repellat delectus. Alias, labore! Lorem ipsum dolor sit amet consectetur adipisicing elit. Vel facilis architecto doloremque quas, mollitia animi quibusdam nobis, suscipit nisi perspiciatis eligendi adipisci a incidunt atque corrupti tempora ducimus praesentium dolorum! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem minima veritatis enim error in, consequuntur quaerat, ex eum voluptate quod maxime. Consectetur consequatur quo illum accusamus itaque! Ad, voluptates nulla.Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti ab eos molestias beatae molestiae voluptate reprehenderit possimus reiciendis sapiente voluptas dolorem laudantium debitis quis, qui architecto repellat delectus. Alias, labore! Lorem ipsum dolor sit amet consectetur adipisicing elit. Vel facilis architecto doloremque quas, mollitia animi quibusdam nobis, suscipit nisi perspiciatis eligendi adipisci a incidunt atque corrupti tempora ducimus praesentium dolorum! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem minima veritatis enim error in, consequuntur quaerat, ex eum voluptate quod maxime. Consectetur consequatur quo illum accusamus itaque! Ad, voluptates nulla.Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti ab eos molestias beatae molestiae voluptate reprehenderit possimus reiciendis sapiente voluptas dolorem laudantium debitis quis, qui architecto repellat delectus. Alias, labore! Lorem ipsum dolor sit amet consectetur adipisicing elit. Vel facilis architecto doloremque quas, mollitia animi quibusdam nobis, suscipit nisi perspiciatis eligendi adipisci a incidunt atque corrupti tempora ducimus praesentium dolorum! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem minima veritatis enim error .", Role: "model"}])
const history = ref([]);
const ststusUser = ref('')
const messageContainer = ref(null);

const testando = async () => {

    nextTick(() => {
        setTimeout(() => {
            if (messageContainer.value) {
                console.log('Rolling to bottom');
                messageContainer.value.scrollTop = messageContainer.value.scrollHeight;
            } else {
                console.log('messageContainer is null');
            }
        }, 100);
    });

    let userMessage = {
        OldText: asking.value,
        Role: 'user'
    };
    history.value.push(userMessage);

    const answear = await PostAsk(asking.value);
    let resp = formatAnswear(answear);

    if (answear) {
        let modelResponse = {
            OldText: resp,
            Role: 'model'
        };
        history.value.push(modelResponse);
    }

    nextTick(() => {
        setTimeout(() => {
            if (messageContainer.value) {
                console.log('Rolling to bottom');
                messageContainer.value.scrollTop = messageContainer.value.scrollHeight;
            } else {
                console.log('messageContainer is null');
            }
        }, 100);
    });
};





const formatAnswear = (resp) => {

    const answear = resp.Parts.join('');
    let formattedAnswerOne = answear.replace(/\n/g, '<br>');
    const formattedAnswer = formattedAnswerOne.replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>');
    return formattedAnswer;
};



// onMounted(async()=>{
//     formatAnswear()
// })
// min-h-screen
</script>

<template>
    <MainLayout>
        <div class="flex flex-col min-h-screen bg-slate-200">
            <div class="flex w-auto md:ml-10 flex-col flex-grow p-1 ">
                <div ref="messageContainer" class="w-full flex-grow bg-slate-200 border flex flex-col scroll-container pb-20"> <!-- Adicionado pb-20 -->
                    <div class="h-12 bg-slate-200">
                        <h2 class="text-4xl">Chat</h2>
                        <p>DashBoard / customers' List</p>
                    </div>
                    <div v-for="a in history" :key="a.OldText" class="h-auto pb-10"> <!-- Alterado de h-screen para h-auto -->
                        <div v-if="a.Role === 'user'"
                            class=" flex-grow card text-left rounded-md mt-10 md:w-3/6 rounded-tr-none p-2 ml-auto md:mr-16 bg-slate-400 h-auto shadow-md">
                            <p class="text-sm">{{ a.OldText }}</p>
                            <p class="mt-5 text-slate-100">you</p>
                        </div>
                        <div v-if="a.Role === 'model'"
                            class="flex-grow card text-left rounded-md  md:w-11/12 m-auto rounded-tl-none p-2 bg-white h-auto shadow-md">
                            <p class="text-sm" v-html="a.OldText"></p>
                            <p class="mt-5 text-slate-500">chat</p>
                        </div>
                    </div>
                </div>
            </div>
            <div class="w-5/6 md:w-auto lg:w-11/12 lg:ml-10 h-auto bg-slate-200 mt-20 fixed bottom-0 p-2">
                <div v-if="ststusUser == 'Gestor'" class="md:gap-4 md:grid-cols-5 hidden md:grid">
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">QUAL A PALAVRA MAIS UTILIZADA?</button>
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">QUAL A MAIOR RECLAMAÇÃO?</button>
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">QUAL O MAIOR ELOGIO?</button>
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">ASSUNTO MAIS COMUM</button>
                    <button class="bg-sky-600 text-sm text-white rounded-xl px-4 py-2">PRINCIPAIS IDEIAS</button>
                </div>
                <div v-else></div>
                <div class="grid grid-cols-12 gap-1 mt-5">
                    <input class="p-2 h-10 border col-span-11 rounded-3xl" v-model="asking" type="text"
                        placeholder="teste">
                    <div class="flex items-center justify-center border">
                        <button @click="testando()"
                            class="bg-slate-800 rounded-3xl hover:bg-slate-900 h-10 hidden md:block">
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

.scroll-container {
    overflow-y: auto;
    height: calc(100vh - 5rem);
    /* Ajuste conforme necessário */
}
</style>