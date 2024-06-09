## 指针传递的好处
传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 return 返回
## defer关键字
关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
defer 和 return 的执行顺序是先为返回值赋值，然后执行 defer，然后 return 到函数调用处。
关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。
关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如
```
// 关闭文件流
defer file.close()
```
```
//解锁一个加锁的资源
mu.Lock()
defer mu.Unlock()
```
```
//关闭数据库连接
defer disconnectFromDB()
```