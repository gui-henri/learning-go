<script setup lang="ts">
import {
  Carousel,
  CarouselContent,
  CarouselItem,
} from '@/components/ui/carousel'
</script>

<template>
  <div class="h-screen flex flex-col">
    <!-- Metade de cima -->
    <div class="flex-1 flex items-center justify-center">
      <h1 class="text-3xl font-bold">Equipe que torna tudo possível</h1>
    </div>

    <!-- Metade de baixo -->
    <div class="flex-1">
      <Carousel opts="{ align: 'start', loop: true }" class="w-full">
        <CarouselContent class="animate-scroll">
          <CarouselItem v-for="i in 10" :key="i" class="basis-1/5">
            <img
              class="w-32 h-32 rounded-full object-cover mx-auto"
              :src="`https://i.pravatar.cc/150?img=${i}`"
              alt="Membro"
            />
          </CarouselItem>
        </CarouselContent>
      </Carousel>
    </div>
  </div>
</template>

<style scoped>
@keyframes scroll {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}

.animate-scroll {
  display: flex;
  animation: scroll 25s linear infinite;
}
</style>
