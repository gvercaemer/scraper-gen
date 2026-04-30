module.exports = {
  name: '{{SITE}}.{{LOCALE}}',
  model: 'productDetail',
  url: 'https://www.{{SITE}}.{{LOCALE}}',
  additionalSkippedResources: [],
  dataSelectors: [
    {
      container: {
        type: 'JSON',
      },
      fields: [
        {
          name: 'ean',
          selector: '.ean',
        },
        {
          name: 'name',
          selector: '.name',
        },
        {
          name: 'productId',
          selector: '.productId',
        },
      ],
    },
  ],
}
