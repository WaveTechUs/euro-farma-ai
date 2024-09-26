<script setup>
import { ref, onMounted, watch } from 'vue';
import '@/assets/index.css';
import MainLayout from './../layout/MainLayout.vue';
import Chart from 'chart.js/auto';
import { GetDataDash } from '@/composables/Dashboard';

// API Call
let numProjetos = ref('');
let studies = ref([]); // Inicialize como array vazio
let arrNames = ref([])
let arrNamesbefore = ref([])

let getDataAPi = async () => {
    let answear = await GetDataDash();
    numProjetos.value = answear.length;
    studies.value = answear;
    console.log(studies.value);
    for (let i = 0; i < studies.value.length; i++) {
        arrNamesbefore.value.push(studies.value[i].Topic)
        console.log(studies.value[i].Topic);
        
    }
    console.log(arrNamesbefore.value);
    
};
function removeDuplicates(arr) {
    return [...new Set(arr)];
}


function getYearFromDate(dateString) {
    // Divide a string da data em partes, separando por espaço
    const [datePart] = dateString.split(' ');
    
    // Divide a parte da data em ano, mês e dia, separando por hífen
    const [year] = datePart.split('-');
    
    // Retorna o ano
    return year;
}
// Chart references
const chartRef5 = ref(null);

// Função para contar os status
const countStatuses = (data) => {
    let statusCount = {
        "Finalizado": 0,
        "Andamento": 0,
        "Analise": 0
    };

    data.forEach(study => {
        if (study.Status === "Completed") {
            statusCount["Finalizado"]++;
        } else if (study.Status === "Ongoing") {
            statusCount["Andamento"]++;
        }
    });

    return statusCount;
};


// Função para criar o gráfico
const createChart = (statusCounts) => {
    if (chartRef5.value) {
        new Chart(chartRef5.value, {
            type: 'doughnut',
            data: {
                labels: ['Finalizado', 'Andamento'],
                datasets: [{
                    label: 'Status dos Estudos',
                    data: [statusCounts.Finalizado, statusCounts.Andamento],
                    backgroundColor: [
                        'rgb(255, 99, 132)', // Cor para 'Finalizado'
                        'rgb(54, 162, 235)', // Cor para 'Andamento'
                        // Cor para 'Análise'
                    ],
                    hoverOffset: 4
                }]
            }
        });
    }
};
const chartRef2 = ref(null);
const chartRef4 = ref(null);
const chartRef6 = ref(null);

// Função para contar os tópicos
const countTopics = (data) => {
    let topicCount = {}; // Objeto para contar tópicos dinamicamente

    data.forEach(study => {
        const topic = study.Topic || 'Unknown'; // Verifica se o campo Topic existe
           if (topicCount[topic]) {
               topicCount[topic]++;
           } else {
               topicCount[topic] = 1;
           }

       
    });

    return topicCount;
};
const countTopicsOnGoing = (data) => {
    let topicCount = {}; // Objeto para contar tópicos dinamicamente


    data.forEach(study => {
        const topic = study.Topic  || 'Unknown'; // Verifica se o campo Topic existe
       if(study.Status !== "Completed"){
           if (topicCount[topic]) {

               topicCount[topic]++;
           } else {
               topicCount[topic] = 1;
           }

       }
    });

    return topicCount;
};

// Função para criar o gráfico
const createChart2 = (topicCounts) => {
    if (chartRef5.value) {
        new Chart(chartRef2.value, {
            type: 'bar', // Altere para 'bar' para barras
            data: {
                labels: Object.keys(topicCounts), // Os diferentes tópicos
                datasets: [{
                    label: 'Quantidade de Estudos por Tópico',
                    data: Object.values(topicCounts), // Quantidade de estudos por tópico
                    backgroundColor: 'rgb(75, 192, 192)', // Cor para as barras
                    hoverOffset: 4
                }]
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true
                    }
                }
            }
        });
    }
};
// Função para contar os Date
const countDate = (data) => {
    let dateCount = {}; // Objeto para contar tópicos dinamicamente

    data.forEach(study => {
        const dates = getYearFromDate(study.CreatedAt) || 'Unknown'; // Verifica se o campo Topic existe

        if (dateCount[dates]) {
            dateCount[dates]++;
        } else {
            dateCount[dates] = 1;
        }
    });

    return dateCount;
};

// Função para criar o gráfico

const createChart3 = (topicCounts) => {
    if (chartRef4.value) {
        new Chart(chartRef4.value, {
            type: 'bar', // Altere para 'bar' para barras
            data: {
                labels: Object.keys(topicCounts), // Os diferentes tópicos
                datasets: [{
                    label: 'Número de pesquisas dos ultimos anos:',
                    data: Object.values(topicCounts), // Quantidade de estudos por tópico
                    backgroundColor: 'rgb(75, 192, 192)', // Cor para as barras
                    hoverOffset: 4
                }]
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true
                    }
                }
            }
        });
    }
};
const createChart4 = (topicCounts) => {
    console.log(topicCounts);
    
    if (chartRef6.value) {
        new Chart(chartRef6.value, {
            type: 'doughnut',
            options:{
                plugins:{
                    legend:{
                        display: false,
                    },
                },
            },
            data: {
                labels: Object.keys(topicCounts),
                datasets: [{
                    label: "",
                    data: Object.values(topicCounts), // Quantidade de estudos por tópico
                    backgroundColor: [
                        'rgb(255, 99, 132)', 
                        'rgb(54, 162, 235)', 
                        'rgb(51, 255, 87)', 
                        'rgb(255, 51, 161)', 
                        'rgb(255, 209, 51)', 
                        'rgb(255, 165, 0)', 
                        'rgb(128, 128, 128)', 
                        'rgb(0, 128, 0)', 
                        ' rgb(139, 69, 19)', 
                        'rgb(50, 205, 50)', 
                      
                    ],
                    hoverOffset: 4
                }]
            }
        });
    }
};
// Watch para detectar quando `studies` é atualizado
watch(studies, (newVal) => {
    if (newVal.length > 0) { // Verifica se há dados
        const statusCounts = countStatuses(newVal);
        const topicCounts = countTopics(newVal); // Conta tópicos dinamicamente
        const topicDate = countDate(newVal); // Conta tópicos dinamicamente
        const topicDateOnGoing = countTopicsOnGoing(newVal); // Conta tópicos dinamicamente

        createChart(statusCounts); // Cria o gráfico quando os dados estiverem disponíveis
        createChart2(topicCounts); // Cria o gráfico quando os dados estiverem disponíveis
        createChart3(topicDate); // Cria o gráfico quando os dados estiverem disponíveis
        createChart4(topicDateOnGoing); // Cria o gráfico quando os dados estiverem disponíveis
        arrNames.value = removeDuplicates(arrNamesbefore.value)
        console.log(arrNames.value);
        
    }
});

// Chamada à API quando o componente é montado
onMounted(() => {
    getDataAPi();
});
</script>




<template>
    <MainLayout>
        <div class="w-full h-full   border bg-slate-200 flex flex-col">
            <div class="h-8 ml-20 mb-20 w-2/3">
                <h2 class="text-4xl mb-2">DashBoard</h2>
                <p>Dashboard / customers' List</p>
            </div>

            <div class="w-8/12 mx-auto h-auto grid grid-rows-2 grid-cols-4 gap-4 pb-4 px-4">
                <!-- Primeira linha -->
                <div class="col-span-1 rounded-2xl bg-white flex  text-center flex-col items-center">
                    <h2 class="text-2xl w-auto  h-auto text-center mt-10 mb-10 font-semibold ">Número de Pesquisas</h2>

                    <h2 class="text-5xl w-auto h-auto font-semibold ">{{ numProjetos }}</h2>
                </div>
                <div class="col-span-1 rounded-2xl bg-white flex justify-center items-center">
                    <canvas class="h-screen" ref="chartRef5"></canvas>
                </div>
                <div class="col-span-2 rounded-2xl bg-white flex justify-center items-center">
                    <canvas style="" ref="chartRef4"></canvas>
                </div>

                <!-- Segunda linha -->
                <div class="col-span-2 rounded-2xl bg-white flex justify-center text-center items-center">
                    <canvas class="h-screen" ref="chartRef2"></canvas>

                </div>
                <div class="col-span-1 rounded-2xl bg-white flex justify-center flex-col items-center">
                    <h2 class="text-xl w-auto h-auto text-center m-2 font-semibold ">Pesquisas por tópicos</h2>

                    <canvas style="" ref="chartRef6"></canvas>

                </div>
                <div class="col-span-1 p-2 rounded-2xl bg-white flex justify-center flex-col items-center">
                    <h2 class="text-xl w-auto h-auto text-center m-2 font-semibold ">Pesquisas</h2>
                    <ul v-for="(item, index) in arrNames.slice(0, 5) " class="">
                        <li class="h-5 py-0 my-0 ">{{index + 1}} - {{item}}</li>
                    </ul>

                </div>
            </div>

        </div>
    </MainLayout>
</template>



<style lang="scss" scoped></style>