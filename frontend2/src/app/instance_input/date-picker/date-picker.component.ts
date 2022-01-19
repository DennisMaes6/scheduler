import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

//import Modal from "../Modal.svelte"
import { ButtonComponent } from "../../model_input/button/button.component"
import { DayHeaderComponent } from "../../schedule_view/day-header/day-header.component";
import { Service } from 'build/openapi' // 




@Component({
  selector: 'app-date-picker',
  templateUrl: './date-picker.component.html',
  styleUrls: ['./date-picker.component.css']
})
export class DatePickerComponent implements OnInit {



  constructor(public dialog: MatDialog) {}

  @Output()
  dateChanged = new EventEmitter<{day: number, month: number, year: number}>();



  openDialog() {
    const dialogRef = this.dialog.open(EditDateDialogContent);
    dialogRef.componentInstance.currentDate = this.currentDate;
    dialogRef.componentInstance.visibleMonth = this.currentDate.month
    dialogRef.componentInstance.visibleYear = this.currentDate.year
    dialogRef.componentInstance.updateMonthDays();
    dialogRef.afterClosed().subscribe(result => {
      
      if(result){
        this.currentDate = dialogRef.componentInstance.currentDate; // TODO: wss nog wel moeten resetten anders in dialog bij close
        this.dateChanged.emit(this.currentDate);
      }
    });
  
  }


  @Input()
  currentDate?: {day: number, month: number, year: number};
  
  @Input()
  handleSubmit: VoidFunction;

  
  setStartDate: (date: Date) => void;

  visibleMonth = this.currentDate?.month
  visibleYear = this.currentDate?.year
  blankDays: Date[];
  monthDays: Date[];


  DAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
  MONTH_NAMES = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];

  updateMonthDays(): void {
    this.blankDays = Array((new Date(this.visibleYear, this.visibleMonth-1).getDay()+6) % 7);
    this.monthDays = new Array(31)
      .fill('')
      .map((_,i) => new Date(this.visibleYear,this.visibleMonth-1,i+1))
      .filter((v) => v.getMonth() === this.visibleMonth-1)
  }

  setStartDay(date: Date){
    
    this.currentDate = {
      day: date.getDate(),
      month: date.getMonth() + 1,
      year: date.getFullYear()
    }
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
    this.visibleMonth = this.currentDate.month
    this.visibleYear = this.currentDate.year
    this.updateMonthDays()
    callback()
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

  dateToValue(): string {
    const returnDate = new Date(this.currentDate.year, this.currentDate.month-1, this.currentDate.day)
    console.log()
    return returnDate.toDateString()
  }



  ngOnInit(): void {
    this.visibleMonth = this.currentDate.month
    this.visibleYear = this.currentDate.year
    this.updateMonthDays()
    console.log("DATE PICKER INITIALIZED")
  }

}





@Component({
  selector: 'edit-date-dialog-content',
  templateUrl: './edit-date-dialog-content.html',
})
export class EditDateDialogContent {
  constructor() {}
  


  @Input()
  currentDate?: {day: number, month: number, year: number};
  
  @Input()
  handleSubmit: VoidFunction;

  
  setStartDate: (date: Date) => void;

  visibleMonth = this.currentDate?.month
  visibleYear = this.currentDate?.year
  blankDays: Date[];
  monthDays: Date[];


  DAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
  MONTH_NAMES = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];

  updateMonthDays(): void {
  
    this.blankDays = Array((new Date(this.visibleYear, this.visibleMonth-1).getDay()+6) % 7);
    this.monthDays = new Array(31)
      .fill('')
      .map((_,i) => new Date(this.visibleYear,this.visibleMonth-1,i+1))
      .filter((v) => v.getMonth() === this.visibleMonth-1)
  }

  setStartDay(date: Date){
    console.log("SET START DAY!!")
    console.log(date)
    console.log("day: ", date.getDate())
    console.log("month: ", date.getMonth())
    console.log("year: ", date.getFullYear())
    this.currentDate = {
      day: date.getDate(),
      month: date.getMonth(),
      year: date.getFullYear()
    }
  }

  incrementMonth(): void {
    console.log("CURRENT DATE: ", this.currentDate)
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
    this.visibleMonth = this.currentDate.month
    this.visibleYear = this.currentDate.year
    this.updateMonthDays()
    callback()
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

  dateToValue(): string {
    const returnDate = new Date(this.currentDate.year, this.currentDate.month-1, this.currentDate.day)
    console.log()
    return returnDate.toDateString()
  }






}
