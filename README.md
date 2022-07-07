# gh-app-adm
simple app for interacting with github apps 


## Generate JWT for Github APP 
You can generate a JWT for your github app by calling the following command 

`./gh-app-adm jwt gen --private-key-path=YOUR_APPS_PRIVATE_KEY --app-id=YOUR_APPS_APPID`

## List Installations 
You can list the installations for your github app by using the jwt generated above and calling the following command 

`./gh-app-adm installations --jwt=YOUR_APPS_JWT` 

## Generate Access Token 
You can generate an access token for your github app by using the above generated as well as an install-id from the output from the installations list command shown above. 

`./gh-app-adm installations accessToken --jwt=YOUR_APPS_JWT --install-id=INSTALLATION_ID`