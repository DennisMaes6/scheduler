import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';



import type { Day } from "build/openapi"

//import Modal from "../Modal.svelte"
import { ButtonComponent } from "../../model_input/button/button.component"




@Component({
  selector: 'app-holiday-picker',
  templateUrl: './holiday-picker.component.html',
  styleUrls: ['./holiday-picker.component.css']
})
export class HolidayPickerComponent implements OnInit {

  MONTH_NAMES = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
  DAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];

  @Input()
  days?: Day[]

  @Output()
  holidaysChanged = new EventEmitter<Day[]>();


  visibleMonth = 0
  visibleYear = 0
  blankDays: number[];
  monthDays: Date[];

  months: { month: number, year: number }[] = [];

  handleSubmit: VoidFunction


  

  openDialog() {
    const dialogRef = this.dialog.open(EditHolidayDialogContent);
    dialogRef.componentInstance.days = this.days
    dialogRef.componentInstance.visibleMonth = this.visibleMonth
    dialogRef.componentInstance.visibleYear = this.visibleYear
    dialogRef.componentInstance.months = this.months
    dialogRef.componentInstance.updateMonthDays();
    dialogRef.afterClosed().subscribe(result => {
      
      if(result){
        this.days = dialogRef.componentInstance.days; // TODO: wss nog wel moeten resetten anders in dialog bij close
        this.holidaysChanged.emit(this.days);
        
      }
    });
  
  }

  toggleHoliday(day: Day): void {
    day.is_holiday = !day.is_holiday
    this.days = this.days
    this.updateMonthDays()
  }

  updateMonthDays(): void {
    this.blankDays = Array((new Date(this.visibleYear, this.visibleMonth - 1).getDay() + 6) % 7);
    this.monthDays = new Array(31)
      .fill('')
      .map((_, i) => new Date(this.visibleYear, this.visibleMonth - 1, i + 1))
      .filter((v) => v.getMonth() === this.visibleMonth - 1);
  }

  incrementMonth(): void {
    if (this.visibleMonth == 12) {
      this.visibleMonth = 1
      this.visibleYear++
    } else {
      this.visibleMonth++
    }
    this.updateMonthDays()
  }


  decrementMonth(): void {
    if (this.visibleMonth == 1) {
      this.visibleMonth = 12;
      this.visibleYear--
    } else {
      this.visibleMonth--
    }
    this.updateMonthDays()
  }


  resetStateAnd(callback: VoidFunction): void {
    this.visibleMonth = this.days[0].date.month
    this.visibleYear = this.days[0].date.year
    this.updateMonthDays()
    callback()
  }

  isPartOfPlanningPeriod(date: Date): boolean {
    return this.days.some(d => d.date.day === date.getDate() && d.date.month === date.getMonth() + 1 && d.date.year === date.getFullYear())
  }

  getDay(date: Date): Day {
    return this.days.find(d => d.date.day === date.getDate() && d.date.month === date.getMonth() + 1 && d.date.year === date.getFullYear())
  }

  closeWindow(): void {
    close()
  }

  openWindow(): VoidFunction {
    return open;
  }

  closeAndSubmit(): void {
    this.handleSubmit()
    close()
  }
  
  constructor(public dialog: MatDialog) { }

  ngOnInit(): void {
    
    this.visibleMonth = this.days[0].date.month  
    this.visibleYear = this.days[0].date.year
    this.months = Array.from(
      new Set(this.days?.map((d) => { return { month: d.date.month, year: d.date.year } }))
    ).sort((a, b) => {
      if (a.year == b.year) {
        return a.month - b.month
      }
      return a.year - b.year
    })
    this.updateMonthDays()
  }

}







@Component({
  selector: 'edit-holiday-dialog-content',
  templateUrl: './edit-holiday-dialog-content.html',
})
export class EditHolidayDialogContent {
  constructor() {}
  
  MONTH_NAMES = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
  DAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];

  days?: Day[]


  visibleMonth = 0
  visibleYear = 0
  blankDays: number[];
  monthDays: Date[];

  months: { month: number, year: number }[] = [];


  toggleHoliday(day: Day): void {
    day.is_holiday = !day.is_holiday
    this.updateMonthDays()
  }

  updateMonthDays(): void {
    this.blankDays = Array((new Date(this.visibleYear, this.visibleMonth - 1).getDay() + 6) % 7);
    this.monthDays = new Array(31)
      .fill('')
      .map((_, i) => new Date(this.visibleYear, this.visibleMonth - 1, i + 1))
      .filter((v) => v.getMonth() === this.visibleMonth - 1);
  }

  incrementMonth(): void {
    if (this.visibleMonth == 12) {
      this.visibleMonth = 1
      this.visibleYear++
    } else {
      this.visibleMonth++
    }
    this.updateMonthDays()
  }


  decrementMonth(): void {
    if (this.visibleMonth == 1) {
      this.visibleMonth = 12;
      this.visibleYear--
    } else {
      this.visibleMonth--
    }
    this.updateMonthDays()
  }


  resetStateAnd(callback: VoidFunction): void {
    this.visibleMonth = this.days[0].date.month
    this.visibleYear = this.days[0].date.year
    this.updateMonthDays()
    callback()
  }

  isPartOfPlanningPeriod(date: Date): boolean {
    return this.days.some(d => d.date.day === date.getDate() && d.date.month === date.getMonth() + 1 && d.date.year === date.getFullYear())
  }

  getDay(date: Date): Day {
    return this.days.find(d => d.date.day === date.getDate() && d.date.month === date.getMonth() + 1 && d.date.year === date.getFullYear())
  }

  






}