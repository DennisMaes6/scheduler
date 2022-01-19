import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModelParametersPageComponent } from './model-parameters-page.component';

describe('ModelParametersPageComponent', () => {
  let component: ModelParametersPageComponent;
  let fixture: ComponentFixture<ModelParametersPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ModelParametersPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ModelParametersPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
