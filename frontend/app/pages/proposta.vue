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
  <Carousel class="w-64 ml-24">
    <CarouselContent>
      <CarouselItem>wait for it</CarouselItem>
      <CarouselItem>wait for them</CarouselItem>
      <CarouselItem>wait for us</CarouselItem>
    </CarouselContent>
    <CarouselPrevious/>
    <CarouselNext />
  </Carousel>
</template>