package main

import "fmt"
import "runtime"
import "os"
import "os/exec"

var ProjectName string
var ProjectPath string

func ClearScreen() {
	os_ := runtime.GOOS

    switch os_ {
        case "windows":
	        cmd := exec.Command("cmd", "/c", "cls")
	        cmd.Stdout = os.Stdout
	        cmd.Run()

        case "linux":
            cmd := exec.Command("clear")
            cmd.Stdout = os.Stdout
            cmd.Run()
            
        case "darwin":
            cmd := exec.Command("clear")
            cmd.Stdout = os.Stdout
            cmd.Run()
    }
}

func SetProjectName() bool {
    var projectName = ""

    fmt.Println("Введите имя проекта:")
    fmt.Scan(&projectName)

    if projectName != "" {
        ProjectName = projectName
        return true
    } else {
        return false
    }
}

func GetProjectName() string {
    return ProjectName
}

func SetProjectPath() bool {
    var projectPath = ""

    fmt.Println("\nВведите путь для создания проекта:")
    fmt.Scan(&projectPath)

    if projectPath != "" {
        ProjectPath = projectPath
        return true
    } else {
        return false
    }
}

func GetProjectPath() string {
    return ProjectPath
}

func CheckDir(path string) bool {
    file, err := os.Stat(path)
    _ = file

    if os.IsNotExist(err) {
        return false
    } else {
        return true
    }
}

func ProjectMkDir(projectPath string, projectName string) bool {

    if !CheckDir(ProjectPath + "/" + ProjectName) {
        err := os.Mkdir(projectPath + "/" + projectName, 0755)

        if err != nil {
            return false
        } else {
            return true
        }
    } else {
        os.RemoveAll(ProjectPath + "/" + ProjectName + "/")
        ProjectMkDir(ProjectPath, ProjectName);

        return false
    }
}

func FilesMkDir() bool {
    err1 := os.Mkdir(ProjectPath + "/" + ProjectName + "/" + "view", 0755)
    err2 := os.Mkdir(ProjectPath + "/" + ProjectName + "/" + "js", 0755)
    err3 := os.Mkdir(ProjectPath + "/" + ProjectName + "/" + "view" + "/" + "images", 0755)

    if err1 != nil && err2 != nil && err3 != nil{
        return false
    } else {
        return true
    }
}

func CreateFiles() bool {
    f1, err1 := os.Create(ProjectPath + "/" + ProjectName + "/" + "index.html")

    f1.WriteString("<!DOCTYPE html>")
    f1.WriteString("\n<html>")
    f1.WriteString("\n<head>")
    f1.WriteString("\n  <meta charset='UTF-8'>")
    f1.WriteString("\n  <meta http-equiv='X-UA-Compatible' content='IE=edge'>")
    f1.WriteString("\n  <meta name='viewport' content='width=device-width, initial-scale=1.0'>")
    f1.WriteString("\n  <title>Document</title>")
    f1.WriteString("\n</head>")
    f1.WriteString("\n<body>")
    f1.WriteString("\n")
    f1.WriteString("\n</body>")
    f1.WriteString("\n</html>")

    f1.Close()

    f2, err2 := os.Create(ProjectPath + "/" + ProjectName + "/" + "view" + "/" + "style.css")
    f2.Close()
    f3, err3 := os.Create(ProjectPath + "/" + ProjectName + "/" + "js" + "/" + "script.js")
    f3.Close()

	if err1 != nil && err2 != nil && err3 != nil {
		return false
	} else {
        return true
    }
}