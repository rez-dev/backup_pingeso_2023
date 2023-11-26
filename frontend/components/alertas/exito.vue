<template>
    <div>
        <c-alert-dialog
        :is-open="openExitoso"
        :least-destructive-ref="$refs.cancelRef"
        >
            <c-alert-dialog-overlay />
            <c-alert-dialog-content id="alerta-exito">
                <c-icon name="check-circle" size="50px" style="padding-top: 10px;"/>
                <c-alert-dialog-header font-size="lg" font-weight="bold">
                    Envío existoso!!
                </c-alert-dialog-header>
                <c-alert-dialog-body>
                    {{descripcion}}
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
        openExitoso: Boolean,
        descripcion: String
    },
  }

</script>

<style>
    #alerta-exito {
    display: flex; 
    align-items: center !important;
    width: 30%;
    }
</style>