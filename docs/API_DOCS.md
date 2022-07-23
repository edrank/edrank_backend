# Edrank Backend Docs

## Test Credentials
### College Admin
``` json
{
    "email":"omgupta1608@gmail.com",
    "password":"om123"
}
```
``` json
{
    "email":"samridhikots@gmail.com",
    "password":"sam123"
}
```
``` json
{
    "email":"rishidholkheria2001@gmail.com",
    "password":"rishirishi"
}
```

## API Docs
### Login API
#### Request Headers
``` json
{
    "tenant_type":"STUDENT"|"TEACHER"|"PARENT"|"COLLEGE_ADMIN"|"SUPER_ADMIN"|"HEIA"
}
```
#### Request Body
``` json
{
    "email":"omgupta1608@gmail.com",
    "password":"om123"
}
```
#### Response
``` json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA",
    "tenant_id": 1,
    "tenant_type": "COLLEGE_ADMIN",
    "user": {
        "id": 1,
        "cid": 1,
        "name": "Om",
        "email": "omgupta1608@gmail.com",
        "is_active": true
    }
}
```