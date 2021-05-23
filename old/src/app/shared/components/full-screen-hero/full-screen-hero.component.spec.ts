import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FullScreenHeroComponent } from './full-screen-hero.component';

describe('FullScreenHeroComponent', () => {
  let component: FullScreenHeroComponent;
  let fixture: ComponentFixture<FullScreenHeroComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FullScreenHeroComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FullScreenHeroComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
