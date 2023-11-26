<template>
  <div>
    <div class="grid-container">
      <div class="sidenavv">
        <sidebar />
      </div>
      <div class="content">
        <div class="right-container">
          <div class="header-form">Nuevo revisor</div>
          <div class="container-form">
            <div class="form-group">
              <c-form-control is-required>
                <c-form-label for="nombre">Nombre</c-form-label>
                <c-input-group>
                  <c-input-left-element
                    ><i class="fa fa-user icono-grande"></i
                  ></c-input-left-element>
                  <c-input
                    type="text"
                    id="nombre"
                    placeholder="Nombre"
                    v-model="nombre"
                    class="inputs"
                  />
                </c-input-group>
              </c-form-control>
            </div>
            <c-divider border-color="#EA7600" />
            <div class="form-group">
              <c-form-control is-required>
                <c-form-label for="rut">Rut</c-form-label>
                <c-input-group>
                  <c-input-left-element
                    ><i class="fa fa-address-card icono-grande"></i
                  ></c-input-left-element>
                  <c-input
                    id="rut"
                    type="text"
                    placeholder="Rut"
                    v-model="rut"
                    @input="formatearRut"
                    class="inputs"
                  />
                </c-input-group>
              </c-form-control>
            </div>
            <c-divider border-color="#EA7600" />
            <div class="form-group">
              <c-form-control is-required>
                <c-form-label for="correo">Correo</c-form-label>
                <c-input-group>
                  <c-input-left-element
                    ><c-icon
                      name="email"
                      color="black"
                      size="24px"
                      style="margin-top: 10px"
                  /></c-input-left-element>
                  <c-input
                    type="text"
                    id="correo"
                    placeholder="Correo electronico"
                    v-model="correo"
                    class="inputs"
                  />
                </c-input-group>
                <div style="margin-top: 3%">
                  <c-form-label
                    for="admin"
                    v-if="usuario.role === 'Superadministrador'"
                    >¿Es administrador?</c-form-label
                  >
                  <c-switch
                    color="vue"
                    id="admin"
                    v-if="usuario.role === 'Superadministrador'"
                    v-model="esAdministrador"
                    style="background-color: #ea7600"
                  />
                </div>
              </c-form-control>
            </div>
            <div
              class="form-contenedor-unidad"
              v-if="usuario.role === 'Superadministrador'"
            >
              <c-form-control is-required style="width: 100%;">
                <c-form-label for="unity">Unidad</c-form-label>
                <c-input-group>
                  <c-input-left-element
                    ><i class="fa fa-address-book icono-grande"></i
                  ></c-input-left-element>
                  <c-select id="unity" v-model="unity" 
                  placeholder="Selecciona la unidad" v-if="!mano">
                    <option
                    v-for="unidad in unidades"
                      :key="unidad"
                      :value="unidad"
                    >
                      {{ unidad }}
                    </option>
                  </c-select>
                  <c-input
                    type="text"
                    id="unity"
                    placeholder="Unidad"
                    v-model="unity"
                    class="inputs"
                    v-if="mano"
                  />
                </c-input-group>
              </c-form-control>
              <c-button class="boton-aniadir" @click="crearUnidad"> 
                    <i class="fa fa-plus" v-if="!mano"></i>
                    <i class="fa fa-minus" v-if="mano"></i>
                </c-button>
            </div>
            <c-divider border-color="#EA7600" />
            <div class="form-group">
              <c-form-control is-required>
                <c-form-label for="password">Contraseña</c-form-label>
                <c-input-group>
                  <c-input-left-element
                    ><i class="fa fa-key icono-grande"></i
                  ></c-input-left-element>
                  <c-input
                    type="password"
                    id="password"
                    placeholder="Contraseña"
                    v-model="password"
                    class="inputs"
                  />
                </c-input-group>
              </c-form-control>
            </div>
            <c-button class="boton-enviar" @click="abrirAdvertencia">
              Registrar revisor
            </c-button
            >
          </div>
        </div>
      </div>
    </div>
    <precaucion
      :openConfirmar="openConfirmar"
      :titulo="'Registrar revisor'"
      :descripcion="'¿Estás seguro que los datos están correctos?'"
      @confirmacion="enviarRegistro"
      @cerrar="cerrarAdvertencia"
    />
    <exito
      :openExitoso="openExitoso"
      :descripcion="'Se hizo el registro con éxito'"
    />
    <fracaso :openFracaso="openFracaso" />
  </div>
</template>

<script>
import sidebar from "../../components/menus/administradores/sidebar.vue";
import precaucion from "../../components/alertas/precaucion.vue";
import exito from "../../components/alertas/exito.vue";
import fracaso from "../../components/alertas/fracaso.vue";
import "font-awesome/css/font-awesome.css";
import axios from "axios";
import { CButton } from "@chakra-ui/vue";
export default {
  components: {
    sidebar,
    precaucion,
    exito,
    fracaso,
    CButton
  },
  data() {
    return {
      mano: false,
      nombre: "",
      rut: "",
      correo: "",
      password: "",
      unity:"",
      unidades: [
        "Dirección de Desarrollo Institucional",
        "Dirección de Administración y Finanzas",
        "Vicerrectoría Académica",
        "Vicerrectoría Investigación,  Desarrollo e Innovación",
        "Vicerrectoría de Vinculación con el Medio",
        "Vicerrectoría de Apoyo al Estudiante",
        "Vicerrectoría de Postgrado",
        "Dirección de Género, Diversidad y Equidad",
      ],
      openConfirmar: false,
      openExitoso: false,
      openFracaso: false,
      esAdministrador: false,
      revisor: {},
      isLogged: false,
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
    crearUnidad(){
      this.unity = "",
      this.mano = !this.mano;
    },
    enviarRegistro() {
      this.cerrarAdvertencia();
      const revisor = {
        name: this.nombre,
        rut: this.rut,
        email: this.correo,
        password: this.password,
        role: this.esAdministrador ? "Administrador" : "Revisor",
        unity: this.unity !== "" ? this.unity : this.usuario.unity,
      };
      axios
        .post("http://localhost:8080/signup", revisor)
        .then((response) => {
          console.log(revisor);
          this.abrirExito();
          setTimeout(() => {
            this.cerrarExito();
          }, 1500);
        })
        .catch((error) => {
          this.abrirFracaso();
          setTimeout(() => {
            this.cerrarFracaso();
          }, 1200);
        });
    },
    confirmarEnvio() {
      this.enviarRegistro();
    },
    abrirAdvertencia() {
      if (
        !this.nombre ||
        !this.rut ||
        !this.correo ||
        !this.password ||
        this.rut.length > 12
      ) {
        if (this.usuario.role === "Superadministrador" && !this.unity) {
          alert("Todos los campos son obligatorios");
        } else {
          alert("Todos los campos son obligatorios");
        }
      } else {
        this.openConfirmar = true;
      }
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
    formatearRut() {
      // Elimina caracteres que no son números ni k/K, excepto el último
      let rutFormateado = this.rut.replace(/[^\dkK]/g, "");

      // Verifica la longitud del rut (incluyendo el dígito verificador)
      if (rutFormateado.length <= 9) {
        // Divide el rut en parte numérica y dígito verificador
        let rutNumerico = rutFormateado.slice(0, -1);
        let digitoVerificador = rutFormateado.slice(-1);

        // Formato N°1: 12.345.678-k
        if (rutNumerico.length >= 7) {
          rutFormateado = rutNumerico
            .replace(/(\d)(?=(\d{3})+$)/g, "$1.")
            .concat("-", digitoVerificador.toLowerCase());
        }
        // Formato N°2: 1.234.567-k
        else {
          rutFormateado = rutNumerico.concat(
            "-",
            digitoVerificador.toLowerCase()
          );
        }
        this.rut = rutFormateado;
      }
    },
  },
};
</script>

<style>
.grid-container {
  display: grid;
  grid-template-columns: 280px 1fr; /* Define el ancho de la barra lateral y el espacio restante */
  height: 100vh; /* Establece la altura completa de la pantalla */
}
.icono-grande {
  font-size: 24px; /* Ajusta el tamaño del icono según tus preferencias */
  padding-top: 5px;
}
.inputs {
  height: 3rem;
  border: 1px solid black;
  box-shadow: 1px 1px;
}
.inputs:hover {
  background-color: rgb(226, 225, 225);
  border: 1px solid black;
}
.inputs:focus {
  background-color: rgb(226, 225, 225);
  border: 1px solid #ea7600;
}
.content {
  color: #000;
  display: grid;
  justify-items: center;
  align-items: center;
}
.right-container {
  width: 90%;
  height: 90%;
  background-color: #d1d3d3;
  justify-content: center;
  align-items: inherit;
  display: flex;
  flex-direction: column;
  border-radius: 10px;
  box-shadow: 5px 5px 4px rgba(0, 0, 0, 0.2);
}
.container-form {
  width: 80%;
  height: 80%;
}
.form-group {
  margin-bottom: 3%; /* Agregar margen inferior para separar los campos */
}
.etiqueta {
  padding-bottom: 15px;
}

.boton-enviar {
  width: 50%;
  margin-left: 25%;
  margin-top: 2%;
  height: 10%;
  padding: 10px 20px;
  background-color: #ea7600;
  border: 2px solid #ea7600;
  border-radius: 10px;
  color: white;
  transition: background-color 0.3s ease;
}
.boton-enviar:hover {
  background-color: #ea7600;
  outline: 1px solid black;
  outline-offset: 3px;
}
.header-form {
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  top: -8%;
  width: 80%;
  height: 8%;
  margin-bottom: 20px;
  background: #394049;
  color: white;
  font-size: xx-large;
  font-family: sans-serif;
}
#unity {
  height: 3rem;
  border: 1px solid black;
  box-shadow: 1px 1px;
  padding-left: 3.5%;
}
.boton-aniadir{
  height: 3rem;
  margin-left: 2%;
  background-color: white;
  align-self: end;
}
.boton-aniadir:hover{
  outline: 1px solid #ea7600;
  outline-offset: 3px;
  background-color: white;
}
.form-contenedor-unidad{
  display: flex;
  flex-direction: row;
  margin-bottom: 3%;
}
</style>
