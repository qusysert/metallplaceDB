const {mainServerHost, mainServerPort} = require("../const")
module.exports.ChartUrl = function (materialIds, propertyId, timeFrame, isBig, type, avgSnap){
    this.materialIds = materialIds
    this.propertyId = propertyId
    this.timeFrame = timeFrame
    this.isBig = isBig
    this.type = type
    this.avgSnap = avgSnap
}

module.exports.FormChartUrl = function (ChartUrl){
    let url = "http://" + mainServerHost + ":" + mainServerPort + "/getChart/"
    const materialIds = ChartUrl.materialIds.join("-")
    if(ChartUrl.group === undefined) ChartUrl.group = 0

    url += materialIds + "_" + ChartUrl.propertyId + "_" + ChartUrl.timeFrame + "_" + ChartUrl.isBig + "_" + ChartUrl.type + "_" + ChartUrl.avgSnap + ".png"
    return url
}