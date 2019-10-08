import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-page-layout-editor',
  templateUrl: './page-layout-editor.component.html',
  styleUrls: ['./page-layout-editor.component.scss']
})
export class PageLayoutEditorComponent implements OnInit {
  public sections = [];

  constructor() { }

  ngOnInit() {
  }

  addSection(sectionType: String) {

    switch (sectionType) {
      case 'full screen hero':
        this.sections.push({
          hero: {
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
          }
        });
        break;
    }

  }

}
