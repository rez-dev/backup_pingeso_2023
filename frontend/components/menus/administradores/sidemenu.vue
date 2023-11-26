<template>
  <div>
    <nav class="sidenav-b">
      <div class="contendor-logo-b">
        <a href="/administrador" class="logo-b"></a>
        <c-button class="boton-abrir-b" @click="isOpen = true">
          <div>
            <i class="fa fa-chevron-right"></i>
          </div>
        </c-button>
      </div>
      <c-divider />
      <div class="opciones-b">
        <div class="icono-inicio-b">
          <div class="icono-b">
            <i class="fa fa-home icono-grande-side"></i>
          </div>
          <a class="lol-b" href="/administrador"
            ><span style="font-size: x-large; color: white; font-weight: 700"
              >Inicio</span
            ></a
          >
        </div>
        <div class="opcion-b">
          <c-button
            class="boton-opcion-b"
            mb="4"
            variant-color="white"
            @click="showServicios = !showServicios"
          >
            <div class="icono-b">
              <i class="fa fa-book icono-grande-side"></i>
            </div>
            <div class="titulo-opcion-b">Servicios</div>
          </c-button>
          <c-collapse id="colapse-b" :is-open="showServicios">
            <div class="opcion-colapse-b">
              <a href="/administrador/revisar">Revisar servicios</a>
            </div>
            <div class="opcion-colapse-b">
              <a href="/administrador/asignar">Asignar servicios</a>
            </div>
          </c-collapse>
        </div>
        <div class="opcion-b">
          <c-button
            class="boton-opcion-b"
            mb="4"
            variant-color="white"
            @click="showRevisores = !showRevisores"
          >
            <div class="icono-b">
              <i class="fa fa-user icono-grande-side"></i>
            </div>
            <div class="titulo-opcion-b">Revisores</div>
          </c-button>
          <c-collapse id="colapse-b" :is-open="showRevisores">
            <div class="opcion-colapse-b">
              <a href="/administrador/registrar">Regristrar revisor</a>
            </div>
            <div class="opcion-colapse">
              <a href="/administrador/revisores">Ver revisores</a>
            </div>
          </c-collapse>
        </div>
        <div class="opcion-b">
          <c-button
            class="boton-opcion-b"
            mb="4"
            variant-color="white"
            @click="showDashboard = !showDashboard"
          >
            <div class="icono-b">
              <i class="fa fa-bar-chart icono-grande-side"></i>
            </div>
            <div class="titulo-opcion-b">Dashboard</div>
          </c-button>
          <c-collapse id="colapse-b" :is-open="showDashboard">
            <div class="opcion-colapse-b">
              <a href="/administrador/registrar">Proximamente...</a>
            </div>
            <div class="opcion-colapse-b">
              <a href="/administrador/asignar">Proximamente...</a>
            </div>
          </c-collapse>
        </div>
      </div>
      <div class="cerrar-sesion-b">
        <c-button
          style="background: none; font-size: x-large"
          @click="cerrarSesion"
        >
          <div class="icono-b">
            <i class="fa fa-power-off icono-medio-side"></i>
          </div>
          <div class="titulo-opcion-b">Cerrar sesión</div>
        </c-button>
      </div>
    </nav>

    <c-drawer
      :isOpen="isOpen"
      placement="left"
      :on-close="closeDrawer"
      :initialFocusRef="() => $refs.inputInsideModal"
    >
      <c-drawer-close-button />
      <sidebarprueba :cerrarDrawer="close" />
    </c-drawer>
  </div>
</template>
<script>
import "font-awesome/css/font-awesome.css";
import sidebar from "../administradores/sidebar.vue"
export default {
  components: {
    sidebar,
  },
  data() {
    return {
      isOpen: false,
      showServicios: false,
      showRevisores: false,
      showDashboard: false,
      sidebarVisible: false,
    };
  },
  methods: {
    cerrarSesion() {
      localStorage.clear();
      this.$router.push("/");
    },
    toggleSidebar() {
      this.sidebarVisible = !this.sidebarVisible;
    },
    close() {
      this.isOpen = false;
    },
    closeDrawer() {
      this.isOpen = false;
    },
  },
};
</script>
<style>
.sidenav-b {
  position: fixed;
  width: 70px;
  height: 100%;
  background-color: #00a490;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}

.opciones-b {
  flex: 1; /* Esto hace que las opciones ocupen el espacio restante */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  margin-top: 20%;
}
.opcion-b {
  width: 100%;
  align-items: center;
  display: flex;
  flex-direction: column;
  margin-bottom: 8%;
}
.boton-opcion-b {
  background: none;
  width: 100%;
  align-items: center;
  justify-content: start;
  font-size: x-large;
  border: none;
}
.boton-opcion-b:focus {
  outline: none;
  outline-offset: no-repeat;
  box-shadow: none;
}
#colapse-b {
  width: 17.5rem;
}
.opcion-colapse-b {
  font-size: large;
  color: white;
  margin-bottom: 10%;
  padding-left: 20%;
}

.opcion-colapse-b:hover {
  background-color: #394049;
}

.contendor-logo-b {
  height: 5%;
  /* align-items: center; */
  justify-content: center;
  font-size: 20px;
  display: flex;
}
.titulo-opcion-b {
  display: none;
}
.logo-b {
  display: none;
}
.icono-b {
  margin-right: 10px; /* Ajusta el margen a tu preferencia */
  color: white;
}

.icono-b i.icono-grande-side {
  font-size: 25px;
  color: white; /* Ajusta el tamaño del icono según tus preferencias */
}

.icono-b i.icono-medio-side {
  font-size: 24px;
  color: black; /* Ajusta el tamaño del icono según tus preferencias */
}

.lol-b {
  display: none; /* Oculta el texto dentro de los botones */
}
.icono-inicio-b {
  display: flex;
  flex-direction: row;
  margin-bottom: 12%;
  margin-right: 0%;
}
.opcion-b.active {
  background-color: #394049;
}
.cerrar-sesion-b {
  align-self: center;
  margin-bottom: 10%;
  font-weight: 500;
  font-size: larger;
  color: white;
  margin-right: 10%;
  background: none;
}
.boton-abrir-b {
  width: 50%;
  /* align-self: center; */
  /* justify-self: center; */
  background: none;
  margin-top: 15%;
}
.boton-abrir-b:hover {
  background: none;
}
.boton-abrir-b:active {
  background: none;
  border: none;
}
</style>
