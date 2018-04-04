<template>
  <div class="about">
	  <template v-if="fetchSucceeded">
		  <div class="container">
			  <div class="avatar" :style="{backgroundImage: 'url(' + user.profile_picture + ')'}"></div>
			  <!-- /.avatar -->
			  <div class="info">
				  <h2>{{user.username}}</h2>
			  </div>
			  <!-- /.info -->
			  <div class="follows">
				  <ul>
					  <li v-for="n in 10" :key="n">
						  <div class="follow">
							  <div class="follow__avatar"></div>
							  <!-- /.follow__avatar -->
							  <div class="follow__username">
								  <h3>mimi {{n}}</h3>
							  </div>
							  <!-- /.follow__username -->
						  </div>
					  </li>
				  </ul>
			  </div>
		  </div>
	  </template>
	  <template v-else>
		  <h3>Please login before</h3>
	  </template>
  </div>
</template>

<script>
import { saveCode, getAuth, authStorage } from '@/api/auth'
export default {
	name: 'about',
	data() {
		return {
			user: {
				profile_picture: '',
				username: ''
			},
			fetchSucceeded: false
		}
	},
	mounted() {
		const code = this.$route.query.code
		saveCode(code)

		getAuth({code})
		.then(result => {
			this.fetchSucceeded = true
			this.user = result.data.user || {}
		})
		.catch(err => {
			this.fetchSucceeded = false
			console.warn(err)
		})
	}
}
</script>

<style lang="stylus" scoped>
.container {
	padding: 10px
	width: 800px
	margin: 0 auto
	display: flex
	flex-direction: column
	align-items: center
}
.avatar {
	background-image: url("https://scontent.cdninstagram.com/vp/2a48e19acf7085d390b362ceaec78015/5B58A47F/t51.2885-19/s150x150/27879917_2043265929252284_7348559739170586624_n.jpg")
	height: 150px
	width: 150px
	border-radius: 50%
	background-size: center center
}

.follows {
	width: 800px
	margin: 0 auto
}

.follow {
	height: 70px
	display: flex
	flex-direction: row
	padding: 10px
	cursor: pointer
	transition: all 0.3s ease-in-out

	&__avatar {
		background-image: url("https://scontent.cdninstagram.com/vp/2a48e19acf7085d390b362ceaec78015/5B58A47F/t51.2885-19/s150x150/27879917_2043265929252284_7348559739170586624_n.jpg")
		height: 70px
		width: 70px
		border-radius: 50%
		background-size: center center
	}

	&__username {
		margin-left: 20px
	}

	&:hover {
		background: #42b983
	}
}
</style>
