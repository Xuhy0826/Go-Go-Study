# Go开发Web应用（2）

## 获取表单数据
通过表单发送数据，数据都是以键值对的形式发送出去的。当服务器接收到请求时，请求走到相应的`Handler`后，从`Handler`的`Request`参数中可以获取表单数据，获取的方法针对不同的表单形式有几种。  
首先要先调用`ParseForm`或`ParseMultipartForm`函数先解析`Request`中的数据，随后便可以使用三个字段来读取表单数据。如下
* Form
* PostForm
* MultipartForm  

前端代码先放一个Form，包含三个`<input>`标签。
```
<form action="http://localhost:8080/register" method="post"
    enctype="application/x-www-form-urlencoded">
    <div class="form-group">
        <label for="emailTxt">Email address: </label>
        <input type="email" class="form-control" name="email" id="emailTxt" placeholder="Email">
    </div>
    <div class="form-group">
        <label for="userNameTxt">User Name</label>
        <input type="text" class="form-control" name="userName" id="userNameTxt" placeholder="User Name">
    </div>
    <div class="checkbox">
        <label>
            <input type="checkbox" name="checkOut"> Check me out
        </label>
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
</form>
```

#### Form
`Form`字段是`map[string][]string`类型，map的key是input标签的name属性，map的value是string切片类型，就以上的前端代码，存放的就是input标签的value属性。  
读取前端提交来的Form中的数据可以在解析完后（r.ParseForm()），访问`Request`的`Form`字段或者使用`Request`的`FormValue()`方法来读取相应的值，`FormValue()`方法会返回切片的第一个值，即`request.Form["userName"][0]`和`request.FormValue("userName")`是等效的。
```
func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		//解析
		r.ParseForm()

		fmt.Fprintln(w, "读值方式1")
		fmt.Fprintln(w, fmt.Sprintf("user name : %v", r.Form["userName"]))
		fmt.Fprintln(w, fmt.Sprintf("email : %v", r.Form["email"]))
		fmt.Fprintln(w, fmt.Sprintf("checkbox checked : %v", r.Form["checkOut"]))

		fmt.Fprintln(w, "读值方式2")
		fmt.Fprintln(w, fmt.Sprintf("user name : %v", r.FormValue("userName")))
		fmt.Fprintln(w, fmt.Sprintf("email : %v", r.FormValue("email")))
		fmt.Fprintln(w, fmt.Sprintf("checkbox checked : %v", r.FormValue("checkOut")))
	})

	http.ListenAndServe("localhost:8080", nil)
}
```
> 提交后看到的输出：
```
读值方式1
user name : [123]
email : [583209544@qq.com]
checkbox checked : [on]
读值方式2
user name : 123
email : 583209544@qq.com
checkbox checked : on
```
最常见的读取表单中的数据就是这样，但是如果在表单提交的地址中加上QueryString，并且QueryString和表单的键值对中存在同样的键时，`request.Form`中的值即string切片中便会存在两个值了，其中表单的值会在前。然而如果读取表单数据用的不是`request.Form`而是`request.PostForm`，便不会读取到除了表单意外的数据。

#### PostForm
如果现将表单的action改为`action="http://localhost:8080/register?userName=Admin"`，即QueryString中增加了一个与表单相同的key即userName。那么使用`request.Form`和`request.PostForm`读取数据的区别可见。
```
func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		//解析
		r.ParseForm()

		fmt.Fprintln(w, "使用Form读取: ")
		fmt.Fprintln(w, r.Form)

		fmt.Fprintln(w, "使用PostForm读取: ")
		fmt.Fprintln(w, r.PostForm)
	})

	http.ListenAndServe("localhost:8080", nil)
}
```
> 提交后看到的输出：
```
使用Form读取: 
map[checkOut:[on] email:[583209544@qq.com] userName:[123 Admin]]
使用PostForm读取: 
map[checkOut:[on] email:[583209544@qq.com] userName:[123]]
```
`r.PostForm`中没有显示QueryString中的值。与`Request`的`FormValue()`方法类似，`Request`中还有一个`PostFormValue()`方法，即是读取PostForm中值的方法。  
再改变一下条件，如果表单数据要上传不仅仅是字符串数据，还要上传文件，那么表单的编码类型即`encType`就不能是默认的`application/x-www-form-urlencoded`而要改成`multipart/form-data`。但是这么一改会发现，是`request.Form`或者`request.PostForm`读取不到任何表单数据，通过`request.PostFormValue()`只能读取string数据，上传的文件数据则无法读取。这种情况下就需要使用`request.MultipartForm`来读取了。

#### MultipartForm
与`request.Form`不同，使用`request.MultipartForm`前的解析方法需换为使用`r.ParseMultipartForm()`来解析。
> ParseMultipartForm的函数签名
```
func (r *Request) ParseMultipartForm(maxMemory int64) error
```
`r.ParseMultipartForm()`需要传入一个int值来设定要读取的数据的长度。`request.MultipartForm`也不再是一个map类型，而是`multipart.Form`的指针类型，`multipart.Form`其实是一个struct。
> formdata.go 中包含 multipart.Form 的类型声明
```
type Form struct {
	Value map[string][]string
	File  map[string][]*FileHeader
}
```
也就是说`request.MultipartForm`可访问两个字段，并且都是map类型。
1. Value：放的都是表单里的字符串数据，那功能就和之前的r.PostForm差不多
2. File：放的是文件数据  
现将`<form>`中的编码类型改为`enctype="multipart/form-data`，并添加一个`<input id="avatarFile" type="file" name="avatar"/>`来上传文件
```
<form action="http://localhost:8080/register?userName=admin" method="post"
enctype="multipart/form-data">
	<div class="form-group">
		<label for="emailTxt">Email address: </label>
		<input id="emailTxt" type="email" class="form-control" name="email" placeholder="Email">
	</div>
	<div class="form-group">
		<label for="userNameTxt">User Name</label>
		<input id="userNameTxt" type="text" class="form-control" name="userName"
			placeholder="User Name">
	</div>
	<div class="checkbox">
		<label>
			<input type="checkbox" name="checkOut"> Check me out
		</label>
	</div>
	<div class="form-group">
		<label for="avatarFile" class="btn btn-secondary">Avatar</label>
		<input id="avatarFile" type="file" name="avatar" class="d-none"/>
	</div>
	<button type="submit" class="btn btn-primary">Submit</button>
</form>
```
在服务端从MultipartForm.File中读取文件可使用Open方法。将表单中的文件数据保存到本地，下面有个示例。
```
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		//先解析
		r.ParseMultipartForm(1024 * 1024)

		// 再试试PostFormValue方法
		fmt.Fprintln(w, fmt.Sprintf("user name : %v", r.PostFormValue("userName")))	//可读取到值
		fmt.Fprintln(w, fmt.Sprintf("avatar : %v", r.PostFormValue("avatar")))		//无法读取

		// 读取MultipartForm
		fmt.Fprintln(w, fmt.Sprintf("user name : %v", r.MultipartForm.Value["userName"]))
		// 取出文件保存到本地
		fh := r.MultipartForm.File["avatar"][0]
		// 创建读取文件的缓冲区
		filebuf := make([]byte, fh.Size)
		// 得到文件
		file, err := fh.Open()
		if err == nil {
			_, err := file.Read(filebuf)
			if err == nil {
				file, err := os.Create("avatar.jpg")
				defer file.Close()
				if err == nil {
					file.Write(filebuf)
				}
			}
		}
	})

	http.ListenAndServe("localhost:8080", nil)
}
```
其实如果单纯的读取上传的单独文件，可以直接使用`r.FormFile()`来获取。即可以将
```
fh := r.MultipartForm.File["avatar"][0]
file, err := fh.Open()
```
换成下面的写法。
```
file, _, err := r.FormFile("avatar")
```
