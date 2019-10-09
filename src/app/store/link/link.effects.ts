import { ILink } from './link';
import { Injectable } from "@angular/core";
import { Effect, ofType, Actions } from '@ngrx/effects';
import { Store, select } from '@ngrx/store';
import { of } from 'rxjs';
import { switchMap, map, withLatestFrom } from 'rxjs/operators';

import { IAppState } from '../app/app.state';
import { PageService } from '../../modules/page/services/page.service';
import { ELinkActions, GetLinks, GetLinksSuccess } from './link.action';

@Injectable()
export class LinkEffect {

  constructor(
    private pagesService: PageService,
    private actions$: Actions,
    private store: Store<IAppState>
  ) { }

  @Effect()
  getLinks = this.actions$.pipe(
    ofType<GetLinks>(ELinkActions.GetLinks),
    switchMap(() => this.pagesService.getLinks()),
    switchMap((links: ILink[]) => {
      return of(new GetLinksSuccess(links));
    })
  );
}
