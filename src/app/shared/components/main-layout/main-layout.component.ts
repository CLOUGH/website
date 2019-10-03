import { Store, select } from '@ngrx/store';
import { Component, OnInit, HostListener, Input } from '@angular/core';

import { ILink } from './../../../store/link/link';
import { IAppState } from './../../../store/app/app.state';
import { selectLinks } from '../../../store/link/link.selector';

@Component({
  selector: 'app-main-layout',
  templateUrl: './main-layout.component.html',
  styleUrls: ['./main-layout.component.scss']
})
export class MainLayoutComponent implements OnInit {
  transparentNavBar = false;
  navbarFixedBehaviour = 'onScroll';
  fixedTop = false;
  links: ILink[] = [];
  constructor(
    private store: Store<IAppState>,
  ) { }

  ngOnInit() {
    this.store.pipe(select(selectLinks)).subscribe((links: ILink[])=>{
      console.log({links});
      this.links = links;
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
