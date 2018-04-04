<template>
	<div class="home">
		<div>
			<button class="button__instagram">
				<a href="http://localhost:3000/login">Login with Instagram</a>
			</button>
		</div>
	</div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

export default {
	name: 'home',
	components: {
		HelloWorld
	}
}
</script>

<style lang="stylus" scoped>
.button {
	&__instagram {
		border-radius: 0
		background: #2c3e50
		padding: 15px 50px
		color: white
		font-size: 20px
		border: 0
		cursor: pointer
		transition: all 0.3s ease-in-out
		&:hover {
			background: #42b983
		}

		a {
			color: white
			text-decoration: none
		}
	}
}
</style>
