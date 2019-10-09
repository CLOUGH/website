import { Component, OnInit } from '@angular/core';
import { HeroSection, Section } from './../../../../core/models/section';

@Component({
  selector: 'app-page-form',
  templateUrl: './page-form.component.html',
  styleUrls: ['./page-form.component.scss']
})
export class PageFormComponent implements OnInit {
  public sections: Section[] = [];

  constructor() { }

  ngOnInit() {
  }

  addSection(sectionType: String) {
    switch (sectionType) {
      case 'full screen hero':

        this.sections.push({
          type: 'hero',
          image: 'assets/placeholder.jpg',
          content: `
            <h1 class="display1">Hero Heading</h1>
            <h3 class="display4">Sub heading. Usually a description of the page or section.</h3>
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
  removeSection(index: number) {
    this.sections.splice(index,1);
  }
}
