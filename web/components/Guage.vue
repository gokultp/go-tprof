<template>
  <div ref="chartdiv" class="guage" />
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
    // Create chart instance
    const chart = am4core.create(this.$refs.chartdiv, am4charts.RadarChart)
    // Add data
    // chart.data = [{
    //   'name': 'Research',
    //   'coverage': 80,
    //   'full': 100
    // }]
    chart.data = this.data
    console.log(this.data)
    // Make chart not full circle
    chart.startAngle = -90
    chart.endAngle = 180
    chart.innerRadius = am4core.percent(20)

    // Set number format
    chart.numberFormatter.numberFormat = '#.#\'%\''

    // Create axes
    const categoryAxis = chart.yAxes.push(new am4charts.CategoryAxis())
    categoryAxis.dataFields.category = 'name'
    categoryAxis.renderer.grid.template.location = 0
    categoryAxis.renderer.grid.template.strokeOpacity = 0
    categoryAxis.renderer.labels.template.horizontalCenter = 'right'
    categoryAxis.renderer.labels.template.fontWeight = 500
    categoryAxis.renderer.labels.template.adapter.add('fill', function (fill, target) {
      return (target.dataItem.index >= 0) ? chart.colors.getIndex(target.dataItem.index) : fill
    })
    categoryAxis.renderer.minGridDistance = 10

    const valueAxis = chart.xAxes.push(new am4charts.ValueAxis())
    valueAxis.renderer.grid.template.strokeOpacity = 0
    valueAxis.min = 0
    valueAxis.max = 100
    valueAxis.strictMinMax = true

    // Create series
    const series1 = chart.series.push(new am4charts.RadarColumnSeries())
    series1.dataFields.valueX = 'full'
    series1.dataFields.categoryY = 'name'
    series1.clustered = false
    series1.columns.template.fill = new am4core.InterfaceColorSet().getFor('alternativeBackground')
    series1.columns.template.fillOpacity = 0.08
    series1.columns.template.cornerRadiusTopLeft = 20
    series1.columns.template.strokeWidth = 0
    series1.columns.template.radarColumn.cornerRadius = 20

    const series2 = chart.series.push(new am4charts.RadarColumnSeries())
    series2.dataFields.valueX = 'coverage'
    series2.dataFields.categoryY = 'name'
    series2.clustered = false
    series2.columns.template.strokeWidth = 0
    series2.columns.template.tooltipText = '{name}: [bold]{coverage}[/]'
    series2.columns.template.radarColumn.cornerRadius = 20

    series2.columns.template.adapter.add('fill', function (fill, target) {
      return chart.colors.getIndex(target.dataItem.index)
    })

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
.guage {
  width: 80%;
  height: 500px;
}
</style>
