import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AssistantInputComponent } from './assistant-input.component';

describe('AssistantInputComponent', () => {
  let component: AssistantInputComponent;
  let fixture: ComponentFixture<AssistantInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AssistantInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AssistantInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
