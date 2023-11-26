<template>
    <div>
        <c-alert-dialog
        :is-open="openFracaso"
        :least-destructive-ref="$refs.cancelRef"
        >
            <c-alert-dialog-overlay />
            <c-alert-dialog-content id="alerta">
                <c-icon name="warning" size="50px" style="padding-top: 10px;"/>
                <c-alert-dialog-header font-size="lg" font-weight="bold">
                    No se pudo enviar
                </c-alert-dialog-header>
                <c-alert-dialog-body>
                    Ocurrió un error inesperado
                </c-alert-dialog-body>
                <c-alert-dialog-footer>
                </c-alert-dialog-footer>
            </c-alert-dialog-content>
        </c-alert-dialog>
    </div>
</template>

<script>
    import 'font-awesome/css/font-awesome.css';
    export default {
        components: {
    },
    props:{
        openFracaso: Boolean
    },
  }

</script>

<style> 
    #alerta {
    display: flex; 
    align-items: center !important;
    width: 30%;
    }
</style>