import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminLeftSideNavComponent } from './admin-left-side-nav.component';

describe('AdminLeftSideNavComponent', () => {
  let component: AdminLeftSideNavComponent;
  let fixture: ComponentFixture<AdminLeftSideNavComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AdminLeftSideNavComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminLeftSideNavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
