import { GetPages } from './store/page/page.action';
import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { IAppState } from './store/app/app.state';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  constructor(
    private store: Store<IAppState>,
  ) { }

  ngOnInit(): void {
    this.store.dispatch(new GetPages());
  }
}
