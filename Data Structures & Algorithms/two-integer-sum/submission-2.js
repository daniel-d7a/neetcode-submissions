class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {

        const hash = {}

        for(let i = 0; i<nums.length; i++){
            hash[nums[i]] = target - nums[i]
        }

        for(let i = 0; i<nums.length; i++){
            if(hash[nums[i]]!= undefined){
                const index = nums.indexOf(hash[nums[i]], i+1)
                if(index === -1) continue
                return [i, index]
            }
        }
        
    }
}
