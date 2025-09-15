<template>
    <div class="logo p-4 flex items-center justify-center gap-6">

 <NuxtLink to="/proposta">
  <img src="/frontend/app/images/Imagem colada2.png" alt="Logo esquerda" class="h-12 cursor-pointer hover:opacity-80">
</NuxtLink>

  <a href="https://interne.com.br" target="_blank" rel="noopener noreferrer">
    <img
      src="https://cdn.bitrix24.com.br/b22035513/landing/2a4/2a44b7c2b4e7e248ec783a2eae611962/Logo_Interne_180_x_60_px_1x.png"
      alt="Logo Interne"
      class="cursor-pointer hover:opacity-80"
    />
  </a>
  <a href="https://www.google.com/maps/place/Interne+Solu%C3%A7%C3%B5es+em+Sa%C3%BAde/@-8.0632204,-34.9326301,13z/data=!4m6!3m5!1s0x7ab18c5bf17521d:0xdae5f9e47e64b2fc!8m2!3d-8.0633074!4d-34.8914301!16s%2Fg%2F1yfdr089g?entry=ttu&g_ep=EgoyMDI1MDkxMC4wIKXMDSoASAFQAw%3D%3D" target="_blank" rel="noopener noreferrer" >
    <img
      src="/frontend/app/images/Imagem colada.png"
      alt="Logo direita"
      class="h-12 cursor-pointer hover:opacity-80"
    />
  </a>
</div>


 
    <nav class="barra-nav">
      <NuxtLink to="/" class="nav-link"> <b>Abrir Tarefa</b></NuxtLink>
      <NuxtLink to="/" class="nav-link"><b>Cadastro de Pacientes</b></NuxtLink>
      <NuxtLink to="/" class="nav-link"><b>W.I.P</b></NuxtLink>
      <NuxtLink to="/tarefas" class="nav-link"><b>Visualizar tarefas</b></NuxtLink>
    </nav>

    <NuxtPage />
    <div class="h-12"></div>
</template>

<style>
  body {
    margin: 0;
    padding: 0;
  }
</style>

<style scoped>

.logo {
  text-align: center;
  margin-bottom: 10px;
}

.barra-nav {
  display: flex;
  background-color: red;
  height: 40px;
  width: 100%;
  justify-content: space-evenly;
  align-items: center;
  margin-bottom: 20px;
}

.nav-link {
  flex: 1;
  text-align: center;
  color: white;
  text-decoration: none;
  padding: 10px 0;
  transition: background-color 0.3s;
}

.nav-link:hover {
  background-color: darkred;
}
</style>
