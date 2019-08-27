<template>
  <div ref="chartdiv" class="hello" />
</template>

<script>
import * as am4core from '@amcharts/amcharts4/core'
import * as am4charts from '@amcharts/amcharts4/charts'
import am4themesAnimated from '@amcharts/amcharts4/themes/animated'

am4core.useTheme(am4themesAnimated)

export default {
  name: 'Charts',
  props: ['data'],
  mounted () {
    const chart = am4core.create(this.$refs.chartdiv, am4charts.PieChart)

    chart.data = this.data
    chart.innerRadius = am4core.percent(50)
    // Add and configure Series
    const pieSeries = chart.series.push(new am4charts.PieSeries())
    pieSeries.dataFields.value = 'passed_cases'
    pieSeries.dataFields.category = 'name'
    pieSeries.slices.template.stroke = am4core.color('#fff')
    pieSeries.slices.template.strokeWidth = 2
    pieSeries.slices.template.strokeOpacity = 1

    // This creates initial animation
    pieSeries.hiddenState.properties.opacity = 1
    pieSeries.hiddenState.properties.endAngle = -90
    pieSeries.hiddenState.properties.startAngle = -90

    this.chart = chart
  },

  beforeDestroy () {
    if (this.chart) {
      this.chart.dispose()
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
.hello {
  width: 70%;
  height: 500px;
}
</style>
