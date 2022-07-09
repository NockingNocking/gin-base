
## go模板语法比较符

* eq  如果 arg1 == arg2 则返回真
* ne  如果 arg1 != arg2 则返回真
* lt  如果 arg1 < arg2 则返回真
* le  如果 arg1 <= arg2 则返回真
* gt  如果 arg1 > arg2 则返回真
* ge  如果 arg1 >= arg2 则返回真

## go模板语法条件判断语法

{{if gt .score 60}}
及格
{{else}}
不及格
{{end}}

## go模板语法循环遍历

{{range $key,$value := .arr}}
{{$key}}----{{$value}}
{{end}}
