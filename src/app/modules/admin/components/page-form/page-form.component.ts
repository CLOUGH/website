import { BsModalRef, BsModalService } from 'ngx-bootstrap';

import { Component, OnInit, TemplateRef, Input, Output, EventEmitter } from '@angular/core';
import { HeroSection, Section, TextSection } from './../../../../core/models/section';
import { faTrash, faArrowsAlt } from '@fortawesome/free-solid-svg-icons';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

@Component({
  selector: 'app-page-form',
  templateUrl: './page-form.component.html',
  styleUrls: ['./page-form.component.scss']
})
export class PageFormComponent implements OnInit {

  @Input() public page: any;
  @Input() public submitButtonLabel: any;
  @Output() public submitPage: EventEmitter<any> = new EventEmitter();

  public sections: Section[] = [];
  public faTrash = faTrash;
  public faArrowsAlt = faArrowsAlt;
  public deleteModalRef: BsModalRef | null;
  public sectionDeletingIndex = -1;

  constructor(private modalService: BsModalService) { }

  ngOnInit() {
  }

  addSection(sectionType: String) {
    switch (sectionType) {
      case 'full screen hero':

        this.page.sections.push({
          order: this.page.sections.length - 1,
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
    this.page.sections.push({
      order: this.page.sections.length - 1,
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
  removeSection(index: number) {

  }

  public submitForm() {
    this.submitPage.emit(this.page);
  }

  public promptDeleteOption(templateRef: TemplateRef<any>, index) {
    this.deleteModalRef = this.modalService.show(templateRef);
    this.sectionDeletingIndex = index;
  }

  public deleteSection(index) {
    this.deleteModalRef.hide();
    this.page.sections[this.sectionDeletingIndex] = {
      ...this.page.sections[this.sectionDeletingIndex],
      updated: true,
      deleted: true
    };

    if (!this.page.sections[this.sectionDeletingIndex].id) {
      this.page.sections.splice(this.sectionDeletingIndex, 1);
    }
    this.sectionDeletingIndex = -1;

  }

  public sectionUpdated(section, index) {
    this.page.sections[index] = {
      ...section,
      updated: true
    };
  }

  drop(event: CdkDragDrop<Section[]>) {
    moveItemInArray(this.page.sections, event.previousIndex, event.currentIndex);
    this.page.sections = this.page.sections.map((section, index) => {
      return { ...section, order: index, updated:true }
    });

    console.log({section: this.page.sections});

  }
}
