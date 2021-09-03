module.exports = (req, res) => {
  const { readFileSync } = require("fs");
  const { join } = require("path");
  const file = readFileSync(join(__dirname, "pages", "ci.yml"), "utf8");
  console.info(join(__dirname, "pages", "ci.yml"));
  res.send(file);
};