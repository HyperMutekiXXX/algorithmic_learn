# 643. 子数组最大平均数 I
```
简单

给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。

请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。

任何误差小于 10-5 的答案都将被视为正确答案。
```


示例 1：
```
输入：nums = [1,12,-5,-6,50,3], k = 4
输出：12.75
解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
```
示例 2：
```
输入：nums = [5], k = 1
输出：5.00000
```

# 解题思路
```
使用滑动窗口，主要问题是在求每一段的和，数据量大的话，直接for循环求和会超时。
设i为右边界，sum[i]则是这一段的和
sum[i] = num[i]+...+num[i-k+1]
sum[i-1] = num[i-1]+...+num[i-k]
sum[i]-sum[i-1] = num[i] - num[i-k]
sum[i] = sum[i-1] + num[i] - num[i-k]

所以先求出第一段的和，再遍历剩下段并进行对比即可得出结果
```