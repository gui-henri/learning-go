<script setup lang="ts">
  import { Button } from '@/components/ui/button'
  import { Input } from '@/components/ui/input'
  import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card';
  import { ref } from "vue"
  import { Combobox, ComboboxAnchor, ComboboxEmpty, ComboboxGroup, ComboboxInput, ComboboxItem, ComboboxList, ComboboxTrigger } from "@/components/ui/combobox"
  import {
    InputGroup,
    InputGroupAddon,
    InputGroupButton,
    InputGroupInput,
  } from '@/components/ui/input-group'
import InputGroupText from '@/components/ui/input-group/InputGroupText.vue';

  const sexo = [
    { value: "m", label: "Masculino" },
    { value: "f", label: "Feminino" },
  ]

  const value = ref<typeof sexo[0]>()

</script>

<template>
  <div class="flex items-center justify-center flex-col p-6">
    <Card class="w-7xl">
      <CardHeader>
        <CardTitle class="text-2xl">Dados gerais</CardTitle>
        <CardDescription>Informações básicas do paciente.</CardDescription>
      </CardHeader>
      <CardContent>
        <span class="text-sm font-semibold">Nome do Paciente</span>
        <Input placeholder="Ex: João Almeida Santos" />

        <div class="flex gap-2.5">
          <div>
            <span class="text-sm font-semibold">Sexo</span>
            <Combobox v-model="value" by="label">
              <ComboboxAnchor as-child>
                <ComboboxTrigger as-child>
                  <Button variant="outline" :class="value?.label ? 'justify-between text-center' : 'justify-between text-center text-neutral-600'">
                    {{ value?.label ?? 'Selecione o sexo...' }}
                  </Button>
                </ComboboxTrigger>
              </ComboboxAnchor>

              <ComboboxList>
                <ComboboxGroup>
                  <ComboboxItem
                    class="hover:bg-gray-200 transition duration-200"
                    v-for="s in sexo"
                    :key="s.value"
                    :value="s"
                  >
                    {{ s.label }}
                  </ComboboxItem>
                </ComboboxGroup>
              </ComboboxList>
            </Combobox>
          </div>
          <div>
            <span class="text-sm font-semibold">Peso</span>
            <InputGroup class="w-20">
              <InputGroupInput placeholder="67" maxlength="3"/>
              <InputGroupAddon>
                <SearchIcon />
              </InputGroupAddon>
              <InputGroupAddon align="inline-end">
                <InputGroupText>kg</InputGroupText>
              </InputGroupAddon>
            </InputGroup>
          </div>
          <div>
            <span class="text-sm font-semibold">Altura</span>
            <InputGroup class="w-24">
              <InputGroupInput placeholder="161" maxlength="3"/>
              <InputGroupAddon>
                <SearchIcon />
              </InputGroupAddon>
              <InputGroupAddon align="inline-end">
                <InputGroupText>cm</InputGroupText>
              </InputGroupAddon>
            </InputGroup>
          </div>
        </div>
      </CardContent>
    </Card>

  </div>
</template>

<style scoped></style>
