<template>
  <div>
    <pkg-view v-for="pkg in pkgs" v-bind:key="pkg.name" :pkg="pkg" />
    <b-table :data="data" :columns="columns" />
    <pie />
    <pie />
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
      pkgs: [
        {
          name: 'fmt'
        },
        {
          name: 'net/http'
        }
      ],
      data: window.data,
      columns: [
        {
          field: 'id',
          label: 'ID',
          width: '40',
          numeric: true
        },
        {
          field: 'first_name',
          label: 'First Name'
        },
        {
          field: 'last_name',
          label: 'Last Name'
        },
        {
          field: 'date',
          label: 'Date',
          centered: true
        },
        {
          field: 'gender',
          label: 'Gender'
        }
      ]
    }
  }
}
</script>
