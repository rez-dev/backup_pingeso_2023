<template>
  <div>
    <c-alert-dialog
      :is-open="openConfirmar"
      :least-destructive-ref="$refs.cancelRef"
      :on-close="cerrarAdvertencia"
    >
      <c-alert-dialog-overlay />
      <c-alert-dialog-content>
        <c-alert-dialog-header font-size="lg" font-weight="bold">
          {{ titulo }}
        </c-alert-dialog-header>
        <c-alert-dialog-body>
          {{ descripcion }}
        </c-alert-dialog-body>
        <c-alert-dialog-footer>
          <c-button ref="cancelRef" @click="cerrarAdvertencia">
            Cancel
          </c-button>
          <c-button variantColor="green" @click="confirmarEnvio" ml="3">
            Enviar
          </c-button>
        </c-alert-dialog-footer>
      </c-alert-dialog-content>
    </c-alert-dialog>
  </div>
</template>

<script>
import "font-awesome/css/font-awesome.css";
export default {
  components: {},
  props: {
    openConfirmar: Boolean,
    titulo: String,
    descripcion: String,
  },

  methods: {
    confirmarEnvio() {
      this.$emit("confirmacion");
    },
    cerrarAdvertencia() {
      this.$emit("cerrar");
    },
  },
};
</script>

<style>
#alerta {
  display: flex;
  align-items: center !important;
}
</style>
