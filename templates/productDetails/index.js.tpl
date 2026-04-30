/**
 * ========================================================
 *                     SCRAPER COMMANDS
 * ========================================================
 *
 * @DEV_USAGE :
 *   scraper run --scraper="{{SCRAPER_PATH}}" --url="https://www.{{SITE}}.{{LOCALE}}/product/example"
 *
 * @PRODUCTION_TEST :
 *   scraper run --metadata='{"model":"productDetail","retailerDataSourceId":"TODO","url":"<url>"}'
 *
 * @NOTES :
 *
 * ========================================================
 */
const config = require('./config')
const pageManager = require('@services/pageManager')(config)

/**
 * @param {Browser} _browser
 * @param {Object} params
 * @returns {string[]}
 */
const getAllUrl = (_browser, params) => [params.url]

/**
 * @param {Browser} browser
 * @param {string} url
 * @param {Object} params
 * @returns {Promise<Object[]>}
 */
const getDataFromUrl = async (browser, url, params) => {
  const page = await pageManager.load(browser, url, params)
  await page.evaluate(() => {
    // TODO: extract product data from the page and populate window.jsonData
    // Example using JSON-LD:
    // const schema = JSON.parse(document.querySelector('script[type="application/ld+json"]').textContent)
    // window.jsonData = [{ ean: schema.gtin13, name: schema.name, productId: schema.sku }]
    window.jsonData = []
  })

  return await pageManager.parse(page)
}

module.exports = { getAllUrl, getDataFromUrl }
