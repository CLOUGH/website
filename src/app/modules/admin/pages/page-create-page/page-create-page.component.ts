import { Component, OnInit } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Subscription } from 'rxjs';
import gql from 'graphql-tag';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

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
        order
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
  constructor(
    private apollo: Apollo,
    private router: Router,
    private toastr: ToastrService
  ) { }

  ngOnInit() {
  }

  public createPage(page) {
    console.log('reachj123adfaa')
    this.apollo
      .mutate<any>({
        mutation: createPageMutation,
        variables: {
          page: {
            ...page,
            sections: page.sections.map(section => {
              const {updated, links, ...sectionData } = section;
              return sectionData;
            })
          }
        }
      })
      .subscribe(({errors, data}) => {
        if (errors) {
          this.toastr.error(
            errors.map(error => error.message).join('. '),
            'Fail Creating Page'
          );
        } else {
          this.page = data.createPage;
          this.toastr.success('Successfully created page', 'Success');
          this.router.navigateByUrl('/admin/pages');
        }
      }, (error) => {
        this.toastr.error(
          error,
          'Fail Creating Page'
        );
      })
  }

}
