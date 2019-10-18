import { Injectable } from '@angular/core';
import { Location } from '@angular/common';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';
import { of } from 'rxjs';
import { map } from 'rxjs/operators';
import { PageService } from './page.service';
import gql from 'graphql-tag';
import { Apollo } from 'apollo-angular';

const PageQuery = gql`
  query getPageByPath($path: String!) {
    pageByPath(path: $path) {
      name
      description
      path
      sections {
        type
        content
        image
      }
    }
  }
`;

@Injectable({
  providedIn: 'root'
})
export class PageResolverService implements Resolve<any> {

  constructor(
    private pageService: PageService,
    private router: Router,
    private location: Location,
    private apollo: Apollo
  ) { }

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    const path = '/' + route.url.join('/');
    return this.apollo.query<any>({
      query: PageQuery,
      variables: { path }
    }).pipe(map(data => data.data.pageByPath))
  }
}
