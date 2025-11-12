<script setup lang="ts">
import type { HTMLAttributes } from "vue"
import { cn } from "@/lib/utils"

const props = defineProps<{
  class?: HTMLAttributes["class"]
}>()
</script>

<template>
  <div
    data-slot="card"
    :class="
      cn(
        'bg-card text-card-foreground flex flex-col gap-6 rounded-xl border py-6 shadow-sm',
        'p-6 border-t-4 border-t-red-600 max-w-3xl',
        props.class,
      )
    "
  >
    <slot />
  </div>
</template>
