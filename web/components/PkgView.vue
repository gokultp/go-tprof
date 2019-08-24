<template>
  <div>
    <b-collapse aria-id="contentIdForA11y2" class="panel" :open.sync="isOpen">
      <div slot="trigger" class="panel-heading" role="button" aria-controls="contentIdForA11y2">
        <strong>{{pkg.name}}</strong>
      </div>
      <div>
        <table-view :functions="pkg.functions"/>
      </div>
    </b-collapse>
  </div>
</template>

<script>
import TableView from './Table'
export default {
  name: 'Charts',
  components: {
    TableView
  },
  props: ['pkg']
}
</script>
