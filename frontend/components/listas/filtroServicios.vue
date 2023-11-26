<template>
  <div>
    <div class="contenedorFiltros">
      <div class="filtrar">
        <select class="select" v-model="nivel1Seleccionado">
          <option value="">Nivel 1</option>
          <option v-for="nivel in niveles1Unicos" :key="nivel" :value="nivel">
            {{ nivel }}
          </option>
        </select>
        <select class="select" v-model="nivel2Seleccionado">
          <option value="">Nivel 2</option>
          <option v-for="nivel in niveles2Unicos" :key="nivel" :value="nivel">
            {{ nivel }}
          </option>
        </select>
        <select class="select" v-model="nivel3Seleccionado">
          <option value="">Nivel 3</option>
          <option v-for="nivel in niveles3Unicos" :key="nivel" :value="nivel">
            {{ nivel }}
          </option>
        </select>
        <select class="select" v-model="estadoSeleccionado">
          <option value="">Estado</option>
          <option v-for="estado in estadosUnicos" :key="estado" :value="estado">
            {{ estado }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  components: {},
  data() {
    return {
      nivel1Seleccionado: "",
      nivel2Seleccionado: "",
      nivel3Seleccionado: "",
      estadoSeleccionado: "",
      servicios: [],
      niveles: [],
      usuario: {},
    };
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    this.usuario = usuarioLog;
    this.obtenerservicios();
    this.obtenerCategorias();
  },
  computed: {
    niveles1Unicos() {
      const nivelesUnicos = new Set();
      this.niveles.forEach((item) => {
        nivelesUnicos.add(item.level1);
      });
      return Array.from(nivelesUnicos);
    },
    niveles2Unicos() {
      const nivelesUnicos = new Set();
      this.niveles.forEach((item) => {
        nivelesUnicos.add(item.level2);
      });
      return Array.from(nivelesUnicos);
    },
    niveles3Unicos() {
      const nivelesUnicos = new Set();
      this.niveles.forEach((item) => {
        nivelesUnicos.add(item.level3);
      });
      return Array.from(nivelesUnicos);
    },
    estadosUnicos() {
      const estadosUnicos = new Set();
      this.servicios.forEach((servicio) => {
        estadosUnicos.add(servicio.state);
      });
      return Array.from(estadosUnicos);
    },
    serviciosFiltrados() {
      const serviciosCompletos = this.servicios;
      if (
        !this.nivel1Seleccionado &&
        !this.nivel2Seleccionado &&
        !this.nivel3Seleccionado &&
        !this.estadoSeleccionado
      ) {
        // Si no se ha seleccionado nada en los filtros, mostrar todos los servicios.
        this.$emit("sinFiltros", serviciosCompletos);
      }
      let serviciosFiltrados = this.servicios;

      if (this.nivel1Seleccionado) {
        serviciosFiltrados = serviciosFiltrados.filter((servicio) => {
          const nivel = this.niveles.find(
            (item) => item.id_wp_term === servicio.id_wp_term
          );
          return nivel && nivel.level1 === this.nivel1Seleccionado;
        });
      }

      if (this.nivel2Seleccionado) {
        serviciosFiltrados = serviciosFiltrados.filter((servicio) => {
          const nivel = this.niveles.find(
            (item) => item.id_wp_term === servicio.id_wp_term
          );
          return nivel && nivel.level2 === this.nivel2Seleccionado;
        });
      }

      if (this.nivel3Seleccionado) {
        serviciosFiltrados = serviciosFiltrados.filter((servicio) => {
          const nivel = this.niveles.find(
            (item) => item.id_wp_term === servicio.id_wp_term
          );
          return nivel && nivel.level3 === this.nivel3Seleccionado;
        });
      }
      if (this.estadoSeleccionado) {
        serviciosFiltrados = serviciosFiltrados.filter(
          (servicio) => servicio.state === this.estadoSeleccionado
        );
      }
      return serviciosFiltrados;
    },
  },
  watch: {
    nivel1Seleccionado: "emitirEventoFiltrado",
    nivel2Seleccionado: "emitirEventoFiltrado",
    nivel3Seleccionado: "emitirEventoFiltrado",
    estadoSeleccionado: "emitirEventoFiltrado",
  },
  methods: {
    obtenerservicios() {
      if (this.usuario.role === "Revisor") {
        axios
          .get("http://localhost:8080/servicesByUser/" + this.usuario.ID)
          .then((response) => {
            this.servicios = response.data.filter(
              (servicio) => servicio.state === "Asignado"
            );
            console.log(this.servicios);
            this.$emit("serviciosCompletos", this.servicios, this.niveles);
          })
          .catch((error) => {
            console.error("Error al obtener personajes:", error);
          });
      }
      if (
        this.usuario.role === "Administrador" ||
        this.usuario.role === "Superadministrador"
      ) {
        axios
          .get("http://localhost:8080/services")
          .then((response) => {
            this.servicios = response.data;
            console.log(this.servicios);
            this.$emit("serviciosCompletos", this.servicios, this.niveles);
          })
          .catch((error) => {
            console.error("Error al obtener personajes:", error);
          });
      }
    },
    obtenerCategorias() {
      axios
        .get("http://localhost:8080/categories")
        .then((response) => {
          this.niveles = response.data;
          console.log(this.niveles);
          this.$emit("nivelesCompletos", this.niveles);
        })
        .catch((error) => {
          console.error("Error al obtener categorías:", error);
        });
    },
    emitirEventoFiltrado() {
      // Emite un evento personalizado con los servicios filtrados
      this.$emit("serviciosFiltradosActualizados", {
        serviciosFiltrados: this.serviciosFiltrados,
        niveles: this.niveles,
      });
    },
  },
};
</script>
