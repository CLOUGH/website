import { Apollo } from 'apollo-angular';
import { Page } from '../../core/models/page';
import { Injectable } from "@angular/core";
import { Effect, ofType, Actions } from '@ngrx/effects';
import { Store, select } from '@ngrx/store';
import { of } from 'rxjs';
import { switchMap, map, withLatestFrom } from 'rxjs/operators';

import { IAppState } from '../app/app.state';
import { PageService } from '../../modules/page/services/page.service';
import { EPageActions, GetPages, GetPagesSuccess } from './page.action';
import gql from 'graphql-tag';

const GetPagesQuery = gql`
  query {
    pages{
      path
      name
      description
    }
  }
`;

@Injectable()
export class PageEffect {

  constructor(
    private pagesService: PageService,
    private actions$: Actions,
    private store: Store<IAppState>,
    private appollo: Apollo
  ) { }

  @Effect()
  getPages = this.actions$.pipe(
    ofType<GetPages>(EPageActions.GetPages),
    switchMap(() => {
      return this.appollo
        .query<any>({query: GetPagesQuery})
        .pipe(map(({data}) => data.pages));
    }),
    switchMap((Pages: Page[]) => {
      return of(new GetPagesSuccess(Pages));
    })
  );
}
