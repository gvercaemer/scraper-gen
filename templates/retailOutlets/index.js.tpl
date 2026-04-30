/**
 * ========================================================
 *                     SCRAPER COMMANDS
 * ========================================================
 *
 * @DEV_USAGE :
 *   scraper run --scraper="{{SCRAPER_PATH}}"
 *
 * @PRODUCTION_TEST :
 *   scraper run --metadata='{"model":"retailOutlet","retailerDataSourceId":"TODO"}'
 *
 * @NOTES :
 *
 * ========================================================
 */
const config = require('./config')
const pageManager = require('@services/pageManager')(config)

/**
 * @param {Browser} browser
 * @param {Object} params
 * @returns {Promise<string[]>}
 */
const getAllUrl = async (browser, params) => {
  const page = await pageManager.load(browser, url, params)
  const urls = await page.evaluate(() => {
    // TODO: build and return the list of URLs to scrape
  })

  // Example: return [`${config.url}/api/stores`]

  return urls
}

/**
 * @param {Browser} browser
 * @param {string} url
 * @returns {Promise<Object[]>}
 */
const getDataFromUrl = async (browser, url) => {
  const page = await pageManager.load(browser, url)
  await page.evaluate(() => {
    // TODO: populate window.jsonData with an array of store objects
    window.jsonData = JSON.parse(document.body.textContent)
  })

  return await pageManager.parse(page)
}

module.exports = { getAllUrl, getDataFromUrl }

