import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InstanceInputComponent } from './instance-input.component';

describe('InstanceInputComponent', () => {
  let component: InstanceInputComponent;
  let fixture: ComponentFixture<InstanceInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InstanceInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(InstanceInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
