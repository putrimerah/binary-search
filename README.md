# binary-search
Perform a binary search in go

This code is one of my coding practice process that I get from educative.io

The goals is to perform a binary search with a low process inside to pass a big O notation.

# 1 Check the mid of an array
Searching mid of array is on of a challange. To get a mid of an array we have to define high and low limit
```golang
var low int = 0
var high int = len(nums) - 1
```
To get the mid position we can use this formula
```golang
midtmp := (high - low) / 2
```
# 2 Get the target position by mid position that we define previously
Because we want to get the index of a target, we have to get a mid value in a loop
```golang
for low <= high {
    mid := low + (high-low)/2
}
```
with this approach, the mid value will be able to set dynamicaly depend on the next logical approach.
Next are we want to get the index num of a target. The logical implementation will be like this

- If nums[mid] is equal to the target value, we return mid.
- If nums[mid] is greater than target, point high to mid-1, and the value of low remains the same. Now we will search the left part of the array.
- If nums[mid] is less than target, point low to mid + 1, and the value of high remains the same. Now we will search the right part of the array.

```golang
if nums[mid] == target {
    return mid
} else if target < nums[mid] {
    high = mid - 1
} else if target > nums[mid] {
    low = mid + 1
}
```
When the value of low is greater than the value of high, this indicates that the target is not present in the array. This is because we’ve checked all possible positions where target might exist. So, -1 is returned.

Good Luck!
