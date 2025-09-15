<template>
  <div>
    <div class="logo p-4">
      <a class="flex items-center justify-center" href="https://interne.com.br" target="_blank" rel="noopener noreferrer">
        <img
          src="https://cdn.bitrix24.com.br/b22035513/landing/2a4/2a44b7c2b4e7e248ec783a2eae611962/Logo_Interne_180_x_60_px_1x.png"
          alt="Logo Interne"
        />
      </a>
    </div>

    <nav class="barra-nav">
      <NuxtLink to="/" class="nav-link"> <b>Abrir Chamado </b></NuxtLink>
      <NuxtLink to="/cadastro" class="nav-link"><b>Cadastro de Pacientes</b></NuxtLink>
      <NuxtLink to="/proposta" class="nav-link"><b>W.I.P</b></NuxtLink>
      <NuxtLink to="/tarefas" class="nav-link"><b>Visualizar tarefas</b></NuxtLink>
    </nav>

    <NuxtPage />
    <div class="h-12"></div>
  </div>
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
