class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {

        const hashArray = Array.from({length: nums.length+1}, ()=>[])

        const hash = {}
        for(let num of nums){
            hash[num] = (hash[num] || 0) + 1
        }

        for (let num in hash){
            hashArray[hash[num]].push(num)
        }

            console.log(hash)
            console.log(hashArray)
        let result = []
        for(let i = hashArray.length - 1; i > 0; i--){
            
            for(let num of hashArray[i]){
                result.push(num)
                if(result.length === k) return result
            }
        }
    }
}
