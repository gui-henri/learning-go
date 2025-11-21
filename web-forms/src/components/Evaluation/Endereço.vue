<script setup>
import { useEnderecoStore } from '@/store/evaluation/endereco';

const enderecoStore = useEnderecoStore();

</script>

<template>

<div class="card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
            <div class="flex flex-col gap-4 w-full">
                
                <h4 class="font-semibold text-xl">Endereço</h4>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="cep">CEP</label>
                        <InputText 
                            id="cep" 
                            v-model="enderecoStore.endereco.cep" 
                            placeholder="55050-550" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="estado">Estado</label>
                        <InputText 
                            id="estado" 
                            v-model="enderecoStore.endereco.estado" 
                            placeholder="Pernambuco" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="cidade">Cidade</label>
                        <InputText 
                            id="cidade" 
                            v-model="enderecoStore.endereco.cidade" 
                            placeholder="Recife" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="bairro">Bairro</label>
                        <InputText 
                            id="bairro" 
                            v-model="enderecoStore.endereco.bairro" 
                            placeholder="Boa vista" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="logradoro">Logradoro</label>
                        <InputText 
                            id="logradoro" 
                            v-model="enderecoStore.endereco.logradoro" 
                            placeholder="Agamenon magalhães" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/6">
                        <label for="numero-logradoro">Número</label>
                        <InputText 
                            id="numero-logradoro" 
                            v-model="enderecoStore.endereco.numero" 
                            placeholder="23-A" 
                            class="w-full"
                        />
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-4">
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="complemento">Complemento</label>
                        <InputText 
                            id="complemento" 
                            v-model="enderecoStore.endereco.complemento" 
                            placeholder="Prédio azul" 
                            class="w-full"
                        />
                    </div>
                    <div class="flex flex-col gap-2 w-full md:w-1/4">
                        <label for="ponto_referencia">Ponto de referência</label>
                        <InputText 
                            id="ponto_referencia" 
                            v-model="enderecoStore.endereco.ponto_referencia" 
                            placeholder="perto do posto de saúde" 
                            class="w-full"
                        />
                    </div>
                </div>
            </div>
        </div>
</template>