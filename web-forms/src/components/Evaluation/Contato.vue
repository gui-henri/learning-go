<script setup>
import { ref, computed } from 'vue';
import { useContatoStore } from '@/store/evaluation/contato';

const contatoStore = useContatoStore();

const forma_contato = ref([
    { name: 'Residencial', code: 'residencial' },
    { name: 'Celular', code: 'celular' },
    { name: 'Whatsapp', code: 'whatsapp' },
    { name: 'Celular e Whatsapp', code: 'celular_whatsapp' }
]);

const emit = defineEmits(['next-step']);
const activeIndex = ref(null);

const isFilled = computed(() => {
    return !!contatoStore.contato.telefone_paciente && 
           contatoStore.contato.telefone_paciente.length > 3;
});

const handleSave = () => {
    activeIndex.value = null;
    
    setTimeout(() => {
        const self = document.getElementById("contato");
        if (self) {
            self.scrollIntoView({ 
                behavior: 'instant', 
                block: 'start',
            });
        }
    }, 0);
    emit('next-step');
};
</script>

<template>
    <Accordion v-model:activeIndex="activeIndex" id="contato" class="scroll-mt-24 card shadow-2xl rounded-2xl w-full p-4 sm:p-8 border-t-8 border-red-600">
        <AccordionTab>
        <template #header>
            <div class="flex items-center gap-3 w-full">
                <i class="pi text-xl" 
                    :class="isFilled ? 'pi-check-circle text-green-600' : 'pi-plus-circle text-gray-400'">
                </i>
                <div class="flex flex-col text-left">
                    <h4 class="font-semibold text-xl p-0 m-0">Contato</h4>
                    <span class="text-xs text-gray-500 font-normal -mt-4">
                        {{ isFilled ? '' : 'Toque para preencher' }}
                    </span>
                </div>
            </div>
        </template>
        <div class="flex flex-col gap-4 w-full">
            <p class="text-red-600">TODO: adicionar mais de um responsável</p>            
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="telefone_residencial">Telefone residencial</label>
                    <InputText 
                        id="telefone_residencial" 
                        v-model="contatoStore.contato.telefone_residencial" 
                        placeholder="(81) 9 9999-9999" 
                        class="w-full"
                    />
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/6">
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
                        placeholder="Selecione" 
                        class="w-full"
                    ></Select>
                </div>
                <div class="flex flex-col gap-2 w-full md:w-1/6">
                    <label for="numero_prioritario">Contato prioritário?</label>
                    <InputSwitch 
                        id="numero_prioritario" 
                        v-model="contatoStore.contato.numero_prioritario" 
                    />
                </div>
            </div>
        </div>
    <Button class="mt-3" v-on:click="handleSave">
        <i class="pi text-xl" :class="'pi-check-circle text-white dark:text-black'" />
        Próximo
    </Button>
    </AccordionTab>
</Accordion>
</template>