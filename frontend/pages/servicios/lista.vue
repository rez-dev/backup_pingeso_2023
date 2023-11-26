<template>
  <div style="height: 100%">
    <Navbar />
    <filtroServicios
      @serviciosFiltradosActualizados="actualizarDatos"
      @serviciosCompletos="mostrarTodosLosServicios"
      @nivelesCompletos="mostrarTodosLosniveles"
    />
    <div class="ed">
      <div class="tablaContainer">
        <table class="tabla">
          <thead class="encabezado">
            <tr>
              <th
                style="
                  text-align: center;
                  text-size: xx-large;
                  padding-right: 1rem;
                "
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
              v-for="servicio in serviciosFiltrados"
              :key="servicio.id"
              class="fila"
            >
              <td style="width: 5%; text-align: center">{{ servicio.id }}</td>
              <td>{{ servicio.name }}</td>
              <td>{{ obtenerLevel1(servicio.id_wp_term) }}</td>
              <td>{{ servicio.state }}</td>
              <td>
                <button>
                  <nuxt-link :to="`/servicios/${servicio.id}`">
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
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from "../../components/menus/revisores/Navbar.vue";
import filtroServicios from "../../components/listas/filtroServicios.vue";
import axios from "axios";
export default {
  components: {
    Navbar,
    filtroServicios,
  },
  data() {
    return {
      servicios: [],
      niveles: [],
      serviciosFiltrados: [],
      usuario: {},
    };
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    this.usuario = usuarioLog;
    this.verificacion();
    this.autorizado();
  },
  methods: {
    obtenerLevel1(id_wp_term) {
      const nivel = this.niveles.find((item) => item.id_wp_term === id_wp_term);
      return nivel ? nivel.level1 : "Categoría no encontrada";
    },
    actualizarDatos({ serviciosFiltrados, niveles }) {
      // Actualiza serviciosFiltrados y niveles con los datos emitidos desde el componente secundario
      this.serviciosFiltrados = serviciosFiltrados;
      this.niveles = niveles;
    },
    mostrarTodosLosServicios(serviciosCompletos) {
      this.serviciosFiltrados = serviciosCompletos;
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
      if (
        !this.isLogged ||
        this.usuario.role == "Administrador" ||
        this.usuario.role == "Superadministrador"
      ) {
        this.$router.push("/");
      }
    },
  },
};
</script>
<style>
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
.ed {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}
.tablaContainer {
  width: 90%;
  height: 100%;
  border-radius: 15px;
  border: 3px solid black;
  box-shadow: 4px 4px 8px gray;
  max-height: 49rem;
  overflow-y: auto;
}
.tablaContainer::-webkit-scrollbar {
  background-color: white;
  border-radius: 15px;
  border: 3px solid black;
  width: 15px;
}
.tablaContainer::-webkit-scrollbar-thumb {
  background-color: #394049;
  border-radius: 15px;
}
.tabla {
  width: 100%;
  height: 100%;
  padding: 8px;
}
.encabezado {
  width: 80%;
  text-align: justify;
  background-color: #ea7600;
  font-size: xx-large;
}
.cuerpo {
  width: 80%;
}
.fila {
  border-top: 3px solid gray;
  font-size: x-large;
}
.fila td {
  padding-top: 15px;
  padding-bottom: 15px;
}
.fila:hover {
  background-color: #00a490;
  color: white;
}

@media (max-width: 768px) {
  .encabezado {
    width: 80%;
    text-align: justify;
    background-color: #ea7600;
    font-size: x-large;
  }
  .fila {
    border-top: 3px solid gray;
    font-size: large;
  }
  .select {
    width: 20%;
    background-color: #394049;
    color: white;
    font-size: large;
    text-align: center;
    box-shadow: 4px 4px 8px #394049;
  }
  .tablaContainer::-webkit-scrollbar {
    background-color: white;
    border-radius: 15px;
    border: 3px solid black;
    width: 0.5rem;
  }
  a Cicon {
    size: 0.3rem;
  }
  .tablaContainer {
    width: 90%;
    height: 100%;
    border-radius: 15px;
    border: 3px solid black;
    box-shadow: 4px 4px 8px gray;
    max-height: 49rem;
    overflow-y: auto;
  }
}

@media (max-height: 900px) {
  .encabezado {
    width: 80%;
    text-align: justify;
    background-color: #ea7600;
    font-size: x-large;
  }
  .fila {
    border-top: 3px solid gray;
    font-size: large;
  }
  .tablaContainer {
    width: 90%;
    border-radius: 15px;
    border: 3px solid black;
    box-shadow: 4px 4px 8px gray;
    max-height: 40rem;
    overflow-y: auto;
  }
  .tablaContainer::-webkit-scrollbar {
    background-color: white;
    border-radius: 15px;
    border: 3px solid black;
    width: 0.5rem;
  }
  .ed {
    padding-bottom: 5%;
  }
}
@media (max-width: 500px) and (max-height: 896px) {
  .encabezado {
    width: 80%;
    text-align: justify;
    background-color: #ea7600;
    font-size: medium;
  }
  .fila {
    border-top: 3px solid gray;
    font-size: small;
  }
  .select {
    width: 20%;
    background-color: #394049;
    color: white;
    font-size: small;
    text-align: center;
    box-shadow: 4px 4px 8px #394049;
  }
  .tablaContainer::-webkit-scrollbar {
    background-color: white;
    border-radius: 15px;
    border: 3px solid black;
    width: 0.2rem;
  }
  a Cicon {
    size: 0.3rem;
  }
  .ed {
    padding-bottom: 5%;
  }
  .tablaContainer {
    width: 90%;
    border-radius: 15px;
    border: 3px solid black;
    box-shadow: 4px 4px 8px gray;
    max-height: 40rem;
    overflow-y: auto;
  }
}
</style>
