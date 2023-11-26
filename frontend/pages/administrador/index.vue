<template>
  <div class="container-import">
    <div style="position: relative">
      <div class="adminn">
        <div class="texto">
          <h1 class="bienvenidoAdmin">
            Bienvenid@ <span style="color: #394049"> {{ usuario.name }} </span>
          </h1>
          <h2 class="subtituloAdmin">
            Como <span style="color: #394049">{{ usuario.role }}</span> tienes
            las siguientes acciones disponibles. <br />¿Cuál necesitas realizar?
          </h2>
        </div>
        <div class="logoo">
          <div class="foto"></div>
        </div>
      </div>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="100%"
        height="574"
        viewBox="0 0 1440 574"
        fill="none"
        preserveAspectRatio="none"
      >
        <path
          d="M-12 -11H1841.54C1841.54 -11 1912.33 536.5 1841.54 497.5C1770.74 458.5 -12 574 -12 574V-11Z"
          fill="#00A499"
        />
      </svg>
    </div>
    <div class="container-revisorlog">
      <cardOption
        :title="'Control de servicios'"
        :icon="'fa-book'"
        @clicIzquierda="prevCard"
        @clicDerecha="nextCard"
        :descripcion="'En esta sección puedes revisar servicios, asignar y ver los que se han aprobado'"
        :textBoton="'Ir a servicios'"
        :ruta="'/administrador/revisar'"
      />
      <cardOption
        :title="'Control de revisores'"
        :icon="'fa-user-plus'"
        @clicIzquierda="prevCard"
        @clicDerecha="nextCard"
        :descripcion="'En esta sección puedes ver los revisores de tu unidad y añadir nuevos revisores'"
        :textBoton="'Ir a revisores'"
        :ruta="'/administrador/revisores'"
      />
      <cardOption
        :title="'Estadísticas'"
        :icon="'fa-bar-chart'"
        @clicIzquierda="prevCard"
        @clicDerecha="nextCard"
        :descripcion="'Próximamente...'"
        :textBoton="''"
        :ruta="''"
      />
    </div>
  </div>
</template>
<script>
import { CDivider } from "@chakra-ui/vue";
import cardOption from "../../components/cardOption.vue";
export default {
  components: {
    CDivider,
    cardOption,
  },
  data() {
    return {
      usuario: {},
      isLogged: false,
    };
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    if (usuarioLog) {
      this.usuario = usuarioLog;
    }
    this.verificacion();
    this.autorizado();
  },
  methods: {
    redirigir() {
      this.$router.push("/lista");
    },
    prevCard() {
      if (this.currentCard > 0) {
        this.currentCard--;
      } else {
        this.currentCard = this.cardOptions.length - 1;
      }
    },
    nextCard() {
      if (this.currentCard < this.cardOptions.length - 1) {
        this.currentCard++;
      } else {
        this.currentCard = 0;
      }
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
      if (!this.isLogged || this.usuario.role == "Revisor") {
        this.$router.push("/");
      }
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
  flex-direction: column;
  height: 100vh;
  overflow: hidden;
}
.container-revisorlog {
  position: absolute;
  width: 100%;
  height: 55%;
  /* margin-top: 6%; */
  top: 22rem;
  display: flex;
}
.bienvenidoAdmin {
  color: white;
  font-size: 60px;
  margin-left: 15%;
  margin-top: 10%;
}
.subtituloAdmin {
  color: white;
  font-size: 30px;
  margin-left: 15%;
}

.adminn {
  width: 100%;
  height: 70%;
  margin-top: 1%;
  position: absolute;
  display: flex;
}

.texto {
  width: 50%;
}
.logoo {
  width: 60%;
  display: flex;
  justify-content: end;
}
.foto {
  background-image: url(../../images/UsachPB.png);
  background-size: contain;
  background-repeat: no-repeat;
  width: 40%;
  height: 100%;
}
</style>
