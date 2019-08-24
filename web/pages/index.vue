<template>
  <div>
    <pie />
    <pie />
    <pkg-view v-for="pkg in pkgs" v-bind:key="pkg.name" :pkg="pkg" />
  </div>
</template>

<script>
import Pie from '../components/Pie'
import PkgView from '../components/PkgView'
export default {
  name: 'HomePage',
  components: {
    Pie,
    PkgView
  },
  data () {
    return {
      pkgs: window.data.pkgs
    }
  }
}
</script>
