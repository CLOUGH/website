import { Component, OnInit, HostListener } from '@angular/core';

@Component({
  selector: 'app-main-layout',
  templateUrl: './main-layout.component.html',
  styleUrls: ['./main-layout.component.scss']
})
export class MainLayoutComponent implements OnInit {
  transparentNavBar = false;

  constructor() { }

  ngOnInit() {
  }

  @HostListener('window:scroll', ['$event'])
  onScroll(event) {
    const rootElement = <Element>(document.documentElement || document.body.parentNode || document.body);

    if (rootElement.scrollTop > 100) {
      this.transparentNavBar = true;
    } else {
      this.transparentNavBar = false;
    }
  }

}
