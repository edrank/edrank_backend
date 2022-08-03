# Edrank Backend Docs
- `tenant_type` can be one of the following types
    - `STUDENT`
    - `TEACHER`
    - `COLLEGE_ADMIN`
    - `SUPER_ADMIN`
    - `PARENT`
    - `HIEA`
- Onboarding File URL - https://omgupta-bucket.s3.ap-south-1.amazonaws.com/edrank/constants/EDRANK_ONBOARDING.xlsx
## Test Credentials

- Find Test Credentials [here](https://github.com/edrank/edrank_backend/blob/master/docs/CREDENTIALS.md)

## API Docs
### Login API (For all tenants)
`POST /login`
#### Request Headers
``` json
{
    "x-edrank-tenant-type":"COLLEGE_ADMIN"
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
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20iLCJjaWQiOjF9.SSgtr0KiyA3-Zm-UKkaq0HrjUAHrpxvFitvW67k5tvE",
        "tenant_id": 1,
        "tenant_type": "COLLEGE_ADMIN",
        "user": {
            "id": 1,
            "cid": 1,
            "name": "Om",
            "email": "omgupta1608@gmail.com",
            "is_active": true
        }
    },
    "message": "Logged In"
}
```
### File Upload API
`POST /file-upload`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Body (FORM DATA)
``` json
{
   "file": "FILE TO UPLOAD",
   "file_type": "COLLEGE_ONBOARDING"
}
```
#### Response
``` json
{
    "data": {
        "filepath": "https://omgupta-bucket.s3.ap-south-1.amazonaws.com/edrank/README.md",
        "file_registry_id": 2
    },
    "message": "File uploaded"
}
```

### Change Password API (For all tenants)
`POST /change-password`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Body (FORM DATA)
``` json
{
   "old_password": "om123",
   "new_password": "om1234"
}
```
#### Response
``` json
{
    "data": {
        "tenant_type": "COLLEGE_ADMIN"
    },
    "message": "Password changed successfully!"
}
```

### Get my college API (For STUDENT, COLLEGE_ADMIN, TEACHER)
`GET /college`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Response
``` json
{
    "data": {
        "college": {
            "id": 1,
            "name": "Maharaja Surajmal Institute",
            "email": "contact@msijanakpuri.com",
            "phone": "011-45656183",
            "website_url": "https://www.msijanakpuri.com/",
            "university_name": "Guru Gobind Singh Indraprastha University (GGSIPU)",
            "college_type": "SEMI_GOVT",
            "city": "New Delhi",
            "state": "Delhi",
            "onboarding_status": "ON_GOING",
            "is_active": true,
            "created_at": "2022-07-23T07:39:19Z",
            "updated_at": "2022-07-23T07:39:19Z"
        }
    },
    "message": "Fetched College"
}
```

### Create new College admin API
`POST /create-college-admin`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Body (FORM DATA)
``` json
{
    "name":"Test Admin",
    "email":"omgupta1608+1@gmail.com"
}
```
#### Response
``` json
{
    "data": {
        "college_admin_id": 6
    },
    "message": "College Admin Created!"
}
```

### Get my profile API
`GET /my-profile`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Response
``` json
{
    "data": {
        "profile": {
            "id": 1,
            "cid": 1,
            "name": "Om",
            "email": "omgupta1608@gmail.com",
            "is_active": true,
            "password": "$2a$14$MfY9CQYjnG7JS8Y8BYKOj.xiBm0DABfROWoYmh489sv8ifNTYqiJW",
            "created_at": "2022-07-23T06:42:01Z",
            "updated_at": "2022-07-24T15:45:52Z"
        }
    },
    "message": "My Profile fetched!"
}
```

### Get Top 3 Teachers API
`POST /top-3-teachers`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Body
- In case of `COLLEGE`, `cid` is important (other keys must be present)
- In case of `NATIONAL`, no other fields are required (keys must be present)
- In case of `REGIONAL`, `city` is important (other keys must be present)
- In case of `STATE`, `state` is important (other keys must be present)
``` json
{
    "request_type": "COLLEGE",
    "cid": 1,
    "city": "",
    "state": ""
}
```
#### Response
``` json
{
    "data": {
        "teachers": [
            {
                "id": 3,
                "name": "Virat Kohli",
                "score": 91.08,
                "rank": 1
            },
            {
                "id": 2,
                "name": "Rishi Dholkheria",
                "score": 83.98,
                "rank": 2
            },
            {
                "id": 1,
                "name": "Kavita Pabreja",
                "score": 61.23,
                "rank": 3
            }
        ]
    },
    "message": "Top 3 Teachers"
}
```

### Get College admins of my college API
`GET /my-college-college-admins`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Response
``` json
{
    "data": {
        "college_admins": [
            {
                "id": 1,
                "cid": 1,
                "name": "Om",
                "email": "omgupta1608@gmail.com",
                "is_active": true,
                "password": "$2a$14$MfY9CQYjnG7JS8Y8BYKOj.xiBm0DABfROWoYmh489sv8ifNTYqiJW",
                "created_at": "2022-07-23T06:42:01Z",
                "updated_at": "2022-07-24T15:45:52Z"
            }
        ]
    },
    "message": "College Admins of your College!"
}
```

### Get teachers of my college API
`GET /my-college-teachers?page=1&size=2`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Query params
``` json
{
    "page":"1",
    "size":"2"
}
```
#### Response
``` json
{
    "data": {
        "teachers": [
            {
                "id": 1,
                "cid": 1,
                "name": "Kavita Pabreja",
                "email": "omgupta1608+2@gmail.com",
                "alt_email": "omgupta1608+3@gmail.com",
                "department": "Computer Science",
                "course_id": 1,
                "designation": "Assistant Professor",
                "score": 61.23,
                "password": "$2a$14$YVHvpWB.FP2SG2qCl/JcXem6QeZD2v79Qkd4b3KMBAkUUYnkkk.XO",
                "is_active": true,
                "created_at": "2022-07-31T19:06:59Z",
                "updated_at": "2022-07-31T19:06:59Z"
            },
            {
                "id": 2,
                "cid": 1,
                "name": "Rishi Dholkheria",
                "email": "omgupta1608+3@gmail.com",
                "alt_email": "omgupta1608+4@gmail.com",
                "department": "Computer Applications",
                "course_id": 1,
                "designation": "Professor",
                "score": 83.98,
                "password": "$2a$14$x8BbpWEU.Vm8Kdypjq/HRu9QZr10AM9RCodg4ERCayflH.gLdnkXm",
                "is_active": true,
                "created_at": "2022-07-31T19:06:59Z",
                "updated_at": "2022-07-31T19:06:59Z"
            }
        ]
    },
    "message": "College Admins of your College!"
}
```


### Get students of my college API
`GET /my-college-students?page=1&size=1`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Query params
``` json
{
    "page":"1",
    "size":"1"
}
```
#### Response
``` json
{
    "data": {
        "students": [
            {
                "id": 1,
                "parent_id": 1,
                "cid": 1,
                "name": "Test Student",
                "email": "test+1@gmail.com",
                "phone": "9643959973",
                "course_id": 1,
                "year": 3,
                "batch": "2019-22",
                "password": "$2a$14$WUc6kfmTOuxK8YWF/SUZ.OFy1/Xcy4T1Zb.2VmGvansv188kr8LHG",
                "enrollment": "04214902019",
                "dob": "2001-08-16T00:29:56Z",
                "fathers_name": "Adam Gilchrist",
                "mother_name": "Kareena Kapoor",
                "guardian_email": "adam@gmail.com",
                "guardian_phone": "9872827722",
                "is_active": true,
                "created_at": "2022-07-30T19:03:55Z",
                "updated_at": "2022-08-02T15:56:52Z"
            }
        ]
    },
    "message": "Students of your College!"
}
```

