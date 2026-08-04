class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {

        const hash = new Map()
        for(let s of strs){

            const key = new Array(26).fill(0);

            for(let i = 0; i < s.length; i++){
                
                const code = s.charCodeAt(i) - "a".charCodeAt(0)
                key[code]++;
            } 
            const hashKey = key.join("-")
            hash.set(hashKey, [...(hash.get(hashKey) || []), s])
        }
        return Object.values(Object.fromEntries(hash.entries()));
    }
}
