<template>
  <div class="grid-container">
    <div class="container-sidebar">
      <sidebar />
    </div>
    <div class="content-revisar">
      <div style="width: 100%">
        <filtroServicios
          @serviciosFiltradosActualizados="actualizarDatos"
          @serviciosCompletos="mostrarTodosLosServicios"
          @nivelesCompletos="mostrarTodosLosniveles"
        />
      </div>

      <div class="servicios-container">
        <div class="tabla-container">
          <table class="tabla">
            <thead class="encabezado">
              <tr>
                <th
                  style="text-align: center;"
                >
                  ID
                </th>
                <th>Nombre servicio</th>
                <th>Unidad</th>
                <th>Estado</th>
                <th></th>
              </tr>
            </thead>
            <tbody class="cuerpo">
              <tr
                class="fila"
                v-for="(servicio, index) in paginatedServicios"
                :key="servicio.id"
              >
                <td style="width: 5%; text-align: center">{{ servicio.id }}</td>
                <td>{{ servicio.name }}</td>
                <td>{{ obtenerLevel1(servicio.id_wp_term) }}</td>
                <td>{{ servicio.state }}</td>
                <td>
                  <button>
                    <nuxt-link :to="`/administrador/servicios/${servicio.id}`">
                      <c-icon
                        name="arrow-forward"
                        size="2em"
                        style="padding-left: 10px"
                      />
                    </nuxt-link>
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
              :disabled="currentPage * itemsPerPage >= servicios.length"
            >
              Siguiente
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import sidebar from "../../components/menus/administradores/sidebar.vue";
import "font-awesome/css/font-awesome.css";
import { CCheckbox, CButton } from "@chakra-ui/vue";
import axios from "axios";
import filtroServicios from "../../components/listas/filtroServicios.vue";
export default {
  components: {
    sidebar,
    CCheckbox,
    CButton,
    filtroServicios,
  },
  data() {
    return {
      servicios: [],
      niveles: [],
      itemsPerPage: 7, // Cantidad de elementos por página
      currentPage: 1,
      isLogged: false,
      usuario: {},
    };
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    if (usuarioLog) {
      this.usuario = usuarioLog;
      this.verificacion();
      this.autorizado();
    }
  },
  computed: {
    paginatedServicios() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.servicios.slice(start, end);
    },
  },
  methods: {
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
    obtenerLevel1(id_wp_term) {
      const nivel = this.niveles.find((item) => item.id_wp_term === id_wp_term);
      return nivel ? nivel.level1 : "Categoría no encontrada";
    },
    actualizarDatos({ serviciosFiltrados, niveles }) {
      // Actualiza serviciosFiltrados y niveles con los datos emitidos desde el componente secundario
      this.servicios = serviciosFiltrados;
    },
    mostrarTodosLosServicios(serviciosCompletos) {
      this.servicios = serviciosCompletos;
    },
    mostrarTodosLosniveles(nivelesCompletos) {
      this.niveles = nivelesCompletos;
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

<style>
.grid-container {
  display: grid;
  grid-template-columns: 280px 1fr;
  height: 100vh;
}
.content-revisar {
  color: #000;
  justify-items: start;
  align-items: center;
  display: flex;
  flex-direction: column;
}
.tabla-container {
  width: 90%;
  border-radius: 5px;
  box-shadow: 6px 6px 3px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.servicios-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  flex-direction: column;
}

.tabla {
  width: 100%;
  padding: 8px;
}
.encabezado {
  width: 80%;
  text-align: justify;
  background-color: #ea7600;
  font-size: x-large;
}
.cuerpo {
  width: 80%;
}
.fila {
  border-top: 3px solid gray;
  font-size: large;
}
.fila td {
  padding-top: 15px;
  padding-bottom: 15px;
}
.fila:hover {
  background-color: #00a490;
  color: white;
}

#botonSeleccionar {
  align-self: center;
  width: 100%;
  height: 100%;
}

.paginacion {
  display: flex;
  justify-content: space-around;
  margin-top: 2%;
  background-color: #394049;
  color: white;
}
.contenedorFiltros {
  width: 100%;
  padding: 0.5rem;
}
.filtrar {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  padding: 20px;
}
.select {
  width: 20%;
  background-color: #394049;
  color: white;
  font-size: x-large;
  text-align: center;
  box-shadow: 4px 4px 8px #394049;
}

@media (max-width: 688px) {
  .grid-container {
    display: grid;
    grid-template-columns: 70px 1fr;
    height: 100vh;
  }
  .container-sidebar {
    widows: 70px;
  }
}
</style>
