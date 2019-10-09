import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminTopNavbarComponent } from './admin-top-navbar.component';

describe('AdminTopNavbarComponent', () => {
  let component: AdminTopNavbarComponent;
  let fixture: ComponentFixture<AdminTopNavbarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AdminTopNavbarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminTopNavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
