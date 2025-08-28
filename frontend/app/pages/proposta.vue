<template>
  <div class="container">
    <div 
      class="box" 
      v-for="(item, index) in boxes" 
      :key="index"
      :style="{ backgroundImage: `url(${item.img})` }"
    >
      <div class="overlay">
        <p>{{ item.text }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
const boxes = [
  {
    img: 'https://definicion.de/wp-content/uploads/2010/08/cuidado-1.jpg',
    text: 'Atenção Humana.',
  },
  {
    img: 'https://pn7.com.br/wp-content/uploads/2022/04/Tecnologia-de-ponta-aplicada-a-cura-de-doencas-deve-crescer-ainda-mais-neste-ano.jpg',
    text: 'Tecnologia de Ponta.',
  },
]
</script>

<style scoped>
.container {
  display: flex;
  width: 100%;
  height: 500px; /* altura de teste */
}

.box {
  flex: 1;
  background-size: cover;
  background-position: center;
  position: relative;
  transition: transform 0.3s ease;
  cursor: pointer;
  overflow: hidden;
  background-color: #ccc; /* backup se imagem falhar */
}

.box:hover {
  transform: scale(1.05);
  z-index: 10;
}

.overlay {
  position: absolute;
  bottom: 20px;
  left: 20px;
  color: white;
  font-size: 1.5rem;
  text-shadow: 0 0 5px black;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.box:hover .overlay {
  opacity: 1;
}
</style>