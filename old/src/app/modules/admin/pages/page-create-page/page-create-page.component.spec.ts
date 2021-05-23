import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PageCreatePageComponent } from './page-create-page.component';

describe('PageCreatePageComponent', () => {
  let component: PageCreatePageComponent;
  let fixture: ComponentFixture<PageCreatePageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PageCreatePageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PageCreatePageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
