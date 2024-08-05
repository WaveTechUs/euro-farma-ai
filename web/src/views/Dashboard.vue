<script setup>
import { ref, onMounted } from 'vue';
import '@/assets/index.css';
import MainLayout from './../layout/MainLayout.vue';
import Chart from 'chart.js/auto';

const chartRef = ref(null);
const chartRef2 = ref(null);
const chartRef3 = ref(null);

onMounted(() => {
    const data1 = [
        { year: 2010, count: 10 },
        { year: 2011, count: 20 },
        { year: 2012, count: 15 },
        { year: 2013, count: 25 },
        { year: 2014, count: 22 },
        { year: 2015, count: 30 },
        { year: 2016, count: 28 },
    ];

    new Chart(
        chartRef.value,
        {
            type: 'polarArea',
            data: {
                labels: data1.map(row => row.year),
                datasets: [
                    {
                        label: 'Acquisitions by year',
                        data: data1.map(row => row.count),
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        borderColor: 'rgba(75, 192, 192, 1)',
                        borderWidth: 1
                    }
                ]
            }
        }
    );



    new Chart(
        chartRef3.value,
        {
            type: 'doughnut',
            data: {
                labels: [
                    'Red',
                    'Blue',
                    'Yellow'
                ],
                datasets: [{
                    label: 'My First Dataset',
                    data: [300, 50, 100],
                    backgroundColor: [
                        'rgb(255, 99, 132)',
                        'rgb(54, 162, 235)',
                        'rgb(255, 205, 86)'
                    ],
                    hoverOffset: 4
                }]
            }

        }
    );

    const data2 = [
        { year: 2010, count: 65 },
        { year: 2011, count: 59 },
        { year: 2012, count: 80 },
        { year: 2013, count: 81 },
        { year: 2014, count: 56 },
        { year: 2015, count: 55 },
        { year: 2016, count: 40 },
    ];

    new Chart(
        chartRef2.value,
        {
            type: 'bar',
            data: {
                labels: data2.map(row => row.year),
                datasets: [
                    {
                        label: 'Acquisitions by year',
                        data: data2.map(row => row.count),
                        backgroundColor: 'rgba(153, 102, 255, 0.2)',
                        borderColor: 'rgba(153, 102, 255, 1)',
                        borderWidth: 1
                    }
                ]
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true
                    }
                }
            },
        }
    );
});
</script>

<template>
    <MainLayout>
        <div class="w-full h-screen   border bg-slate-200 flex flex-col">
            <div class="h-12 ml-20 mb-20 w-2/3">
                <h2 class="text-4xl mb-2">DashBoard</h2>
                <p>Dashboard / customers' List</p>
            </div>

            <div class="w-4/6 mx-auto h-full grid grid-rows-3 grid-cols-2 gap-4 p-4">
                <!-- Container superior dividido em 4 pequenos -->
                <div class="col-span-2 row-span-1 grid grid-cols-4 gap-4">
                    <div class="col-span-1 rounded-2xl bg-white flex justify-center text-center items-center">
                        <h2 class="text-5xl w-auto h-auto font-semibold ">230</h2>
                    </div>
                    <div class="col-span-1 rounded-2xl bg-white flex justify-center items-center">
                        <h2 class="text-5xl w-auto h-auto font-semibold ">9.80</h2>
                    </div>
                    <div class="col-span-1 rounded-2xl bg-white flex justify-center items-center">
                        <h2 class="text-5xl w-auto h-auto font-semibold ">25.690</h2>
                    </div>
                    <div class="col-span-1 rounded-2xl bg-white flex justify-center items-center"> <canvas
                            class="h-screen" ref="chartRef3"></canvas>
                    </div>
                </div>
                <!-- Segundo container ocupando 1 coluna e 2 linhas -->
                <div class="col-span-1 row-span-2 rounded-2xl bg-white flex justify-center items-center">
                    <canvas class="h-screen" ref="chartRef2"></canvas>
                </div>
                <!-- Terceiro container ocupando 1 coluna e 2 linhas -->
                <div class="col-span-1 row-span-2 rounded-2xl bg-white flex justify-center items-center">
                    <canvas style="" ref="chartRef"></canvas>
                </div>
            </div>
        </div>
    </MainLayout>
</template>



<style lang="scss" scoped></style>