import { hasUncaughtExceptionCaptureCallback } from "process";

const puppeteer = require('puppeteer');
const fse = require('fs-extra');
const uuid = require('uuid');
const Multer = require('multer');
const mqtt = require('mqtt');
const config = require('config');

const { format } = require("util");
const { Storage } = require("@google-cloud/storage");

// message type
class Message {
    url: string;

    constructor(message: string) {
        let json = JSON.parse(message)
        this.url = json.url;
    }
}

// Instantiate a storage client with credentials
const storage = new Storage({ keyFilename: config.gcp.key_file });
// A bucket is a container for objects (files).
const bucket = storage.bucket(config.gcp.storage_bucket);

// Multer is required to process file uploads and make them available via
// req.files.
const multer = Multer({
    storage: Multer.memoryStorage(),
    limits: {
      fileSize: 5 * 1024 * 1024, // no larger than 5mb, you can change as needed.
    },
  });

const scrapeFile = async (msg: Message) => {
  let page_id = await uuid.v4(); 

  let url = msg.url;
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
};


(async () => {

    var options = {
        host: config.mqtt.host,
        port: config.mqtt.port,
        protocol: config.mqtt.protocol,
        username: config.mqtt.username,
        password: config.mqtt.password
    }

    //initialize the MQTT client
    var client = mqtt.connect(options);

    //setup the callbacks
    client.on('connect', function () {
        console.log('Connected');
    });

    client.on('error', function (error) {
        console.log(error);
    });

    client.on('message', function (topic, message) {
        //Called each time a message is received
        console.log('Received message:', topic, message.toString());
        var msg = new Message(message.toString());
        scrapeFile(msg);
    });

    // subscribe to scraper topic 
    client.subscribe(config.mqtt.topics.scraper);

})();