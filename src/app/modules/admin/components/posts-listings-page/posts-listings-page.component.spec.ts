import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PostsListingsPageComponent } from './posts-listings-page.component';

describe('PostsListingsPageComponent', () => {
  let component: PostsListingsPageComponent;
  let fixture: ComponentFixture<PostsListingsPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PostsListingsPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PostsListingsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
