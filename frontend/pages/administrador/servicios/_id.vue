<template>
  <div class="grid-container-ad">
    <div class="sidenavv">
      <sidebar />
    </div>
    <div class="content-asignar-ad">
      <div class="containerPadre-ad">
        <div class="ultima" v-if="this.servicio.state === 'Aprobado'">
          <i
            class="fa fa-check"
            aria-hidden="true"
            style="margin-top: 1%; color: green"
          ></i>
          <h1 style="font-family: cursive; margin-left: 2%; font-size: large">
            Última aprobación realizada por:
            <span style="color: #ea7600">{{ AutorAprobacion }}</span>
          </h1>
          <i
            class="fa fa-clock-o"
            aria-hidden="true"
            style="margin-top: 1%; color: #394049; margin-left: 3%"
          ></i>
          <h1 style="font-family: cursive; margin-left: 2%; font-size: large">
            Fecha última aprobación:
            <span style="color: #ea7600">{{ servicio.last_approval }}</span>
          </h1>
        </div>
        <form @submit.prevent="submitForm">
          <div class="containerForm-ad">
            <h1 class="Serviciotitulo">{{ servicio.name }}</h1>
            <div class="containerDescripcion-ad">
              <c-button
                class="descripcion-revisar-ad"
                mb="4"
                @click="toggleItem('descripcion')"
              >
                Descripción
              </c-button>
              <c-collapse :is-open="showItem === 'descripcion'">
                <div class="contenidoDescripcion">
                  <c-textarea
                    :value="servicio.description"
                    v-model="servicio.description"
                    
                  />
                </div>
              </c-collapse>
            </div>
            <div v-for="item in itemsOrdenadosOriginal" :key="item.position">
              <div class="containerContenido-ad">
                <c-button
                  class="descripcion-revisar-ad"
                  mb="4"
                  @click="toggleItem(item.position)"
                >
                  {{ item.title }}
                </c-button>
                <c-collapse :is-open="showItem === item.position">
                  <div class="autor-entrada">
                    <span style="color: #ea7600">Autor: </span>
                    {{ item.author }}
                  </div>
                  <div class="contenido-ad">
                    <c-textarea
                      :value="contenidoLimpio(item.content)"
                      v-model="item.content"
                    />
                    <label class="labeloculta" for="selectlabel-ad"
                      >Selecciona una opción:</label
                    >
                    <c-select
                      id="selectlabel-ad"
                      class="selectID-ad"
                      variant="outline"
                      placeholder="Orden"
                      v-model="item.order"
                    >
                      <option
                        v-for="i in información.length"
                        :value="i"
                        :key="i"
                      >
                        {{ i }}
                      </option>
                    </c-select>
                  </div>
                </c-collapse>
              </div>
            </div>
          </div>
          <precaucion
            :openConfirmar="openConfirmar"
            :titulo="'Aprobar servicio'"
            :descripcion="'¿Estás seguro que los datos están correctos?'"
            @confirmacion="confirmSubmit"
            @cerrar="closeAdvertencia"
          />
          <div class="botones-ad">
            <c-button variant-color="green" @click="previsualizar"
              >Previsualizar</c-button
            >
            <c-button variant-color="green" type="submit">Enviar</c-button>
          </div>
          <Previsualizacion
            v-if="showPrevisualizacion"
            :items="itemsOrdenados"
            :descripcion="servicio.description"
            @cerrar-previsualizacion="cerrarPrevisualizacion"
          />
        </form>
        <exito
          :openExitoso="openExitoso"
          :descripcion="'Se hizo el registro con éxito'"
        />
      </div>
    </div>
  </div>
</template>

<script>
import sidebar from "~/components/menus/administradores/sidebar.vue"
import Previsualizacion from "~/components/Previsualizacion.vue";
import precaucion from "../../../components/alertas/precaucion.vue";
import exito from "../../../components/alertas/exito.vue";
import axios from "axios";
import { CButton, CTextarea, CSelect } from "@chakra-ui/vue";
export default {
  components: {
    CButton,
    CTextarea,
    CSelect,
    Previsualizacion,
    sidebar,
    precaucion,
    exito,
  },
  data() {
    return {
      showItem: {},
      servicio: {},
      información: [],
      openConfirmar: false,
      openExitoso: false,
      showPrevisualizacion: false,
      itemsOrdenados: [],
      isLogged: false,
      AutorAprobacion: "",
      usuario: {},
    };
  },
  watch: {
    contenido() {
      this.actualizarFilas();
    },
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    if (usuarioLog) {
      this.usuario = usuarioLog;
    }
    this.verificacion();
    this.autorizado();
    this.obtenerservicio(this.$route.params.id);
    this.getAutorAprobacion(this.$route.params.id);
  },
  methods: {
    toggleItem(itemKey) {
      if (this.showItem === itemKey) {
        this.showItem = null;
      } else {
        this.showItem = itemKey;
      }
    },
    actualizarFilas() {
      // Calcular la altura deseada del textarea basada en el contenido
      const alturaDeseada = this.servicio.description.split("\n").length;
      // Establecer el número de filas del textarea
      this.filas = alturaDeseada;
    },
    previsualizar() {
      this.itemsOrdenados = [...this.información].sort(
        (a, b) => a.order - b.order
      );
      this.showPrevisualizacion = true;
    },
    cerrarPrevisualizacion() {
      this.showPrevisualizacion = false;
    },
    submitForm() {
      this.openAdvertencia();
    },
    confirmSubmit() {
      this.closeAdvertencia();
      const servicioNameTextareaValue = this.servicio.description;
      const servicioNameSelectValue = this.servicio.SelectedOption;

      let json = {
        id_user: this.servicio.id_user,
        description: servicioNameTextareaValue,
      };
      let jsonWordpress = {
        description: servicioNameTextareaValue,
      };

      this.editarServicio(json, this.servicio.id_wp_term);
      this.editarServicioWordpress(jsonWordpress, this.servicio.id_wp_term);

      this.información.forEach((item) => {
        const ContendidoLimpio = this.contenidoLimpio(item.content);

        let json = {
          author: item.author,
          content: ContendidoLimpio,
          id: item.id,
          id_author: item.id_author,
          id_post_meta: item.id_post_meta,
          id_wp_post: item.id_wp_post,
          id_wp_term: item.id_wp_term,
          order: parseInt(item.order),
          position: parseInt(item.order),
          title: item.title,
        };
        this.editarEntradas(json, item.id);
        this.editarEntradasWordpress(json, item.id_wp_post);
      });
      this.openMensaje();
      setTimeout(() => {
        this.$router.push("/administrador/revisar");
      }, 1500);
    },

    obtenerservicio(id) {
      axios
        .get("http://localhost:8080/services/" + id)
        .then((response) => {
          const servicioaux = response.data;
          this.servicio = servicioaux[0];
          this.formatearFecha();
          this.obtenerInformacion(this.servicio.id_wp_term);

          if (this.servicio.id_user === 0) {
            this.servicio.id_user = this.usuario.ID;
          }

          if (
            this.servicio.state === "Asignado" &&
            this.servicio.id_user !== this.usuario.ID
          ) {
            this.servicio.id_user = this.usuario.ID;
          }
        })
        .catch((error) => {
          console.error("Error al obtener los servicios", error);
        });
    },
    obtenerInformacion(id) {
      axios
        .get("http://localhost:8080/informationByTerm/" + id)
        .then((response) => {
          this.información = response.data;
          console.log(this.información);
        })
        .catch((error) => {
          console.error("Error al obtener las entradas:", error);
        });
    },
    editarEntradas(json, id) {
      axios
        .put("http://localhost:8080/informations/" + id, json)
        .then((response) => {
          console.log("bien", response.data);
        })
        .catch((error) => {
          console.error("Error al editar las entradas:", error);
        });
    },
    editarServicio(json, id) {
      axios
        .put("http://localhost:8080/approveService/" + id, json)
        .then((response) => {
          console.log("bien", response.data);
        })
        .catch((error) => {
          console.error("Error al editar el servicio", error);
        });
    },
    editarEntradasWordpress(json, id) {
      axios
        .put("http://localhost:8080/approveInformationWP/" + id, json)
        .then((response) => {
          //console.log("bien wordpress entradas", response.data)
          //console.log("bien wordpress entradas", response.data)
        })
        .catch((error) => {
          //console.error("Error al editar las entradas:", error);
        });
    },
    editarServicioWordpress(json, id) {
      axios
        .put("http://localhost:8080/approveServiceWP/" + id, json)
        .then((response) => {
          //console.log("bien wordpress servicio", response.data)
        })
        .catch((error) => {
          console.error("Error al editar el servicio", error);
        });
    },
    closeAdvertencia() {
      this.openConfirmar = false;
    },
    openAdvertencia() {
      this.openConfirmar = true;
    },
    openMensaje() {
      this.openExitoso = true;
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
    getAutorAprobacion(id) {
      axios
        .get("http://localhost:8080/getUserByService/" + id)
        .then((response) => {
          this.AutorAprobacion = response.data.name;
          console.log(this.AutorAprobacion);
        });
    },
    formatearFecha() {
      let fecha = new Date(this.servicio.last_approval);
      let dia = fecha.getDate();
      let mes = fecha.getMonth() + 1;
      let ano = fecha.getFullYear();
      let fechaFormateada = `${dia}/${mes}/${ano}`;
      this.servicio.last_approval = fechaFormateada;
    },
  },
  computed: {
    contenidoLimpio() {
      return (content) => {
        const contenidoSinEtiquetas = content
          .replace(/<!--.*?-->/g, "") // Eliminar etiquetas wp: y su contenido
          .replace(/<[^>]*>/g, "") // Eliminar etiquetas HTML
          
          .replace(/\n+/g, "\n") // Eliminar saltos de lineas multiples
          .replace(/\n/g, "")
          .replace(/&nbsp;/g, "\n");
          
        return contenidoSinEtiquetas;
      };
    },
    itemsOrdenadosOriginal() {
      return [...this.información].sort((a, b) => a.position - b.position);
    },
  },
}
</script>
<style>
.grid-container-ad {
  display: grid;
  grid-template-columns: 280px 1fr;
  height: 100vh;
}
.content-asignar-ad {
  color: #000;
  display: grid;
  justify-items: start;
  align-items: center;
}
.containerPadre-ad {
  height: 100%;
  align-items: center;
  display: flex;
  flex-direction: column;
  padding-bottom: 2em;
  width: 100%;
}
.autor-entrada {
  margin-left: 1%;
  font-size: x-large;
  margin-bottom: 2%;
}
.containerForm-ad {
  background-color: white;
  width: 100%;
  border-radius: 10px;
  border: 3px solid #394049;
  display: flex;
  flex-direction: column;
}
.containerContenido-ad {
  padding: 1rem;
  padding-top: 2rem;
}
.containerDescripcion-ad {
  padding: 1rem;
  padding-top: 2rem;
}
.contenidoDescripcion{
  height: 8rem;
}
.containerDescripcion-ad textarea{
  font-size: 18px;
  height: 100%;
}
.containerDescripcion-ad textarea:focus,
.containerDescripcion-ad textarea:hover {
  border: 1px solid #ea7600;
}
.descripcion-revisar-ad {
  background-color: white;
  justify-content: stretch;
  font-size: 35px;
  color: #00a499;
  white-space: normal;
}
.botones-ad {
  width: 80%;
  display: flex;
  justify-content: space-around;
  padding-top: 2em;
}
.contenido-ad {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 10rem;
}
.selectID-ad {
  width: 8rem !important;
  height: 5rem !important;
  border-radius: 5px;
  border: 1px solid #394049;
}
#selectlabel-ad {
  height: 100% !important;
}
.labeloculta {
  display: none;
}

.selectID-ad:hover {
  border-radius: 5px;
  border: 1px solid #ea7600;
}
.selectID-ad:focus {
  border-radius: 5px;
  border: 1px solid #ea7600;
}
.selectID-ad select {
  height: 100%;
  border: 1px solid #394049;
}
.selectID-ad select:focus {
  height: 100%;
  border: 1px solid #ea7600;
}
.contenido-ad textarea {
  width: 70%;
  height: 100%;
  font-size: 18px
}
.contenido-ad textarea:hover {
  border: 1px solid #ea7600;
}
.contenido-ad textarea:focus {
  border: 1px solid #ea7600;
}
.Serviciotitulo {
  font-size: xxx-large;
  font-weight: 600;
  text-align: center;
  color: #ea7600;
}
form {
  width: 90%;
  justify-content: center;
  align-items: center;
  display: flex;
  flex-direction: column;
  margin-top: 3%;
}
.ultima {
  width: 30%;
  display: flex;
  align-self: start;
  margin-left: 5%;
  margin-top: 2%;
}
</style>
