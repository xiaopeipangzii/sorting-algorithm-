# 归并排序
归并排序是分治法的典型应用

通过递归或者迭代将序列分成一个个的子序列, 再对子序列进行排序

### 策略
1. 分 把原问题分解成若干子问题
2. 治 (递归)解决子问题
3. 合 合并子问题的解


### 优化点
除了自顶向下的递归实现, 还有一种代码量比较少的自底向上的循环方法实现
首先进行两两归并, 然后是四四归并, 然后是八八归并...
每一轮归子数组的大小会翻倍, 子数组的数量会减少

### 复杂度
|  平均时间复杂度  |  最好情况   |   最坏情况   | 空间复杂度 |
|     ----      |    ----    |    ----     | ---- |
|    n log n    |   n log n  |   n log n   |   n  |

### 稳定性
1. 稳定

