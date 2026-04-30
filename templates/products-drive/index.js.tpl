/**
 * ========================================================
 *                     SCRAPER COMMANDS
 * ========================================================
 *
 * @DEV_USAGE :
 *   scraper run --scraper="{{SCRAPER_PATH}}" --query="brandname" --remoteId="1"
 *
 * @PRODUCTION_TEST :
 *   scraper run --metadata='{"model":"product","retailerDataSourceId":"TODO","retailOutletRemoteId":"1"}' --query="brandname"
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
  // TODO: build and return the list of URLs to scrape
  // Example: return [`${config.url}/api/search?query=${params.query}`]

  return urls
}

/**
 * @param {Browser} browser
 * @param {string} url
 * @param {Object} params
 * @returns {Promise<Object[]>}
 */
const getDataFromUrl = async (browser, url, params) => {
  const page = await pageManager.load(browser, url, params)
  // TODO: build and return the list of products to recover
  await page.evaluate((products) => (window.jsonData = products), products)

  return await pageManager.parse(page)
}

module.exports = { getAllUrl, getDataFromUrl }
