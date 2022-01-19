import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import {MatDialog} from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatFormField } from '@angular/material/form-field';


import { InstanceData, Assistant, Service } from "build/openapi";
import { AssistantType } from "build/openapi";

import { IconButtonComponent } from "../../model_input/icon-button/icon-button.component"
import { ButtonComponent } from "../../model_input/button/button.component";
//import Modal from "../Modal.svelte";
import { DayHeaderComponent } from "../../schedule_view/day-header/day-header.component";
import { NewAssistantComponent } from "../new-assistant/new-assistant.component";




@Component({
  selector: 'app-assistant-input',
  templateUrl: './assistant-input.component.html',
  styleUrls: ['./assistant-input.component.css']
})
export class AssistantInputComponent implements OnInit {

  @Input()
  data?: InstanceData;
  assistantTypes: AssistantType[] = Object.values(AssistantType);

  id_array = this.data?.assistants.map((a) => a.id)

  new_assistant: Assistant = {
    id: this.id_array?.reduce(function (a, b) {
      return Math.max(a, b);
    }, 0) + 1,
    name: "",
    type: AssistantType.Ja,
    free_days: []
  }

  removeAssistant(assistant: Assistant): void {
    
    this.data.assistants = this.data.assistants.filter((a) => a.id != assistant.id)
  }

  addAssistant(): void {
    this.data.assistants.push(this.new_assistant)
    this.new_assistant = {
      id: this.id_array?.reduce(function (a, b) {
        return Math.max(a, b);
      }, 0) + 1,
      name: "",
      type: AssistantType.Ja,
      free_days: []
    }

    this.data.assistants = this.data.assistants
  }

  toggle(assistant: Assistant, day_nb: number): void {
    if (assistant.free_days.includes(day_nb)) {
      assistant.free_days = assistant.free_days.filter((d: number) => d != day_nb)
    } else {
      assistant.free_days.push(day_nb)
    }
    this.data.assistants = this.data.assistants
  }
  /*
  setWeek(assistant: Assistant, weekNb: number, selected: any): void {
    if (selected.checked) {
      [7 * weekNb + 1, 7 * weekNb + 2, 7 * weekNb + 3, 7 * weekNb + 4, 7 * weekNb + 5, 7 * weekNb + 6, 7 * weekNb + 7]
        .forEach((day) => {
          if (!assistant.free_days.includes(day)) assistant.free_days.push(day)
        })
    } else {
      [7 * weekNb + 1, 7 * weekNb + 2, 7 * weekNb + 3, 7 * weekNb + 4, 7 * weekNb + 5, 7 * weekNb + 6, 7 * weekNb + 7]
        .forEach((day) => {
          assistant.free_days = assistant.free_days.filter((d: number) => d != day)
        })
    }
    this.data.assistants = this.data.assistants
  } */

  freeWeek(assistant: Assistant, weekNb: number): boolean {
    return [7 * weekNb + 1, 7 * weekNb + 2, 7 * weekNb + 3, 7 * weekNb + 4, 7 * weekNb + 5, 7 * weekNb + 6, 7 * weekNb + 7]
      .every((day) => assistant.free_days.includes(day))
  }

  getTypeString(type: AssistantType): string {
    switch (type) {
      case AssistantType.Ja: return "JA"
      case AssistantType.JaF: return "JA_F"
      case AssistantType.Sa: return "SA"
      case AssistantType.SaF: return "SA_F"
      case AssistantType.SaNeo: return "SA_N"
      case AssistantType.SaFNeo: return "SA_FN"
    }
    return ""
  }

  changeNewAssistantName(newName: string){
    this.new_assistant.name = newName;
  }

  constructor(public dialog: MatDialog) {}

  openDialog() {
    const dialogRef = this.dialog.open(AssistantInputDialogContent);
    dialogRef.componentInstance.data = this.data
    dialogRef.afterClosed().subscribe(result => {
      if(result){
        this.handleSubmit();
      }
    });
  }


  ngOnInit(): void {
    console.log("init assistant input")
    console.log("data: ", this.data)
    let isSyncingLeftScroll: boolean = false;
    let isSyncingTopScroll: boolean = false;
    let isSyncingCenterScroll: boolean = false;
    let leftDiv: HTMLElement = document.getElementById('left');
    let topDiv: HTMLElement = document.getElementById('top');
    let centerDiv: HTMLElement = document.getElementById('center');

    /*
    leftDiv.onscroll = function () {
      if (!isSyncingLeftScroll) {
        isSyncingCenterScroll = true;
        centerDiv.scrollTop = leftDiv.scrollTop;
      }
      isSyncingLeftScroll = false;
    }

    topDiv.onscroll = function () {
      if (!isSyncingTopScroll) {
        isSyncingCenterScroll = true;
        centerDiv.scrollLeft = topDiv.scrollLeft;
      }
      isSyncingTopScroll = false;
    }

    centerDiv.onscroll = function () {
      if (!isSyncingCenterScroll) {
        isSyncingLeftScroll = true;
        isSyncingTopScroll = true;
        leftDiv.scrollTop = centerDiv.scrollTop;
        topDiv.scrollLeft = centerDiv.scrollLeft;
      }
      isSyncingCenterScroll = false;

    } */
  }
  /*
  openWindow(): void {
    open()
  }

  closeWindow(): void {
    close()
  } 

  addAndClose(): void {
    this.addAssistant()
    close()
  } */
  handleSubmit(): void {
    Service.postInstanceData(this.data)
  }
}

@Component({
  selector: 'assistant-input-dialog-content',
  templateUrl: 'assistant-input-dialog-content.html',
  providers: [ Service]
})
export class AssistantInputDialogContent implements OnInit {

  constructor(public dialog: MatDialog) {}

  data?: InstanceData;
  assistantTypes: AssistantType[] = Object.values(AssistantType);

  id_array = this.data?.assistants.map((a) => a.id)

  new_assistant: Assistant = {
    id: this.id_array?.reduce(function (a, b) {
      return Math.max(a, b);
    }, 0) + 1,
    name: "",
    type: AssistantType.Ja,
    free_days: []
  }

  removeAssistant(assistant: Assistant): void {
    this.data.assistants = this.data.assistants.filter((a) => a.id != assistant.id)
  }

  addAssistant(): void {
  
    this.data.assistants.push(this.new_assistant)
    this.new_assistant = {
      id: this.id_array?.reduce(function (a, b) {
        return Math.max(a, b);
      }, 0) + 1,
      name: "",
      type: AssistantType.Ja,
      free_days: []
    }

    this.data.assistants = this.data.assistants
  }

  toggle(assistant: Assistant, day_nb: number): void {

    if (assistant.free_days.includes(day_nb)) {
      assistant.free_days = assistant.free_days.filter((d: number) => d != day_nb)
    } else {
      assistant.free_days.push(day_nb)
    }
    //this.data.assistants = this.data.assistants
  }

  setWeek(assistant: Assistant, weekNb: number, selected: any): void {
    console.log("SET WEEK: ", selected)
    if (selected.checked) {
      [7 * weekNb + 1, 7 * weekNb + 2, 7 * weekNb + 3, 7 * weekNb + 4, 7 * weekNb + 5, 7 * weekNb + 6, 7 * weekNb + 7]
        .forEach((day) => {
          if (!assistant.free_days.includes(day)) assistant.free_days.push(day)
        })
    } else {
      [7 * weekNb + 1, 7 * weekNb + 2, 7 * weekNb + 3, 7 * weekNb + 4, 7 * weekNb + 5, 7 * weekNb + 6, 7 * weekNb + 7]
        .forEach((day) => {
          assistant.free_days = assistant.free_days.filter((d: number) => d != day)
        })
    }
    this.data.assistants = this.data.assistants
  }

  freeWeek(assistant: Assistant, weekNb: number): boolean {
    return [7 * weekNb + 1, 7 * weekNb + 2, 7 * weekNb + 3, 7 * weekNb + 4, 7 * weekNb + 5, 7 * weekNb + 6, 7 * weekNb + 7]
      .every((day) => assistant.free_days.includes(day))
  }

  getTypeString(type: AssistantType): string {
    switch (type) {
      case AssistantType.Ja: return "JA"
      case AssistantType.JaF: return "JA_F"
      case AssistantType.Sa: return "SA"
      case AssistantType.SaF: return "SA_F"
      case AssistantType.SaNeo: return "SA_N"
      case AssistantType.SaFNeo: return "SA_FN"
    }
    return ""
  }

  changeNewAssistantName(newName: string){
    this.new_assistant.name = newName;
  }

  ngOnInit(): void {
    console.log("init assistant input")
    console.log("data: ", this.data)
    let isSyncingLeftScroll: boolean = false;
    let isSyncingTopScroll: boolean = false;
    let isSyncingCenterScroll: boolean = false;
    let leftDiv: HTMLElement = document.getElementById('left');
    let topDiv: HTMLElement = document.getElementById('top');
    let centerDiv: HTMLElement = document.getElementById('center');
    
    leftDiv.onscroll = function () {
      if (!isSyncingLeftScroll) {
        isSyncingCenterScroll = true;
        centerDiv.scrollTop = leftDiv.scrollTop;
      }
      isSyncingLeftScroll = false;
    }

    topDiv.onscroll = function () {
      if (!isSyncingTopScroll) {
        isSyncingCenterScroll = true;
        centerDiv.scrollLeft = topDiv.scrollLeft;
      }
      isSyncingTopScroll = false;
    }

    centerDiv.onscroll = function () {
      if (!isSyncingCenterScroll) {
        isSyncingLeftScroll = true;
        isSyncingTopScroll = true;
        leftDiv.scrollTop = centerDiv.scrollTop;
        topDiv.scrollLeft = centerDiv.scrollLeft;
      }
      isSyncingCenterScroll = false;

    } 
  }
  /*
  closeWindow(): void {
    close()
  }

  addAndClose(): void {
    this.addAssistant()
    close()
  } */

  handleSubmit(): void {
    console.log("POST INSTANCE DATA")
    console.log("Data:", this.data)
    
  
    Service.postInstanceData(this.data)
  }

  openDialog() {
    this.id_array= this.data?.assistants.map((a) => a.id)
    this.new_assistant.id = this.id_array?.reduce(function (a, b) {
      return Math.max(a, b);
    }, 0) + 1
    console.log("id_array: ", this.id_array)
    console.log("data: ", this.data)
    const dialogRef = this.dialog.open(NewAssistantDialogComponent);
    dialogRef.componentInstance.assistantTypes = this.assistantTypes;
    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
      console.log("name: ", dialogRef.componentInstance.value);
      console.log("type: ", dialogRef.componentInstance.assistantType);
      if(result){
        console.log("id: ", this.new_assistant.id)
        console.log("type: ", this.new_assistant.type)
        this.changeNewAssistantName(dialogRef.componentInstance.value);
        this.new_assistant.type = dialogRef.componentInstance.assistantType;
        this.addAssistant();
        this.handleSubmit();
      }
    });
    
  }

  

}



@Component({
  selector: 'app-new-assistant-dialog',
  templateUrl: './new-assistant-dialog.component.html',
  styleUrls: ['./new-assistant-dialog.component.css']
})
export class NewAssistantDialogComponent implements OnInit {


  assistantName: string;
  assistantType: AssistantType;
  assistantTypes: AssistantType[] = Object.values(AssistantType);

  @Input()
  value: string;

  @Output()
  onAssistantNameChange = new EventEmitter<string>();


  @Output()
  onAssistantTypeChange = new EventEmitter<AssistantType>();

  assistantNameChanged(newAssistantName: string){

    this.value = newAssistantName;
    this.onAssistantNameChange.emit(newAssistantName);
  }

  assistantTypeChanged(newAssistantTypeString: string){
    console.log(newAssistantTypeString)
    const newAssistantType = this.getStringType(newAssistantTypeString);
    this.assistantType = this.getStringType(newAssistantTypeString);


    this.onAssistantTypeChange.emit(this.getStringType(newAssistantTypeString));

    
  }


  getStringType(typeString: string): AssistantType {
    switch (typeString) {
      case "JA": return AssistantType.Ja
      case "JA_F": return AssistantType.JaF
      case "SA": return AssistantType.Sa
      case "SA_F": return AssistantType.SaF
      case "SA_NEO": return AssistantType.SaNeo
      case "SA_F_NEO": return AssistantType.SaFNeo
    }
    return AssistantType.Ja
  }

  constructor() { }

  ngOnInit(): void {
  }

}
