<template>
	<div>
		<h1>Rotate Matrix</h1>
		<input type="number" placeholder="Enter n" v-model="n">
		<button @click="fillMatrix">Create matrix</button>
		<button @click="rotate90deg">Rotate matrix</button>

		<table style="margin: 20px auto">
			<tr v-for="(row, index) in matrix" :key="index">
				<td v-for="val in row" :key="val">
					<span>{{val}}</span>
				</td>
			</tr>
		</table>

	</div>
</template>

<script>
export default {
	data() {
		return {
			n: 5,
			matrix: []
		}
	},
	methods: {
		createMatrix() {
			return Array.from({
				length: this.n
			}, () => new Array(this.n).fill(0))
		},
		fillMatrix() {
			const tmp = this.createMatrix()
			let count = 0
			for(let i = 0; i < this.n; i++) {
				for(let j = 0; j < this.n; j++) {
					count++
					tmp[i][j] = count
				}
			}
			this.matrix = tmp
		},
		rotate90deg() {
			this.matrix = this.matrix.reverse()
			for(let i = 0; i < this.matrix.length; i++) {
				for(let j = 0; j < i; j++) {
					const top = this.matrix[i][j]
					this.matrix[i][j] = this.matrix[j][i]
					this.matrix[j][i] = top
				}
			}
		}
	}
}
</script>

<style lang="stylus" scoped>
input {
	border: 1px solid #2c3e50
	padding: 10px
	font-size: 20px
	margin: 10px
}

button {
	font-size: 20px
	padding: 10px
	color: white
	background: #42b983
	border: none
	margin: 10px
	cursor: pointer
}
span {
	font-size: 20px
}
</style>
