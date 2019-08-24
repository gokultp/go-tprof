<template>
  <div>
    <nav
      class="navbar header has-shadow is-primary"
      role="navigation"
      aria-label="main navigation"
    >
      <div class="navbar-brand">
        <div class="navbar-burger">
          <span />
          <span />
          <span />
        </div>
      </div>
    </nav>

    <section class="main-content columns">
      <div class="container column is-full">
        <nuxt />
      </div>
    </section>
  </div>
</template>

<script>
export default {
  data () {
    window.data = { 'pkgs': [] }
    return {
      items: [
      ]
    }
  }
}
</script>
