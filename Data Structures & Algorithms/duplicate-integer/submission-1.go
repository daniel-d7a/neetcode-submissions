func hasDuplicate(nums []int) bool {
    
    valueMap := make(map[int]bool)
    duplicateFlag := false

    for _, val := range nums {
        if !valueMap[val] {
            valueMap[val] = true
        }else{
            duplicateFlag = true
            break
        }
    }

    return duplicateFlag

}
