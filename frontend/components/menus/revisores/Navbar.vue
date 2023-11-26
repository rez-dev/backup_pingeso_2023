<template>
  <div class="lol">
    <a href="/servicios/lista" class="logo"></a>
    <div class="contenedor-menu">
      <Menupeq class="containerMenupe" />
      <c-menu v-slot="{ isOpen }">
        <c-menu-button
          :is-active="isOpen"
          :right-icon="isOpen ? 'chevron-up' : 'chevron-down'"
          class="aaa"
        >
          {{ isOpen ? "Menu" : "Menu" }}
        </c-menu-button>
        <c-menu-list>
          <c-menu-item @click="RedirigirLista" class="menu-item">
            Servicios
          </c-menu-item>
          <c-menu-divider />
          <c-menu-item @click="cerrarSesion" class="menu-item">
            Cerrar sesión
            <i class="fa fa-sign-out icono-grande-side" aria-hidden="true"></i>
          </c-menu-item>
        </c-menu-list>
      </c-menu>
    </div>
  </div>
</template>

<script>
import Menupeq from "../revisores/Menupeq.vue";
import Sidebar from "../revisores/Sidebar.vue";
export default {
  components: {
    Menupeq,
    Sidebar,
  },
  methods:{
    RedirigirLista(){
      this.$router.push("/servicios/lista");
    },
    cerrarSesion(){
      localStorage.clear();
      this.$router.push("/");
    }
  }
};
</script>

<style scoped>
.lol {
  background-color: #00a499;
  width: 100%;
  height: 5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.contenedor {
  display: flex;
  align-items: end;
  justify-content: end;
  width: 100%;
  height: 35%;
}

.containerMenu {
  margin-right: 2%;
  border: white;
  border-radius: 5px;
  color: #394049;
  height: 100%;
}
.containerMenupe {
  display: none;
}
.icono-grande-side {
  font-size: 20px;
  margin-left: 5%;
  color: black;
}
.logo {
  background-image: url(../../../images/UsachSB.png);
  background-size: contain;
  background-repeat: no-repeat;
  width: 350px;
  height: 100%;
  background-position: 8px;
}
.containerMenu CButtonIcon {
  margin-left: 10%;
}
.menu-item:hover{
  background-color: #00a499;
}
.contenedor-menu {
  margin-right: 4%;
}
.aaa {
  color: black;
}

.aaa:hover{
  outline: 1px solid white; /* Ancho de 2 píxeles y color rojo para el contorno */
  outline-offset: 2px;
}
.aaa:active{
  outline: 1px solid white; /* Ancho de 2 píxeles y color rojo para el contorno */
  outline-offset: 2px;
}
@media (max-width: 688px) {
  .containerMenu {
    display: none; /* Muestra el menú de hamburguesa */
  }
  .containerMenupe {
    display: block;
    background-color: white;
    border: white;
    border-radius: 5px;
    color: #394049;
    margin-right: 2%;
    height: 100%;
  }
  .logo {
    background-image: url(../../../images/UsachSB.png);
    background-size: contain;
    background-repeat: no-repeat;
    width: 15rem;
    height: 100%;
    background-position: center;
  }
}
@media (max-width: 414px) {
  .containerMenu {
    display: none; /* Muestra el menú de hamburguesa */
  }
  .containerMenupe {
    display: block;
    border: white;
    border-radius: 5px;
    color: #394049;
    margin-right: 2%;
    height: 100%;
  }
  .logo {
    background-image: url(../../../images/UsachSB.png);
    background-size: contain;
    background-repeat: no-repeat;
    width: 10rem;
    height: 100%;
    background-position: center;
  }
  .contenedor {
    height: 30%;
  }
}
</style>
