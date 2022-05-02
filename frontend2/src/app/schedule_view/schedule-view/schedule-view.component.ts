import { Component, OnInit } from '@angular/core';



import type { Schedule, InstanceData, Assistant, IndividualSchedule, Assignment, Day, DBFile} from 'build/openapi';
import { AssistantType } from 'build/openapi';
import { DefaultService } from 'build/openapi';
import { Service } from 'build/openapi/services/Service';

import { AssistantHeaderComponent } from '../assistant-header/assistant-header.component';
import { DayHeaderComponent } from '../day-header/day-header.component';
import { AssignmentBoxComponent } from '../assignment-box/assignment-box.component';




@Component({
  selector: 'app-schedule-view',
  templateUrl: './schedule-view.component.html',
  styleUrls: ['./schedule-view.component.css']
})
export class ScheduleViewComponent implements OnInit {

  schedule?: Schedule = {
    fairness_score: 0,
    balance_score: 0,
    jaev_fairness_score: 0,
    jaev_balance_score: 0,
    individual_schedules: []
  };
  data: InstanceData = {
    assistants: [],
    days: []
  };

  dbfile: DBFile = {
    filename: "demo_dummy.db"
  }

  
  workload_array = this.schedule?.individual_schedules?.map((a) => a.workload)

  max_workload?: number = this.workload_array?.reduce(function (a, b) {
    return Math.max(a, b);
  }, 0);

  min_workload?: number = this.workload_array?.reduce(function (a, b) {
    return Math.min(a, b);
  }, 0);


  types: AssistantType[] = Object.values(AssistantType)


  getAssistant(id: number): Assistant | undefined{
    return this.data.assistants.find(a => a.id === id);
  }

  getIs(id: number): IndividualSchedule | undefined {
    //console.log(this.schedule?.individual_schedules?.find(s => s.assistant_id === id))
    if(this.schedule?.individual_schedules?.find(s => s.assistant_id === id) == undefined){
      console.log("UNDEFINDED GET IS: ", id );
    }
    return this.schedule?.individual_schedules?.find(s => s.assistant_id === id);
  }

  getAssignmentsOnDay(assistant: Assistant, day: Day): Assignment | undefined { 
      return this.getIs(assistant.id).assignments.find(a => a.day_nb === day.id)
  }

  getInstanceInput(): void {
    
    this.instanceInputService.instanceDataGetGet().subscribe(data => (this.data = data));
    
    console.log("DATA ON GET= ", this.data)
  }

  getScheduleInput(): void {
    this.instanceInputService.dbScheduleGet().subscribe(schedule => (this.schedule = schedule));
  }

  constructor(private instanceInputService: DefaultService) { }

  ngOnInit(): void {

    console.log("Init Schedule View")

    

    console.log(this.types)
    this.getInstanceInput();
    this.getScheduleInput();

    Service.postDBFile(this.dbfile)


    let isSyncingLeftScroll: boolean = false;
    let isSyncingTopScroll: boolean = false;
    let isSyncingCenterScroll: boolean = false;
    let leftDiv: HTMLElement | null = document.getElementById('assistantlist');
    let topDiv: HTMLElement | null  = document.getElementById('dayheaders');
    let centerDiv: HTMLElement | null = document.getElementById('schedule');

    if(leftDiv !== null){
      leftDiv.onscroll = function () {
        if (!isSyncingLeftScroll && leftDiv !== null && centerDiv !== null) {
          isSyncingCenterScroll = true;
          centerDiv.scrollTop = leftDiv.scrollTop;
        }
        isSyncingLeftScroll = false;
      }
    }

    if(topDiv !== null){
      topDiv.onscroll = function () {
        if (!isSyncingTopScroll && centerDiv !== null && topDiv !== null) {
          isSyncingCenterScroll = true;
          centerDiv.scrollLeft = topDiv.scrollLeft;
        }
        isSyncingTopScroll = false;
      }
    }
    
    if(centerDiv !== null){
      centerDiv.onscroll = function () {
        if (!isSyncingCenterScroll && leftDiv !== null && topDiv !== null && centerDiv !== null) {
          isSyncingLeftScroll = true;
          isSyncingTopScroll = true;
          leftDiv.scrollTop = centerDiv.scrollTop;
          topDiv.scrollLeft = centerDiv.scrollLeft;
        }
        isSyncingCenterScroll = false;
  
      }
    }
  }

}
