import { Component, OnInit, Input, TemplateRef, ViewChild } from '@angular/core';
import { faTrash, faArrowsAlt, faExpand } from '@fortawesome/free-solid-svg-icons';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';
import { BsModalRef, BsModalService, ModalDirective } from 'ngx-bootstrap';

import { HeroSection, Section, TextSection } from './../../../../core/models/section';
@Component({
  selector: 'app-page-layout-editor',
  templateUrl: './page-layout-editor.component.html',
  styleUrls: ['./page-layout-editor.component.scss']
})
export class PageLayoutEditorComponent implements OnInit {
  @Input() sections: Array<Section>;
  public deleteModalRef: BsModalRef | null;
  public fullScreenModalRef: BsModalRef | null;
  @ViewChild(ModalDirective, { static: false }) fullScreenModal: ModalDirective;
  public faTrash = faTrash;
  public faExpand = faExpand;
  public faArrowsAlt = faArrowsAlt;
  public sectionDeletingIndex = -1;
  public isFullScreen = false;

  constructor(private modalService: BsModalService) {
  }

  ngOnInit() {
  }

  addSection(sectionType: String) {
    switch (sectionType) {
      case 'full screen hero':

        this.sections.push({
          order: this.sections.length - 1,
          type: 'hero',
          image: 'assets/placeholder.jpg',
          content: `
            <h1 class="display-1">Hero Heading</h1>
            <h4 class="display-5">Sub heading. Usually a description of the page or section.</h4>
          `,
          links: [
            {
              text: 'Link 1',
              link: '#'
            }
          ]
        } as HeroSection);
        break;
    }
  }

  public addTextSection() {
    this.sections.push({
      order: this.sections.length - 1,
      type: 'text',
      content: `
        <div class="bg-light pt-5 pb-5">
          <div class="container">
            <h2 class="display-2">Section Title</h2>
            <p>Nulla pariatur laborum dolor laborum eiusmod occaecat. Ullamco ipsum culpa eu pariatur labore. Consectetur nulla pariatur voluptate ut qui consectetur anim qui commodo aliquip fugiat deserunt ipsum officia. Ullamco in nulla aliquip tempor officia ipsum in amet eu aliquip sunt eiusmod mollit. Tempor non cillum non adipisicing. Laborum excepteur aliquip laborum nostrud dolor deserunt pariatur fugiat ipsum dolore. Ea qui in sit velit ea culpa dolore ea velit et.</p>
            <p>Tempor excepteur ex est occaecat excepteur minim minim nisi nisi eu. Consequat veniam veniam velit ut ex commodo. Nostrud laboris velit ad commodo culpa ut incididunt adipisicing irure anim ea reprehenderit.</p>
          </dv>
        </div>
      `
    } as TextSection);
  }
  public promptDeleteOption(templateRef: TemplateRef<any>, index) {
    this.deleteModalRef = this.modalService.show(templateRef);
    this.sectionDeletingIndex = index;
  }

  public deleteSection(index) {
    this.deleteModalRef.hide();
    this.sections[this.sectionDeletingIndex] = {
      ...this.sections[this.sectionDeletingIndex],
      updated: true,
      deleted: true
    };

    if (!this.sections[this.sectionDeletingIndex].id) {
      this.sections.splice(this.sectionDeletingIndex, 1);
    }
    this.sectionDeletingIndex = -1;

  }

  public sectionUpdated(section, index) {
    this.sections[index] = {
      ...section,
      updated: true
    };
  }

  private drop(event: CdkDragDrop<Section[]>) {
    moveItemInArray(this.sections, event.previousIndex, event.currentIndex);
    this.sections = this.sections.map((section, index) => {
      return { ...section, order: index, updated: true }
    });
  }

  public expandToFullScreen(templateRef: TemplateRef<any>) {
    this.fullScreenModal.show();

    this.isFullScreen = true;
    this.fullScreenModal.onHide.subscribe((event) => {
      this.isFullScreen = false;

    });
  }

  public removeSection(index: number) {

  }
}
