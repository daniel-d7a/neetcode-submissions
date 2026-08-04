class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        const map = new Map()
        for(let i of nums){
            map.set(i, map.has(i) ? map.get(i) +1 : 1)
        }

        for(let [key, val] of map.entries()){
            if(val > 1) return true
        }
        return false
    }
}
