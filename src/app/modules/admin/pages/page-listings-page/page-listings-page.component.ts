import { map } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { Apollo } from 'apollo-angular';
import { Component, OnInit, OnDestroy } from '@angular/core';
import gql from 'graphql-tag';
import { Subscription } from 'apollo-client/util/Observable';

const PagesQuery = gql`
  query{
    pages{
      id
      name
      path
      description
    }
  }
`;

@Component({
  selector: 'app-page-listings-page',
  templateUrl: './page-listings-page.component.html',
  styleUrls: ['./page-listings-page.component.scss']
})
export class PageListingsPageComponent implements OnInit, OnDestroy {
  public data: Observable<any>;

  constructor(private apollo: Apollo) { }

  ngOnInit() {
    this.data = this.apollo
      .watchQuery({query: PagesQuery})
      .valueChanges
      .pipe(map(({data}) => (data as any).pages));
  }

  ngOnDestroy() {
    // this.querySubscription.unsubscribe();
  }

}
