/* eslint-disable no-unused-vars */
const Service = require('./Service');

/**
* Returns a generated schedule.
*
* returns Schedule
* */
const scheduleGET = () => new Promise(
  async (resolve, reject) => {
    try {
      resolve(Service.successResponse({
      }));
    } catch (e) {
      reject(Service.rejectResponse(
        e.message || 'Invalid input',
        e.status || 405,
      ));
    }
  },
);

module.exports = {
  scheduleGET,
};
