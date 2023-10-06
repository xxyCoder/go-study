# HTML渲染
- 使用LoadHTMLGlob(pattern string)或者LoadHTMLFiles(files ...string)
- 对于要使用不同目录下的模板，需要在首尾添加一定内容，不然无法找到对应文件

# JSONP
- 如果查询参数存在回调函数，则将回调函数添加到响应体中