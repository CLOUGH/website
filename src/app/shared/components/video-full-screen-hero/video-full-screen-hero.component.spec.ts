import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VideoFullScreenHeroComponent } from './video-full-screen-hero.component';

describe('VideoFullScreenHeroComponent', () => {
  let component: VideoFullScreenHeroComponent;
  let fixture: ComponentFixture<VideoFullScreenHeroComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VideoFullScreenHeroComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VideoFullScreenHeroComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
