import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ScheduleViewPageComponent } from './schedule-view-page.component';

describe('ScheduleViewPageComponent', () => {
  let component: ScheduleViewPageComponent;
  let fixture: ComponentFixture<ScheduleViewPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ScheduleViewPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ScheduleViewPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
