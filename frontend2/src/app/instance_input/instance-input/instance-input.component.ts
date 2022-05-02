import { Component, OnInit } from '@angular/core';


import type { InstanceData, Day } from "build/openapi";

import { Service } from 'build/openapi/services/Service';
import { DefaultService } from 'build/openapi';

// import Modal from "../Modal.svelte";
import { ButtonComponent } from "../../model_input/button/button.component";
import { AssistantInputComponent } from "../assistant-input/assistant-input.component";
import { DatePickerComponent } from "../date-picker/date-picker.component";
import { HolidayPickerComponent } from "../holiday-picker/holiday-picker.component";

interface CustomDate {
  day: number,
  month: number,
  year: number
}



@Component({
  selector: 'app-instance-input',
  templateUrl: './instance-input.component.html',
  styleUrls: ['./instance-input.component.css'],
  providers: [ DefaultService, Service]
})
export class InstanceInputComponent implements OnInit {

  data: InstanceData = {
    assistants: [],
    days: []
  };

  hello: string = "ins-inp";

  updateAndSubmit(date: any) {
    console.log("DATE GOT: ", date)
    const returnDate = new Date(date.year, date.month, date.day)
    this.setStartDate(returnDate);
    this.handleSubmit();
  }

  holidaysUpdateAndSubmit(days: any){
    this.data.days = days
    this.handleSubmit();
  }

  handleSubmit(): void {
    console.log("POST INSTANCE DATA")
    console.log("Data:", this.data)
    
    Service.postInstanceData(this.data)
  }
  
  addDays(date: CustomDate, days: any): CustomDate {
    var result: Date = new Date(date.year, date.month - 1, date.day);
    result.setDate(result.getDate() + days);
    return { day: result.getDate(), month: result.getMonth() + 1, year: result.getFullYear() };
  }
  
  removeWeek(): void {
    console.log("REMOVE WEEK")
    this.data.days = this.data.days.slice(0, this.data.days.length - 7)
    this.handleSubmit()
  }
  
  addWeek(): void  {
    console.log("ADD WEEK! ")    
    console.log("test str: ", this.hello)
    console.log("data: ", this.data) // WERKT NIET WANT THIS.DATA WORDT GECALLD OP BUTTON

    let newDays = new Array(7)
      .fill(0)
      .map((_, i) => {
        return {
          id: this.data.days.length + i + 1,
          date: this.addDays(this.data.days[this.data.days.length - 1].date, i + 1),
          is_holiday: false,
        }
      })
    this.data.days = this.data.days.concat(newDays)
    this.handleSubmit()
  
  }
  
  setStartDate(startDate: Date): void  {
    console.log("SET START DAY IN PARENT")
    let newDays = this.data.days.map((_, i) => {
      let date = this.addDays({ day: startDate.getDate(), month: startDate.getMonth() + 1, year: startDate.getFullYear() }, i)
      return {
        id: i + 1,
        date,
        is_holiday: this.data.days.some(d => d.date.day === date.day && d.date.month == date.month && d.date.year === date.year) && this.data.days.find(d => d.date.day === date.day && d.date.month == date.month && d.date.year === date.year).is_holiday
      }
    })
    this.data.days = newDays
  }
  
  openWindow(): void {
    open()
  }

  closeWindow(): void {
    close()
  }

  closeAndSubmit(): void {
    this.handleSubmit()
    close()
  }

  dateToValue(day:number, month: number, year: number): string {
    const returnDate = new Date(year, month, day)
    console.log()
    return returnDate.toDateString()
  }
  
  constructor(private instanceInputService: DefaultService) { }

  

  getInstanceInput(): void {
    
    this.instanceInputService.instanceDataGetGet().subscribe(data => (this.data = data));
    console.log("DATA ON GET= ", this.data)
  }

  ngOnInit(): void {
    console.log("INIT INSTANCE INPUT")
    this.getInstanceInput();
  
    console.log("DATA ON INIT= ", this.data)
  }

}
