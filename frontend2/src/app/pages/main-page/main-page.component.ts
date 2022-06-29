import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';


import { DefaultService, ModelParameters } from 'build/openapi';
import { Service } from 'build/openapi/services/Service';
import { delay } from 'rxjs';



@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.css']
})
export class MainPageComponent implements OnInit {

  constructor(private defaultService: DefaultService) {}

  

  IsWait = false;
  

  reader:FileReader = new FileReader();
  generateSchedule() {



      console.log("Button to generate schedule clicked.")
      this.IsWait = true;
      this.defaultService.scheduleGenerateGet().subscribe( () => this.IsWait = false )
      
      console.log("Default service done")
      
  }

  modelParameters: ModelParameters = {};




  getCoverageChanged(evt: any, new_weight: number){
    for (const prms of this.modelParameters.weights){
        prms.coverage = evt.value;
    }
    console.log("mpe: ", this.modelParameters)
    Service.postModelParams(this.modelParameters);
  }

  getBalanceChanged(evt: any, new_weight: number){
    for (const prms of this.modelParameters.weights){
        prms.balance = evt.value;
    }
    Service.postModelParams(this.modelParameters);
  }

  getFairnessChanged(evt: any, new_weight: number){
    for (const prms of this.modelParameters.weights){
        prms.fairness = evt.value;
    }
    Service.postModelParams(this.modelParameters);
  }



  ngOnInit(): void {
    this.defaultService.modelParametersGetGet().subscribe( modelParameters => ( this.modelParameters = modelParameters));
    console.log("mp: ", this.modelParameters)
  }

}
