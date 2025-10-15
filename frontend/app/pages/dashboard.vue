<template>
  <div class="container">
    <h1>📝 Planejamento e Lista de Tarefas do Projeto</h1>
    <p class="subtitle">
      Um guia com as próximas funcionalidades, correções e ideias para a aplicação.
      <br>
      Última atualização: {{ new Date().toLocaleDateString('pt-BR') }}
    </p>

    <div v-for="(task, index) in projectTasks" :key="index" class="task-card" :class="`border-${task.category.toLowerCase()}`">
      <span class="category-badge" :class="`bg-${task.category.toLowerCase()}`">
        {{ task.category }}
      </span>
      
      <h2>{{ task.title }}</h2>
      
      <p class="details">{{ task.details }}</p>

      <ul v-if="task.items && task.items.length > 0">
        <li v-for="item in task.items" :key="item">{{ item }}</li>
      </ul>
      
      <div class="status">
        <strong>Status:</strong> <span>{{ task.status }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

/**
 * @description
 * Array contendo todas as tarefas, ideias e planos para o projeto.
 * Cada objeto representa um item de trabalho.
 * - category: 'Bug', 'Feature', 'Dashboard', 'Ideia'
 * - title: Um título curto e descritivo.
 * - details: Uma explicação mais completa do que precisa ser feito.
 * - items: (Opcional) Uma lista de sub-itens ou requisitos.
 * - status: 'Pendente', 'Em Andamento', 'Concluído'
 */
const projectTasks = ref([
  {
    category: 'Bug',
    title: 'Validação da Data de Nascimento',
    details: 'O sistema permite que a data de nascimento seja uma data futura. É preciso corrigir a validação no formulário de cadastro para bloquear datas posteriores ao dia de hoje.',
    status: 'Pendente',
  },
  {
    category: 'Feature',
    title: 'Módulo de Gerenciamento de Técnicos',
    details: 'Implementar uma nova área no sistema para gerenciar "Técnicos", que serão perfis de usuários com permissões específicas.',
    items: [
      'Criar tela de Cadastro de Técnicos.',
      'Criar tela de Listagem e Edição.',
      'Definir permissões de acesso para este perfil.'
    ],
    status: 'Pendente',
  },
  {
    category: 'Feature',
    title: 'Sistema de Gerenciamento de Tarefas',
    details: 'Desenvolver um módulo completo para que os usuários possam criar e atribuir tarefas, acompanhando seu ciclo de vida.',
    items: [
      'Definir o Custo de uma tarefa.',
      'Associar um Profissional (Técnico) responsável.',
      'Categorizar por Setor (ex: Administrativo, Clínico).',
      'Definir um nível de Prioridade (Baixa, Média, Alta).'
    ],
    status: 'Pendente',
  },
  {
    category: 'Dashboard',
    title: 'Desenvolvimento do Dashboard de Pacientes',
    details: 'Criar painéis visuais para fornecer insights sobre os dados dos pacientes. O objetivo é responder rapidamente às seguintes perguntas-chave (KPIs):',
    items: [
      'Quantos pacientes novos foram cadastrados esta semana?',
      'Qual é a média de idade dos nossos pacientes?',
      'Temos mais pacientes homens ou mulheres?',
      'Qual o perfil dos pacientes por localidade? (Requer coleta de dados de endereço)'
    ],
    status: 'Em Andamento',
  },
  {
    category: 'Ideia',
    title: 'Brainstorming de Novas Funcionalidades',
    details: 'Reservar um tempo para pensar em novas funcionalidades que possam agregar mais valor ao sistema e aos usuários.',
    items: [],
    status: 'Contínuo',
  }
]);
</script>

<style scoped>
.container {
  font-family: sans-serif;
  padding: 2rem;
  background-color: #f4f4f9;
  max-width: 900px;
  margin: 2rem auto;
  border-radius: 8px;
}

h1 {
  color: #2c3e50;
  border-bottom: 2px solid #e0e0e0;
  padding-bottom: 0.5rem;
}

.subtitle {
  color: #666;
  margin-bottom: 2rem;
}

.task-card {
  background-color: #fff;
  border-radius: 8px;
  padding: 1.5rem;
  margin-bottom: 1.5rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  border-left: 5px solid;
  position: relative;
}

.details {
  color: #333;
  line-height: 1.6;
}

ul {
  margin-top: 1rem;
  padding-left: 1.5rem;
  color: #555;
}

li {
  margin-bottom: 0.5rem;
}

.status {
  margin-top: 1rem;
  font-size: 0.9rem;
  color: #666;
}

.status span {
  font-weight: bold;
  color: #333;
}

.category-badge {
  position: absolute;
  top: 1rem;
  right: 1rem;
  padding: 0.3rem 0.8rem;
  border-radius: 12px;
  font-size: 0.8rem;
  font-weight: bold;
  color: #fff;
}

/* Cores por Categoria */
.border-bug { border-color: #ef4444; }
.bg-bug { background-color: #ef4444; }

.border-feature { border-color: #3b82f6; }
.bg-feature { background-color: #3b82f6; }

.border-dashboard { border-color: #10b981; }
.bg-dashboard { background-color: #10b981; }

.border-ideia { border-color: #f97316; }
.bg-ideia { background-color: #f97316; }
</style>