# Go开发Web应用（2）


| enctype | | |
| :-----| :-----| :-----|
| application/x-www-form-urlencoded |Form数据以name-val对形式编码至QueryString | |
| multipart/form-data |Form数据的name-val转换成MIME消息| |
| text/plain | | |

## 获取表单数据
通过表单发送数据，数据都是以键值对的形式发送出去的。当服务器接收到请求时，请求走到相应的`Handler`后，从`Handler`的`Request`参数中可以获取表单数据，获取的方法针对不同的表单形式有几种。  
首先要先调用`ParseForm`或`ParseMultipartForm`函数先解析`Request`中的数据，随后便可以使用三个字段来读取表单数据。如下
* Form
* PostForm
* MultipartForm

#### Form
