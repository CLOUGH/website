import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminPageHeadingComponent } from './admin-page-heading.component';

describe('AdminPageHeadingComponent', () => {
  let component: AdminPageHeadingComponent;
  let fixture: ComponentFixture<AdminPageHeadingComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AdminPageHeadingComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminPageHeadingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
