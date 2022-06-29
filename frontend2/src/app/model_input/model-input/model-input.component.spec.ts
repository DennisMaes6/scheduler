import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModelInputComponent } from './model-input.component';

describe('ModelInputComponent', () => {
  let component: ModelInputComponent;
  let fixture: ComponentFixture<ModelInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ModelInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ModelInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
