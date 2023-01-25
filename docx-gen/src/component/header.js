const docx = require("docx");
const paragraph = require("../atom/paragraph")
const {TableCellMarginNil, FontFamily, HeaderFooterMargin} = require("../const");

module.exports = function (title){
    return new docx.Header({
        children: [
            new docx.Table({
                width: {
                    size: 100,
                    type: docx.WidthType.PERCENTAGE,
                },
                borders: {
                    top: {size: 0},
                    right: {size: 0},
                    left: {size: 0},
                    bottom: {style: docx.BorderStyle.DASHED, size: 20, color: "#d3d3d3"},
                },
                rows: [
                    new docx.TableRow({
                        children: [
                            new docx.TableCell({
                                margins: TableCellMarginNil,
                                children: [
                                    paragraph({
                                        alignment: docx.AlignmentType.JUSTIFIED,
                                        children: [new docx.TextRun({text: title,  font: FontFamily})],
                                        spacing: {
                                            after: HeaderFooterMargin
                                        }
                                    })
                                ],
                                borders: {
                                    top: {size: 0},
                                    right: {size: 0},
                                    left: {size: 0},
                                    bottom: {style: docx.BorderStyle.DASHED, size: 20, color: "#d3d3d3"},
                                },
                            })
                        ],

                    })
                ]
            }),
        ],
    });
}
