package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		//*************读取 Form 数据 ：********************
		// r.ParseForm()
		// fmt.Fprintln(w, "读值方式1")
		// fmt.Fprintln(w, fmt.Sprintf("user name : %v", r.Form["userName"]))
		// fmt.Fprintln(w, fmt.Sprintf("email : %v", r.Form["email"]))
		// fmt.Fprintln(w, fmt.Sprintf("checkbox checked : %v", r.Form["checkOut"]))

		// fmt.Fprintln(w, "读值方式2")
		// fmt.Fprintln(w, fmt.Sprintf("user name : %v", r.FormValue("userName")))
		// fmt.Fprintln(w, fmt.Sprintf("email : %v", r.FormValue("email")))
		// fmt.Fprintln(w, fmt.Sprintf("checkbox checked : %v", r.FormValue("checkOut")))
		//**********************************************

		//*************读取 PostForm 数据 ：********************
		// r.ParseForm()
		// fmt.Fprintln(w, "使用Form读取: ")
		// fmt.Fprintln(w, r.Form)
		// fmt.Fprintln(w, "使用PostForm读取: ")
		// fmt.Fprintln(w, r.PostForm)
		//**********************************************

		//*************读取 MultipartForm 数据 ：********************
		//先解析
		r.ParseMultipartForm(1024 * 1024)
		// 再试试PostFormValue方法
		fmt.Fprintln(w, fmt.Sprintf("user name : %v", r.PostFormValue("userName")))
		fmt.Fprintln(w, fmt.Sprintf("avatar : %v", r.PostFormValue("avatar")))
		// 读取MultipartForm
		fmt.Fprintln(w, fmt.Sprintf("user name : %v", r.MultipartForm.Value["userName"]))
		// 取出文件保存到本地
		if avatar, ok := r.MultipartForm.File["avatar"]; ok {
			fh := avatar[0]
			file, err := fh.Open()

			filebuf := make([]byte, fh.Size)
			if err == nil {
				_, err := file.Read(filebuf)
				if err == nil {
					file, err := os.Create("avatar.jpg")
					defer file.Close()
					if err == nil {
						file.Write(filebuf)
						fmt.Fprintln(w, "avatar : saved")
					}
				}
			}
		} else {
			fmt.Fprintln(w, "avatar : no file")
		}

		//**********************************************
	})

	http.ListenAndServe("localhost:8080", nil)
}
