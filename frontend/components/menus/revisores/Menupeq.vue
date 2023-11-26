<template>
  <div class="conteiner">
    <c-menu>
      <c-menu-button class="botonpeq" right-icon="chevron-down">
      </c-menu-button>
      <c-menu-list>
        <c-menu-item class="opcion" as="a" href="/servicios/lista">
          Servicios
        </c-menu-item>
        <c-menu-item class="opcion" as="a" href="/">
          Cerrar sesión
        </c-menu-item>
      </c-menu-list>
    </c-menu>
  </div>
</template>

<script>
import {
  CMenu,
  CMenuButton,
  CMenuList,
  CMenuItem,
  CMenuGroup,
  CMenuDivider,
  CMenuOptionGroup,
  CMenuItemOption,
} from "@chakra-ui/vue";
export default {
  components: {
    CMenu,
    CMenuButton,
    CMenuList,
    CMenuItem,
    CMenuGroup,
    CMenuDivider,
    CMenuOptionGroup,
    CMenuItemOption,
  },
};
</script>

<style scoped>
.container {
  align-items: end;
  align-self: flex-end;
  height: 100%;
}
.botonpeq {
  width: 1.5rem;
}
.opcion:hover {
  background-color: rgb(160, 157, 157);
}
@media (max-width: 688px) {
  .botonpeq {
    height: 100%;
    width: 0.5rem;
    background-color: white;
  }
  .botonpeq:hover {
    background-color: gray.100;
  }
  .container {
    align-items: end;
    align-self: flex-end;
    height: 100%;
  }
  .opcion:hover {
    background-color: rgb(160, 157, 157);
  }
}
</style>
