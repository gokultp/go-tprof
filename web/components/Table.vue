<template>
  <section>
    <b-table
      :data="data"
      ref="table"
      detailed
      hoverable
      custom-detail-row
      :opened-detailed="['Board Games']"
      detail-key="name"
      :show-detail-icon="showDetailIcon"
    >
      <template slot-scope="props">
        <b-table-column
          field="name"
          label="Name"
          width="300"
        >
          <template >
            <a @click="toggle(props.row)">{{ props.row.name }}</a>
          </template>
        </b-table-column>

        <b-table-column
          field="time"
          label="Time"
          centered
        >{{ props.row.time }}</b-table-column>

        <b-table-column
          field="msg"
          label="Messages"
        >{{ props.row.error }}</b-table-column>

        <b-table-column
          label="Status"
          centered
        >
          <span
            :class="
                            [
                                'tag',
                                getTagClass(props.row.status)
                            ]"
          >{{ props.row.status }}</span>
        </b-table-column>
      </template>

      <template slot="detail" slot-scope="props">
        <tr v-for="(item, index) in props.row.items" :key="index">
          <td v-if="showDetailIcon"></td>
          <td>&nbsp;&nbsp;&nbsp;&nbsp;{{ item.name }}</td>
          <td class="has-text-centered">{{ item.time }}</td>
          <td
            class="has-text-centered"
          >{{ item.error }}</td>
          <td  class="has-text-centered">
            <span
              :class="
                            [
                                'tag',
                                getTagClass(item.status)

                            ]"
            >{{ item.status }}</span>
          </td>
        </tr>
      </template>
    </b-table>
  </section>
</template>

<script>
export default {
  data () {
    return {
      data: this.functions,
      showDetailIcon: false
    }
  },
  props: ['functions'],
  methods: {
    toggle (row) {
      this.$refs.table.toggleDetails(row)
    },
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
