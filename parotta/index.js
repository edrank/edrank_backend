const express = require("express")
const app = express();
const fs = require('fs')
const csv = require('csv')
const { default: axios } = require('axios')
const multer = require('multer')
const upload = multer({ dest: './tmp' })
app.use(express.json())
const bcrypt = require('bcrypt')



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


    let mainStudents, mainTeachers, mainParents = {};
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
            // fileRows.forEach(async (fr, i) => {
            let i = 0;
            for (let fr of fileRows) {
                // console.log(fr)
                if (i == 0) {
                    fr.forEach(d => {
                        keys.push(d)
                    })
                } else {
                    fr.forEach((d, i) => {
                        // console.log(keys, keys[i], i, fr.length, fileRows.length)
                        student[studentsMapping[keys[i]].trim()] = d.trim();
                    })

                    let hash1 = await bcrypt.hash(`${student['name']}@123`, 10)
                    student.course_id = Courses[student.course]
                    student.password = hash1
                    student.is_active = true
                    student['dob'] = new Date(student['dob'])
                    student['year'] = parseInt(student['year'])
                    student['course_id'] = parseInt(student['course_id'])
                    delete student['course']
                    // console.log(student)
                    students.push(student)
                    let hash = await bcrypt.hash(`${student['fathers_name']}@123`, 10)
                    mainParents[student['enrollment_number']] = {
                        name: student['fathers_name'],
                        email: student['guardian_email'],
                        phone: student['guardian_phone'],
                        password: hash,
                        is_active: true
                    }
                }
                i++;

            }
            mainStudents = students

            // teachers
            let fileRows1 = []
            fs.createReadStream(teacherFile.path)
                .pipe(csv.parse({ headers: true }))
                .on("error", (error) => {
                    res.send(error.message);
                })
                .on("data", (row) => {
                    // console.log(row)
                    fileRows1.push(row)
                })
                .on("end", async () => {
                    let keys1 = [], student = {}, teachers = []
                    // fileRows1.forEach(async (fr, i) => 
                    let i = 0
                    for (let fr of fileRows1) {
                        if (i == 0) {
                            fr.forEach(d => {
                                keys1.push(d)
                            })
                        } else {
                            fr.forEach((d, i) => {
                                // console.log("11111111111111111", fr, keys1, keys1[i])
                                student[teachersMapping[keys1[i]].trim()] = d.trim();
                            })

                            let hash2 = await bcrypt.hash(`${student['name']}@123`, 10)
                            student.course_id = Courses[student.course]
                            student.is_active = true
                            student['course_id'] = parseInt(student['course_id'])
                            student['password'] = hash2
                            delete student['course']
                            // console.log(student)
                            teachers.push(student)
                        }
                        i++
                    }

                    mainTeachers = teachers;

                    console.log({
                        students: mainStudents,
                        teachers: mainTeachers
                    })

                    try {
                        let data = await axios.post('http://localhost:5000/api/v1/onboard/js', {
                            students: mainStudents,
                            teachers: mainTeachers,
                            parents: mainParents,
                            cid: req.body.cid
                        }, {
                            headers: {
                                "hello": "edrank"
                            }
                        })

                        if (data && data.data) {
                            // console.log(data.data)
                            res.send(data.data)
                        } else {
                            res.send(data)
                        }
                    } catch (err) {
                        console.log(err)
                        res.send(err)
                    }

                });

        });
})

app.listen(5003, () => {
    console.log('Server Up')
})