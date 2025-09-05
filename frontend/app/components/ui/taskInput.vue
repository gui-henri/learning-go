<template>
    <div class="items-center justify-center gap-2">
        <Input class="p-2 rounded-full border-2 border-red-500" type="text" name="description" placeholder="Digite a tarefa" />
        <Button class="bg-red-600 hover:bg-red-500 p-5 rounded-3xl">Enviar</Button>
    </div>
</template>

<style lang="css" scoped>
    
    div {
        display: flex;
        width: 100%;
        max-width: 520px;
    }
    input {
        width: 100%;
    }

</style>