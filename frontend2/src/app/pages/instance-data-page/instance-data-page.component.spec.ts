import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InstanceDataPageComponent } from './instance-data-page.component';

describe('InstanceDataPageComponent', () => {
  let component: InstanceDataPageComponent;
  let fixture: ComponentFixture<InstanceDataPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InstanceDataPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(InstanceDataPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
