import { Component, Input, OnInit, EventEmitter, Output} from '@angular/core';

import { DefaultService, ModelParameters } from 'build/openapi';
import { Service } from 'build/openapi/services/Service';

@Component({
  selector: 'app-input-field',
  templateUrl: './input-field.component.html',
  styleUrls: ['./input-field.component.css'],
  providers: [ DefaultService ]
})
export class InputFieldComponent implements OnInit {
  
  @Input()
  value: number = 0;
  
  @Input()
  step: number = 0;

  @Output()
  onChange = new EventEmitter<number>();

  callback($event: any) {
    console.log($event, "II")
    console.log(typeof $event, "II")

    
    // TODO: value krijgen uit box
    // TODO: ook nog linken met echte modelparams ipv lokaal hier.
    this.onChange.emit(this.value);
    console.log(this.value)
  };

  valChanged(newVal: string){
    const newValNumber = parseInt(newVal)
    this.onChange.emit(newValNumber);
    console.log("VALCHANGEDDDDD")
    console.log(newValNumber)
  }
  
  
  constructor() { }

  ngOnInit(): void {
  }





}
