<script setup lang="ts">
import { Bandage, Briefcase } from 'lucide-vue-next'; // Mudei para um ícone que representa melhor "tarefas"
</script>

<template>
    <div class="p-4 flex items-center justify-center gap-6">
        <NuxtLink to="/proposta">
            <img src="/images/Imagem colada2.png" alt="Logo Proposta" class="h-12 cursor-pointer hover:opacity-80 transition-opacity">
        </NuxtLink>

        <a href="https://interne.com.br" target="_blank" rel="noopener noreferrer">
            <img
                src="https://cdn.bitrix24.com.br/b22035513/landing/2a4/2a44b7c2b4e7e248ec783a2eae611962/Logo_Interne_180_x_60_px_1x.png"
                alt="Logo Interne"
                class="h-12 cursor-pointer hover:opacity-80 transition-opacity"
            />
        </a>

        <a href="https://www.google.com/maps/place/Interne+Solu%C3%A7%C3%B5es+em+Sa%C3%BAde/@-8.0632204,-34.9326301,13z/data=!4m6!3m5!1s0x7ab18c5bf17521d:0xdae5f9e47e64b2fc!8m2!3d-8.0633074!4d-34.8914301!16s%2Fg%2F1yfdr089g?entry=ttu" target="_blank" rel="noopener noreferrer" >
            <img
                src="/images/Imagem colada.png"
                alt="Logo Localização"
                class="h-12 cursor-pointer hover:opacity-80 transition-opacity"
            />
        </a>
    </div>
 
   
    <nav class="flex bg-red-600 h-12 w-full items-center justify-evenly mb-5">
        <NuxtLink to="/" class="flex flex-1 items-center justify-center gap-2 h-full text-white font-semibold transition-colors hover:bg-red-700">
            <Briefcase class="h-5 w-5" />
            <span>Tarefas</span>
        </NuxtLink>

        <NuxtLink to="/pacientes" class="flex flex-1 items-center justify-center gap-2 h-full text-white font-semibold transition-colors hover:bg-red-700">
            <Bandage class="h-5 w-5" />
            <span>Pacientes</span>
        </NuxtLink>
    </nav>


    <main class="container mx-auto px-4">
        <NuxtPage />
    </main>

    <div class="h-12"></div>
</template>

<style>

  body {
    margin: 0;
    padding: 0;
    font-family: sans-serif;
    background-color: #f9fafb; 
  }

  
  .router-link-exact-active {
    background-color: #b91c1c; 
    font-weight: bold;
  }
</style>
