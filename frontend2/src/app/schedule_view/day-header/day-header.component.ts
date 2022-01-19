import { Component, Input, OnInit } from '@angular/core';

import type { Day } from 'build/openapi';




@Component({
  selector: 'app-day-header',
  templateUrl: './day-header.component.html',
  styleUrls: ['./day-header.component.css']
})
export class DayHeaderComponent implements OnInit {

  @Input()
  day: Day;


  day_of_week(day: Day): string {
    switch (day.id % 7) {
        case 0:
            return "Thu";
        case 1:
            return "Fri";
        case 2:
            return "Sat";
        case 3:
            return "Sun";
        case 4:
            return "Mon";
        case 5:
            return "Tue";
        case 6:
            return "Wed";
    }
    return "";
}


  constructor() { }

  ngOnInit(): void {
    //   console.log("INIT DAY HEADER")
      
  }

}
