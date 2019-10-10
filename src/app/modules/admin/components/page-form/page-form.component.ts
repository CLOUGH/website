import { BsModalRef, BsModalService } from 'ngx-bootstrap';

import { Component, OnInit, TemplateRef } from '@angular/core';
import { HeroSection, Section, TextSection } from './../../../../core/models/section';
import { faTrash } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-page-form',
  templateUrl: './page-form.component.html',
  styleUrls: ['./page-form.component.scss']
})
export class PageFormComponent implements OnInit {
  public sections: Section[] = [];
  public faTrash = faTrash;
  public deleteModalRef: BsModalRef | null;

  constructor(private modalService: BsModalService) { }

  ngOnInit() {
  }

  addSection(sectionType: String) {
    switch (sectionType) {
      case 'full screen hero':

        this.sections.push({
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

  public addTextSection () {
    this.sections.push({
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

  public submitForm(){
  }

  public promptDeleteOption(templateRef: TemplateRef<any>) {
    this.deleteModalRef = this.modalService.show(templateRef);
  }

  public deleteSection(index) {
    this.deleteModalRef.hide();
    this.sections.splice(index,1);
  }
}
