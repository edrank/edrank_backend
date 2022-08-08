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
<details>

<summary style="font-size:20px">Login API</summary>

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
</details>


<details>

<summary style="font-size:20px">File Upload API</summary>

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

</details>

<details>

<summary style="font-size:20px">Change Password API (For all tenants)</summary>

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

</details>


<details>

<summary style="font-size:20px">Get my college API (For STUDENT, COLLEGE_ADMIN, TEACHER)</summary>

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

</details>

<details>

<summary style="font-size:20px">Create new College admin API</summary>

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

</details>

<details>

<summary style="font-size:20px">Get my profile API</summary>

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

</details>


<details>

<summary style="font-size:20px">Get Top N Teachers API</summary>

`POST /top-n-teachers`
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
- Send "n" = -1 for getting all the data
``` json
{
    "request_type": "COLLEGE",
    "cid": 1,
    "city": "",
    "state": "",
    "n": 3
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
                "rank": 1,
                "college_name": "Maharaja Surajmal Institute"
            },
            {
                "id": 2,
                "name": "Rishi Dholkheria",
                "score": 83.98,
                "rank": 2,
                "college_name": "Maharaja Surajmal Institute"
            },
            {
                "id": 1,
                "name": "Kavita Pabreja",
                "score": 61.23,
                "rank": 3,
                "college_name": "Maharaja Surajmal Institute"
            }
        ]
    },
    "message": "Top 3 Teachers"
}
```

</details>


<details>

<summary style="font-size:20px">Get College Entity of my college API</summary>

`GET /my-college-entity/:entity?page=1&size=2`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Params
- Entity should be one of `teachers`, `parents`, `students`, `college_admins`

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

</details>


<details>

<summary style="font-size:20px">Get Feedback Questions API</summary>

`GET /feedback-questions/:type`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Params
- Type should be one of `ST`, `SC`, `PC`, `HC`


#### Response
``` json
{
    "data": {
        "questions": [
            {
                "id": 1,
                "title": "How was the teacher's subject knowledge?",
                "option_1": "Excellent",
                "option_2": "Very Good",
                "option_3": "Fair",
                "option_4": "Poor",
                "Option5": "Very Poor",
                "type": "ST",
                "is_active": "1",
                "created_at": "2022-08-06T19:29:40Z",
                "updated_at": "0001-01-01T00:00:00Z"
            },
            {
                "id": 2,
                "title": "How was the teacher's communication skills?",
                "option_1": "Excellent",
                "option_2": "Good",
                "option_3": "Satisfactory",
                "option_4": "Bad",
                "Option5": "Very Bad",
                "type": "ST",
                "is_active": "1",
                "created_at": "2022-08-06T19:29:40Z",
                "updated_at": "0001-01-01T00:00:00Z"
            }
        ],
        "type": "ST"
    },
    "message": "Feedback Questions"
}
```

</details>