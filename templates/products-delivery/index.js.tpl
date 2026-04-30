/**
 * ========================================================
 *                     SCRAPER COMMANDS
 * ========================================================
 *
 * @DEV_USAGE :
 *   scraper run --scraper="{{SCRAPER_PATH}}" --query="brandname"
 *
 * @PRODUCTION_TEST :
 *   scraper run --metadata='{"model":"product","retailerDataSourceId":"TODO"}' --query="brandname"
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
  // Example: return [`${config.url}/api/search?query=${params.query}`]

  return urls
}

/**
 * @param {.Browser} browser
 * @param {string} url
 * @returns {Promise<Object[]>}
 */
const getDataFromUrl = async (browser, url) => {
  const page = await pageManager.load(browser, url, params)
  // TODO: build and return the list of products to recover
  await page.evaluate((products) => (window.jsonData = products), products)

  return await pageManager.parse(page)
}

module.exports = { getAllUrl, getDataFromUrl }
