import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BlogListingsPageComponent } from './blog-listings-page.component';

describe('BlogListingsPageComponent', () => {
  let component: BlogListingsPageComponent;
  let fixture: ComponentFixture<BlogListingsPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BlogListingsPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BlogListingsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
