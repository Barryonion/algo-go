package main

import (
	"fmt"
)

//对于一组不同重量、不可分割的物品，我们需要选择一些装入背包，在满足背包最大重量限
//制的前提下，背包中物品总重量的最大值是多少呢？
//weight = {1,2,3,4,5} , n = 5 , w = 16

//weight:物品重量，n:物品个数，w:背包可承受重量
//二维数组解法
func knapsack(weight []int, n, w int) int {
	states := make([][]bool, n)
	for i := 0; i < n; i++ {
		states[i] = make([]bool, w+1)
	}
	//初始化第一行数据，决策第i=0个物品
	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}
	//第i次循环，决策第i个物品 ：i可以理解成编号，而不是通常语义下的次序
	for i := 1; i < n; i++ {
		//j表示，当前阶段决策下背包的状态（其中物品的总重量）。
		for j := 0; j <= w; j++ { //不把第i个物品放入背包
			if states[i-1][j] == true { //上一阶段决策时的背包总重量为j，因不放入，所以背包总质量仍然为j
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ { //把第i个物品放入背包，注意不能超过背包重量w
			if states[i-1][j] == true {
				states[i][j+weight[i]] = true
			}
		}
	}
	for j := w; j >= 0; j-- {
		if states[n-1][j] == true {
			return j
		}
	}
	return 0
}

//仅用一维数组的解法
func knapsack2(weight []int, n, w int) int {
	states := make([]bool, w+1)
	//对第i=0个物品进行决策，不会将该物品放入背包
	states[0] = true
	//如果该物品的重量不超过背包总承重，则将物品放入背包
	if weight[0] < w {
		states[weight[0]] = true
	}
	// 依次对后续的物品进行决策
	for i := 1; i < n; i++ {
		//如果从左到右遍历的话，排在前面的计算结果也会被纳入排在后面的计算中来，
		//导致出现过多下标j对应的数组值states[j]为true的问题
		for j := w - weight[i]; j >= 0; j-- {
			if states[j] == true { //对上一个物品的决策结果中，包含背包重量是j的情况
				states[j+weight[i]] = true //把第i个物品放入背包
			}
		}
	}
	for j := w; j >= 0; j-- {
		if states[j] == true {
			return j
		}
	}
	return 0
}

//我们刚刚讲的背包问题，只涉及背包重量和物品重量。我们现在引入物品价值这一变量。对
//于一组不同重量、不同价值、不可分割的物品，我们选择将某些物品装入背包，在满足背包
//最大重量限制的前提下，背包中可装入物品的总价值最大是多少呢？

func knapsack3(value, weights []int, capacity int) int {
	states := make([][]int, len(weights))
	for i := 0; i < len(weights); i++ {
		states[i] = make([]int, capacity+1)
	}
	for i := 0; i < len(weights); i++ {
		for j := 0; j < capacity+1; j++ {
			states[i][j] = -1
		}
	}
	//对第i=0个物品进行决策，不放入背包
	states[0][0] = 0
	//判断是否放得下，若是，则放入背包
	if weights[0] <= capacity {
		states[0][weights[0]] = value[0]
	}
	//对剩下的物品进行决策
	for i := 1; i < len(weights); i++ {
		//不放入背包
		for j := 0; j <= capacity; j++ {
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}
		//放入背包
		for j := 0; j <= capacity-weights[i]; j++ {
			if states[i-1][j] >= 0 {
				v := states[i-1][j] + value[i]
				if v > states[i][j+weights[i]] {
					states[i][j+weights[i]] = v
				}
			}
		}
	}
	//找出最大值
	maxValue := -1
	for j := 0; j <= capacity; j++ {
		if states[len(weights)-1][j] > maxValue {
			maxValue = states[len(weights)-1][j]
		}
	}
	return maxValue
}

//带路径状态的问题 (0-1背包问题升级版（路径打印）)
//你要去环游世界了！如果可能的话，你肯定想往背包里塞无数东西，但航空公司有规定，行李不能超过一定重量。
//为了确保带上旅行中最有价值的物品，你决定给所有物品打分，来表示这些物品的价值。你想要让带的东西具备最大的价值。
//实现一个函数，来计算**背包能携带的最大物品价值**。
//函数参数由三个数组组成，
//第一个是得分，第二个是重量。两个参数总是等长的有效数组，因此不用验证输入。第三个参数是背包不能超过的最大重量。
//例如，给定这些输入:
//
//scores = [15, 10, 9, 5]
//weights = [1, 5, 3, 4]
//capacity = 8
//
//最高分为29，来自于第1、3和4项。

func main() {
	weight := []int{1, 2, 3, 4, 5}
	n := 5
	w := 16
	res := knapsack2(weight, n, w)
	fmt.Println(res)
}
