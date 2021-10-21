package main

import ("fmt")

func main() {
    ClearScreen()
    SetProjectName()
    SetProjectPath()

    if CheckDir("'" + GetProjectPath() + "'") {
        ProjectMkDir(GetProjectPath(), GetProjectName())
        if FilesMkDir() && CreateFiles() {
            fmt.Println("Проект успешно создан!")
        } else {
            fmt.Println("Что-то пошло не так!")
        }
    } else {
        fmt.Println("Директория не найдена!")

        i := 1

        for i < 3 {
            SetProjectName()

            if CheckDir("'" + GetProjectPath() + "'") {
                break
            } else {
                fmt.Println("Директория не найдена!")
            }
        }
    }
}