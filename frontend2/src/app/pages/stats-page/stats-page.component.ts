import { Component, OnInit } from '@angular/core';
import { Assistant, DefaultService, IndividualSchedule, InstanceData, Schedule } from 'build/openapi';


export interface globalStatsDataSource {
  Coverage: number;
  Balance: number;
  Fairness: number;
  TotalNbShifts: number;
  TotalNbShiftsAssigned: number;
}


@Component({
  selector: 'app-stats-page',
  templateUrl: './stats-page.component.html',
  styleUrls: ['./stats-page.component.css']
})


export class StatsPageComponent implements OnInit {

  constructor(private instanceInputService: DefaultService) { }

  schedule?: Schedule = {
    fairness_score: 0,
    balance_score: 0,
    jaev_fairness_score: 0,
    jaev_balance_score: 0,
    individual_schedules: [],
    Coverage: 0,
    Balance: 0,
    Fairness: 0,
    TotalNbShifts: 0,
    TotalNbShiftsAssigned: 0
  };


  

  globalStats: globalStatsDataSource[] = [{
    Coverage: 0,
    Balance: 0,
    Fairness: 0,
    TotalNbShifts: 0,
    TotalNbShiftsAssigned: 0
  }]

  

  data: InstanceData = {
    assistants: [],
    days: []
  };
  
  getIs(id: number): IndividualSchedule | undefined {
    //console.log(this.schedule?.individual_schedules?.find(s => s.assistant_id === id))
    if(this.schedule?.individual_schedules?.find(s => s.assistant_id === id) == undefined){
      console.log("UNDEFINDED GET IS: ", id );
    }

    return this.schedule?.individual_schedules?.find(s => s.assistant_id === id);
  }


  getScheduleInput(): void {
    this.instanceInputService.dbScheduleGet().subscribe(schedule => {
      this.schedule = schedule 
      this.globalStats[0].Coverage = schedule.Coverage
      this.globalStats[0].Balance = schedule.Balance
      this.globalStats[0].Fairness = schedule.Fairness
      this.globalStats[0].TotalNbShifts = schedule.TotalNbShifts
      this.globalStats[0].TotalNbShiftsAssigned = schedule.TotalNbShiftsAssigned
    });
  }



  getInstanceInput(): void {
    this.instanceInputService.instanceDataGetGet().subscribe(data => (this.data = data));
  }

  getAssistant(id: number): Assistant | undefined{
    return this.data.assistants.find(a => a.id === id);
  }

  displayedColumns: string[] = ['assistant_id', 'workload', 'absolute_workload', 'days_available', 'days_worked', 'days_vacation', 'avg_days_rest'];
  displayedColumnsGlobal: string[] = ['Coverage', 'Balance', 'Fairness', 'TotalNbShifts', 'TotalNbShiftsAssigned'];

  ngOnInit(): void {
    this.getInstanceInput();
    this.getScheduleInput();

  }

}
