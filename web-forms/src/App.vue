<script setup lang="ts">
  import { Button } from '@/components/ui/button'
  import { Input } from '@/components/ui/input'
import { Card } from './components/ui/card';
import { Switch } from './components/ui/switch';
import { ref } from "vue"
import { Combobox, ComboboxAnchor, ComboboxEmpty, ComboboxGroup, ComboboxInput, ComboboxItem, ComboboxItemIndicator, ComboboxList, ComboboxTrigger } from "@/components/ui/combobox"

const frameworks = [
  { value: "next.js", label: "Next.js" },
  { value: "sveltekit", label: "SvelteKit" },
  { value: "nuxt", label: "Nuxt" },
  { value: "remix", label: "Remix" },
  { value: "astro", label: "Astro" },
]

const value = ref<typeof frameworks[0]>()

</script>

<template>
  <Button>You did it!</Button>
  <p class="text-red-500">
    Visit <a href="https://vuejs.org/" target="_blank" rel="noopener">vuejs.org</a> to read the
    documentation.
  </p>
  <Card>
    <h2>Card box</h2>
    <Input />
    <div class="flex items-center gap-2 ">
      <span>Should be free? </span><Switch />
    </div>
    <Combobox v-model="value" by="label">
    <ComboboxAnchor as-child>
      <ComboboxTrigger as-child>
        <Button variant="outline" class="justify-between">
          {{ value?.label ?? 'Escolha uma das opções' }}
        </Button>
      </ComboboxTrigger>
    </ComboboxAnchor>

    <ComboboxList>
      <div class="relative w-full max-w-sm items-center">
        <ComboboxInput class="focus-visible:ring-0 border-0 border-b rounded-none h-10" placeholder="Digite..." />
      </div>

      <ComboboxEmpty>
        No framework found.
      </ComboboxEmpty>

      <ComboboxGroup>
        <ComboboxItem
          v-for="framework in frameworks"
          :key="framework.value"
          :value="framework"
        >
          {{ framework.label }}
        </ComboboxItem>
      </ComboboxGroup>
    </ComboboxList>
  </Combobox>
  </Card>
</template>

<style scoped></style>
