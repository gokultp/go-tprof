<template>
  <div>
    <div class="columns">
      <div class="column">
        <h2> No of Testcases </h2>
        <pie :data="pkgs" />
      </div>
      <div class="column">
        <h2> Test Coverage </h2>
        <guage :data="pkgs" />
      </div>
    </div>
    <pkg-view v-for="pkg in pkgs" v-bind:key="pkg.name" :pkg="pkg" />
  </div>
</template>

<script>
import Pie from '../components/Pie'
import PkgView from '../components/PkgView'
import Guage from '../components/Guage'
export default {
  name: 'HomePage',
  components: {
    Pie,
    PkgView,
    Guage
  },
  data () {
    return {
      pkgs: window.data.pkgs
    }
  }
}
</script>
