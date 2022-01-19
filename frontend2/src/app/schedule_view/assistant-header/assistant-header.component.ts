import { Component, Input, OnInit } from '@angular/core';


import type { Assistant } from 'build/openapi'
import { AssistantType } from 'build/openapi'

//export type assistantType =  AssistantType;


@Component({
  selector: 'app-assistant-header',
  templateUrl: './assistant-header.component.html',
  styleUrls: ['./assistant-header.component.css']
})
export class AssistantHeaderComponent implements OnInit {


  @Input()
  assistant: Assistant;

  @Input()
  workload: number;

  max_workload: number;
  min_workload: number;

  assistantType = AssistantType;
  

  getWorkloadTextStyle(): string {
    if (this.workload == this.max_workload) {
      return "text-red-500 font-bold"
    } else if (this.workload == this.min_workload) {
      return "text-green-500 font-bold"
    } else {
      return "text-gray-700"
    }
  }

  constructor() { }

  ngOnInit(): void {
    
  }

}
