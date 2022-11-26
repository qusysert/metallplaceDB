const docx = require("docx");
const chartBlock = require("./chart_block")

const {TableCellMarginNil} = require("../const");
module.exports = async function twoChart(url1, url2, avgGroup){
    if(avgGroup === undefined) avgGroup = 1
    const block1 = await chartBlock(url1, false, avgGroup)
    const block2 = await chartBlock(url2, false, avgGroup)
    return new docx.Table({
        width: {
            size: 100,
            type: docx.WidthType.PERCENTAGE,
        },
        columnWidths: [20, 1, 20],
        borders: docx.TableBorders.NONE,
        rows: [
            new docx.TableRow({
                children: [
                    new docx.TableCell({
                        margins: TableCellMarginNil,
                        children: [
                            block1
                        ],
                    }),
                    new docx.TableCell({children: []}),
                    new docx.TableCell({
                        margins: TableCellMarginNil,
                        children: [
                            block2
                        ],
                    }),
                ],
            }),
        ],
    })
}