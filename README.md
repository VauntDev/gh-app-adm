# gh-app-adm

simple app for interacting with github apps

## Access Token

```
curl -i -X POST \
-H "Authorization: Bearer YOUR_JWT" \
-H "Accept: application/vnd.github+json" \
https://api.github.com/app/installations/:installation_id/access_tokens
```

## Webhook config

```Bash
curl \
  -H "Accept: application/vnd.github+json" \ 
  -H "Authorization: Bearer <YOUR-TOKEN>" \
  https://api.github.com/app/hook/config
```
