const express = require('express')
const request = require('request-promise')
const bodyParser = require('body-parser')
const cors = require('cors')

const env = require('./env')
const app = express()

app.use(cors())
app.use(bodyParser.urlencoded({ extended: false }))

app.get('/login', (req, res) => {
	res.redirect(env.INSTAGRAM.AuthUrl)
})

app.get('/auth', async (req, res, next) => {
	try {
		const insUser = await request({
			url: 'https://api.instagram.com/oauth/access_token',
			method: 'POST',
			form: {
				client_id: env.INSTAGRAM.ClientID,
				client_secret: env.INSTAGRAM.ClientSecret,
				grant_type: 'authorization_code',
				redirect_uri: env.INSTAGRAM.RedirectURL,
				code: req.query.code
			},
			json: true
		})

		res.json(insUser)
	} catch (err) {
		next(err.message)
	}
})

app.get('/follows', async (req, res, next) => {
	try {
		const accessToken = req.session.accessToken
		if(!accessToken) {
			res.send('Missing access token')
		}
		const result = await request({
			url: `https://api.instagram.com/v1/users/self/follows?access_token=${accessToken}`,
			method: 'GET',
			json: true
		})
		console.log(result)
		res.send(result)
	} catch (err) {
		next(err.message)
	}
})

app.use((err, req, res, next) => {
	res.status(404).json({
		message: err.toString()
	})
})

app.listen(env.PORT, () => {
	console.log(`App is running on port ${env.PORT}`)
})
