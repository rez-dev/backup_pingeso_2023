<template>
  <div>
    <c-alert-dialog
      :is-open="openLista"
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
        <div class="titulo-tabla" v-if="servicios.length === 0">
          Actualmente
          <span style="color: #ea7600"> {{ revisor.name }} </span> no tiene
          servicios asociados
        </div>
        <div class="titulo-tabla" v-if="servicios.length > 0">
          Servicios a cargo de
          <span style="color: #ea7600"> {{ revisor.name }} </span>
        </div>
        <div class="container-servicios-tabla">
          <div class="tablaContainer">
            <table class="tabla">
              <tbody class="cuerpo">
                <tr
                  v-for="(servicio, index) in paginatedServicios"
                  :key="servicio.id"
                  class="fila"
                >
                  <td>{{ servicio.name }}</td>
                  <td class="boton-des">
                    <c-button variant-color="red" @click="abrirAdvertencia"
                      >Desasignar</c-button
                    >
                  </td>
                  <precaucion
                    :openConfirmar="openConfirmar"
                    :titulo="'Desasignar servicios'"
                    :descripcion="'¿Estás seguro que quieres desasignar el servicio seleccionado?'"
                    @confirmacion="desasignar(servicio)"
                    @cerrar="cerrarAdvertencia"
                  />
                </tr>
              </tbody>
            </table>
            <div class="paginacion" v-if="servicios.length > 0">
              <button @click="previousPage" :disabled="currentPage === 1">
                Anterior
              </button>
              <span>Página {{ currentPage }}</span>
              <button
                @click="nextPage"
                :disabled="currentPage * itemsPerPage >= servicios.length"
              >
                Siguiente
              </button>
            </div>
          </div>
        </div>
      </c-alert-dialog-content>
    </c-alert-dialog>

    <exito
      :openExitoso="openExitoso"
      :descripcion="'Se hizo el registro con éxito'"
    />
  </div>
</template>
<script>
import "font-awesome/css/font-awesome.css";
import { CButton } from "@chakra-ui/vue";
import axios from "axios";
import precaucion from "../alertas/precaucion.vue";
import exito from "../alertas/exito.vue";
export default {
  components: {
    CButton,
    precaucion,
    exito,
  },
  data() {
    return {
      servicios: [],
      openExitoso: true,
      itemsPerPage: 5, // Cantidad de elementos por página
      currentPage: 1,
      openConfirmar: false,
      openExitoso: false,
    };
  },
  props: {
    revisor: Object,
    openLista: Boolean,
  },
  computed: {
    paginatedServicios() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.servicios.slice(start, end);
    },
  },
  watch: {
    revisor: "obtenerServicios",
  },
  mounted() {
    this.obtenerServicios();
  },
  methods: {
    obtenerServicios() {
      axios
        .get("http://localhost:8080/servicesByUser/" + this.revisor.ID)
        .then((response) => {
          this.servicios = response.data.filter(
            (servicio) => servicio.state === "Asignado"
          );
          console.log(this.servicios);
        })
        .catch((error) => {
          console.log("Error al obtener los servicios");
        });
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
      const lastPage = Math.ceil(this.servicios.length / this.itemsPerPage);
      if (this.currentPage < lastPage) {
        this.currentPage++;
      }
    },
    cerrarTabla() {
      this.$emit("cerrar");
    },
    desasignar(servicioLista) {
      axios
        .put("http://localhost:8080/desassignService/" + servicioLista.id)
        .then((response) => {
          this.servicios = this.servicios.filter(
            (servicio) => servicio.id !== servicioLista.id
          );
          console.log("se actualizó el estado a pendiente");
          this.cerrarAdvertencia();
          setTimeout(() => {
            this.cerrarExito();
          }, 1500);
        });
    },
    abrirAdvertencia() {
      this.openConfirmar = true;
    },
    cerrarAdvertencia() {
      this.openConfirmar = false;
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
.container-servicios-tabla {
  width: 90%;
  margin-top: 2%;
  margin-bottom: 2%;
}
#boton-cerrar-tabla {
  align-self: end;
  background: none;
}
.titulo-tabla {
  font-size: xx-large;
}
.boton-des {
  width: 15%;
}
</style>
