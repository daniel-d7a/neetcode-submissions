class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {

        const hash = {};
        const freq = Array.from({length: nums.length + 1},()=>[])

        for(let num of nums){
            hash[num] = (hash[num] || 0) + 1
        }
        
        for(let [n, q] of Object.entries(hash)){
            freq[q].push(n)
        }


        const res = []
        for(let i = freq.length-1; i>0; i--){
            for(let n of freq[i]){
                res.push(n)
                if(res.length === k) return res
            }
        }

    }
}
