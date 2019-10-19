import { VideoFullScreenHeroSection } from './../../../core/models/section';
import { BsModalRef, BsModalService } from 'ngx-bootstrap';
import { Component, OnInit, Input, Output, EventEmitter, TemplateRef, ViewChild, ElementRef } from '@angular/core';
import { faTrash, faVideo, faPause, faPlay } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-video-full-screen-hero',
  templateUrl: './video-full-screen-hero.component.html',
  styleUrls: ['./video-full-screen-hero.component.scss']
})
export class VideoFullScreenHeroComponent implements OnInit {
  public faTrash = faTrash;
  public faVideo = faVideo;
  public faPause = faPause;
  public faPlay = faPlay;
  public modalRef: BsModalRef | null;
  public videoUrl: string;
  public isVideoPlaying = true;
  @ViewChild('video', { static: true }) video: ElementRef;

  @Input() public section: VideoFullScreenHeroSection;
  @Input() public readonly: boolean;
  @Output() public sectionChange: EventEmitter<VideoFullScreenHeroSection> = new EventEmitter();
  @Output() public removeSection: EventEmitter<boolean> = new EventEmitter();

  constructor(private modalService: BsModalService) { }
  public updateContent(content) {
    this.sectionChange.emit({
      ...this.section,
      updated: true,
      content: content
    });
  }

  ngOnInit() {
  }

  public saveVideoUrl() {
    this.sectionChange.emit({
      ...this.section,
      video: this.videoUrl,
      updated: true,
    });
    this.modalRef.hide();
  }

  public addVideo(templateRef: TemplateRef<any>) {
    this.modalRef = this.modalService.show(templateRef);
    this.videoUrl = this.section.video;
  }

  public toggleVideoPlay() {
    if (this.isVideoPlaying) {
      this.video.nativeElement.pause();
    } else {
      this.video.nativeElement.play();
    }
    this.isVideoPlaying = !this.isVideoPlaying;
  }
}
