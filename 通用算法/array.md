# leetcode 刷题记录之数组操作
## 给定一个非负整数数组 A，返回一个由 A 的所有偶数元素组成的数组，后面跟 A 的所有奇数元素。

- 你可以返回满足此条件的任何数组作为答案。示例：

  ```
    输入：[3,1,2,4]
    输出：[2,4,3,1]
    输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。
   ```

- 提示：
    ```
    1 <= A.length <= 5000
    0 <= A[i] <= 5000
    ```
- 我最开始的思路
    ```php
    function sortArrayByParity($A) {
        $len = count($A);
        $res = new SplFixedArray($len);
        $j = 0;
        $k = 0;
        for($i = 0; $i< $len ; $i++ ){
            if($A[$i] % 2 === 0){
                $res[$k] = $A[$i];
                $k++;
            } else {
                $res[$len-1-$j] = $A[$i];
                $j++;
            }
        }
        return (array)$res;
    }
    // 执行时间 48ms
    ```
- 但是碰到了两个 比这个还快的代码 一个用了 `array_merge` 我就不说什么了，没意思， 但是另一个 提供了 一种判定 奇偶数的新方式
- 按位运算：与1按位运算等于0，为偶数。如果等于1，为奇数
    ```php
    function sortArrayByParity($A) {
        $len = count($A);
        $res = new SplFixedArray($len);
        $j = 0;
        $k = 0;
        for($i = 0; $i< $len ; $i++ ){
            if(($A[$i] & 1) === 0){
                $res[$k] = $A[$i];
                $k++;
            } else {
                $res[$len-1-$j] = $A[$i];
                $j++;
            }
        }
        return (array)$res;
    }
    // 执行时间 40ms
    ```
- 继续这次我们使用双指针进行来回同时遍历
    ```php
    function sortArrayByParity($A) {
        $start = 0;
        $end = count($A) - 1;
        while($start < $end){
            $is_open = 1;

            if(($A[$start] & 1) === 0){
                $start++;
                $is_open = 0;
            }

            if(($A[$end] & 1) === 1){
                $end--;
                $is_open = 0;
            }

            if($is_open === 1){
                $tem = $A[$start];
                $A[$start] = $A[$end];
                $A[$end] = $tem;
                $start++;
                $end--;
            }
        }
        return $A;
    }
    // 执行时间 40ms;
    ```
- 对双指针进行进一步优化,上面的单次只能移动一格，对于连续的偶数或者奇数提升不大
    ```php
    function sortArrayByParity($A) {
        $start = 0;
        $end = count($A) - 1;
        while($start < $end){

            while ($start < $end && ($A[$start] & 1) === 0){
                $start++;
            }

            while($start < $end && ($A[$end] & 1) === 1){
                $end--;
            }

            $tem = $A[$start];
            $A[$start] = $A[$end];
            $A[$end] = $tem;
            $start++;
            $end--;
        }
        return $A;
    }
    // 执行时间 36ms;
    ```
- 总结
    - 对于索引数组 操作来说 使用 这个 `new SplFixedArray($len)` 要比普通数组快
    - 对于奇偶性 判定 ($num%2 === 0) 实际上要比 ($num & 1 === 0) 慢一些 但是 好理解一点
    - 对于 数组的遍历来说 有时候可以考虑 头尾 同时开始

- golang 代码示例
    ```golang
    func sortArrayByParity(A []int) []int {
        start := 0
        end := len(A) - 1
        for start < end {

            for start < end && A[start]%2 == 0 {
                start++;
            }

            for start < end && A[end]%2 == 1 {
                end--;
            }

            A[start], A[end] = A[end], A[start]
            start++;
            end--;
        }
        return A
    }
    ```
- 好久没看golang 都快忘的差不多了 (0~0)
    ```golang
    func sortArrayByParity(A []int) []int {
        caps := len(A)
        arr := make([]int, caps)
        i, j := 0, 0
        for _, v := range A {
            if v%2 == 0 {
                arr[i] = v
                i++;
            } else {
                arr[caps-1-j] = v
                j++
            }
        }
        return arr
    }
    ```
- 通过 leetcode 测试 结果不准确 因为 我拿 第一的代码 也只能跑 192ms 但是 上面两个 一个 是 140ms 一个是 152ms

## 给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

- 对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

- 你可以返回任何满足上述条件的数组作为答案。

- 示例：
    ```
    输入：[4,2,5,7]
    输出：[4,5,2,7]
    解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
    ```

- 提示：
    ```
    2 <= A.length <= 20000
    A.length % 2 == 0
    0 <= A[i] <= 1000
    ```
- php 双指针写法 需要注意边界条件
    ```php
    function sortArrayByParityII($A)
    {
        $start = 0;
        $mid = 1;
        $end = count($A)-1;
        while ($start < $end ) {

            while (($start + 2 <= $end+1) && ($A[$start] & 1) === 0) {
                $start += 2;
            }

            while (($mid + 2 <= $end) && ($A[$mid] & 1) === 1) {
                $mid += 2;
            }

            if($start !== $end+1){
                $tem = $A[$start];
                $A[$start] = $A[$mid];
                $A[$mid] = $tem;
            }
            $start += 2;
            $mid += 2;

        }
        return $A;
    }
    ```
- php 基本改良写法
    ```php
    function sortArrayByParityII($A)
    {
        $j = 0;
        $k = 1;
        $end = count($A);
        $res = new SplFixedArray($end);
        for ($i = 0; $i < $end; $i++) {
            if (($A[$i] & 1) === 0) {
                $res[$j] = $A[$i];
                $j += 2;
            } else {
                $res[$k] = $A[$i];
                $k += 2;
            }
        }

        return (array)$res;
    }
    ```
- golang 基础写法
    ```golang
    func sortArrayByParityII(A []int) []int {
        arr := make([]int, len(A))
        i, j := 0, 1
        for _, v := range A {
            if v%2 == 0 {
                arr[i] = v
                i = i + 2
            } else {
                arr[j] = v
                j = j + 2
            }
        }
        return arr
    }
    ```