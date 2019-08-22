<template>
  <div>
    <b-collapse aria-id="contentIdForA11y2" class="panel" :open.sync="isOpen">
      <div slot="trigger" class="panel-heading" role="button" aria-controls="contentIdForA11y2">
        <strong>{{pkg.name}}</strong>
      </div>
      <div>
        {{pkg.name}}
        {{pkg.name}}
      </div>
    </b-collapse>
  </div>
</template>

<script>

export default {
  name: 'Charts',
  props: ['pkg']
}
</script>
