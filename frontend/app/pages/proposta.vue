<script setup lang="ts">
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from '@/components/ui/carousel'
</script>

<template>
  <div class="flex items-start justify-center h-screen">
    <Carousel class="w-150 h-150">
      <CarouselContent>
        
        <!-- Slide 1 -->
        <CarouselItem>    
          <div class="flex flex-col items-center">
            <img 
              src="https://cdn.bitrix24.com.br/b22035513/landing/750/750d952c46e82fa4ef81b1103248c565/Atendimento_Domiciliar_-_HomeCare_1x.png" 
              class="w-full h-auto rounded-lg"
            > 
            <p class="mt-2 text-center text-red-700 text-sm">
              Atendimento 24H para nossos pacientes.
            </p>
          </div>
        </CarouselItem>

        <!-- Slide 2 -->
        <CarouselItem>
          <div class="flex flex-col items-center">
            <img 
              src="https://telemedicinamorsch.com.br/wp-content/uploads/2017/02/iStock-592647828.jpg" 
              class="w-full h-auto rounded-lg"
            >
            <p class="mt-2 text-center text-blue-700 text-sm">
              Tecnologia em medicina.
            </p>
          </div>
        </CarouselItem>

        <!-- Slide 3 -->
        <CarouselItem>
          <div class="flex flex-col items-center">
            <img 
              src="https://cdn.prod.website-files.com/647e410aa35e6ce70fed96cd/67449063bc02321b87ce595a_bom-atendimento-medico.jpeg" 
              class="w-full h-auto rounded-lg"
            >
            <p class="mt-2 text-center text-green-700 text-sm">
              Atendimento humanizado com qualidade
            </p>
          </div>
        </CarouselItem>

      </CarouselContent> 

      <CarouselPrevious />
      <CarouselNext />
    </Carousel>
  </div>
</template>
