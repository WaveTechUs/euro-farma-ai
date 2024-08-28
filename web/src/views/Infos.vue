<script setup>
import { ref, onMounted } from 'vue';
import '@/assets/index.css';
import MainLayout from './../layout/MainLayout.vue'
</script>

<template>
        <MainLayout>
            <div class="w-full h-screen overflow-x-hidden border  flex flex-col">
                <div class="h-12 ml-20 mb-20 w-2/3">
                    <h2 class="text-4xl mb-2">Informações</h2>
                    <p>Infos / versions</p>
                </div>
                <div class="ml-20 w-10/12 h-auto border  flex flex-row">
                    <div class="w-6/12 flex justify-center items-center">
                        <img src="../assets/images/Science-cuate.svg" alt="" class="w-11/12 h-3/3 border ">
                    </div>
                    <div class="w-6/12 flex flex-col justify-center items-center border h-auto  p-4">
                        <h2 class="text-4xl  h-auto mb-4">Sobre a Aplicação</h2>
                        <p class="h-auto ">Nossa aplicação foi desenvolvida para auxiliar a Eurofarma a navegar pelo vasto universo de pesquisas científicas que já foram realizadas. Com a crescente quantidade de dados disponíveis, a necessidade de uma ferramenta eficiente para acessar, analisar e compreender essas informações se torna essencial. Nossa plataforma foi projetada para oferecer essa eficiência, proporcionando uma visão clara e detalhada dos resultados obtidos em estudos anteriores, permitindo que a Eurofarma tome decisões informadas e embasadas.</p>
                    </div>
                </div>
                <div class="ml-20 w-10/12 h-auto border flex flex-row">
                    <div class="w-6/12 flex flex-col justify-center items-center border h-auto  p-4">
                        <h2 class="text-4xl  h-auto mb-4">Inteligência Artificial de Ponta
                        </h2>
                        <p class="h-auto "> 
                            Utilizando a poderosa IA do Gemini, nossa aplicação vai além da simples consulta de dados. A IA é capaz de analisar os resultados das pesquisas, identificar padrões e insights valiosos que poderiam passar despercebidos em uma análise manual. Isso não apenas acelera o processo de pesquisa, mas também garante que a Eurofarma esteja sempre à frente, utilizando as descobertas mais recentes e relevantes em seu campo de atuação.
                        </p>
                    </div>
                    <div class="w-6/12 flex justify-center items-center">
                        <img src="../assets/images/Remedy-rafiki.svg" alt="" class="w-11/12 h-3/3 border ">
                    </div>
                   
                </div>
                <div class="ml-20 w-10/12 h-auto border  flex flex-row">
                     <div class="w-6/12 flex justify-center items-center">
                        <img src="../assets/images/Business Plan-amico.svg" alt="" class="w-11/12 h-3/3 border ">
                    </div>
                    <div class="w-6/12 flex flex-col justify-center items-center border h-auto  p-4">
                        <h2 class="text-4xl  h-auto mb-4">Visualização de Dados Avançada
                        </h2>
                        <p class="h-auto "> 
                            Compreender os resultados das pesquisas é fundamental para a inovação. Por isso, nossa aplicação oferece gráficos detalhados que ilustram de forma clara e acessível os dados das pesquisas anteriores. Esses gráficos são interativos e permitem uma exploração profunda dos resultados, facilitando a identificação de tendências e conclusões importantes. Seja para comparar estudos, visualizar a progressão de resultados ao longo do tempo ou analisar correlações entre diferentes variáveis, nossa ferramenta de visualização é um recurso indispensável.
                        </p>
                    </div>
                   
                   
                </div>
                <div class="ml-20 w-10/12 h-auto border flex flex-row">
                  
                    <div class="w-6/12 flex flex-col justify-center items-center border h-auto  p-4">
                        <h2 class="text-4xl  h-auto mb-4">Ferramenta Estratégica para o Futuro
                        </h2>
                        <p class="h-auto "> 
Com a combinação de pesquisa avançada, inteligência artificial e visualização de dados, nossa aplicação se torna uma ferramenta estratégica para a Eurofarma. Ela não apenas facilita a pesquisa sobre estudos anteriores, mas também contribui diretamente para o desenvolvimento de novas estratégias de pesquisa e inovação. Ao utilizar nossa aplicação, a Eurofarma garante que cada decisão seja informada, precisa e fundamentada nos melhores dados disponíveis.                        </p>
                    </div>
                       <div class="w-6/12 flex justify-center items-center">
                        <img src="../assets/images/Science-cuate.svg" alt="" class="w-11/12 h-3/3 border ">
                    </div>
                   
                   
                </div>
                
            </div>
            

        </MainLayout>
</template>

<style lang="scss"></style>