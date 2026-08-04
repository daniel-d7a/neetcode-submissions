class Solution {


    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
        
        function getFrequency(s){
            const map = new Map()

            for(let c of s){
                const count = map.has(c) ? map.get(c) : 0
                map.set(c, count + 1)
            }
            return map
        }

        const sMap = getFrequency(s)
        const tMap = getFrequency(t)


        const longerEntries = sMap.size > tMap.size ? sMap.entries() : tMap.entries()

        for(let [key] of longerEntries){
            if(sMap.get(key) !== tMap.get(key)) return false
        }
        return true
    }

}
