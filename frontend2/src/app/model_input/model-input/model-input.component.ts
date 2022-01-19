import { Component, OnInit, Output, EventEmitter } from '@angular/core';

import type { ModelParameters} from "build/openapi";
  
import { DefaultService } from 'build/openapi';
import { Service } from 'build/openapi' // 
import { Observable } from 'rxjs';
//import { Assignment } from "../assignment-box/assignment-box.component"




@Component({
  selector: 'app-model-input',
  templateUrl: './model-input.component.html',
  styleUrls: ['./model-input.component.css'],
  providers: [ DefaultService, Service ]
})
export class ModelInputComponent implements OnInit {

  modelParameters: ModelParameters = {};
  
  buttonText: string = "Submit";
  constructor(private modelParamService: DefaultService) { }

  ngOnInit(): void {
    console.log("MODEL INPUT INITIALIZED")
    this.getModelParameters();
    console.log("model params: ", this.modelParameters)
  }

  // opgeroepen wanneer op button gedrukt wordt
  getValueChanged(evt: any, shift_type: string){
    if(typeof this.modelParameters.shift_type_parameters !== 'undefined'){
      for (const prms of this.modelParameters.shift_type_parameters){
        if(prms.shift_type == shift_type){
          prms.shift_workload = evt;
        }
      }
    }
    Service.postModelParams(this.modelParameters);
  }

  getValueBufferChanged(evt: any, shift_type: string){
    if(typeof this.modelParameters.shift_type_parameters !== 'undefined'){
      for (const prms of this.modelParameters.shift_type_parameters){
        if(prms.shift_type == shift_type){
          prms.max_buffer = evt;
        }
      }
    }
    Service.postModelParams(this.modelParameters);
  }


  handleSubmit(): void {
    console.log("HANDLE SUBMIT CALLED")
    console.log(this.modelParameters)
    Service.postModelParams(this.modelParameters);
  
  }

  balanceChanged(new_bal: number) {
      this.modelParameters.min_balance = new_bal
      this.handleSubmit()
  }
  balanceJaevChanged(new_bal: number) {
    this.modelParameters.min_balance_jaev = new_bal
    this.handleSubmit()
}


  // DIT IS AL NODIG OM GEEN UNDEFINED BODY TE KRIJGEN
  // TODO: MOMENTEEL 2X SUBMIT, NAAR 1 GAAN
  wrapSubmit(evt: any, shift_type: string): void {
    this.getValueChanged(evt, shift_type)
    //this.handleSubmit() // TODO: zien of dit nodig is?
  }


  getModelParameters(): void {
    this.modelParamService.modelParametersGetGet().subscribe( modelParameters => ( this.modelParameters = modelParameters));
  }
  

}
