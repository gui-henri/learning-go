<!-- frontend/app/app.vue -->
<template>
  <div>
    <div class="logo">
      <a href="https://interne.com.br" target="_blank" rel="noopener noreferrer">
        <img
          src="https://cdn.bitrix24.com.br/b22035513/landing/2a4/2a44b7c2b4e7e248ec783a2eae611962/Logo_Interne_180_x_60_px_1x.png"
          alt="Logo Interne"
        />
      </a>
    </div>

    <nav class="barra-nav">
      <NuxtLink to="/" class="nav-link">Página inicial</NuxtLink>
      <NuxtLink to="/equipe" class="nav-link">Equipe</NuxtLink>
      <NuxtLink to="/proposta" class="nav-link">Nossa proposta</NuxtLink>
      <NuxtLink to="/tarefas" class="nav-link">Visualizar tarefas</NuxtLink>
    </nav>

    <!-- Aqui é onde será renderizado o conteúdo da rota -->
    <NuxtPage />
  </div>
</template>

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
