import { Component, OnInit } from '@angular/core';


import { DefaultService } from 'build/openapi';

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.css']
})
export class MainPageComponent implements OnInit {

  constructor(private defaultService: DefaultService) { }

  generateSchedule() {
      console.log("Button to generate schedule clicked.")
      this.defaultService.scheduleGenerateGet().subscribe();
      
  }



  ngOnInit(): void {
  }

}
