<template>
  <div class="grid-container">
    <div class="container-sidebar">
      <sidebar />
    </div>

    <c-grid class="grid-revisores">
      <c-grid-item col-span="2" bg="red.300" style="display: flex">
        <div class="cant-revisores">
          <div class="numero">XX</div>
          <h1>Revisores en la unidad</h1>
        </div>
        <div class="grafico-revisores">
            <h1> Disponibildiad revisores</h1>
          <div class="grafico-posicion">
            <apexchart
              type="pie"
              :options="chartOptions"
              :series="series"
              class="aaa"
            ></apexchart>
          </div>
        </div>
      </c-grid-item>
      <c-grid-item col-span="2" bg="red.300" style="display: flex">
        <div class="cant-revisores2">POner cantidad total de servicios</div>
        <div class="cant-revisores3">
          POner grafico de dona de los servicios
        </div>
      </c-grid-item>
      <c-grid-item col-span="4" bg="blue.300" />
    </c-grid>
  </div>
</template>

<script>
import sidebar from "../components/menus/administradores/sidebar.vue";
import "font-awesome/css/font-awesome.css";
import { CCheckbox, CButton, CGrid, CGridItem } from "@chakra-ui/vue";
import axios from "axios";
export default {
  components: {
    sidebar,
    CGrid,
    CGridItem,
  },
  data() {
    return {
        chartOptions: {
        chart: {
          type: 'pie',
          width: '100%',
          height: '100%',
        },
        colors: ['#EA7600', '#A04949'],
        labels: ["Ocupados", "Libres"],
        tooltip: {
          y: {
            formatter: function (val) {
              return val + " unidades";
            }
          }
        },
        legend: {
          position: 'bottom',
          height: 50,
        }
      },
      // Datos de la serie
      series: [5, 15],
    };
  },
  mounted() {
    console.log("Componente montado");
  },
  computed: {},
  methods: {},
};
</script>

<style>
.grid-container {
  display: grid;
  grid-template-columns: 280px 1fr;
  height: 100vh;
}
.content-revisar {
  color: #000;
  justify-items: start;
  align-items: center;
  display: flex;
  flex-direction: column;
}
.grid-revisores {
  width: 100%;
  height: 100%;
  gap: 0.25rem;
}
.cant-revisores {
  background-color: white;
  height: 100%;
  width: 45%;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.grafico-revisores {
  height: 100%;
  background-color: aqua;
  width: 55%;
  display: flex;
  flex-direction: column;
align-items: center;
justify-content: center;
}
.numero {
  width: 30%;
  height: 20%;
  background-color: #ea7600;
  color: black;
  font-size: x-large;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 1px solid black;
  border-radius: 10px;
  margin-top: 15%;
}
.grafico-posicion{
    width: 100%;
    height: 80%;

}
.aaa{
    height: 15rem;
    width: 100%;
}
</style>
