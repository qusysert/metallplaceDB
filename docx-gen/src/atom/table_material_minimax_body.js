const getChange = require("../utils/get_change");
const paragraphCentred = require("../atom/paragraph_centred")
const textTd = require("../atom/text_td")
const docx = require("docx");

module.exports = function (body){
    let rows = [];
    body.forEach(m =>{
        const changeUnits1 = getChange(m.Week1Med.price_feed, 0, m.Week1Med.prev_price, false);
        const changePercents1 = getChange(m.Week1Med.price_feed, 0, m.Week1Med.prev_price, true);
        const changeUnits2 = getChange(m.Week2Med.price_feed, 0, m.Week2Med.prev_price, false);
        const changePercents2 = getChange(m.Week2Med.price_feed, 0, m.Week2Med.prev_price, true);

        rows.push(
            new docx.TableRow({
                children: [
                    new docx.TableCell({
                        children: [textTd(m.CountryAndType)]
                    }),
                    new docx.TableCell({
                        children: [textTd(m.Delivery)]
                    }),

                    new docx.TableCell({
                        children: [textTd(m.Week1Min.price_feed[0].value)]
                    }),
                    new docx.TableCell({
                        children: [textTd(m.Week1Max.price_feed[0].value)]
                    }),
                    new docx.TableCell({
                        children: [textTd(m.Week1Med.price_feed[0].value)]
                    }),
                    new docx.TableCell({
                        children: [textTd(changeUnits1.Text, changeUnits1.Color)]
                    }),
                    new docx.TableCell({
                        children: [textTd(changePercents1.Text, changePercents1.Color)]
                    }),

                    new docx.TableCell({
                        children: [textTd(m.Week2Min.price_feed[0].value)]
                    }),
                    new docx.TableCell({
                        children: [textTd(m.Week2Max.price_feed[0].value)]
                    }),
                    new docx.TableCell({
                        children: [textTd(m.Week2Med.price_feed[0].value)]
                    }),
                    new docx.TableCell({
                        children: [textTd(changeUnits2.Text, changeUnits1.Color)]
                    }),
                    new docx.TableCell({
                        children: [textTd(changePercents2.Text, changePercents1.Color)]
                    }),
                ]
            })
        )
    })
    return rows
}