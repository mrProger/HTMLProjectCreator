package main

import "fmt"

func main() {
    ClearScreen()
    SetProjectName()
    SetProjectPath()

    if CheckDir("'" + GetProjectPath() + "'") {
        ProjectMkDir(GetProjectPath(), GetProjectName())
        if FilesMkDir() && CreateFiles() {
            fmt.Println("\nПроект успешно создан!")
            fmt.Println("\nНажмите любую клавишу, чтобы продолжить...")
            fmt.Scanln()
        } else {
            fmt.Println("Что-то пошло не так!")
            fmt.Println("\nНажмите любую клавишу, чтобы продолжить...")
            fmt.Scanln()
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

    fmt.Scanln()
}