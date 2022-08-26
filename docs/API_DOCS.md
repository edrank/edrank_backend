# Edrank Backend Docs
- `tenant_type` can be one of the following types
    - `STUDENT`
    - `TEACHER`
    - `COLLEGE_ADMIN`
    - `SUPER_ADMIN`
    - `PARENT`
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

`POST /feedback-questions/:type`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Params
- Type should be one of `ST`, `SC`, `PC`, `HC`


#### Request Body
- In case of `PC` and `HC` - pass `cid` in body
- Not important in case of `SC` and `ST`

```json
{
    "cid": 1
}
```

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
                "option_5": "Very Poor",
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
                "option_5": "Very Bad",
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


<details>

<summary style="font-size:20px">Submit ST Feedback Form API</summary>

`POST /submit-feedback/ST`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Params
- Type should be one of `ST`, `SC`, `PC`, `HC`


#### Request Body
```json 
{
  "drive_id": 1,
  "feedbacks": [
    {
      "teacher_id": 1,
      "mcq": [
        {
          "question_id": 2,
          "answer_id": 4
        },
        {
          "question_id": 1,
          "answer_id": 3
        }
      ],
      "text_feedback": "Nice teaching."
    },
    {
      "teacher_id": 2,
      "mcq": [
        {
          "question_id": 1,
          "answer_id": 3
        }
      ],
      "text_feedback": "Poor teaching."
    }
  ]
}
```

#### Response
``` json
{
    "data": {},
    "message": "Feedback submitted!"
}
```

</details>




<details>

<summary style="font-size:20px">Get teachers for feedback API</summary>

`POST /get-feedback-teachers`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```

#### Request Body
```json 
{
    "course_id":1
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
                "password": "$2a$14$/L8ug6Gb3Wh8G6/LT5AGYeuS.72Tmb6cLjUAot9T02DvGkt2R0miS",
                "is_active": true,
                "created_at": "2022-07-31T19:06:59Z",
                "updated_at": "2022-08-08T04:20:30Z"
            }
        ]
    },
    "message": "Teachers"
}
```

</details>


<details>

<summary style="font-size:20px">Get my college's rank API</summary>

`POST /get-my-colleges-rank`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```

#### Request Body
- In case of `NATIONAL`, no other fields are required (keys must be present)
- In case of `REGIONAL`, `city` is important (other keys must be present)
- In case of `STATE`, `state` is important (other keys must be present)
```json 
{
    "request_type": "REGIONAL",
    "cid": 2,
    "city": "New Delhi",
    "state": ""
}
```

#### Response
- if returned `-1` or `any errors` , check if input is correct, or else show can't get rank
``` json
{
    "data": {
        "rank": 1
    },
    "message": "My college rank"
}
```

</details>


<details>

<summary style="font-size:20px">Get students for parent API</summary>

`GET /get-my-students`
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
        "students": [
            {
                "cid": 1,
                "name": "Test Student",
                "id": 1
            },
            {
                "cid": 1,
                "name": "Akshay Kumar",
                "id": 2
            }
        ]
    },
    "message": "My Students"
}
```

</details>


<details>

<summary style="font-size:20px">Get my text feedbacks API</summary>

`GET /get-my-text-feedbacks`
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
        "feedbacks": [
            {
                "text_feedback": "Very good teaching. Excellent teacher.",
                "sa_score": "88"
            }
        ]
    },
    "message": "Your feeedbacks"
}
```

</details>


<details>

<summary style="font-size:20px">Get course info API</summary>

`GET /get-course/1`

### Query params
- `course id` in params
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
        "course": {
            "id": 1,
            "name": "Bachelors of Computer Applications",
            "abbreviation": "BCA",
            "duration_in_years": 3
        }
    },
    "message": "Course"
}
```

</details>



<details>

<summary style="font-size:20px">Get Top N Colleges API</summary>

`POST /top-n-colleges`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Body
- In case of `NATIONAL`, no other fields are required (keys must be present)
- In case of `REGIONAL`, `city` is important (other keys must be present)
- In case of `STATE`, `state` is important (other keys must be present)
- Send "n" = -1 for getting all the data
``` json
{
    "request_type": "NATIONAL",
    "city": "",
    "state": "",
    "n": 3
}
```
#### Response
``` json
{
    "data": {
        "colleges": [
            {
                "id": 3,
                "score": 99.01,
                "name": "Sona College of Technology",
                "city": "Salem",
                "state": "Tamil Nadu",
                "rank": 1
            },
            {
                "id": 5,
                "score": 96.11,
                "name": "Indian Institute of Technology Delhi",
                "city": "New Delhi",
                "state": "Delhi",
                "rank": 2
            },
            {
                "id": 2,
                "score": 83.98,
                "name": "Testing College of Engineering",
                "city": "Indore",
                "state": "Madhya Pradesh",
                "rank": 3
            }
        ]
    },
    "message": "Top 3 Colleges"
}
```

</details>


<details>

<summary style="font-size:20px">Submit SC Feedback Form API</summary>

`POST /submit-feedback/SC`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Params
- Type should be one of `ST`, `SC`, `PC`, `HC`


#### Request Body
```json 
{
    "feedback": {
        "drive_id":1,
        "mcq":[
            {
                "question_id":1,
                "answer_id":2
            },
            {
                "question_id":2,
                "answer_id":3
            },
            {
                "question_id":3,
                "answer_id":1
            },
            {
                "question_id":4,
                "answer_id":1
            },
            {
                "question_id":5,
                "answer_id":5
            },
            {
                "question_id":6,
                "answer_id":4
            }
        ],
        "text_feedback":"Moderate college. Not very good. Poor staff and management but good infrastructure and facilities"
    }
}
```

#### Response
``` json
{
    "data": {},
    "message": "Feedback submitted!"
}
```

</details>



<details>

<summary style="font-size:20px">Get college_admin dashboard metrics API</summary>

`GET /dashboard-metrics`
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
        "college_feedbacks": 4,
        "drives": 0,
        "students": 2,
        "teachers": 3
    },
    "message": "College Metrics"
}
```

</details>




<details>

<summary style="font-size:20px">Submit PC Feedback Form API</summary>

`POST /submit-feedback/PC`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Params
- Type should be one of `ST`, `SC`, `PC`, `HC`


#### Request Body
```json 
{
    "feedback": {
        "drive_id":1,
        "college_id":1,
        "mcq":[
            {
                "question_id":1,
                "answer_id":2
            },
            {
                "question_id":2,
                "answer_id":3
            },
            {
                "question_id":3,
                "answer_id":1
            },
            {
                "question_id":4,
                "answer_id":1
            },
            {
                "question_id":5,
                "answer_id":5
            },
            {
                "question_id":6,
                "answer_id":4
            }
        ],
        "text_feedback":"Moderate college. Not very good. Poor staff and management but good infrastructure and facilities"
    }
}
```

#### Response
``` json
{
    "data": {},
    "message": "Feedback submitted!"
}
```

</details>



<details>

<summary style="font-size:20px">Submit Feedback Form API for ANDROID</summary>

`POST /android/submit-feedback/:type`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Params
- Type should be one of `ST`, `SC`, `PC`


#### Request Body
- for `SC` and `PC` - `college_id` is important
- for `ST`, `teacher_id` is important
```json 
{
    "drive_id": 1,
    "teacher_id": 4,
    "college_id": 1,
    "feedback": {
        "mcq": [
            {
                "question_id": 2,
                "answer_id": 4
            },
            {
                "question_id": 1,
                "answer_id": 3
            }
        ],
        "text_feedback": "Teaching was moderate. I didn't understand many concepts yet the teacher didn't bother."
    }
}
```

#### Response
``` json
{
    "data": {},
    "message": "Feedback submitted!"
}
```

</details>


<details>

<summary style="font-size:20px">Get my rank API</summary>

`GET /get-my-rank/:rank_type`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Params
- `rank_type` can be one of - `NATIONAL`, `REGIONAL`, `STATE` or `COLLEGE`

#### Response
``` json
{
    "data": {
        "rank": 6
    },
    "message": "My Rank"
}
```

</details>


<details>

<summary style="font-size:20px">KBC GRAPH API</summary>

`POST /kbc-graph`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```
#### Request Body
- Incase of teacher - no need to pass `teacher_id`
- Incase of college_admin - `teacher_id` is mandatory
``` json
{
    "question_id":7,
    "teacher_id": 1
}
```
#### Response
``` json
{
    "data": {
        "graph_data": [
            {
                "OptionName": "Excellent",
                "ResponseCount": 3
            },
            {
                "OptionName": "Very Good",
                "ResponseCount": 0
            },
            {
                "OptionName": "Fair",
                "ResponseCount": 1
            },
            {
                "OptionName": "Poor",
                "ResponseCount": 0
            },
            {
                "OptionName": "Very Poor",
                "ResponseCount": 1
            }
        ]
    },
    "message": "Graph Fetched!"
}
```

</details>



<details>

<summary style="font-size:20px">SA GRAPH API</summary>

`POST /sa-graph/:type`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```

#### Request params
- `type` can be one of `teacher` or `college`
- pass `id` accordingly
#### Request Body
- Incase of teacher - no need to pass `college_id`
- Incase of college - `college_id` is mandatory
``` json
{
    "college_id": 1,
    "teacher_id": 2
}
```
#### Response
``` json
{
    "data": {
        "sa_graph": {
            "Negative": 6,
            "Neutral": 0,
            "Positive": 0
        }
    },
    "message": "Sentimental Analysis Graph"
}
```

</details>



<details>

<summary style="font-size:20px">Grivenancencen Submit API</summary>

`POST /submit-gc`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```

#### Request Body
``` json
{
    "complaint_for":"college",
    "subject":"Very brutal teacher behavior",
    "is_cc": true,
    "cc_response": "Its okay. Leave it.",
    "proof_file_id": 2,
    "description":"Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

}
```
#### Response
``` json
{
    "data": {
        "id": 1
    },
    "message": "Grievance Submitted. Will be reviwed shortly"
}
```

</details>


<details>

<summary style="font-size:20px">Progress graph API</summary>

`GET /progress-graph/:tid`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOjEsInRlbmFudF90eXBlIjoiQ09MTEVHRV9BRE1JTiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJvbWd1cHRhMTYwOEBnbWFpbC5jb20ifQ.UFnQCWw_9lsD6bDqHx4RJalvNxwuTmSkeVzuCsQ_TlA"
}
```

#### Request params
- `tid` is the teacher Id

#### Response
``` json
{
    "data": {
        "progress": [
            {
                "drive_id": 1,
                "started_at": "2022-08-26T02:44:54Z",
                "rank": 2
            },
            {
                "drive_id": 2,
                "started_at": "2022-08-26T05:38:21Z",
                "rank": 11
            },
            {
                "drive_id": 3,
                "started_at": "2022-08-26T05:38:21Z",
                "rank": 7
            }
        ]
    },
    "message": "Progress"
}
```

</details>



<details>

<summary style="font-size:20px">Linking API</summary>

`GET /get-teacher-linking-data`
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
        "courses": [
            {
                "id": 1,
                "name": "Bachelors of Computer Applications",
                "abbreviation": "BCA",
                "duration_in_years": 3
            },
            {
                "id": 2,
                "name": "Bachelors of Technology",
                "abbreviation": "B Tech.",
                "duration_in_years": 4
            },
            {
                "id": 3,
                "name": "Bachelors of Business Administration",
                "abbreviation": "BBA",
                "duration_in_years": 3
            },
            {
                "id": 4,
                "name": "Bachelors of Commerce",
                "abbreviation": "B Com.",
                "duration_in_years": 3
            },
            {
                "id": 5,
                "name": "Bachelors of Education",
                "abbreviation": "B Ed.",
                "duration_in_years": 2
            },
            {
                "id": 6,
                "name": "Masters of Computer Applications",
                "abbreviation": "MCA",
                "duration_in_years": 3
            },
            {
                "id": 7,
                "name": "Masters of Technology",
                "abbreviation": "M Tech.",
                "duration_in_years": 2
            },
            {
                "id": 8,
                "name": "Masters of Business Administration",
                "abbreviation": "MBA",
                "duration_in_years": 2
            },
            {
                "id": 9,
                "name": "Bachelor of Architecture",
                "abbreviation": "B Arch",
                "duration_in_years": 5
            },
            {
                "id": 10,
                "name": "Bachelor of Pharmacy",
                "abbreviation": "B Pharm",
                "duration_in_years": 4
            },
            {
                "id": 11,
                "name": "Master of Architecture",
                "abbreviation": "M Arch",
                "duration_in_years": 2
            },
            {
                "id": 12,
                "name": "Bachelor of Arts",
                "abbreviation": "BA",
                "duration_in_years": 3
            },
            {
                "id": 13,
                "name": "Bachelor of Science",
                "abbreviation": "B.Sc.",
                "duration_in_years": 3
            },
            {
                "id": 14,
                "name": "Bachelor of Tourism and Travel Management",
                "abbreviation": "BTTM",
                "duration_in_years": 4
            },
            {
                "id": 15,
                "name": "Bachelor of Journalism and Mass Communication",
                "abbreviation": "BJMC",
                "duration_in_years": 3
            }
        ],
        "students": [
            {
                "id": 1,
                "parent_id": 1,
                "cid": 1,
                "name": "Sumit Verma",
                "email": "test+1@gmail.com",
                "phone": "9643959973",
                "course_id": 1,
                "year": 3,
                "batch": "2019-22",
                "password": "$2a$14$74TX4CB2Q4acP2/K3XtQ0.FlPPk3aW0.TFvsxWd55JEemBfBiStuq",
                "enrollment": "04214902019",
                "dob": "2001-08-16T00:29:56Z",
                "fathers_name": "Brijesh Verma",
                "mother_name": "Kareena Kapoor",
                "guardian_email": "adam@gmail.com",
                "guardian_phone": "9872827722",
                "is_active": true,
                "created_at": "2022-07-30T19:03:55Z",
                "updated_at": "2022-08-14T12:26:16Z"
            },
            {
                "id": 2,
                "parent_id": 1,
                "cid": 1,
                "name": "Akash Verma",
                "email": "akshay@gmail.com",
                "phone": "9878616222",
                "course_id": 1,
                "year": 2,
                "batch": "2019-22",
                "password": "$2a$14$A0kz4GVhr/K.2XZr6A0zeOaNaKZx.HCcvaXZgHC8xQs.msDru.bZS",
                "enrollment": "04914902019",
                "dob": "2022-08-11T17:10:48Z",
                "fathers_name": "Brijesh Verma",
                "mother_name": "Kareena Kapoor",
                "guardian_email": "adam@gmail.com",
                "guardian_phone": "9878616271",
                "is_active": true,
                "created_at": "2022-08-11T17:12:21Z",
                "updated_at": "2022-08-14T12:26:12Z"
            }
        ],
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
                "password": "$2a$14$/L8ug6Gb3Wh8G6/LT5AGYeuS.72Tmb6cLjUAot9T02DvGkt2R0miS",
                "is_active": true,
                "created_at": "2022-07-31T19:06:59Z",
                "updated_at": "2022-08-08T04:20:30Z"
            },
            {
                "id": 2,
                "cid": 1,
                "name": "Seema Kidwai",
                "email": "omgupta1608+3@gmail.com",
                "alt_email": "omgupta1608+4@gmail.com",
                "department": "Computer Applications",
                "course_id": 1,
                "designation": "Professor",
                "score": 83.98,
                "password": "$2a$14$x8BbpWEU.Vm8Kdypjq/HRu9QZr10AM9RCodg4ERCayflH.gLdnkXm",
                "is_active": true,
                "created_at": "2022-07-31T19:06:59Z",
                "updated_at": "2022-08-14T12:24:25Z"
            },
            {
                "id": 3,
                "cid": 1,
                "name": "Preeti Jain",
                "email": "omgupta1608+7@gmail.com",
                "alt_email": "omgupta1608+6@gmail.com",
                "department": "Computer Science",
                "course_id": 1,
                "designation": "Assistant Professor",
                "score": 91.08,
                "password": "$2a$14$aR1M8X2QDZHyrT4lRFx1eu5tZpyTbm65BXYch6kcipS2S3P/nbQlK",
                "is_active": true,
                "created_at": "2022-07-31T19:09:33Z",
                "updated_at": "2022-08-14T12:23:25Z"
            }
        ]
    },
    "message": "Linking Data"
}```

</details>
