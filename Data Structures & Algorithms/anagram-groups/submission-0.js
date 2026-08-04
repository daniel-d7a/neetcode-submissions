class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {

        /**
         * @param {string} str
         * @return {string[]}
         */
        function createStrKey(str){
            let key = Array(26).fill(0)

            for(let i = 0; i<str.length; i++){
                key[str.charCodeAt(i)-97]++
            }
            return key.map(String).join('-')
        }

        let hash = {}
        for(let str of strs){
            const key = createStrKey(str)
            if(hash[key]){
                hash[key].push(str)
            }else{
                hash[key] = [str]
            }
        }
        return Object.values(hash)
    }
}
