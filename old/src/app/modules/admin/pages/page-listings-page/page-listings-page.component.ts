import { map, filter } from 'rxjs/operators';
import { Observable, Subscription } from 'rxjs';
import { Apollo } from 'apollo-angular';
import { Component, OnInit, OnDestroy } from '@angular/core';
import gql from 'graphql-tag';
import { BsModalRef, BsModalService } from 'ngx-bootstrap';
import { faTrash, faPen } from '@fortawesome/free-solid-svg-icons';

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
const DeletePageMutation = gql`
  mutation ($id: ID!){
    deletePage(id: $id) {
      id
    }
  }
`;

@Component({
  selector: 'app-page-listings-page',
  templateUrl: './page-listings-page.component.html',
  styleUrls: ['./page-listings-page.component.scss']
})
export class PageListingsPageComponent implements OnInit, OnDestroy {
  public pages: Array<any>;
  public loading: any;
  private pagesQuerySubscription: Subscription;
  public modalRef: BsModalRef | null;
  public pageDeletingIndex = -1;
  public faTrash = faTrash;
  public faPencil = faPen;

  constructor(private apollo: Apollo, private modalService: BsModalService) { }

  ngOnInit() {
    this.pagesQuerySubscription = this.apollo
      .watchQuery<any>({ query: PagesQuery })
      .valueChanges
      .subscribe(({ data, loading }) => {
        this.pages = data.pages;
        this.loading = loading;
      });
  }

  ngOnDestroy() {
    this.pagesQuerySubscription.unsubscribe();
  }

  promptDeleteModal(templateRef, pageIndex) {
    this.modalRef = this.modalService.show(templateRef);
    this.pageDeletingIndex = pageIndex;
  }

  deletePage() {
    this.modalRef.hide();
    this.apollo
      .mutate({mutation: DeletePageMutation, variables: { id: this.pages[this.pageDeletingIndex].id}})
      .subscribe(() => {
        this.pages = this.pages.filter((page, index) => index !== this.pageDeletingIndex);
        this.pageDeletingIndex = -1;
      });
  }
}
