module.exports = {
  name: '{{SITE}}.{{LOCALE}}',
  model: 'retailOutlet',
  url: 'https://www.{{SITE}}.{{LOCALE}}',
  additionalSkippedResources: [],
  dataSelectors: [
    {
      container: {
        type: 'JSON',
      },
      fields: [
        {
          name: 'address',
          selector: '.address',
        },
        {
          name: 'city',
          selector: '.city',
        },
        {
          name: 'latitude',
          selector: '.latitude',
        },
        {
          name: 'longitude',
          selector: '.longitude',
        },
        {
          name: 'name',
          selector: '.name',
        },
        {
          name: 'openingHours',
          selector: '.openingHours',
        },
        {
          name: 'phone',
          selector: '.phone',
        },
        {
          name: 'remoteId',
          selector: '.remoteId',
        },
        {
          name: 'url',
          selector: '.url',
        },
        {
          name: 'zip',
          selector: '.zip',
        }
      ],
    },
    {
      container: {
        type: 'CONST',
      },
      fields: [
        {
          name: 'active',
          value: true,
        },
      ],
    },
  ],
}
