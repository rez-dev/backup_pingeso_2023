<template>
  <div class="grid-container">
    <div class="sidenavv">
      <sidebar />
    </div>
    <div class="content-asignar">
      <div class="container-primero">
        <div class="revisor-container">
          <div class="nombre-revisor" v-if="revisorSeleccionado">
            <div class="close-preview">
              <c-close-button @click="deseleccionarRevisor" color="black" />
            </div>
            <h1 style="margin-top: -3%">{{ revisorSeleccionado.name }}</h1>
            <h1>{{ revisorSeleccionado.Email }}</h1>
          </div>
          <c-button
            id="botonSeleccionar"
            v-else-if="!mostrarListaRevisores"
            @click="mostrarListaDeRevisores"
          >
            <i class="fa fa-plus icono-medio-side-a"></i>
            <h1>Seleccionar revisor</h1>
            
          </c-button>
          <listaRevisores
            v-else-if="mostrarListaRevisores"
            @revisor-seleccionado="seleccionarRevisor"
            @cerrar="cerrarListaRevisores"
          />
        </div>
        <div class="container-cant-servicios" v-if="revisorSeleccionado">
          <h1>{{ contarServiciosAsignados() }}</h1>
          <h1>Servicios asignados</h1>
        </div>
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
                <td>VRAE</td>
                <td>Pendiente</td>
                <td>
                  <CCheckbox variant-color="green" v-model="servicio.isChecked">
                    Asignar
                  </CCheckbox>
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
        <div
          class="boton"
          v-if="
            revisorSeleccionado &&
            servicios.some((servicio) => servicio.isChecked)
          "
        >
          <c-button
            variant-color="green"
            type="submit"
            style="width: 100%"
            @click="abrirAdvertencia"
            >Enviar</c-button
          >
        </div>
        <precaucion
          :openConfirmar="openConfirmar"
          :titulo="'Asignar servicios'"
          :descripcion="'¿Estás seguro que quieres asignar los servicio seleccionados?'"
          @confirmacion="enviarServiciosSeleccionados"
          @cerrar="cerrarAdvertencia"
        />
        <exito
          :openExitoso="openExitoso"
          :descripcion="'Se hizo el registro con éxito'"
        />
      </div>
    </div>
  </div>
</template>

<script>
import sidebar from "../../components/menus/administradores/sidebar.vue";
import "font-awesome/css/font-awesome.css";
import { CCheckbox, CButton } from "@chakra-ui/vue";
import axios from "axios";
import listaRevisores from "../../components/listas/listaRevisores.vue";
import precaucion from "../../components/alertas/precaucion.vue";
import exito from "../../components/alertas/exito.vue";
export default {
  components: {
    sidebar,
    CCheckbox,
    listaRevisores,
    CButton,
    exito,
    precaucion,
  },
  data() {
    return {
      isChecked: false,
      servicios: [],
      revisores: [],
      usuario: {},
      mostrarListaRevisores: false,
      revisorSeleccionado: null,
      itemsPerPage: 7,
      currentPage: 1,
      openConfirmar: false,
      openExitoso: false,
      openFracaso: false,
      isLogged: false,
    };
  },
  computed: {
    serviciosAsignados() {
      // Calcula la cantidad de servicios asignados
      return this.servicios.filter((servicio) => servicio.isChecked).length;
    },
    paginatedServicios() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.servicios.slice(start, end);
    },
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    if (usuarioLog) {
      this.usuario = usuarioLog;
    }
    this.verificacion();
    this.autorizado();
    this.obtenerservicios();
  },
  methods: {
    obtenerservicios() {
      axios
        .get("http://localhost:8080/servicesByState/Pendiente")
        .then((response) => {
          this.servicios = response.data;
          isChecked: false;
        })
        .catch((error) => {
          console.error("Error al obtener personajes:", error);
        });
    },
    mostrarListaDeRevisores() {
      this.mostrarListaRevisores = !this.mostrarListaRevisores;
      this.revisorSeleccionado = null; // Limpia la selección cuando se muestra la lista
    },
    seleccionarRevisor(revisor) {
      this.revisorSeleccionado = revisor;
      this.mostrarListaRevisores = false;
    },
    deseleccionarRevisor() {
      this.revisorSeleccionado = null;
      this.mostrarListaRevisores = false;
      this.servicios.forEach((servicio) => {
        servicio.isChecked = false;
      });
    },
    contarServiciosAsignados() {
      return this.serviciosAsignados;
    },
    enviarServiciosSeleccionados() {
      this.cerrarAdvertencia();
      const serviciosSeleccionados = this.servicios.filter(
        (servicio) => servicio.isChecked
      );
      const serviciosIDs = serviciosSeleccionados.map(
        (servicio) => servicio.id
      );

      let json = {
        id_user: this.revisorSeleccionado.ID,
        servicios: serviciosIDs,
      };

      if (serviciosSeleccionados.length > 0 && this.revisorSeleccionado) {
        axios
          .put("http://localhost:8080/assignServices", json)
          .then((reponse) => {
            console.log("Se asignaron con exito");
            this.abrirExito();
            setTimeout(() => {
              this.cerrarExito();
            }, 1500);
          });
      } else {
        console.log("No se han seleccionado servicios o revisor.");
      }
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
    abrirAdvertencia() {
      this.openConfirmar = true;
    },
    cerrarAdvertencia() {
      this.openConfirmar = false;
    },
    abrirExito() {
      this.openExitoso = true;
    },
    cerrarExito() {
      this.openExitoso = false;
      window.location.reload();
    },
    abrirFracaso() {
      this.openFracaso = true;
    },
    cerrarFracaso() {
      this.openFracaso = false;
      window.location.reload();
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
    cerrarListaRevisores() {
      this.mostrarListaRevisores = false;
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
.content-asignar {
  color: #000;
  display: grid;
  justify-items: start;
  align-items: center;
}
.container-primero {
  display: flex;
  width: 100%;
  height: 60%;
}
.container-cant-servicios {
  background-color: beige;
  display: flex;
  border-radius: 10px;
  box-shadow: 5px 5px 4px rgba(0, 0, 0, 0.2);
  margin-left: 5%;
  width: 30%;
  align-items: center;
  color: #ea7600;
  justify-content: center;
  flex-direction: column;
  font-size: xx-large;
}

.tabla-container {
  width: 90%;

  border-radius: 5px;
  box-shadow: 6px 6px 3px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.revisor-container {
  width: 50%;
  height: 100%;
  background-color: #d1d3d3;
  display: flex;
  border-radius: 10px;
  box-shadow: 5px 5px 4px rgba(0, 0, 0, 0.2);
  margin-left: 5%;
}
.servicios-container {
  display: flex;
  align-items: center;
  width: 100%;
  height: 38rem;
  flex-direction: column;
}
.nombre-revisor {
  width: 100%;
  background-color: beige;
  display: flex;
  align-items: center;
  flex-direction: column;
  font-size: xx-large;
  color: #ea7600;
  border-radius: 10px;
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
    display: flex;
    flex-direction: column;
    background-color: beige;
}
#botonSeleccionar:hover {
    background-color: rgb(218, 218, 200);
}
.close-preview {
  align-self: end;
  height: 20%;
  margin-right: 20px;
}

.boton {
  align-self: start;
  margin-left: 5%;
  width: 10%;
  margin-top: 3%;
}
.paginacion {
  display: flex;
  justify-content: space-around;
  margin-top: 2%;
  background-color: #394049;
  color: white;
}
.icono-medio-side-a{
  font-size: 30px;
  color: #394049;
  margin-bottom: 2%;
}
</style>
