/* eslint-disable no-unused-vars */
const Service = require('./Service');

const { promisify } = require('util');
const exec = promisify(require('child_process').exec)


/**
* Returns a generated schedule.
*
* returns Schedule
* */
const scheduleGET = () => new Promise(
  async (resolve, reject) => {
    try {
      resolve(Service.successResponse(
        await getScheduleString()
      ));
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

let schedule = {};
let cached = false;

async function getScheduleString() {
  if (!cached) {
    const scheduleString = await (await exec('minizinc --solver gurobi minizinc/scheduler.mzn minizinc/data.dzn')).stdout
    const lines = scheduleString.split('\n')
    const nb_days = lines[0].split(' ').slice(-2)[0].split(":")[1]
    const shift_types = ["JA_E", "JA_NW", "JA_WH", "SA_EW", "SA_WH", "TS", "C", "F"]
    const assistants = []
    const individual_schedules = []
    lines.slice(1, -2).forEach((line) => {
      assistants.push({
        id: line.split(' ')[0].split(':')[1],
        type: line.split(' ')[1].split(':')[1]
      })
  
      const individual_schedule = {assistant_id: line.split(' ')[0].split(':')[1], assignments: []}
      line.split(' ').slice(2, -1).forEach((shift)=>{
        individual_schedule.assignments.push(shift)
      })
  
      individual_schedules.push(individual_schedule)
    })
  
    schedule = {
      nb_days,
      assistants,
      shift_types,
      individual_schedules
    }
    cached = true
  }
  return schedule
}