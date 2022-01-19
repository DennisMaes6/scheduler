import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { InstanceDataPageComponent } from './pages/instance-data-page/instance-data-page.component';
import { MainPageComponent } from './pages/main-page/main-page.component';
import { ModelParametersPageComponent } from './pages/model-parameters-page/model-parameters-page.component';
import { ScheduleViewPageComponent } from './pages/schedule-view-page/schedule-view-page.component';


const routes: Routes = [
  {path: '', redirectTo: 'home-page', pathMatch: 'full'}, 
  {path: 'home-page', component: HomePageComponent},
  {path: 'main', component: MainPageComponent},
  {path: 'model-parameters', component: ModelParametersPageComponent },
  {path: 'instance-data', component: InstanceDataPageComponent },
  {path: 'schedule-view', component: ScheduleViewPageComponent }   
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
