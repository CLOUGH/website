import { HeroSection } from './../../../core/models/section';
import { Component, OnInit, Input, Output, EventEmitter, TemplateRef } from '@angular/core';
import { faImage,faTrash } from '@fortawesome/free-solid-svg-icons';
import { BsModalService, BsModalRef } from 'ngx-bootstrap';

@Component({
  selector: 'app-full-screen-hero',
  templateUrl: './full-screen-hero.component.html',
  styleUrls: ['./full-screen-hero.component.scss']
})
export class FullScreenHeroComponent implements OnInit {
  public faImage = faImage;
  public faTrash = faTrash;
  public tinymceConfig = {
    base_url: '/tinymce', // Root for resources
    suffix: '.min',
  };
  public modalRef: BsModalRef | null;
  public deleteModalRef: BsModalRef | null;
  public imageUrl: string;

  @Input() public section: HeroSection;
  @Output() public sectionChange: EventEmitter<HeroSection> = new EventEmitter();
  @Output() public removeSection: EventEmitter<boolean> = new EventEmitter();

  constructor(private modalService: BsModalService) { }

  ngOnInit() {
  }


  public updateContent(event) {

    this.sectionChange.emit({
      ...this.section,
      content: event.editor.getContent()
    });
  }

  public openGallary(templateRef: TemplateRef<any>) {
    this.modalRef = this.modalService.show(templateRef);
    this.imageUrl = this.section.image;
  }
  public promptDeleteOption(templateRef: TemplateRef<any>) {
    this.deleteModalRef = this.modalService.show(templateRef);
  }

  public saveImageUrl() {
    this.sectionChange.emit({
      ...this.section,
      image: this.imageUrl
    });
    this.modalRef.hide();
  }

  public deleteSection() {
    this.removeSection.emit(true);
    this.deleteModalRef.hide();
  }
}
