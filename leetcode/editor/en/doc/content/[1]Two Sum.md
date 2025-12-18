<p>Given an array of integers <code>nums</code>&nbsp;and an integer <code>target</code>, return <em>indices of the two numbers such that they add up to <code>target</code></em>.</p>

<p>You may assume that each input would have <strong><em>exactly</em> one solution</strong>, and you may not use the <em>same</em> element twice.</p>

<p>You can return the answer in any order.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,7,11,15], target = 9
<strong>Output:</strong> [0,1]
<strong>Explanation:</strong> Because nums[0] + nums[1] == 9, we return [0, 1].
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,2,4], target = 6
<strong>Output:</strong> [1,2]
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,3], target = 6
<strong>Output:</strong> [0,1]
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li><code>2 &lt;= nums.length &lt;= 10<sup>4</sup></code></li> 
 <li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li> 
 <li><code>-10<sup>9</sup> &lt;= target &lt;= 10<sup>9</sup></code></li> 
 <li><strong>Only one valid answer exists.</strong></li> 
</ul>

<p>&nbsp;</p> 
<strong>Follow-up:&nbsp;</strong>Can you come up with an algorithm that is less than 
<code>O(n<sup>2</sup>)</code>
<font face="monospace">&nbsp;</font>time complexity?

<details><summary><strong>Related Topics</strong></summary>Array | Hash Table</details><br>

<div>👍 65907, 👎 2448<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/issues' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**Note: English content is improving...**



<p><strong><a href="https://labuladong.online/algo/en/practice-in-action/nsum/" target="_blank">⭐️Detailed Explanation</a></strong></p>
<details><summary><strong>Brief Thoughts</strong></summary>





<div id="labuladong_solution_en">

## Thoughts

Everyone loves a person with a sense of humor. If you want to joke about your chronic procrastination, you can joke about yourself like this (manual dog head):

I've been studying vocabulary for half a year and I'm still at abandon, abandon. I've been practicing coding problems for half a year and I'm still at two sum, two sum...

But jokes aside, this problem isn't difficult, but it's quite famous because it's the first problem on LeetCode. There are various ways to solve it, and I'll mention two of the most common approaches.

The first approach is based on sorting. After sorting the `nums` array, you can use the left and right pointers technique mentioned in [Summary of Array Two-pointer Techniques](https://labuladong.online/algo/en/essential-technique/array-two-pointers-summary/) to find two numbers that sum up to `target`.

However, since the problem asks us to return the indices of the elements and sorting would destroy the original indices, we need to map the values to their original indices.

Furthermore, if the problem extends to finding the sum of three numbers, four numbers, and so on, you can still use the two-pointer technique. I've written a function that can solve all N-number sum problems in [One Function to Solve All nSum Problems](https://labuladong.online/algo/en/practice-in-action/nsum/).

The second approach is to use a hash table for assistance. For an element `nums[i]`, you want to know if there is another element `nums[j]` with the value of `target - nums[i]`. This is simple. We use a hash table to record the mapping from each element's value to its index, allowing us to quickly determine if there is an element with the value of `target - nums[i]` in the array.

In short, the array can be understood as a hash table mapping from 'index -> value', and we just need to establish a mapping from 'value -> index' to solve this problem.

**Detailed Solution: [One Method to Solve All nSum Problems](https://labuladong.online/algo/en/practice-in-action/nsum/)**

Please note that the links in the markdown should remain unchanged as per your instruction.

</div>


<div id="solution">

## Solution



<div class="tab-panel"><div class="tab-nav">
<button data-tab-item="cpp" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">cpp🤖</button>

<button data-tab-item="python" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">python🤖</button>

<button data-tab-item="java" class="tab-nav-button btn active" data-tab-group="default" onclick="switchTab(this)">java🟢</button>

<button data-tab-item="go" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">go🤖</button>

<button data-tab-item="javascript" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">javascript🤖</button>
</div><div class="tab-content">
<div data-tab-item="cpp" class="tab-item " data-tab-group="default"><div class="highlight">

```cpp
// Note: This cpp code is translated by chatGPT🤖 based on my java code.
// This code has passed all the test cases, should be accepted by LeetCode.

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        // maintain the mapping of val -> index
        unordered_map<int, int> valToIndex;
        for (int i = 0; i < nums.size(); i++) {
            // check the table to see if there is an element that can sum up with nums[i] to the target
            int need = target - nums[i];
            if (valToIndex.find(need) != valToIndex.end()) {
                return {valToIndex[need], i};
            }
            // store the mapping of val -> index
            valToIndex[nums[i]] = i;
        }
        return {};
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# Note: This python code is translated by chatGPT🤖 based on my java code.
# This code has passed all the test cases, should be accepted by LeetCode.

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # maintain the mapping of val -> index
        val_to_index = {}
        for i in range(len(nums)):
            # check the table to see if there is an element that can sum up with nums[i] to the target
            need = target - nums[i]
            if need in val_to_index:
                return [val_to_index[need], i]
            # store the mapping of val -> index
            val_to_index[nums[i]] = i
        return []
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public int[] twoSum(int[] nums, int target) {
        // maintain the mapping of val -> index
        HashMap<Integer, Integer> valToIndex = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            // check the table to see if there is an element that can sum up with nums[i] to the target
            int need = target - nums[i];
            if (valToIndex.containsKey(need)) {
                return new int[]{valToIndex.get(need), i};
            }
            // store the mapping of val -> index
            valToIndex.put(nums[i], i);
        }
        return null;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// Note: This go code is translated by chatGPT🤖 based on my java code.
// This code has passed all the test cases, should be accepted by LeetCode.

func twoSum(nums []int, target int) []int {
    // maintain the mapping of val -> index
    valToIndex := make(map[int]int)
    for i, num := range nums {
        // check the table to see if there is an element that can sum up with nums[i] to the target
        need := target - num
        if j, ok := valToIndex[need]; ok {
            return []int{j, i}
        }
        // store the mapping of val -> index
        valToIndex[num] = i
    }
    return nil
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// Note: This javascript code is translated by chatGPT🤖 based on my java code.
// This code has passed all the test cases, should be accepted by LeetCode.

var twoSum = function(nums, target) {
    // maintain the mapping of val -> index
    let valToIndex = new Map();
    for (let i = 0; i < nums.length; i++) {
        // check the table to see if there is an element that can sum up with nums[i] to the target
        let need = target - nums[i];
        if (valToIndex.has(need)) {
            return [valToIndex.get(need), i];
        }
        // store the mapping of val -> index
        valToIndex.set(nums[i], i);
    }
    return null;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>👾👾 Algo Visualize 👾👾</strong></summary><div id="data_two-sum"  en="true" category="leetcode" ></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_two-sum"></div></div>
</details><hr /><br />

</div>
</details>
</div>

