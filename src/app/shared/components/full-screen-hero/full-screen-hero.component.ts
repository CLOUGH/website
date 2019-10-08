import { Component, OnInit, Input } from '@angular/core';
import * as InlineEditor from '@ckeditor/ckeditor5-build-inline';
import * as ClassicEditor from '@ckeditor/ckeditor5-build-classic';
import { CKEditor5 } from '@ckeditor/ckeditor5-angular';
// import * as Alignmnent from '@ckeditor/ckeditor5-alignment/src/alignment';

@Component({
  selector: 'app-full-screen-hero',
  templateUrl: './full-screen-hero.component.html',
  styleUrls: ['./full-screen-hero.component.scss']
})
export class FullScreenHeroComponent implements OnInit {
  @Input() public section: any;

  public Editor = InlineEditor;
  public editorConfig:any;

  constructor() { }

  ngOnInit() {
  }

}
