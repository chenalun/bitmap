# bitmap
GoLang bitmap 位图算法

## 基本原理
用一个bit来标记某个元素对应的value，key就是该元素，使用bit存储大大节省空间

## 适用场景
1. 20亿个int数据去重排序
2. 判断指定的数据是否在20亿数据中
3. 海量密集型数据

> 注1：换算方式，GB=>MB=>KB=>Byte=>Bit
>
> 注2：存储20亿int数据所需内存：(2000000000 * 4)/(1024^3) ≈ 7.45G，1(int) = 4(byte)
>
> 注3：使用bit存储 2000000000 / 8 /(1024^3)≈ 0.23G，1(byte)=8(bit) 1byte可以存8位二进制
