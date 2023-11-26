<template>
  <div class="container-import">
    <div class="container-revisorlog">
      <h1 class="Bienvenida">
        Bienvenid@ <br /><span style="color: #394049">
          {{ this.revisor.name }}
        </span>
      </h1>
      <c-divider
        border-color="black"
        style="margin-bottom: 8%; margin-left: 15%; width: 50%"
      />
      <h2 class="subtituloRevisor">
        Actualmente tienes
        <span style="color: #394049">{{ this.cantServicios }} <br /></span>
        servicios que necesitan tu <br />aprobación
      </h2>
      <c-button class="boton-pag" size="md" @click="redirigir">
        Ver servicios
      </c-button>
    </div>

    <div style="position: relative">
      <div class="overlay"></div>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="1130"
        height="1024"
        viewBox="0 0 1130 1024"
        fill="none"
      >
        <path d="M0 0H1130V1024H475.89L0 0Z" fill="#00A499" />
      </svg>
    </div>
  </div>
</template>
<script>
import { CDivider } from "@chakra-ui/vue";
import axios from "axios";
export default {
  components: {
    CDivider,
  },
  data() {
    return {
      revisor: {},
      cantServicios: 0,
      isLogged: false,
    };
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    if (usuarioLog) {
      this.revisor = usuarioLog;
    }
    this.getCantServicios();
    this.verificacion();
    this.autorizado();
  },
  methods: {
    redirigir() {
      this.$router.push("/servicios/lista");
    },
    verificacion() {
      const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
      if (usuarioLog) {
        this.isLogged = true;
      } else {
        this.isLogged = false;
      }
    },
    autorizado() {
      if (!this.isLogged) {
        this.$router.push("/");
      }
    },
    getCantServicios() {
      axios
        .get("http://localhost:8080/countServicesByUser/" + this.revisor.ID)
        .then((response) => {
          this.cantServicios = response.data.count;
        });
    },
  },
};
</script>
<style scoped>
.boton-pag {
  margin-left: 15%;
  margin-top: 15%;
  border-radius: 30px;
  background-color: #ea7600;
  color: white;
  width: 25%;
}
.boton-pag:hover {
  transform: scale(1.1, 1.1);
  background-color: #f1a252;
}
.container-import {
  display: flex;
  height: 100vh;
  overflow: hidden;
}
.container-revisorlog {
  width: 100%;
  display: flex;
  flex-direction: column;
  height: 100%;
  margin-bottom: 15%;
}
.Bienvenida {
  color: #ea7600;
  font-size: 60px;
  margin-left: 15%;
  margin-top: 15%;
}
.subtituloRevisor {
  color: #ea7600;
  font-size: 35px;
  margin-left: 15%;
}
.overlay {
  position: absolute;
  top: 10rem;
  left: 25rem;
  width: 50%;
  height: 50%;
  background-image: url(../../images/UsachPB.png);
  background-size: contain;
  background-repeat: no-repeat;
}
</style>
