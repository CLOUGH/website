import { Store, select } from '@ngrx/store';
import { Component, OnInit, HostListener, Input } from '@angular/core';

import { Page } from '../../../core/models/page';
import { IAppState } from '../../../store/app/app.state';
import { selectPages } from '../../../store/page/page.selector';

@Component({
  selector: 'app-main-layout',
  templateUrl: './main-layout.component.html',
  styleUrls: ['./main-layout.component.scss']
})
export class MainLayoutComponent implements OnInit {
  transparentNavBar = false;
  navbarFixedBehaviour = 'onScroll';
  fixedTop = false;
  pages: Page[] = [];
  constructor(
    private store: Store<IAppState>,
  ) { }

  ngOnInit() {
    this.store.pipe(select(selectPages)).subscribe((pages: Page[])=>{
      this.pages = pages;

    });
  }
  @Input('navbarFixedBehaviour')
  set setTopType(value: string) {
    this.navbarFixedBehaviour = value;
  }

  @HostListener('window:scroll', ['$event'])
  onScroll(event) {
    const rootElement = <Element>(document.documentElement || document.body.parentNode || document.body);

    if (rootElement.scrollTop > 100) {
      this.transparentNavBar = true;
      switch (this.navbarFixedBehaviour) {
        case 'onScroll':
          this.fixedTop = false;
          break;
        case 'always':
          this.fixedTop = true;
          break;
      }

    } else {
      this.transparentNavBar = false;
      switch (this.navbarFixedBehaviour) {
        case 'onScroll':
          this.fixedTop = true;
          break;
        case 'always':
          this.fixedTop = true;
          break;
      }
    }
  }

}
