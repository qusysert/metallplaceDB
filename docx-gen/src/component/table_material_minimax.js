const docx = require("docx");
const {TableNoOuterBorders, TableCellMarginNil} = require("../const");
const paragraphCentred = require("../atom/paragraph_centred");
const textTh = require("../atom/text_th")

function priceBlock(unit){
    return new docx.Table({
        width: {
            size: 100,
            type: docx.WidthType.PERCENTAGE,
        },
        borders: TableNoOuterBorders,
        rows:[
            new docx.TableRow({
                children: [
                    new docx.TableCell({columnSpan: 3, children:[textTh(`Цена, ${unit}`)]})
                ]
            }),
            new docx.TableRow({
                children: [
                    new docx.TableCell({children:[textTh(`мин`)]}),
                    new docx.TableCell({children:[textTh(`макс`)]}),
                    new docx.TableCell({children:[textTh(`сред`)]}),
                ]
            })
        ]
    })
}

function headerMaterial(title, unit){
    return  new docx.Table({
        width: {
            size: 100,
            type: docx.WidthType.PERCENTAGE,
        },
        borders: TableNoOuterBorders,
        rows: [
            new docx.TableRow({
                children: [new docx.TableCell({columnSpan: 3, margins: TableCellMarginNil, children: [paragraphCentred(title)]})]
            }),
            new docx.TableRow({
                children: [
                    new docx.TableCell({margins: TableCellMarginNil, children: [priceBlock(unit)]})
                ],
            })
        ]
    })
}

module.exports = async function tableMaterialMinimax(materialIds, dateTitle, date) {
    return new docx.Table({
        width: {
            size: 100,
            type: docx.WidthType.PERCENTAGE,
        },
        columnWidths: [4,4,10,10],
        rows:[
            new docx.TableRow({
                children: [
                    new docx.TableCell({ margins: TableCellMarginNil, children: [paragraphCentred("Страна/вид")]}),
                    new docx.TableCell({ margins: TableCellMarginNil, children: [paragraphCentred("Условия поставки")]}),
                    new docx.TableCell({ margins: TableCellMarginNil, children: [headerMaterial("неделя1", "USD/т")]}),
                    new docx.TableCell({ margins: TableCellMarginNil, children: [headerMaterial("неделя2", "USD/т")]}),
                ],
            })
        ]
    })
}