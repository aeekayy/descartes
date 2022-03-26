import { hasUncaughtExceptionCaptureCallback } from "process";

const puppeteer = require('puppeteer');
const fse = require('fs-extra');
const uuid = require('uuid');
const Multer = require('multer');

const { format } = require("util");
const { Storage } = require("@google-cloud/storage");


// Instantiate a storage client with credentials
const storage = new Storage({ keyFilename: "gcp.json" });
// A bucket is a container for objects (files).
const bucket = storage.bucket(process.env.GCLOUD_STORAGE_BUCKET);

// Multer is required to process file uploads and make them available via
// req.files.
const multer = Multer({
    storage: Multer.memoryStorage(),
    limits: {
      fileSize: 5 * 1024 * 1024, // no larger than 5mb, you can change as needed.
    },
  });

(async () => {
  let page_id = await uuid.v4(); 
  console.log(process.argv.slice(2));
  let args = process.argv.slice(2);
  if (args.length < 1) {
    console.log(`Invalid argument`);
    process.exit(1);
  }

  let url = args[0];
  let outputDir = "./output";
  await fse.ensureDir(outputDir);

  const browser = await puppeteer.launch({
    //args: [ '--proxy-server=200.73.128.156:3128' ]
  });
  const page = await browser.newPage();
  await page.setRequestInterception(true);
  page.on('request', (request) => {
    request.continue();
  });
  await page.goto(url, {
    waitUntil: 'networkidle2',
  });
  await page.screenshot({ path: `${outputDir}/${page_id}.png` });
  await page.pdf({ path: `${outputDir}/${page_id}.pdf`, format: 'a4' });
  const html_content = await page.content();
  // With a callback:
  await fse.outputFile(`${outputDir}/${page_id}.html`, html_content);

  await browser.close();

  let res = await storage.bucket(bucket.name).upload(`${outputDir}/${page_id}.png`);
  console.log(res.id);
  res = await storage.bucket(bucket.name).upload(`${outputDir}/${page_id}.pdf`);
  console.log(res.id);
  res = await storage.bucket(bucket.name).upload(`${outputDir}/${page_id}.html`);
  console.log(res.id);
})();