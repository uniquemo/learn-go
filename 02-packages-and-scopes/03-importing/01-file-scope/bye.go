package main

// This file cannot see main.go's imported names ("fmt").
// Because the imported names belong to file scope.

// 报错：这里没有导入fmt会报错，因为导入操作的scope是文件级别的，即使main.go中导入了，这里也取不到
// func bye() {
// 	fmt.Println("Bye!")
// }
