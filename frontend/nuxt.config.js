require("dotenv").config();
export default {

  env: {
    version: process.env.VERSION
  },

  // Global page headers: https://go.nuxtjs.dev/config-head
  server: {
    host: process.env.NUXT_HOST,  // Escucha en todas las interfaces de red o la configuración proporcionada en NUXT_HOST
    port: process.env.NUXT_PORT  // Puerto en el que escucha o la configuración proporcionada en NUXT_PORT
  },

  // // Axios module configuration (https://go.nuxtjs.dev/config-axios)
  //axios: {
  //  baseURL: process.env.BASE_URL_BACKEND
  //},

  head: {
    title: 'frontend',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    { src: '~/plugins/vue-apexcharts', ssr: false },
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/chakra
    '@chakra-ui/nuxt',
    // https://go.nuxtjs.dev/emotion
    '@nuxtjs/emotion',
    '@nuxtjs/axios',
    '@nuxtjs/dotenv',
  ],
  fontAwesome: { // Configuración de Font Awesome
    icons: {
      solid: ['faHome', 'faUser', 'faCog', 'faAdressCard', 'faKey', 'faChartSimple', 'faPen'],
      regular: ['faEnvelope'],
      brands: ['faGithub', 'faWhatsapp']
    }
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  }
}
