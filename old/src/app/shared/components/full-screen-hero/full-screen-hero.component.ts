import { HeroSection } from './../../../core/models/section';
import { Component, OnInit, Input, Output, EventEmitter, TemplateRef } from '@angular/core';
import { faImage, faTrash,faVideo } from '@fortawesome/free-solid-svg-icons';
import { BsModalService, BsModalRef } from 'ngx-bootstrap';

@Component({
  selector: 'app-full-screen-hero',
  templateUrl: './full-screen-hero.component.html',
  styleUrls: ['./full-screen-hero.component.scss']
})
export class FullScreenHeroComponent implements OnInit {
  public faImage = faImage;
  public faTrash = faTrash;
  public faVideo = faVideo;
  public modalRef: BsModalRef | null;
  public imageUrl: string;
  public videoUrl: string;

  @Input() public section: HeroSection;
  @Input() public readonly: boolean;
  @Output() public sectionChange: EventEmitter<HeroSection> = new EventEmitter();
  @Output() public removeSection: EventEmitter<boolean> = new EventEmitter();

  constructor(private modalService: BsModalService) { }

  ngOnInit() {
  }


  public updateContent(content) {
    this.sectionChange.emit({
      ...this.section,
      updated: true,
      content: content
    });
  }

  public openGallery(templateRef: TemplateRef<any>) {
    this.modalRef = this.modalService.show(templateRef);
    this.imageUrl = this.section.image;
  }
}
