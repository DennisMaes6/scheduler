import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { GoogleChartsModule } from 'angular-google-charts';

import { FlexLayoutModule } from '@angular/flex-layout';

import { MatSliderModule } from '@angular/material/slider';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatDialogModule } from '@angular/material/dialog';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatTabsModule } from '@angular/material/tabs';
import {MatFormFieldModule} from '@angular/material/form-field'
import { MatOptionModule } from '@angular/material/core';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatTableModule} from '@angular/material/table';
import {MatDividerModule} from '@angular/material/divider';



import { AssignmentBoxComponent } from './schedule_view/assignment-box/assignment-box.component';
import { ModelInputComponent } from './model_input/model-input/model-input.component';
import { InputFieldComponent } from './model_input/input-field/input-field.component';
import { ButtonComponent } from './model_input/button/button.component';
import { DayHeaderComponent } from './schedule_view/day-header/day-header.component';
import { AssistantHeaderComponent } from './schedule_view/assistant-header/assistant-header.component';
import { ScheduleViewComponent } from './schedule_view/schedule-view/schedule-view.component';
import { NewAssistantComponent } from './instance_input/new-assistant/new-assistant.component';
import { IconButtonComponent } from './model_input/icon-button/icon-button.component';
import { AssistantInputComponent } from './instance_input/assistant-input/assistant-input.component';
import { DatePickerComponent } from './instance_input/date-picker/date-picker.component';
import { HolidayPickerComponent } from './instance_input/holiday-picker/holiday-picker.component';
import { InstanceInputComponent } from './instance_input/instance-input/instance-input.component';
import { ModalComponent } from './modal/modal.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { MainPageComponent } from './pages/main-page/main-page.component';
import { ModelParametersPageComponent } from './pages/model-parameters-page/model-parameters-page.component';
import { InstanceDataPageComponent } from './pages/instance-data-page/instance-data-page.component';
import { ScheduleViewPageComponent } from './pages/schedule-view-page/schedule-view-page.component';
import { EditDateDialogContent } from './instance_input/date-picker/date-picker.component';
import { EditHolidayDialogContent } from './instance_input/holiday-picker/holiday-picker.component';
import { AssistantInputDialogContent } from './instance_input/assistant-input/assistant-input.component';
import { EditAssistantComponent } from './instance_input/edit-assistant/edit-assistant.component';
import { AssistantPipePipe } from './assistant-pipe.pipe';
import { NewAssistantDialogComponent } from './instance_input/assistant-input/assistant-input.component';
import { SchedulePickerPageComponent } from './pages/schedule-picker-page/schedule-picker-page.component';
import { StatsPageComponent } from './pages/stats-page/stats-page.component';



@NgModule({
  declarations: [
    AppComponent,
    AssignmentBoxComponent,
    ModelInputComponent,
    InputFieldComponent,
    ButtonComponent,
    DayHeaderComponent,
    AssistantHeaderComponent,
    ScheduleViewComponent,
    NewAssistantComponent,
    IconButtonComponent,
    AssistantInputComponent,
    DatePickerComponent,
    HolidayPickerComponent,
    InstanceInputComponent,
    ModalComponent,
    HomePageComponent,
    MainPageComponent,
    ModelParametersPageComponent,
    InstanceDataPageComponent,
    ScheduleViewPageComponent,
    EditDateDialogContent,
    EditHolidayDialogContent,
    AssistantInputDialogContent,
    EditAssistantComponent,
    AssistantPipePipe,
    NewAssistantDialogComponent,
    SchedulePickerPageComponent,
    StatsPageComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    GoogleChartsModule,
    MatSliderModule,
    MatButtonModule,
    FlexLayoutModule,
    MatInputModule,
    MatDialogModule,
    MatToolbarModule,
    MatTabsModule,
    MatFormFieldModule,
    MatOptionModule,
    MatCheckboxModule,
    MatProgressBarModule,
    MatProgressSpinnerModule,
    MatGridListModule,
    MatTableModule,
    MatDividerModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
