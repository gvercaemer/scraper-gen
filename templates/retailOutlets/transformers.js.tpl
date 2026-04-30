module.exports = () => {
  this.addressTransformer = (data) => {
    return data
  }

  this.cityTransformer = (data) => {
    return data
  }

  this.latitudeTransformer = (data) => {
    return data
  }

  this.longitudeTransformer = (data) => {
    return data
  }

  this.nameTransformer = (data) => {
    return data
  }

  this.openingHoursTransformer = (data) => {
    // Return a JSON string: { monday: ['09:00-18:00'], tuesday: [...], ... }
    return JSON.stringify(data)
  }

  this.phoneTransformer = (data) => {
    return data
  }

  this.remoteIdTransformer = (data) => {
    return String(data)
  }

  this.urlTransformer = (data) => {
    return data
  }

  this.zipTransformer = (data) => {
    return data
  }

  return this
}
