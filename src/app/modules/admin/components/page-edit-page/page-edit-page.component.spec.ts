import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PageEditPageComponent } from './page-edit-page.component';

describe('PageEditPageComponent', () => {
  let component: PageEditPageComponent;
  let fixture: ComponentFixture<PageEditPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PageEditPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PageEditPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
