<template>
  <div>
    <c-alert-dialog
      :is-open="openExitoso"
      :least-destructive-ref="$refs.cancelRef"
    >
      <c-alert-dialog-overlay />
      <c-alert-dialog-content id="alerta">
        <c-icon-button
          aria-label="Search database"
          icon="close"
          id="boton-cerrar-tabla"
          @click="cerrarTabla"
        >
        </c-icon-button>
        <div class="container-revisores-tabla">
          <div class="tablaContainer">
            <table class="tabla">
              <tbody class="cuerpo">
                <tr
                  v-for="(revisor, index) in paginatedServicios"
                  :key="revisor.id"
                  class="fila"
                >
                  <td>{{ revisor.name }}</td>
                  <td>{{ revisor.Email }}</td>
                  <td>
                    <button @click="seleccionarRevisor(revisor)">
                      Seleccionar
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
            <div class="paginacion">
              <button @click="previousPage" :disabled="currentPage === 1">
                Anterior
              </button>
              <span>Página {{ currentPage }}</span>
              <button
                @click="nextPage"
                :disabled="currentPage * itemsPerPage >= revisores.length"
              >
                Siguiente
              </button>
            </div>
          </div>
        </div>
      </c-alert-dialog-content>
    </c-alert-dialog>
  </div>
</template>
<script>
import "font-awesome/css/font-awesome.css";
import axios from "axios";
export default {
  components: {},
  data() {
    return {
      revisores: [],
      openExitoso: true,
      itemsPerPage: 5, // Cantidad de elementos por página
      currentPage: 1,
      usuario: {},
    };
  },
  computed: {
    paginatedServicios() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.revisores.slice(start, end);
    },
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    this.usuario = usuarioLog;
    this.obtenerRevisores();
  },
  methods: {
    obtenerRevisores() {
      if (this.usuario.role === "Administrador") {
        axios
          .get(
            "http://localhost:8080/usersByUnity/" +
              this.usuario.unity +
              "?role=revisor"
          )
          .then((response) => {
            this.revisores = response.data.filter(
              (revisor) => revisor.Email !== this.usuario.Email
            );
            console.log(this.revisores);
          })
          .catch((error) => {
            console.log("Error al obtener los revisores");
          });
      }
      if (this.usuario.role === "Superadministrador") {
        axios
          .get("http://localhost:8080/usersByRole/revisor")
          .then((response) => {
            this.revisores = response.data.filter(
              (revisor) => revisor.Email !== this.usuario.Email
            );
            console.log(this.revisores);
          })
          .catch((error) => {
            console.log("Error al obtener los revisores");
          });
      }
    },
    seleccionarRevisor(revisor) {
      this.$emit("revisor-seleccionado", revisor);
    },
    previousPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      }
    },
    nextPage() {
      const lastPage = Math.ceil(this.revisores.length / this.itemsPerPage);
      if (this.currentPage < lastPage) {
        this.currentPage++;
      }
    },
    cerrarTabla() {
      this.$emit("cerrar");
    },
  },
};
</script>
<style>
#alerta {
  display: flex;
  align-items: center !important;
  max-width: 70%;
  justify-content: center;
  border-radius: 10px;
  background-color: beige;
}
.container-revisores-tabla {
  width: 90%;
  margin-top: 2%;
  margin-bottom: 2%;
}
</style>
