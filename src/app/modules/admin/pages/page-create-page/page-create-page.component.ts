import { Component, OnInit } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Subscription } from 'rxjs';
import gql from 'graphql-tag';

const createPageMutation = gql`
  mutation updatePage( $page: PageInput!) {
    createPage(page: $page) {
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
  selector: 'app-page-create-page',
  templateUrl: './page-create-page.component.html',
  styleUrls: ['./page-create-page.component.scss']
})
export class PageCreatePageComponent implements OnInit {
  public page = {
    name: '',
    path: '',
    description: '',
    sections: []
  };
  constructor(private apollo: Apollo) { }

  ngOnInit() {
  }

  public createPage(page) {
    this.apollo
      .mutate<any>({
        mutation: createPageMutation,
        variables: { page }
      })
      .subscribe(({ data, errors }) => {
        this.page = data.createPage;
      })
      .unsubscribe();
  }

}
