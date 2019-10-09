import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-page-create-page',
  templateUrl: './page-create-page.component.html',
  styleUrls: ['./page-create-page.component.scss']
})
export class PageCreatePageComponent implements OnInit {

  constructor() { }

  public onReady(editor) {
    editor.ui.getEditableElement().parentElement.insertBefore(
      editor.ui.view.toolbar.element,
      editor.ui.getEditableElement()
    );
  }

  ngOnInit() {
  }

  createPage() {

  }

}