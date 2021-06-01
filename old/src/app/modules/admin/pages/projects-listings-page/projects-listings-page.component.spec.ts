import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProjectsListingsPageComponent } from './projects-listings-page.component';

describe('ProjectsListingsPageComponent', () => {
  let component: ProjectsListingsPageComponent;
  let fixture: ComponentFixture<ProjectsListingsPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProjectsListingsPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProjectsListingsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
