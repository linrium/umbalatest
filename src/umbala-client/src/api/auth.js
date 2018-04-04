import axios from 'axios'

export const authStorage = {
	saveCurrentUser: '',
	code: ''
}

export const api = axios.create({
	baseURL: 'http://ec2-18-144-10-33.us-west-1.compute.amazonaws.com:3000'
})

export const saveCode = (code) => {
	if(code) {
		authStorage.code = code
		localStorage.setItem('@umbala/code', code)
	}
}

export const getCode = () => {
	const code = localStorage.getItem('@umbala/code')
	authStorage.accessToken = code
	return code
}

export const saveCurrentUser = (currentUser) => {
	if(currentUser) {
		authStorage.currentUser = currentUser
		localStorage.setItem('@umbala/currentUser', JSON.stringify(currentUser))
	}
}

export const getCurrentUser = () => {
	const currentUser = localStorage.getItem('@umbala/currentUser')
	if(currentUser) {
		authStorage.currentUser = JSON.parse(currentUser)
	}
	return currentUser
}

export const getAuth = (payload = {}) => {
	console.log(payload)
	return api.get('/auth', {
		params: {
			code: payload.code || getCode()
		}
	})
}
