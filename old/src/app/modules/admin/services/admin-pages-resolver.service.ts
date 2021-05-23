import { map } from 'rxjs/operators';
import { Apollo } from 'apollo-angular';
import { PageService } from './../../page/services/page.service';
import { Router, Resolve, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { Injectable } from '@angular/core';
import gql from 'graphql-tag';

const PageQuery = gql`
  query getPages($path: String!) {
    pages(path: $path) {
      name
      description
      path
      published
      status
      sections {
        id
        type
        content
        image
        video
        order
      }
    }
  }
`;

@Injectable({
  providedIn: 'root'
})
export class AdminPagesResolverService implements Resolve<any> {

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
