import { readdir } from "fs/promises"
import { join } from "path"

async function* ls (path = ".")
{ yield path
  for (const dirent of await readdir(path, { withFileTypes: true }))
    if (dirent.isDirectory())
      yield* ls(join(path, dirent.name))
    else
      yield join(path, dirent.name)
}

async function* empty () {}

async function toArray (iter = empty())
{ 
  console.info('test');
  let r = []
  for await (const x of iter)
    console.log(x);
    r.push(x)
  return r
}


module.exports = (req, res) => {
  const { readFileSync } = require("fs");
  const { join } = require("path");
  const file = readFileSync(join(__dirname, "pages", "viz-2021-9-3.html"), "utf8");
  console.info(join(__dirname, "pages", "viz-2021-9-3.html"));
  toArray(ls(".")).then(console.log, console.error);
  res.send(file);
};