const express = require("express")
const app = express();
const fs = require('fs')
const csv = require('csv')
const { default: axios } = require('axios')
const multer = require('multer')
const upload = multer({ dest: '.' })
app.use(express.json())



const teachersMapping = Object.freeze({
    "Name": "name",
    "Official Email": "email",
    "Personal Email": "alt_email",
    "Department": "department",
    "Course": "course",
    "Designation": "designation"
});

const studentsMapping = Object.freeze({
    "Name": "name",
    "Email": "email",
    "Phone": "phone",
    "Course": "course",
    "Year": "year",
    "Batch": "batch",
    "Section": "section",
    "Enrollment Number": "enrollment_number",
    "Date of Birth": "dob",
    "Fathers Name": "fathers_name",
    "Mother Name": "mothers_name",
    "Guardian Email": "guardian_email",
    "Guardian Phone": "guardian_phone",
});

const Courses = Object.freeze({
    "BCA": "1",
    "B Tech.": "2",
    "BBA": "3",
    "B Com.": "4",
    "B Ed.": "5",
    "MCA": "6",
    "M Tech.": "7",
    "MBA": "8",
})

app.post("/onboard-college", upload.any(), async (req, res) => {
    // students
    // console.log(req.files)
    let fileRows = [];
    let teacherFile = req['teacher_file']
    let studentFile = req['student_file']

    req.files.forEach(f => {
        if (f.fieldname == "teacher_file") {
            teacherFile = f;
        } else if (f.fieldname == "student_file") {
            studentFile = f
        }
    });


    let mainStudents, mainTeachers;
    fs.createReadStream(studentFile.path)
        .pipe(csv.parse({ headers: true }))
        .on("error", (error) => {
            res.send(error.message);
        })
        .on("data", (row) => {
            // console.log(row)
            fileRows.push(row)
        })
        .on("end", async () => {
            let keys = [], student = {}, students = []
            fileRows.forEach((fr, i) => {
                if (i == 0) {
                    fr.forEach(d => {
                        keys.push(d)
                    })
                } else {
                    fr.forEach((d, i) => {
                        student[studentsMapping[keys[i]].trim()] = d.trim();
                    })
                    student.course_id = Courses[student.course]
                    delete student['course']

                    students.push(student)
                }

            })
            mainStudents = students



            // teachers
            fs.createReadStream(teacherFile.path)
                .pipe(csv.parse({ headers: true }))
                .on("error", (error) => {
                    res.send(error.message);
                })
                .on("data", (row) => {
                    // console.log(row)
                    fileRows.push(row)
                })
                .on("end", async () => {
                    let keys = [], student = {}, teachers = []
                    fileRows.forEach((fr, i) => {
                        if (i == 0) {
                            fr.forEach(d => {
                                keys.push(d)
                            })
                        } else {
                            fr.forEach((d, i) => {
                                student[teachersMapping[keys[i]].trim()] = d.trim();
                            })
                            student.course_id = Courses[student.course]
                            // delete student['course']
                            console.log(student)
                            teachers.push(student)
                        }

                    })
                    mainTeachers = teachers;

                    console.log({
                        students: mainStudents,
                        teachers: mainTeachers
                    })

                    let data = await axios.post('http://localhost:5000/api/v1/onboard/js', {
                        students: mainStudents,
                        teachers: mainTeachers
                    }, {
                        headers: {
                            "hello": "edrank"
                        }
                    })

                    if (data && data.data) {
                        console.log(data.data)
                        res.send(data.data)
                    } else {
                        res.send(data)
                    }
                });
        });

})

app.listen(5003, () => {
    console.log('Server Up')
})