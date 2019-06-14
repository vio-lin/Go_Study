# Go_Study
A repository for study go language

# 2.顺序编程
# 2.1 变量
```
// 声明
var v1 int
var v2 string
var v3 [10]int // 数组
var v4 []int // 数组切片
var v5 struct {
 f int
}
var v6 *int // 指针
var v7 map[string]int // map，key为string类型，value为int类型
var v8 func(a int) int
// 赋值
var v1 int = 10 // 正确的使用方式1
var v2 = 10 // 正确的使用方式2，编译器可以自动推导出v2的类型
v3 := 10 // 正确的使用方式3，编译器可以自动推导出v3的类型
// 多值交换
t = i; i = j; j = t; 
// 匿名变量
func GetName() (firstName, lastName, nickName string) {
 return "May", "Chan", "Chibi Maruko"
} 
_, _, nickName := GetName() 
```
# 2.2 常量
```
// 字面的常量
-12
3.14159265358979323846 // 浮点类型的常量
3.2+12i // 复数类型的常量
true // 布尔类型的常量
"foo" // 字符串常量

// 常量的定义
const Pi float64 = 3.14159265358979323846
const zero = 0.0 // 无类型浮点常量
const (
 size int64 = 1024
 eof = -1 // 无类型整型常量
)
const u, v float32 = 0, 3 // u = 0.0, v = 3.0，常量的多重赋值
const a, b, c = 3, 4, "foo"
// a = 3, b = 4, c = "foo", 无类型整型和字符串常量
```
# 2.2.3 预定义常量
Go语言预定义了这些常量：true、false和iota。 
iota：在const中进行值进行充值 在之间每次出现都会导致+1
```
const ( // iota被重设为0
 c0 = iota // c0 == 0
 c1 = iota // c1 == 1
 c2 = iota // c2 == 2
) 

//单个后移
const (
 a = 1 << iota // a == 1 (iota在每个const开头被重设为0)
 b = 1 << iota // b == 2
 c = 1 << iota // c == 4
)
const (
 u = iota * 42 // u == 0
 v float64 = iota * 42 // v == 42.0 
 w = iota * 42 // w == 84
)
 
// 重置了
const x = iota // x == 0 (因为iota又被重设为0了)
const y = iota // y == 0 (同上) 
 
// 多个可以省略
const ( // iota被重设为0
 c0 = iota // c0 == 0
 c1 // c1 == 1
 c2 // c2 == 2
)
const (
 a = 1 <<iota // a == 1 (iota在每个const开头被重设为0)
 b // b == 2
 c // c == 4
) 
```
# 2.2.3 枚举
```
// 大写的变量名可以在范围之外访问
const (
 Sunday = iota
 Monday
 Tuesday
 Wednesday
 Thursday
 Friday
 Saturday
 numberOfDays // 这个常量没有导出
) 
```
## 2.3 类型

类型定义 | 描述
---|---
布尔类型 | bool
整型 | int8、byte、int16、int、uint、uintptr
浮点类型 | float32、float64
复数类型 | complex64、complex128
字符串 | string
字符类型 | rune
错误类型 | error
指针| pointer
数组|array
切片|slice
字典|map
通道|chan
结构体|struct
接口|interface

### 2.3.2 整形
类 型|长度（字节）| 值 范 围
---|---|---
int8 |1 |128 ~ 127
uint8（即byte）| 1| 0 ~ 255
int16 |2| 32 768 ~ 32 767
uint16 |2| 0 ~ 65 535
int32 |4| 2 147 483 648 ~ 2 147 483 647
uint32 |4| 0 ~ 4 294 967 295
int64 |8| 9 223 372 036 854 775 808 ~ 9 223 372 036 854 775 807
uint64 |8| 0 ~ 18 446 744 073 709 551 615
int |平台相关| 平台相关
uint| 平台相关| 平台相关
uintptr |同指针 |在32位平台下为4字节，64位平台下为8字节

运算符号+、、*、/和%  比较符号 >、<、==、>=、<=和!=
位运算|结果
---|---
x << y 左移 |124 << 2 // 结果为496
x >> y 右移 |124 >> 2 // 结果为31
x ^ y 异或 |124 ^ 2 // 结果为126
x & y 与 |124 & 2 // 结果为0
x|y 或 |124 | 2 // 结果为126
^x |取反 ^2 // 结果为3 

### 2.3.3 浮点
1. 浮点数表示
Go语言定义了两个类型float32和float64，其中float32等价于C语言的float类型，
float64等价于C语言的double类型。
> == 会出现问题 最好还是 a-b<p(误差值)

### 2.3.4 复数类型
```
//定义的方法
var value1 complex64 // 由2个float32构成的复数类型
value1 = 3.2 + 12i
value2 := 3.2 + 12i // value2是complex128类型
value3 := complex(3.2, 12) // value3结果同 value2 
```
### 2.3.6 字符类型
关于rune相关的操作，可查阅Go标准库的unicode包。另外unicode/utf8包也提供了
UTF8和Unicode之间的转换。、
### 2.3.7 数组
```
[32]byte // 长度为32的数组，每个元素为一个字节
[2*N] struct { x, y int32 } // 复杂类型数组
[1000]*float64 // 指针数组
[3][5]int // 二维数组
[2][2][2]float64 // 等同于[2]([2]([2]float64)) 

//数组遍历的
for i, v := range array {
 fmt.Println("Array element[", i, "]=", v)
} 
```
