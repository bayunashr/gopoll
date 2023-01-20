## Gopoll

- **Sign Up API Call**

http://localhost:3000/signup

Body:

```json
{
    "Email": "email@email.com",
	"Password": "password",
	"Name": "firstname lastname"
}
```

Response: 

```json
{
	"message": "success created new user"
}
```

- **Log In API Call**

http://localhost:3000/login

Body:

```json
{
    "Email": "email@email.com",
	"Password": "password"
}
```

Response: 

```json
{
	"message": "success logged in"
}
```