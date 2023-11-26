<template>
  <div>
    <nav class="sidenav">
      <div class="contendor-logo">
        <a href="/administrador" class="logo"></a>
      </div>
      <c-divider />
      <div class="opciones">
        <div class="icono-inicio">
          <div class="icono"><i class="fa fa-home icono-grande-side"></i></div>
          <a class="lol" href="/administrador"
            ><span style="font-size: x-large; color: white; font-weight: 700"
              >Inicio</span
            ></a
          >
        </div>
        <div class="opcion">
          <c-button
            id="boton-opcion-side"
            mb="4"
            variant-color="white"
            @click="showServicios = !showServicios"
          >
            <div class="icono">
              <i class="fa fa-book icono-grande-side"></i>
            </div>
            <div class="titulo-opcion">Servicios</div>
          </c-button>
          <c-collapse id="colapse" :is-open="showServicios">
            <div class="opcion-colapse">
              <a href="/administrador/revisar">Revisar servicios</a>
            </div>
            <div class="opcion-colapse">
              <a href="/administrador/asignar">Asignar servicios</a>
            </div>
          </c-collapse>
        </div>
        <div class="opcion">
          <c-button
            id="boton-opcion-side"
            mb="4"
            variant-color="white"
            @click="showRevisores = !showRevisores"
          >
            <div class="icono">
              <i class="fa fa-user icono-grande-side"></i>
            </div>
            <div class="titulo-opcion">Revisores</div>
          </c-button>
          <c-collapse id="colapse" :is-open="showRevisores">
            <div class="opcion-colapse">
              <a href="/administrador/registrar">Registrar revisor</a>
            </div>
            <div class="opcion-colapse">
              <a href="/administrador/revisores">Ver revisores</a>
            </div>
          </c-collapse>
        </div>
        <div class="opcion">
          <c-button
            id="boton-opcion-side"
            mb="4"
            variant-color="white"
            @click="showDashboard = !showDashboard"
          >
            <div class="icono">
              <i class="fa fa-bar-chart icono-grande-side"></i>
            </div>
            <div class="titulo-opcion">Dashboard</div>
          </c-button>
          <c-collapse id="colapse" :is-open="showDashboard">
            <div class="opcion-colapse">
              <a href="/administrador/registrar">Proximamente...</a>
            </div>
            <div class="opcion-colapse">
              <a href="/administrador/asignar">Proximamente...</a>
            </div>
          </c-collapse>
        </div>
      </div>
      <button class="cerrar-drawer" @click="caaa">Cerrar Drawer</button>
      <div class="cerrar-sesion">
        <c-button
          style="background: none; font-size: x-large"
          @click="cerrarSesion"
        >
          <div class="titulo-opcion">Cerrar sesión</div>
          <div class="icono">
            <i class="fa fa-sign-out icono-medio-side"></i>
          </div>
        </c-button>
      </div>
    </nav>
  </div>
</template>
<script>
import "font-awesome/css/font-awesome.css";
export default {
  components: {},
  data() {
    return {
      isOpen: false,
      showServicios: false,
      showRevisores: false,
      showDashboard: false,
      sidebarVisible: false,
    };
  },
  props: {
    cerrarDrawer: Function,
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
    caaa() {
      this.cerrarDrawer();
    },
  },
};
</script>
<style>
.sidenav {
  position: fixed;
  width: 280px;
  height: 100%;
  background-color: #00a490;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}

.opciones {
  flex: 1; /* Esto hace que las opciones ocupen el espacio restante */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  margin-top: 20%;
}
.opcion {
  width: 100%;
  align-items: center;
  display: flex;
  flex-direction: column;
  margin-bottom: 8%;
}
#boton-opcion-side {
  background: none;
  width: 100%;
  align-items: center;
  justify-content: start;
  font-size: x-large;
  border: none;
}
#boton-opcion-side:focus {
  outline: none;
  outline-offset: no-repeat;
  box-shadow: none;
}
#colapse {
  width: 17.5rem;
}
.opcion-colapse {
  font-size: large;
  color: white;
  margin-bottom: 10%;
  padding-left: 20%;
}

.opcion-colapse:hover {
  background-color: #394049;
}

.contendor-logo {
  width: 100%;
  height: 10%;
  display: flex;
}
.logo {
  background-image: url(../../../images/UsachSB.png);
  background-size: contain;
  background-repeat: no-repeat;
  width: 350px;
  height: 100%;
  background-position: 8px;
  margin-right: 20px;
}
.icono {
  margin-right: 10px; /* Ajusta el margen a tu preferencia */
  color: white;
}

.icono i.icono-grande-side {
  font-size: 30px;
  color: white; /* Ajusta el tamaño del icono según tus preferencias */
}

.icono i.icono-medio-side {
  font-size: 24px;
  color: black; /* Ajusta el tamaño del icono según tus preferencias */
  margin-left: 30%;
}

.icono-inicio {
  display: flex;
  flex-direction: row;
  margin-right: 52%;
  margin-bottom: 12%;
}
.opcion.active {
  background-color: #394049;
}
.cerrar-sesion {
  align-self: center;
  margin-bottom: 10%;
  font-weight: 500;
  font-size: larger;
  color: white;
  margin-right: 10%;
  background: none;
}
.cerrar-drawer {
  display: none;
}
@media (max-width: 768px) {
  .cerrar-drawer {
    display: flex;
  }
}
</style>
