import { Component, OnInit, Input } from '@angular/core';


import type { Assignment, ModelParameters } from "build/openapi";

import { ShiftType } from 'build/openapi';

import { DefaultService } from 'build/openapi';


export let free_day: boolean

@Component({
  selector: 'app-assignment-box',
  templateUrl: './assignment-box.component.html',
  styleUrls: ['./assignment-box.component.css'],
  providers: [ DefaultService ]
})
export class AssignmentBoxComponent implements OnInit {

  modelParameters: ModelParameters;


  @Input()
  assignment: Assignment;

  

  constructor(private modelParamService: DefaultService) { }
  ngOnInit(): void {
    //this.getShiftTypes();

  }

  getColor(st: ShiftType): string {
    switch (st) {
      case ShiftType.Jaev:
          return "bg-yellow-400";
      case ShiftType.Sanw:
          return "bg-green-400";
      case ShiftType.Jawe:
          return "bg-red-500";
      case ShiftType.Jaho:
          return "bg-blue-500";
      case ShiftType.Saew:
          return "bg-green-500";
      case ShiftType.Sawe:
          return "bg-red-600";
      case ShiftType.Saho:
          return "bg-blue-600";
      case ShiftType.Call:
          return "bg-green-600";
      case ShiftType.Tpwe:
          return "bg-red-700";
      case ShiftType.Tpho:
          return "bg-blue-700";
      case ShiftType.Free:
          if (free_day) return "bg-gray-200";
          return "bg-gray-100";
  }
  return "";
}
  
  getShiftTypes(): void {
  
    this.modelParamService.modelParametersGetGet().subscribe(modelParameters => (this.modelParameters = modelParameters));
  }

}
