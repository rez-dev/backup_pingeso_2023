<template>
  <div class="contenedor-tarjeta">
    <div class="tarjeta">
      <div class="icono-card">
        <i :class="'fa ' + icon + ' icono-grande-sides'"></i>
      </div>
      <div class="titulo-tarjeta">
        <div class="titulo">{{ title }}</div>
      </div>
      <div class="cuerpo-tarjeta">
        <div class="descripcion-tarjeta">{{ descripcion }}</div>
        <div class="boton-tarjeta">
          <c-button
            class="adsa"
            variant-color="blue"
            variant="outline"
            @click="redirigir"
          >
            {{ textBoton }}
          </c-button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import "font-awesome/css/font-awesome.css";
export default {
  components: {},
  props: {
    title: String,
    descripcion: String,
    visible: Boolean,
    icon: String,
    textBoton: String,
    ruta: String,
  },
  methods: {
    clicIzquierda() {
      this.$emit("clicIzquierda"); // Emitir un evento cuando se hace clic en el botón izquierdo
    },
    clicDerecha() {
      this.$emit("clicDerecha"); // Emitir un evento cuando se hace clic en el botón derecho
    },
    redirigir() {
      this.$router.push(this.ruta);
    },
  },
};
</script>

<style scoped>
.icono-card {
  color: beige;
  position: relative;
  left: 2rem;
  top: 2.4rem;
  background-color: #7ee8fa;
  width: 15%;
  align-items: center;
  display: flex;
  justify-content: center;
  height: 20%;
  border-radius: 50% 50% 50% 50%;
}
.icono-card i.icono-grande-sides {
  font-size: 40px;
  color: #394049;
}
.contenedor-tarjeta {
  height: 100%;
  display: flex;
  align-items: center;
  margin-left: 3%;
  width: 30%;
}
.tarjeta {
  height: 80%;
  display: flex;
  flex-direction: column;
  border-radius: 10px;
  width: 100%;
}
.tarjeta:hover {
  transform: scale(1.02, 1.02);
}

.titulo-tarjeta {
  width: 100%;
  height: 30%;
  background-color: white;
  border-radius: 10px 10px 0 0;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 5px 5px 4px rgba(0, 0, 0, 0.2);
}
.titulo {
  font-size: xx-large;
  color: #ea7600;
}
.cuerpo-tarjeta {
  width: 100%;
  height: 70%;
  background-color: #7ee8fa;
  border-radius: 0 0 10px 10px;
  box-shadow: 5px 5px 4px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  justify-content: space-around;
}
.botoness {
  height: 80%;
  display: flex;
  align-items: center;
  width: 5%;
  justify-content: center;
  margin-top: 10%;
}
.botoness:hover {
  transform: scale(1.15, 1.15);
}
#bot-sgte {
  background: none;
  font-size: 30px;
}
#bot-prev {
  background: none;
  font-size: 30px;
}
.descripcion-tarjeta {
  font-size: x-large;
  margin-left: 2%;
}
.boton-tarjeta {
  width: 30%;
  margin-left: 2%;
}

.adsa:hover {
  outline: 1px solid black; /* Ancho de 2 píxeles y color rojo para el contorno */
  outline-offset: 2px;
}
</style>
