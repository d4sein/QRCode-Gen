<template>
  <div id="app">
    <Container/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Container from './components/Container.vue'
import 'bulma/css/bulma.css'

@Component({
  components: {
    Container
  }
})
export default class App extends Vue {}
</script>
