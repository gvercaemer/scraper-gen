module.exports = {
  name: '{{SITE}}.{{LOCALE}}',
  model: 'product',
  url: 'https://www.{{SITE}}.{{LOCALE}}',
  additionalSkippedResources: [],
  dataSelectors: [
    {
      container: {
        type: 'JSON',
      },
      fields: [
        {
          name: 'available',
          selector: '.available',
        },
        {
          name: 'brand',
          selector: '.brand',
        },
        {
          name: 'ean',
          selector: '.ean',
        },
        {
          name: 'imageUrl',
          selector: '.imageUrl',
        },
        {
          name: 'name',
          selector: '.name',
        },
        {
          name: 'price',
          selector: '.price',
        },
        {
          name: 'productId',
          selector: '.productId',
        },
        {
          name: 'productUrl',
          selector: '.productUrl',
        },
      ],
    },
  ],
}
