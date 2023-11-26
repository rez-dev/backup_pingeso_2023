<template>
  <div class="grid-container">
    <div class="sidenavv">
      <sidebar class="uno" />
      <sidemenu class="dos" />
      <div class="tres">
        <c-button @click="isOpen = true">
          <div>
            <i class="fa fa-chevron-right"></i>
          </div>
        </c-button>
        <c-drawer
          :isOpen="isOpen"
          placement="left"
          :on-close="closeDrawer"
          :initialFocusRef="() => $refs.inputInsideModal"
        >
          <c-drawer-close-button />
          <sidebar :cerrarDrawer="close" />
        </c-drawer>
      </div>
    </div>
    <div class="content-revisores">
      <div class="titulo-revisores">Revisores de la unidad</div>
      <div class="revisores-container">
        <div class="tabla-container-revisores">
          <table class="tabla">
            <thead class="encabezado">
              <tr>
                <th
                  style="text-align: center;"
                >
                  ID
                </th>
                <th>Nombre revisor</th>
                <th>Correo</th>
                <th>Servcios</th>
              </tr>
            </thead>
            <tbody class="cuerpo">
              <tr
                class="fila"
                v-for="(revisor, index) in paginatedServicios"
                :key="revisor.ID"
              >
                <td style="width: 5%; text-align: center">{{ revisor.ID }}</td>
                <td>{{ revisor.name }}</td>
                <td>{{ revisor.Email }}</td>
                <td class="contenedor-boton-servicios">
                  <c-button
                    id="boton-servicios"
                    variant-color="green"
                    @click="abrirLista(revisor)"
                    >Servicios asignados</c-button
                  >
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
          <listaServicios
            :openLista="openLista"
            :revisor="revisorSeleccionado"
            @cerrar="cerrarAdvertencia"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import listaServicios from "../../components/listas/listaServicios.vue";
import sidemenu from "../../components/menus/administradores/sidemenu.vue";
import sidebar from "../../components/menus/administradores/sidebar.vue";
import "font-awesome/css/font-awesome.css";
import { CCheckbox, CButton } from "@chakra-ui/vue";
import axios from "axios";
import filtroServicios from "../../components/listas/filtroServicios.vue";
export default {
  components: {
    sidemenu,
    sidebar,
    CCheckbox,
    CButton,
    filtroServicios,
    listaServicios,
  },
  data() {
    return {
      usuario: {},
      revisores: [],
      itemsPerPage: 7, // Cantidad de elementos por página
      currentPage: 1,
      correoPropio: "",
      isLogged: false,
      openLista: false,
      isOpen: false,
      revisorSeleccionado: {},
    };
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    this.usuario = usuarioLog;
    this.verificacion();
    this.autorizado();
    this.getRevisores();
  },
  computed: {
    paginatedServicios() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.revisores.slice(start, end);
    },
  },
  methods: {
    getRevisores() {
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
          });
      }
      if (this.usuario.role === "Superadministrador") {
        axios
          .get("http://localhost:8080/usersByRole/Revisor")
          .then((response) => {
            this.revisores = response.data.filter(
              (revisor) => revisor.Email !== this.usuario.Email
            );
          });
      }
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
    abrirLista(revisor) {
      this.revisorSeleccionado = revisor;

      this.$nextTick(() => {
        this.openLista = true;
      });
    },
    cerrarAdvertencia() {
      this.openLista = false;
    },
    closeDrawer() {
      this.isOpen = false;
    },
    close() {
      this.isOpen = false;
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
.content-revisores {
  color: #000;
  justify-items: start;
  align-items: center;
  display: flex;
  flex-direction: column;
}
.tabla-container-revisores {
  width: 90%;
  border-radius: 5px;
  box-shadow: 6px 6px 3px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.revisores-container {
  display: flex;
  justify-content: center;
  width: 100%;
  margin-top: 5%;
}

.titulo-revisores {
  font-size: xxx-large;
  color: #ea7600;
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
  font-size: 20px;
}
.fila:hover {
  background-color: #00a490;
  color: white;
}

.paginacion {
  display: flex;
  justify-content: space-around;
  margin-top: 2%;
  background-color: #394049;
  color: white;
}
.contenedor-boton-servicios {
  width: 15%;
}
.dos {
  display: none;
}
.tres {
  display: none;
}

@media (max-width: 768px) {
  .uno {
    display: none;
  }
  .dos {
    display: block;
  }
  .tres {
    display: none;
  }
  .grid-container {
    display: grid;
    grid-template-columns: 80px 1fr;
    height: 100vh;
  }
  #boton-servicios {
    font-size: 12px;
  }
  .encabezado {
    width: 80%;
    text-align: justify;
    background-color: #ea7600;
    font-size: large;
  }
  .fila {
    border-top: 3px solid gray;
    font-size: medium;
  }
  .titulo-revisores {
    font-size: xx-large;
    color: #ea7600;
  }
  .contenedor-boton-servicios {
    width: 27%;
  }
}
@media (max-width: 480px) {
  .uno {
    display: none;
  }
  .dos {
    display: none;
  }
  .tres {
    display: block;
  }
  .grid-container {
    display: grid;
    grid-template-columns: 80px 1fr;
    height: 100vh;
  }
  #boton-servicios {
    font-size: 12px;
  }
  .encabezado {
    width: 80%;
    text-align: justify;
    background-color: #ea7600;
    font-size: large;
  }
  .fila {
    border-top: 3px solid gray;
    font-size: medium;
  }
  .titulo-revisores {
    font-size: xx-large;
    color: #ea7600;
  }
  .contenedor-boton-servicios {
    width: 27%;
  }
}
</style>
