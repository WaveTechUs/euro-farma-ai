<script setup>
import { ref, onMounted, watch } from 'vue';
import '@/assets/index.css';
import MainLayout from './../layout/MainLayout.vue';
import Chart from 'chart.js/auto';
import { GetDataDash } from '@/composables/Dashboard';

let numProjetos = ref('');
let studies = ref([]);
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
    const [datePart] = dateString.split(' ');
    
    const [year] = datePart.split('-');

    return year;
}

const chartRef5 = ref(null);

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
                        'rgb(255, 99, 132)',
                        'rgb(54, 162, 235)',
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


const countTopics = (data) => {
    let topicCount = {};

    data.forEach(study => {
        const topic = study.Topic || 'Unknown';
           if (topicCount[topic]) {
               topicCount[topic]++;
           } else {
               topicCount[topic] = 1;
           }

       
    });

    return topicCount;
};
const countTopicsOnGoing = (data) => {
    let topicCount = {};


    data.forEach(study => {
        const topic = study.Topic  || 'Unknown';
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

const createChart2 = (topicCounts) => {
    if (chartRef5.value) {
        new Chart(chartRef2.value, {
            type: 'bar',
            data: {
                labels: Object.keys(topicCounts),
                datasets: [{
                    label: 'Quantidade de Estudos por Tópico',
                    data: Object.values(topicCounts),
                    backgroundColor: 'rgb(75, 192, 192)',
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

const countDate = (data) => {
    let dateCount = {};

    data.forEach(study => {
        const dates = getYearFromDate(study.CreatedAt) || 'Unknown';

        if (dateCount[dates]) {
            dateCount[dates]++;
        } else {
            dateCount[dates] = 1;
        }
    });

    return dateCount;
};

const createChart3 = (topicCounts) => {
    if (chartRef4.value) {
        new Chart(chartRef4.value, {
            type: 'bar',
            data: {
                labels: Object.keys(topicCounts),
                datasets: [{
                    label: 'Número de pesquisas dos ultimos anos:',
                    data: Object.values(topicCounts),
                    backgroundColor: 'rgb(75, 192, 192)',
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
                    data: Object.values(topicCounts),
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

watch(studies, (newVal) => {
    if (newVal.length > 0) {
        const statusCounts = countStatuses(newVal);
        const topicCounts = countTopics(newVal);
        const topicDate = countDate(newVal);
        const topicDateOnGoing = countTopicsOnGoing(newVal);

        createChart(statusCounts);
        createChart2(topicCounts);
        createChart3(topicDate);
        createChart4(topicDateOnGoing);
        arrNames.value = removeDuplicates(arrNamesbefore.value)
        console.log(arrNames.value);
        
    }
});

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
