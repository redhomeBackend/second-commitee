# 2048
## 主要逻辑说明
### 初始化游戏
程序开始时，创建一个4x4的数组来表示游戏界面，并将所有元素初始化为0。接着，在随机位置生成两个包含数字2或4的初始块，进行操作。
### 游戏循环
游戏进入主循环，等待用户输入指令。根据用户的输入，执行相应的操作（如向上、向下、向左或向右移动），然后更新游戏界面。在每次更新后，检查游戏是否已经胜利（即是否出现了“2048”方块）或者游戏是否无法继续进行（即无法再移动或合并方块），若满足这些条件之一，则结束游戏.
## 关键函数
`Creatnumber()`: 在空位置上生成一个随机生成的数字（2或4）。

`TwoorFour(a int) int`: 基于概率生成一个随机数，可能是2或4,2为90%,4为10%。

`mergeUp()、mergeDown()、mergeLeft()、mergeRight()`: 这些函数实现了每个方向的合并和移动。

`Action()`: 用户输入并执行相应的移动操作。

`canPerformUp()、canPerformDown()、canPerformLeft()、canPerformRight()`: 这些函数检查当前游戏状态下是否可以执行移动。

`check(arr [4][4]int) bool`: 通过验证是否存在可以移动或是否已达到2048来检查游戏是否结束。
## 运行效果
### 失败案例
1.1 当无法执行该操作,但其他方向可进行时

<img width="174" alt="屏幕截图 2024-01-04 122233" src="https://github.com/Z-eng-zi/second-commitee/assets/146848093/0a2987a9-eb50-42d3-b09b-4b92e1db00bb">

 1.2 当游戏失败时


<img width="190" alt="屏幕截图 2024-01-04 122352" src="https://github.com/Z-eng-zi/second-commitee/assets/146848093/568fd148-5551-4f63-9ee9-6e5ef99d0bc7">

### 成功案例
由于2048数字太大,故改为出现32方块时游戏胜利

<img width="181" alt="屏幕截图 2024-01-04 122533" src="https://github.com/Z-eng-zi/second-commitee/assets/146848093/2c8098f3-1458-4cee-b1a3-c5899734e630">
