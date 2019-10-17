import { ActivatedRoute } from '@angular/router';
import { Apollo } from 'apollo-angular';
import { Subscription } from 'rxjs';
import { Component, OnInit, OnDestroy } from '@angular/core';
import gql from 'graphql-tag';

const PagesQuery = gql`
  query getPage($id: ID!){
    page(id: $id){
      id
      name
      path
      description
      sections {
        id
        type
        content
        image
      }
    }
  }
`;

const updatePageQuery = gql`
  mutation updatePage($id: ID!, $page: PageInput!) {
    updatePage(id: $id, page: $page) {
      id
      name
      path
      description
      sections {
        id
        type
        content
        image
      }
    }
  }
`;
@Component({
  selector: 'app-page-edit-page',
  templateUrl: './page-edit-page.component.html',
  styleUrls: ['./page-edit-page.component.scss']
})
export class PageEditPageComponent implements OnInit, OnDestroy {

  public page: any;
  private querySubscription: Subscription;
  constructor(private apollo: Apollo, private route: ActivatedRoute) { }

  ngOnInit() {
    this.querySubscription = this.apollo
      .watchQuery<any>({
        query: PagesQuery,
        variables: { id: this.route.snapshot.paramMap.get('id') }
      })
      .valueChanges
      .subscribe(({ data, loading, errors }) => {
        this.page = data.page;
      });
  }

  ngOnDestroy(): void {
    this.querySubscription.unsubscribe();
  }

  public updatePage(event) {
    const { __typename, id, ...page } = event;

    page.sections = page.sections.map((data) => {
      const { __typename, links, ...section } = data;
      return section;
    });
    this.apollo
      .mutate<any>({
        mutation: updatePageQuery,
        variables: { id: this.page.id, page }
      })
      .subscribe(({ data, errors }) => {
        this.page = data.updatePage;
      });
  }
}
