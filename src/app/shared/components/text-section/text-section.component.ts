import { TextSection } from './../../../core/models/section';
import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-text-section',
  templateUrl: './text-section.component.html',
  styleUrls: ['./text-section.component.scss']
})
export class TextSectionComponent implements OnInit {
  @Input() public section: TextSection;
  @Input() public readonly: Boolean;
  @Output() public sectionChange: EventEmitter<TextSection> = new EventEmitter();
  @Output() public removeSection: EventEmitter<boolean> = new EventEmitter();

  constructor() { }



  ngOnInit() {
  }

}
