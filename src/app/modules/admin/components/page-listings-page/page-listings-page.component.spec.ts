import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PageListingsPageComponent } from './page-listings-page.component';

describe('PageListingsPageComponent', () => {
  let component: PageListingsPageComponent;
  let fixture: ComponentFixture<PageListingsPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PageListingsPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PageListingsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
