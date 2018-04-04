<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
	    <router-link to="/test1">Test 1</router-link> |
      <router-link to="/test2">Test 2</router-link> |
      <router-link to="/about">Test 3</router-link>
    </div>
    <router-view/>
  </div>
</template>

<script>
import { getCode } from './api/auth'

export default {
	mounted() {
		getCode()
//		getCurrentUser()
	}
}
</script>

<style lang="stylus">
#app
  font-family 'Avenir', Helvetica, Arial, sans-serif
  -webkit-font-smoothing antialiased
  -moz-osx-font-smoothing grayscale
  text-align center
  color #2c3e50

#nav
  padding 30px
  a
    font-weight bold
    color #2c3e50
    &.router-link-exact-active
      color #42b983

ul {
	list-style none
}
</style>
