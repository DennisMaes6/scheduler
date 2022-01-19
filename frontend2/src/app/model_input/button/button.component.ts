import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-button',
  templateUrl: './button.component.html',
  styleUrls: ['./button.component.css']
})
export class ButtonComponent implements OnInit {
  
  @Input()
  callback(): void {};

  @Input()
  buttonText: string;
  
  hello: string = "button";
  primary: boolean = true

  constructor() { }



  ngOnInit(): void {
  }

}
