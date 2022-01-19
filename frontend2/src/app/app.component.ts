import { Component } from '@angular/core';

import { Service } from 'build/openapi';

import { ScheduleViewComponent } from './schedule_view/schedule-view/schedule-view.component';
import { ModelInputComponent } from './model_input/model-input/model-input.component';
import { InstanceInputComponent } from './instance_input/instance-input/instance-input.component';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Scheduler';
}
