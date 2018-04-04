const env = {}

env.PORT = process.env.PORT || 3000
env.CLIENT_PORT = process.env.CLIENT_PORT || 8080
env.DB_URL = 'mongodb://localhost/umbala'
env.INSTAGRAM = {}
env.INSTAGRAM.ClientID = 'f39bfe0995e04f298c94dc5eab37fc76'
env.INSTAGRAM.ClientSecret = '366364d74374444eab7e0ca7d3e89370'
env.INSTAGRAM.RedirectURL = `http://ec2-13-56-78-27.us-west-1.compute.amazonaws.com/about`
env.INSTAGRAM.AuthUrl = `https://api.instagram.com/oauth/authorize/?client_id=${env.INSTAGRAM.ClientID}&redirect_uri=${env.INSTAGRAM.RedirectURL}&response_type=code&scope=follower_list+basic`

module.exports = env
