import Vue from 'vue'
import Router from 'vue-router'
import About from './views/About.vue'
import Home from './views/Home.vue'
import Test1 from './views/Test1.vue'
import Test2 from './views/Test2.vue'

Vue.use(Router)

export default new Router({
	mode: 'history',
	routes: [
		{
			path: '/',
			name: 'home',
			component: Home
		}, {
			path: '/test1',
			name: 'test1',
			component: Test1
		}, {
			path: '/test2',
			name: 'test2',
			component: Test2
		}, {
			path: '/about',
			name: 'about',
			component: About
		}
	]
})
