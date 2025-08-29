# 🧩 Interne Tarefas

Interne Tarefas é uma aplicação web full-stack usando Nuxt para o frontend e Go no backend. Tudo dentro de containers do Docker. Tarefas podem ser criadas, listadas e concluídas, bem como há um histórico de todas as tarefas. Atualmente, utiliza um banco de dados em memória.

---

## 📁 Estrutura

├── backend/                  # Projeto em Go
│ └── main.go
├── frontend/                 # App Nuxt
| └── app
│ |  ├── pages/
| |  ├── components/
| |  └── images/
| ├── package.json
│ └── nuxt.config.ts
├── docker-compose.yml    |
├── Dockerfile.backend    ├── # Arquivos do Docker
├── Dockerfile.frontend   |
└── readme.md

## 🛠️ Setup

1. Baixe este projeto
```bash
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
```

2. Crie um arquivo ```.env``` com o seguinte conteúdo
```
API_BASE=http://backend:8090
NUXT_PUBLIC_API_BASE=http://localhost:8090
```

3. Rode usando o Docker

```bash
docker-compose up --build
```

Ao finalizar, ambas as aplicações já estarão rodando:
* Backend Go em: ```http://localhost:8090```
* Frontend Nuxt em: ```http://localhost:8091```


## 👥 Autores

* Guilherme Henrique Freitas da Silva
* Samuel Vitor Melo do Nascimento 