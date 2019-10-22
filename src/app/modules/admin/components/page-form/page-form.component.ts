import { BsModalRef, BsModalService, BsDaterangepickerConfig } from 'ngx-bootstrap';

import { Component, OnInit, TemplateRef, Input, Output, EventEmitter } from '@angular/core';
import { HeroSection, Section, TextSection } from './../../../../core/models/section';
import { faTrash, faArrowsAlt, faExpand } from '@fortawesome/free-solid-svg-icons';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';
import { BsDatepickerInlineConfig } from 'ngx-bootstrap/datepicker/public_api';

@Component({
  selector: 'app-page-form',
  templateUrl: './page-form.component.html',
  styleUrls: ['./page-form.component.scss']
})
export class PageFormComponent implements OnInit {
  public currentDate = new Date();
  @Input() public page: any;
  @Input() public submitButtonLabel: any;
  @Output() public submitPage: EventEmitter<any> = new EventEmitter();
  public datePickerConfig: Partial<BsDatepickerInlineConfig> = {
    isAnimated: true,
    containerClass: 'theme-default'
  };
  constructor(private modalService: BsModalService) { }

  ngOnInit() {
  }

  public submitForm() {
    this.submitPage.emit(this.page);
  }


}
