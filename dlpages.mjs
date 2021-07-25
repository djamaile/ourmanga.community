#!/usr/bin/env zx

await $`cd pages && find . -type f ! -name 'BUILD.bazel' -delete`

const date = new Date();
const year = date.getUTCFullYear();
const month = date.getMonth() + 1;
const day = date.getUTCDate();
const monthName = date.toLocaleDateString('en-US', {month: 'long'});

console.log(chalk.blue(`getting files for: ${year}-${month}-${day}`));

const sites = [
  {url: "https://yenpress.com/new-releases/", name: "yenpress"},
  {url: "https://sevenseasentertainment.com/release-dates/", name: "sevenseas"},
  {url:`https://www.darkhorse.com/Books/Browse/Manga---${monthName}+${year}-${monthName}+${year}/P9wdwkt8`, name:"darkhorse"},
  {url:"https://kodansha.us/manga/calendar", name:"kodansha"},
  {url:`https://www.viz.com/calendar/${year}/${month}`, name:"viz"},
];

await Promise.all(sites.map(async site => {
  let resp = await fetch(site.url);
  const page = await resp.text();
  await $`echo ${page} > pages/${site.name}-${year}-${month}-${day}.html`;
}));