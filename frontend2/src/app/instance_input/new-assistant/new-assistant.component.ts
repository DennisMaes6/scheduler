import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

import type { Assistant } from "build/openapi";
import { AssistantType } from "build/openapi";




@Component({
  selector: 'app-new-assistant',
  templateUrl: './new-assistant.component.html',
  styleUrls: ['./new-assistant.component.css']
})
export class NewAssistantComponent implements OnInit {


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
    this.onAssistantNameChange.emit(newAssistantName);
  }

  assistantTypeChanged(newAssistantTypeString: string){
    const newAssistantType = this.getStringType(newAssistantTypeString);
    this.onAssistantTypeChange.emit(newAssistantType);
  }


  getStringType(typeString: string): AssistantType {
    switch (typeString) {
      case "JA": return AssistantType.Ja
      case "JA_F": return AssistantType.JaF
      case "SA": return AssistantType.Sa
      case "SA_F": return AssistantType.SaF
      case "SA_N": return AssistantType.SaNeo
      case "SA_FN": return AssistantType.SaFNeo
    }
    return AssistantType.Ja
  }

  constructor() { }

  ngOnInit(): void {
  }

}
