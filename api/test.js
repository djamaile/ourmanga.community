module.exports = (req, res) => {
  const { readFileSync } = require("fs");
  const { join } = require("path");
  const file = readFileSync(join(__dirname, "pages", "viz-2021-9-3.html"), "utf8");
  console.info(join(__dirname, "pages", "viz-2021-9-3.html"));
  res.send(file);
};