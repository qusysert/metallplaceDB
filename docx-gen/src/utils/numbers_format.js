module.exports = function (num, fixed){
    let numStr = num.toString()
    if (numStr.split(".").length - 1 > 1 || numStr.indexOf("'") !== -1){
        return numStr
    }
    if(num >= 1000){
        let afterComma = ""
        let beforeComma = num
        if(num.toString().indexOf(".") !== -1){
            const numArr = num.toString().split(".")
            beforeComma = numArr[0]
            afterComma = numArr[1]
        }
        const after = beforeComma.toString().slice(-3)
        const before = beforeComma.toString().slice(0, beforeComma.toString().length - 3)
        numStr = before + " " + after
        if(afterComma !== ""){
            numStr += "." + afterComma
        }
    }
    numStr = numStr.replace(".", ",")
    if (fixed !== 0 && fixed !== undefined){
        if (numStr.indexOf(",") === -1){
            numStr += "," + "0".repeat(fixed)
        } else {
            numStr += "0".repeat(fixed - numStr.substring(numStr.indexOf(",") + 1).length)
        }
    }
    return numStr
}