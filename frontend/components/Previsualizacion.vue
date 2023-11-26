<template>
  <div>
    <div class="containerMain">
      <c-alert-dialog
        :is-open="openExitoso"
        :least-destructive-ref="$refs.cancelRef"
      >
        <c-alert-dialog-overlay />
        <c-alert-dialog-content id="alerta-preview">
          <c-alert-dialog-body style="width: 100%">
            <div class="containerForm-preview">
              <div class="close-preview">
                <c-close-button @click="closePrevisualizacion" />
              </div>
              <h1 id="h1-preview">Previsualización</h1>
              <div class="containerDescripcion-preview">
                <c-button
                  class="descripcion-preview"
                  mb="4"
                  @click="toggleItem('descripcion')"
                >
                  Descripción
                </c-button>
                <c-collapse :is-open="showItem === 'descripcion'">
                  <div class="contenidoDescripcion-preview">
                    <c-textarea :value="descripcion" :readonly="true" />
                  </div>
                </c-collapse>
              </div>
              <div v-for="(item, index) in items" :key="index">
                <div class="containerContenido-preview">
                  <c-button
                    class="descripcion-preview"
                    mb="4"
                    @click="toggleItem(index)"
                  >
                    {{ item.title }}
                  </c-button>
                  <c-collapse :is-open="showItem === index">
                    <div class="contenido-preview">
                      <c-textarea
                        :value="contenidoLimpio(item.content)"
                        v-model="item.content"
                        :readonly="true"
                      />
                    </div>
                  </c-collapse>
                </div>
              </div>
            </div>
          </c-alert-dialog-body>
        </c-alert-dialog-content>
      </c-alert-dialog>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showItem: null,
      openExitoso: true,
    };
  },
  props: {
    items: Array,
    descripcion: String,
  },
  methods: {
    closePrevisualizacion() {
      this.$emit("cerrar-previsualizacion");
    },
    toggleItem(itemKey) {
      if (this.showItem === itemKey) {
        this.showItem = null; // Ocultar si se hizo clic nuevamente en el mismo botón
      } else {
        this.showItem = itemKey; // Mostrar si se hizo clic en un botón diferente
      }
    },
  },
  computed: {
    contenidoLimpio() {
      return (content) => {
        const contenidoSinEtiquetas = content
          .replace(/<!--.*?-->/g, "") // Eliminar etiquetas wp: y su contenido
          .replace(/<[^>]*>/g, "") // Eliminar etiquetas HTML
          .replace(/\n+/g, "\n"); // Eliminar saltos de lineas multiples
        return contenidoSinEtiquetas;
      };
    },
  },
};
</script>

<style>
.containerMain {
  height: 100%;
  max-width: 100%;
}
.containerForm-preview {
  background-color: white;
  width: 100%;
  border-radius: 10px;
  border: 3px solid #394049;
  display: flex;
  flex-direction: column;
}
.containerContenido-preview {
  padding: 1rem;
  padding-top: 2rem;
}
.containerDescripcion-preview {
  padding: 1rem;
  padding-top: 2rem;
}
.containerDescripcion-preview textarea:focus,
.containerDescripcion-preview textarea:hover {
  border: 1px solid #00a499;
}

#h1-preview {
  font-size: xxx-large;
  font-weight: 600;
  text-align: center;
  color: #00a499;
}
.descripcion-preview {
  background-color: white;
  justify-content: stretch;
  font-size: xxx-large;
  color: #ea7600;
  white-space: normal;
}
.contenido-preview {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.contenido-preview textarea {
  width: 100%;
}
.contenido-preview textarea:hover {
  border: 1px solid #00a499;
}
.contenido-preview textarea:focus {
  border: 1px solid #00a499;
}
#alerta-preview {
  display: flex;
  width: 80%;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  max-width: 80%;
  border: 2px solid white;
  border-radius: 5px;
}
.close-preview {
  align-self: flex-end;
}
</style>
