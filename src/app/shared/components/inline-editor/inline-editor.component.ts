import { Component, OnInit,Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-inline-editor',
  templateUrl: './inline-editor.component.html',
  styleUrls: ['./inline-editor.component.scss']
})
export class InlineEditorComponent implements OnInit {
  @Input() public content: string;
  @Output() public contentChange: EventEmitter<string> = new EventEmitter();

  public tinymceConfig = {
    base_url: '/tinymce', // Root for resources
    suffix: '.min',
    plugins: 'code'
  };

  constructor() { }

  ngOnInit() {
  }

  public updateContent(event) {
    this.contentChange.emit(event.editor.getContent() as string);
  }

}
