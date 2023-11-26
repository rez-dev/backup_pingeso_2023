<template>
  <div>
    <c-button class="logo-menu" @click="isOpen = true"
      ><i class="fa fa-bars" aria-hidden="true"></i
    ></c-button>
    <c-drawer placement="right" :on-close="close" :isOpen="isOpen">
      <c-drawer-overlay />
      <c-drawer-content class="container-drawer">
        <c-drawer-header borderBottomWidth="1px"> Menu </c-drawer-header>
        <c-drawer-body>
          <router-link to="/servicios/lista" class="boton-drawer">
            <p>Servicios</p>
          </router-link>
          <c-menu-divider />
          <router-link to="/" class="boton-drawer" @click="cerrarSesion">
            <p>Cerrar sesión</p>
          </router-link>
        </c-drawer-body>
      </c-drawer-content>
    </c-drawer>
  </div>
</template>

<script>
import "font-awesome/css/font-awesome.css";
export default {
  data() {
    return {
      isOpen: false,
      placement: "right",
    };
  },
  methods: {
    setPlacement(value) {
      this.placement = value;
    },
    close() {
      this.isOpen = false;
    },
    cerrarSesion() {
      localStorage.clear();
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
.container-drawer {
  border: 2px solid white;
  border-radius: 2px;
}
.boton-drawer {
  width: 100%;
  text-decoration: none;
  color: black;
  display: block;
  transition: background-color 0.3s;
}

@media (max-width: 688px) {
  .container-drawer s {
    width: 20%;
  }
}
.logo-menu {
  background: none;
  font-size: 30px;
  color: black;
}
</style>
