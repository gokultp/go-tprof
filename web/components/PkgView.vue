<template>
  <div>
    <b-collapse aria-id="contentIdForA11y2" class="panel" :open.sync="isOpen">
      <div slot="trigger" class="panel-heading columns" role="button" aria-controls="contentIdForA11y2">
        <div class="column">
            <strong>{{pkg.name}}</strong>
        </div>
        <div class="column">
          Coverage : {{pkg.coverage}}%
        </div>
        <div class="column">
          Time : {{pkg.time}}
        </div>
        <div class="column">
          <span
            :class="
                    [
                        'tag',
                        getTagClass(pkg.status)
                    ]"
          >{{ pkg.status }}</span>
        </div>
      </div>
      <div class="table">
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
  props: ['pkg'],
  methods: {
    getTagClass (status) {
      if (status === 'PASS') {
        return 'is-success'
      }
      if (status === 'FAIL') {
        return 'is-danger'
      }
    }
  }
}
</script>

<style >
  .table{
    margin-top: 20px;
  }
</style>