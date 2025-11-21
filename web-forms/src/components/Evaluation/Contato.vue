<script setup>

import { useContatoStore } from '@/store/evaluation/contato';

const contatoStore = useContatoStore();

</script>

<template>

<div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Contato</h4>
                <p class="text-red-600">TODO: adicionar mais de um responsável</p>
                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="telefone_residencial">Telefone residencial</label>
                        <InputText 
                            id="telefone_residencial" 
                            v-model="contatoStore.contato.telefone_residencial" 
                            placeholder="(81) 9 9999-9999" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/8">
                        <label for="telefone_paciente">Telefone paciente</label>
                        <InputText 
                            id="telefone_paciente" 
                            v-model="contatoStore.contato.telefone_paciente" 
                            placeholder="(81) 9 9999-9999" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="email_paciente">E-mail paciente</label>
                        <InputText 
                            id="email_paciente" 
                            v-model="contatoStore.contato.email_paciente" 
                            placeholder="exemplo@dominio.com" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/2">
                        <label for="nome_responsavel">Nome do responsável</label>
                        <InputText 
                            id="nome_responsavel" 
                            v-model="contatoStore.contato.nome_responsavel" 
                            placeholder="João albério" 
                            class="w-full"
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="parentesco_responsavel">Parentesco do responsável</label>
                        <InputText 
                            id="parentesco_responsavel" 
                            v-model="contatoStore.contato.parentesco_responsavel" 
                            placeholder="Pai, mãe, irmão etc" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="email_responsavel">E-mail do responsável</label>
                        <InputText 
                            id="email_responsavel" 
                            v-model="contatoStore.contato.email_responsavel" 
                            placeholder="responsavel@gmail.com" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="telefone_responsavel">Telefone do responsável</label>
                        <InputText 
                            id="telefone_responsavel" 
                            v-model="contatoStore.contato.telefone_responsavel" 
                            placeholder="(81) 9 9999-9999" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 md:w-1/6">
                        <label for="forma_contato">Forma de contato</label>
                        <Select 
                            id="forma_contato" 
                            v-model="contatoStore.contato.forma_contato" 
                            :options="forma_contato" 
                            optionLabel="name" 
                            placeholder="Selecione forma de contato" 
                            class="w-full"
                        ></Select>
                    </div>
                    <div class="flex flex-col gap-2 py-2 w-">
                        <label for="previsao_alta">Número prioritário</label>
                        <InputSwitch 
                            id="previsao_alta" 
                            v-model="contatoStore.contato.numero_prioritario" 
                        />
                    </div>
                </div>
            </div>
        </div>

</template>