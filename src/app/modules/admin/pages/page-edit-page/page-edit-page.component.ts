import { ActivatedRoute } from '@angular/router';
import { Apollo } from 'apollo-angular';
import { Subscription } from 'rxjs';
import { Component, OnInit, OnDestroy } from '@angular/core';
import gql from 'graphql-tag';
import { ToastrService } from 'ngx-toastr';

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
        video
        order
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
        video
        order
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
  constructor(
    private apollo: Apollo,
    private route: ActivatedRoute,
    private toastr: ToastrService
  ) { }

  ngOnInit() {
    this.apollo
      .query<any>({
        query: PagesQuery,
        variables: { id: this.route.snapshot.paramMap.get('id') }
      })
      .subscribe(({ data, loading, errors }) => {
        this.page = data.page;
      });
  }

  ngOnDestroy(): void {
    // this.querySubscription.unsubscribe();
  }

  public updatePage(event) {
    const { __typename, id, ...page } = event;
    page.sections = page.sections
      .filter(section => !section.id || section.updated === true)
      .map((data) => {
        const { __typename, updated, links, ...section } = data;
        return section;
      });

    this.apollo
      .mutate<any>({
        mutation: updatePageQuery,
        variables: { id: this.page.id, page }
      })
      .subscribe(({ data, errors }) => {
        if (errors) {
          this.toastr.error(
            errors.map(error => error.message).join('. '),
            'Fail Updating Page'
          );
        } else {
          this.toastr.success('Successfully updated page', 'Success');
          this.page = data.updatePage;
        }
      }, (error) => {
        this.toastr.error(
          error,
          'Fail Updating Page'
        );
      });
  }
}
