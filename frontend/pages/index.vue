<template>
  <section class="authContainer">
    <div class="leftSection">
      <div class="logo2"></div>
    </div>
    <div class="rightSection">
      <div class="login">
        <div class="headerLogin">
          <h1 class="tituloLogin">Bienvenido</h1>
        </div>
        <c-form-control class="textAreaLogin" is-required>
          <c-input
            id="fname"
            class="textLogin"
            placeholder="Correo"
            v-model="correo"
          />
        </c-form-control>
        <c-form-control class="textAreaLogin" is-required>
          <c-input
            id="fname"
            class="textLogin"
            type="password"
            placeholder="Contraseña"
            v-model="password"
          />
        </c-form-control>
        <button @click="redirigir" class="btnLogin">
          <h2 class="textBotonLogin">Ingresar</h2>
        </button>
      </div>
    </div>
  </section>
</template>
<script>
import axios from "axios";
import { CFormControl, CFormLabel, CFormHelperText } from "@chakra-ui/vue";
export default {
  name: "IndexPage",
  components: {
    CFormControl,
    CFormLabel,
    CFormHelperText,
  },
  data() {
    return {
      correo: "",
      password: "",
      usuario: {},
      userId: "",
    };
  },
  methods: {
    redirigir() {
      const datos = {
        email: this.correo,
        password: this.password,
      };
      axios.post("http://localhost:8080/login", datos).then((response) => {
        let tokenUser = this.parseJwt(response.data.token);
        this.userId = tokenUser.sub;
        this.getUser(this.userId);
      });
    },
    parseJwt(token) {
      var base64Url = token.split(".")[1];
      var base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
      var jsonPayload = decodeURIComponent(
        atob(base64)
          .split("")
          .map(function (c) {
            return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
          })
          .join("")
      );

      return JSON.parse(jsonPayload);
    },
    getUser(idUser) {
      axios.get("http://localhost:8080/users/" + idUser).then((response) => {
        this.usuario = response.data;
        const { password, ...usuarioSinPassword } = this.usuario;
        localStorage.setItem("UsuarioLog", JSON.stringify(usuarioSinPassword));
        if (this.usuario.role === "Revisor") {
          this.$router.push("/revisor");
        }
        if (
          this.usuario.role === "Administrador" ||
          this.usuario.role === "Superadministrador"
        ) {
          this.$router.push("/administrador");
        }
      });
    },
  },
};
</script>

<style>
.authContainer {
  display: flex;
  width: 100vw;
  height: 100vh;
  background: #ccd1d1;
  overflow: hidden;
}

.leftSection {
  position: relative;
  top: -10vw;
  left: -15vw;
  width: 200%;
  height: 200%;
  flex: 1;
  display: flex;
  align-items: top;
  justify-content: center;
  transform: rotate(15.925deg);
  background: #00a499;
}

.rightSection {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo2 {
  position: absolute;
  top: 20vw;
  left: 15vw;
  background-image: url(../images/UsachPB.png);
  background-size: contain;
  background-repeat: no-repeat;
  width: 35.1vw;
  height: 50.2vh;
  background-position: 8px;
  transform: rotate(-15.925deg);
}
.headerLogin {
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  top: -3.4vw;
  width: 24vw;
  height: 8.2vh;
  margin-bottom: 20px;
  background: #394049;
}
.tituloLogin {
  font-size: 2.9vh;
  font-style: normal;
  font-weight: 300;
  line-height: normal;
  color: #ffffff;
}
.login {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 35.2vw;
  height: 56.7vh;
  border-radius: 26px;
  background: #fff;
  box-shadow: 10px 10px 10px 10px rgba(0, 0, 0, 0.25);
}
.btnLogin {
  width: 17.9vw;
  height: 8.5vh;
  border-radius: 16px;
  background: #ea7600;
}
.textAreaLogin {
  width: 28.4vw;
  height: 6.7vh;
  margin-top: 4vh;
  margin-bottom: 6vh;
}
.textLogin {
  width: 28.4vw;
  height: 6.7vh;
  background: #ccd1d1;
  color: #000000;
}
.textBotonLogin {
  font-size: 20px;
  font-style: normal;
  font-weight: 300;
  line-height: normal;
  color: #ffffff;
}
</style>
