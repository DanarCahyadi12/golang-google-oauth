
# Golang Google OAuth 

Example authentication with google.  

Make sure you have google client id, client secret, redirect url. 
# Get client id, client secret & redirect url tutorial
https://www.balbooa.com/help/gridbox-documentation/integrations/other/google-client-id
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`CLIENT_ID`=<YOUR_CLIENT_ID>

`CLIENT_SECRET`=<YOUR_CLIENT_SECRET>

`REDIRECT_URL`=<YOUR_REDIRECT_URL>


## Framework & Libary

- Fiber (HTTP framework) [https://docs.gofiber.io/]
- Viper (Configuration) [https://github.com/spf13/viper]
- Oauth2 (Google oauth library) [https://pkg.go.dev/golang.org/x/oauth2]

## How to auth with google?
- Run project by running "go run main.go" command
- Visit "http://localhost:8080/oauth/google" link on your browser
- Choose your google account
- Google automatically redirect to the URL you set before and return the user information.



