const docx = require("docx");

module.exports.TableCellMarginNil = {top: 0, left: 0, bottom: 0, right: 0, marginUnitType: docx.WidthType.NIL};
module.exports.SpacingDefault = 2500;
module.exports.FontFamily = "Arial";
module.exports.WeeklyHeaderTitle = "Обзор рынка сырья, стали и ферросплавов";
module.exports.MonthlyHeaderTitle = "Обзор рынка сырья и стали";
module.exports.ShortHeaderTitle = "Мировой рынок металлургического сырья: итоги"
module.exports.RusMonth = ["января","фервраля","марта","апреля","мая","июня","июля","августа","сентября","октября","ноября","декабря"];
module.exports.LineWidth = 20;
module.exports.LineColor = "#d3d3d3";
module.exports.MinPriceId = 2;
module.exports.MaxPriceId = 3;
module.exports.MedPriceId = 1;
module.exports.StockId = 4;
module.exports.LabelOffset = 5;
module.exports.Green = "#00ff00"
module.exports.Red = "#ff0000"
module.exports.ColorDefault = "#000000"
module.exports.FontSizeTd = 18
module.exports.FontSizeTdSmall = 8
module.exports.FontSizeTh = 18
module.exports.FontSizeInfoRow = 14
module.exports.FontSizeHeading3 = 24
module.exports.HeadingColor = "#2E74B5"
module.exports.mainServerHost = "localhost"
module.exports.mainServerPort = 8080
const NONE_BORDER = {
    style: docx.BorderStyle.NONE,
    size: 0,
    color: "auto",
};
module.exports.TableNoOuterBorders = {
    top: NONE_BORDER,
    bottom: NONE_BORDER,
    left: NONE_BORDER,
    right: NONE_BORDER,
};
