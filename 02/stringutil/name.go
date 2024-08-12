package stringutil //表示该代码属于 stringutil 这个包。包的名称通常与其功能或内容有关。s

// MyName will be exported because it starts with a capital letter.
var MyName = "Todd" //定义了一个名为 MyName 的变量，并赋值为 "Todd"
//变量前的注释说明了这个变量的可导出性：由于它的名称首字母是大写的，MyName 将被导出。这意味着 MyName 可以被其他包访问

/*拓展：在 Go 语言中，任何首字母大写的标识符
（如变量、函数、类型等）都被视为导出的，
其他包可以使用它。而首字母小写的标识符则是包内部私有的
不能被外部包访问*/
