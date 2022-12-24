const docx = require("docx");
const footer = require("../component/footer");
const header = require("../component/header");
const h1 = require("../atom/heading1");
const h2 = require("../atom/heading2");
const h3 = require("../atom/heading3");
const h3Fake = require("../atom/heading3_fake");
const paragraph = require("../atom/paragraph");
const twoChart = require("../component/two_chart");
const {MonthlyHeaderTitle, MedPriceId, StockId, RusMonth} = require("../const");
const oneChartText = require("../component/one_chart_text");
const oneChart = require("../component/one_chart");
const singleTable = require("../component/table_single");
const singleTableMinimax = require("../component/table_single_minimax");
const tableDoubleAvg = require("../component/table_double_avg");
const tableDouble = require("../component/table_double");
const tableMaterialMinimax = require("../component/table_material_minimax");
const doubleTableMinimax = require("../component/table_double_minimax")
const tableMaterialGrouped = require("../component/table_material_grouped")
const {GetWeekDates, GetWeekNumber, GetLastDayOfMonth, Get2LastFridays, Get2LastThursdays} = require("../utils/date_operations");
const {GetMonthRange, Get2WeekRange, GetYearRange, Get2YearRange} = require("../utils/date_ranges")
const {ChartUrl, FormChartUrl} = require("../utils/form_chart_url")
const fs = require("fs");

function getFooterTitle(date) {

    const weekDates = GetWeekDates(date)
    return `Отчетный период: ${weekDates.first.day} ${RusMonth[weekDates.first.month]} - ` +
        `${weekDates.last.day} ${RusMonth[weekDates.last.month]} ${weekDates.last.year} года (${GetWeekNumber(date)} неделя)`
}

module.exports = class MonthlyReport {
    async generate(date) {
        date = GetLastDayOfMonth(date)
        return new docx.Document({
            features: {
                updateFields: true,
            },
            sections: [
                {
                    properties: {
                        page: {
                            margin: {
                                top: 0,
                                right: 0,
                                bottom: 0,
                                left: 0,
                            },
                        },
                    },
                    children: [
                        new docx.Paragraph({
                            children: [
                                new docx.ImageRun({
                                    data: fs.readFileSync("/home/ivan/Pictures/Screenshots/pic2.png"),
                                    transformation: {
                                        width: 1000,
                                        height: 2000,
                                    },
                                }),
                            ]
                        }),
                    ]
                },
                {
                    footers: {
                        default: footer(getFooterTitle(date)),
                    },
                    headers: {
                        default: header(MonthlyHeaderTitle)
                    },
                    children: [
                        h3Fake("Содержание"),

                        paragraph("Дисклеймер: Информация, представленная на портале metallplace.ru предназначена только для справки и\n" +
                            "не предназначена для торговых целей или для удовлетворения ваших конкретных требований. Контент\n" +
                            "включает факты, взгляды и мнения отдельных лиц, а не веб-сайта или его руководства.\n"),
                        paragraph("Пользователи/посетители должны принимать собственные решения на основе собственных независимых\n" +
                            "запросов, оценок, суждений и рисков. Портал metallplace.ru не несет ответственность за какие-либо убытки,\n" +
                            "затраты или действия, возникающие в результате использования распространяемых цен."),

                        new docx.Paragraph({children: [new docx.PageBreak()]}),
                        h2("Краткая сводка новостей по мировову рынку"),
                        new docx.Paragraph({children: [new docx.PageBreak()]}),

                        h2("Краткая сводка цен по мировому рынку"),
                        h3Fake("Сырьевые материалы"),
                        paragraph({
                            children: [
                                await twoChart( // ЖРС62 ЛОМ hms
                                    FormChartUrl(new ChartUrl([1], MedPriceId, Get2YearRange(date), 0, "line", "month")),
                                    FormChartUrl(new ChartUrl([4], MedPriceId, Get2YearRange(date), 0, "line", "month")),
                                    undefined,
                                    "м/м"
                                )
                            ]
                        }),
                    ],
                },
            ],
        });
    }
}

function getRangeArr(first, last) {
    let arr = []
    for (let i = first; i <= last; i++) {
        arr.push(i)
    }
    return arr
}
