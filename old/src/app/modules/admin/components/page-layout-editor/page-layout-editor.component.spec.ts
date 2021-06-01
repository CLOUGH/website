import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PageLayoutEditorComponent } from './page-layout-editor.component';

describe('PageLayoutEditorComponent', () => {
  let component: PageLayoutEditorComponent;
  let fixture: ComponentFixture<PageLayoutEditorComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PageLayoutEditorComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PageLayoutEditorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
