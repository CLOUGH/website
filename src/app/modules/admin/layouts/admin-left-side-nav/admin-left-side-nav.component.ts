import { Component, OnInit } from '@angular/core';
import { faTachometerAlt, faCamera } from '@fortawesome/free-solid-svg-icons';
import { faFile } from '@fortawesome/free-regular-svg-icons';
import { faBloggerB } from '@fortawesome/free-brands-svg-icons';

@Component({
  selector: 'app-admin-left-side-nav',
  templateUrl: './admin-left-side-nav.component.html',
  styleUrls: ['./admin-left-side-nav.component.scss']
})
export class AdminLeftSideNavComponent implements OnInit {
  faTachometerAlt = faTachometerAlt;
  faFile = faFile;
  faBloggerB = faBloggerB;
  faCamera = faCamera;

  constructor() { }

  ngOnInit() {
  }

}
