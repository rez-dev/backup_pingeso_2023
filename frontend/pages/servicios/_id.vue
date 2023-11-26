<template>
  <div>
    <Navbar />
    <div class="containerPadre-rev">
      <form @submit.prevent="submitForm">
        <div class="containerForm-rev">
          <h1 class="Serviciotitulo-rev">{{ servicio.name }}</h1>
          <div class="containerDescripcion-rev">
            <c-button
              class="descripcion-serevisor"
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
                  :rows="3"
                />
              </div>
            </c-collapse>
          </div>
          <div v-for="item in itemsOrdenadosOriginal" :key="item.position">
            <div class="containerContenido-rev">
              <c-button
                class="descripcion-serevisor"
                mb="4"
                @click="toggleItem(item.position)"
              >
                {{ item.title }}
              </c-button>
              <c-collapse :is-open="showItem === item.position">
                <div class="autor-entrada-rev">
                  <span style="color: #ea7600">Autor: </span> {{ item.author }}
                </div>
                <div class="contenido-rev">
                  <c-textarea
                    :value="contenidoLimpio(item.content)"
                    v-model="item.content"
                  />
                  <c-select
                    class="selectID-revisor"
                    variant="outline"
                    placeholder="Orden"
                    v-model="item.order"
                  >
                    <option v-for="i in información.length" :value="i" :key="i">
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
        <div class="botones">
          <c-button variant-color="green" @click="previsualizar"
            >Previsualizar</c-button
          >
          <c-button variant-color="green" type="submit" >Enviar   </c-button>
        </div>
        <previsualizacion
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
</template>

<script>
import Navbar from "../../components/menus/revisores/Navbar.vue";
import Previsualizacion from "~/components/Previsualizacion.vue";
import precaucion from "../../components/alertas/precaucion.vue";
import exito from "../../components/alertas/exito.vue";
import axios from "axios";
import { CButton, CTextarea, CSelect } from "@chakra-ui/vue";
export default {
  components: {
    Navbar,
    CButton,
    CTextarea,
    CSelect,
    Previsualizacion,
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
      usuario: {},
    };
  },
  mounted() {
    const usuarioLog = JSON.parse(localStorage.getItem("UsuarioLog"));
    this.usuario = usuarioLog;
    this.verificacion();
    this.autorizado();
    this.obtenerservicio(this.$route.params.id);
  },
  methods: {
    toggleItem(itemKey) {
      if (this.showItem === itemKey) {
        this.showItem = null;
      } else {
        this.showItem = itemKey;
      }
    },
    previsualizar() {
      this.itemsOrdenados = [...this.información].sort(
        (a, b) => a.order - b.order
      );
      this.showPrevisualizacion = true;
    },
    devolver() {
      window.history.go(-1);
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
        this.$router.push("/servicios/lista");
      }, 1500);
    },

    obtenerservicio(id) {
      axios
        .get("http://localhost:8080/services/" + id)
        .then((response) => {
          const servicioaux = response.data;
          this.servicio = servicioaux[0];

          this.obtenerInformacion(this.servicio.id_wp_term);
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
          //console.error("Error al editar las entradas:", error);
        });
    },
    editarServicio(json, id) {
      axios
        .put("http://localhost:8080/approveService/" + id, json)
        .then((response) => {
          //console.log("bien", response.data)
        })
        .catch((error) => {
          //console.error("Error al editar el servicio", error);
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
      if (
        !this.isLogged ||
        this.usuario.role == "Administrador" ||
        this.usuario.role == "Superadministrador"
      ) {
        this.$router.push("/");
      }
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
};
</script>
<style>
.containerPadre-rev {
  height: 100%;
  justify-content: center;
  align-items: center;
  display: flex;
  flex-direction: column;
  padding-top: 3em;
  padding-bottom: 2em;
}
.containerForm-rev {
  background-color: white;
  width: 100%;
  border-radius: 10px;
  border: 3px solid #394049;
  display: flex;
  flex-direction: column;
}
.containerContenido-rev {
  padding: 1rem;
  padding-top: 2rem;
}
.containerDescripcion-rev {
  padding: 1rem;
  padding-top: 2rem;
}
.containerDescripcion-rev textarea{
  font-size: 18px;
}
.containerDescripcion-rev textarea:focus,
.containerDescripcion-rev textarea:hover {
  border: 1px solid #ea7600;
}


.Serviciotitulo-rev {
  font-size: xxx-large;
  font-weight: 600;
  text-align: center;
  color: #ea7600;
}
.descripcion-serevisor {
  background-color: white;
  justify-content: stretch;
  font-size: 35px;
  color: #00a499;
  white-space: normal;
}
.botones {
  width: 80%;
  display: flex;
  justify-content: space-around;
  padding-top: 2em;
}
.contenido-rev {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 12rem;
}
.selectID-revisor {
  width: 8rem !important;
  height: 5rem;
  border-radius: 5px;
  border: 1px solid #394049;
}

.selectID-revisor:hover {
  border-radius: 5px;
  border: 1px solid #ea7600;
}
.selectID-revisor:focus {
  border-radius: 5px;
  border: 1px solid #ea7600;
}
.selectID-revisor select {
  height: 100%;
  border: 1px solid #394049;
}
.autor-entrada-rev {
  margin-left: 1%;
  margin-bottom: 2%;
  font-size: x-large;
}
.contenido-rev textarea {
  width: 70%;
  height: 100%;
  font-size: 18px;
}
.contenido-rev textarea:hover {
  border: 1px solid #ea7600;
}
.contenido-rev textarea:focus {
  border: 1px solid #ea7600;
}
form {
  width: 90%;
  justify-content: center;
  align-items: center;
  display: flex;
  flex-direction: column;
}
.atras {
  font-size: 30px;
  align-self: start;
  margin-left: 4%;
}

@media (max-width: 900px) {
  .descripcion-serevisor {
    background-color: white;
    justify-content: stretch;
    font-size: x-large;
    color: #00a499;
    white-space: normal;
  }
  .selectID-revisor {
    width: 7rem !important;
    height: 3rem;
    border-radius: 5px;
    border: 1px solid #394049;
    margin-top: 2rem;
  }
  .contenido-rev textarea {
    width: 100%;
  }
  .containerContenido-rev {
    padding: 1rem;
    padding-top: 2rem;
  }
  .containerDescripcion-rev {
    padding: 1rem;
    padding-top: 2rem;
  }
  .containerDescripcion-rev textArea {
    width: 100%;
  }
  .contenido-rev {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }
  h1 {
    font-size: xx-large;
  }
}
@media (max-height: 800px) {
  .descripcion-serevisor {
    background-color: white;
    justify-content: stretch;
    font-size: xx-large;
    color: #00a499;
    white-space: normal;
  }
  .selectID-revisor {
    width: 7rem !important;
    height: 3rem;
    border-radius: 5px;
    border: 1px solid #394049;
    margin-top: 2rem;
  }
}
</style>
