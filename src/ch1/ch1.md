# 数据类型

## 一、值类型
> 1. 内存通常在栈中分配
> 2. 变量指向内存的值 
> 3. 赋值时是进行值拷贝

- int系列
    - uint8、uint16、uint32、uint64
        - uint8:无符号 8 位整型 (0 到 255)
    - int8、int16、int32、int64
        - 有符号 8 位整型 (-128 到 127)
- float系列
  - float32、float64
        - IEEE-754  32/64位浮点型数
  - complex64、complex128
        - 32 /64 位实数和虚数
- bool
- string
- 数组
- 结构体

## 二、值类型

> 1. 变量指向地址，地址对应的空间里存储值 
> 2. 内存地址称之为指针 
> 3. 内存通常在堆中分配 
> 4. 通过与符号&获取地址（&num）

- 指针
- 切片slice
- 管道channel
- 接口interface
- map
- 函数func

## 三、其它数字类型

- byte, 类似 uint8
- rune, 类似 int32
- uint, 32 或 64 位
- int, 与 uint 一样大小
- uintptr, 无符号整型，用于存放一个指针
